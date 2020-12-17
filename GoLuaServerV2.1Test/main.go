// test server入口文件
package main

import (
	"GoLuaServerV2.1Test/Core/Lua"
	"GoLuaServerV2.1Test/Core/NetWork"
	"GoLuaServerV2.1Test/Core/Utils/zLog"
	"flag"
	//"util/logs"
	//"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/go-ini/ini"
	oldLog "log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"sync"
	"time"
)


var GameManagerLua *Lua.MyLua // 公共部分lua脚本

var GameServerAddress string
//var GameServerWebSocketAddress string

var WebSocketPort int
var SocketPort int
var UdpPort int
var ClientStart int
var ClientEnd int

var ShowLog int
var UDPSocket = true
var WebSocketServer = true	// websocket 开启
var SocketServer = true		// socket 开启


var udpClients []*NetWork.UDPClient
var tcpClients []*NetWork.TCPClient
var wsClients []*NetWork.WSClient

var GlobalMutex sync.Mutex // 全局互斥锁
//var GlobalClients map[*Client] interface{}  // 全局client

// 程序入口
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//远程获取pprof数据打开浏览器http://localhost:8080/debug/pprof/
	go func() {
		oldLog.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("配置文件出错")
		return
	}

	//LoginServer := f.Section("Server").Key("LoginServer").Value()
	GameServerAddress = f.Section("Server").Key("GameServerAddress").Value()
	//GameServerWebSocketAddress = f.Section("Server").Key("GameServerWebSocketAddress").Value()
	//ClientStart,err   := f.Section("Server").Key("ClientStart").Int()
	//ClientEnd ,err   := f.Section("Server").Key("ClientEnd").Int()
	ShowLog ,err   = f.Section("Server").Key("ShowLog").Int()
	UDPSocket,err   = f.Section("Server").Key("UDPSocket").Bool()
	WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()

	// -------------------------读取命令行参数--------------------------
	wsPort := flag.Int("WebSocketPort", 0, "")
	sPort := flag.Int("SocketPort", 0, "")
	uPort := flag.Int("UdpPort", 0, "")
	start := flag.Int("ClientStart", 0, "")
	end := flag.Int("ClientEnd", 0, "")
	flag.Parse()
	WebSocketPort = *wsPort
	SocketPort = *sPort
	UdpPort = * uPort
	ClientStart = * start
	ClientEnd = *end
	zLog.ServerPort = SocketPort

	if WebSocketPort == 0 || SocketPort == 0 ||  UdpPort ==0 || ClientStart == 0 || ClientEnd == 0 {

		for{
			fmt.Println("缺少命令行参数！ 参数要设置类似 -WebSocketPort=8089 -SocketPort=8123 -UdpPort=8124 -ClientStart=1 -ClientEnd=100")
			time.Sleep(time.Second)

		}
	}
	fmt.Println("-------------------服务器初始化---------------------------")
	initVar()
	fmt.Println("-------------------Lua逻辑处理器---------------------------")
	GameManagerInit()

	fmt.Println("max conn start :", ClientStart, "--------", ClientEnd)
	StartClient()

	for{
		//GetStaticPrint()
		time.Sleep(time.Second)
	}
}


//----------------------------变量的初始化---------------------------------------------------------------
func initVar()  {
	//Client.AllClientsList = make(map[*Client.Client]struct{})
	//Client.AllUserClientList = make(map[uint32]*Client.Client)
	//Games.AllGamesList = make(map[int]*Games.Games)
	Lua.InitGlobalVar()
}

//-----------------------------------游戏公共逻辑处理-------------------------------------------------------
func GameManagerInit() {

	GameManagerLua = Lua.NewMyLua()

	GameManagerLua.Init() // 绑定lua脚本
	//Lua.GoCallLuaTest(GameManagerLua.L,1)

	Lua.GameManagerLuaHandle = GameManagerLua // 把句柄传递给lua保存一份
	//GameManagerLuaReloadTime = GlobalVar.LuaReloadTime
	GameManagerLua.GoCallLuaSetStringVar("ServerIP_Port", GameServerAddress+ ":" + strconv.Itoa(SocketPort)) //把服务器地址传递给lua

}

//-----------------------------------建立服务器的网络功能---------------------------------------------------------------
func StartClient() {
	//GlobalClients = make(map[*Client]interface{},0)
	Lua.ClientStart = ClientStart
	//UDPSocket := false
	if SocketServer {
		// socket client----------------------------------------------------------
		client := new(NetWork.TCPClient)
		client.Addr = GameServerAddress+":"+ strconv.Itoa(SocketPort)
		client.ConnNum = 1  //废了
		client.ConnectInterval = 3 * time.Second	// 客户端自动重连
		client.PendingWriteNum = 1000	// 发送缓冲区
		client.LenMsgLen = 4
		client.MaxMsgLen = math.MaxUint32
		client.NewAgent = func(conn *NetWork.TCPConn,index int) NetWork.Agent {
			a := Lua.NewMyTcpServer(conn,GameManagerLua) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}

		fmt.Println("开始连接", client.Addr)
		client.Start(ClientStart, ClientEnd)
		tcpClients = append(tcpClients, client)
	}
	if WebSocketServer {
		// websocket client------------------------------------------------------------------


		wsclient := new(NetWork.WSClient)
		wsclient.Addr = "ws://"+GameServerAddress+":"+ strconv.Itoa(WebSocketPort)+"/"
		wsclient.ConnNum = 1
		wsclient.ConnectInterval = 3 * time.Second// 客户端自动重连
		wsclient.PendingWriteNum = 1000 	// 发送缓冲区
		wsclient.HandshakeTimeout = 10 * time.Second
		wsclient.MaxMsgLen = math.MaxUint32
		wsclient.NewAgent = func(conn *NetWork.WSConn,index int) NetWork.Agent {
			a := Lua.NewMyTcpServer(conn,GameManagerLua)
			return a
		}

		fmt.Println("开始连接",wsclient.Addr)
		wsclient.Start(ClientStart, ClientEnd)
		wsClients = append(wsClients, wsclient)
	}
	if UDPSocket {
		// socket client----------------------------------------------------------
		client := new(NetWork.UDPClient)
		client.Addr = GameServerAddress+":"+ strconv.Itoa(UdpPort)
		client.ConnectInterval = 3 * time.Second	// 客户端自动重连
		client.PendingWriteNum = 1000	// 发送缓冲区
		client.LenMsgLen = 4
		client.MaxMsgLen = math.MaxUint32
		client.NewAgent = func(conn *NetWork.UdpConn,index int) NetWork.Agent {
			a := Lua.NewMyUdpServer(conn,GameManagerLua) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}

		fmt.Println("开始连接", client.Addr)
		client.Start(ClientStart, ClientEnd)
		udpClients = append(udpClients, client)
	}

}


