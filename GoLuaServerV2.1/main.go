//------------------------------------------------------------------
// game server入口文件
//------------------------------------------------------------------

package main

import (
	"GoLuaServerV2.1/Core/Lua"
	"GoLuaServerV2.1/Core/NetWork"
	"GoLuaServerV2.1/Core/Utils/ip"
	"GoLuaServerV2.1/Core/Utils/zLog"
	"GoLuaServerV2.1/Core/Utils/ztimer"
	"fmt"
	"github.com/go-ini/ini"
	"math"
	"strconv"
	"time"

	"flag"
	oldLog "log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)


// ---------------------------程序入口-----------------------------------
var wsServer *NetWork.WSServer
var tcpServer *NetWork.TCPServer
var udpServer *NetWork.UDPServer

var WebSocketServer = true	// websocket 开启
var SocketServer = true		// socket 开启
var UDPServer = true		// udp socket 开启

var WebSocketPort int
var SocketPort int
var UdpPort int
var GameRoomServerID int	// 游戏原有的ServerID(来源于GameRoomInfo中对应的)
var ServerAddress string    // ServerAddress 服务器地址
var ServerTypeName string    // ServerAddress 服务器地址

var err error
var LuaReloadTime int //lua脚本当前最新版本的时间戳，后台设置的，保存在服务器中，定期去更新一次


var GameManagerLua *Lua.MyLua    // 公共部分lua脚本
var GameManagerLuaReloadTime int // 公共逻辑处理的lua更新时间


func main() {

	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	//远程获取pprof数据打开浏览器http://localhost:8081/debug/pprof/
	go func() {
		oldLog.Println(http.ListenAndServe("localhost:8081", nil))
	}()

	fmt.Println("------------------	首先读取命令行参数	---------------------------")
	wsPort := flag.Int("WebSocketPort", 0, "")
	iPort := flag.Int("SocketPort", 0, "")
	uPort := flag.Int("UdpPort", 0, "")
	iGameServerID := flag.Int("ServerID", 0, "")
	sServerTypeName := flag.String("ServerTypeName", "", "定义游戏服务器的类型")
	flag.Parse()
	WebSocketPort = *wsPort
	SocketPort = *iPort
	UdpPort = * uPort
	GameRoomServerID = *iGameServerID
	ServerTypeName = *sServerTypeName

	fmt.Println("WebSocketPort=",WebSocketPort,"SocketPort=",SocketPort,"GameServerID=",GameRoomServerID)
	if WebSocketPort==0 || UdpPort ==0 ||SocketPort==0{
		for{
			fmt.Println("缺少命令行参数！ 参数要设置类似 -SocketPort=9001 -UdpPort=10001 -WebSocketPort=11001   -ServerTypeName=Game")
			time.Sleep(time.Second)
		}
	}
	zLog.ServerPort = SocketPort // 传递给log日志，让日志记录的时候区分服务器端口
	fmt.Println("-------------------	读取本地配置文件	---------------------------")
	initSetting()

	//fmt.Println("-------------------	服务器初始化		---------------------------")
	initVar()
	Lua.QueueInit()

	fmt.Println("-------------------	Lua逻辑处理器		---------------------------")
	GameManagerInit()
	//fmt.Println("Lua 代码初始化完成")


	//fmt.Println("-------------------多核房间逻辑处理器---------------------------")
	//CreateGoroutineForLuaGameTable()

	//fmt.Println("-------------------	启动gameManager	---------------------------")
	GameManagerLua.GoCallLuaLogic("Main","GoCallLuaStartAllServers")
	//StartMultiThreadChannelPlayerToGameManager()

	TimerCommonLogicStart()

	//fmt.Println("-------------------	读取数据库设置---------------------------")
	//UpdateLuaReload()

	fmt.Println("---------------- 游戏服务器开始建立连接	  -------------------------")
	NetWorkServerStart()


	// ----------------------主循环计时器----------------------------------------
	//tickerCheckUpdateData := time.NewTicker(time.Second * 5)		// 每5秒触发一次计时器，用于定期更新数据库的通用配置信息
	//defer tickerCheckUpdateData.Stop()

	for {
		//select {
		//case <-tickerCheckUpdateData.C:
		//	// 定期更新后台的配置数据
		//	UpdateLuaReload()
		//}
		//start:= ztimer.GetOsTimeMillisecond()
		//GameManagerLua.GoCallLuaLogic("MultiThreadChannelPlayerToGameManager") //公共逻辑处理循环

		//UpdateDBSetting()
		//ztimer.CheckRunTimeCost(func() {
		//		GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 房间的run
		//	}, "房间循环GoCallLuaGoRoutineForLuaGameTable"		)
		//startTime := ztimer.GetOsTimeMillisecond()
		//GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 房间的run
		//if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		//	zLog.PrintfLogger("--------!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!![ 警告 ] 房间循环 消耗时间过长: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
		//}
		runtime.GC()
		time.Sleep(time.Millisecond * 1000)                                //给其他协程让出10秒的时间， 这个可以后期调整

		//fmt.Printf("go run  %d \n" , ztimer.GetOsTimeMillisecond() )
		//end:= ztimer.GetOsTimeMillisecond()
		//if end - start > 120 {
			//fmt.Println("一个循环用时", end-start)
		//}
	}
}
//-----------------------------本地配置文件---------------------------------------------------
func initSetting()  {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("读取配置文件出错")
		return
	}

	zLog.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()
	WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()
	UDPServer,err   = f.Section("Server").Key("UDPSocket").Bool()
	ServerAddress = f.Section("Server").Key("ServerAddress").String()

	zLog.CheckError(err)

	ServerAddress = string(ip.GetInternal(0)) // 获取本机内网ip
	fmt.Println("本机内网ip :",ServerAddress)
}

//----------------------------变量的初始化---------------------------------------------------------------
func initVar()  {
	Lua.InitGlobalVar()
}


//-------------------------定期更新后台的配置数据---------------------------------------------------------
func UpdateLuaReload() {
	// 以后增加读取数据库的代码
	//...

	// 从服务器更新lua热更新的时间戳
	LuaReloadTime = 1111
	GameManagerLuaReloadCheck() //共有逻辑检查一下是否需要更新, 玩家部分每个连接自己检查
	//GoroutineTableLuaReloadCheck()
}

//-----------------------------------建立服务器的网络功能---------------------------------------------------------------
func NetWorkServerStart()  {

	ServerAddress = "0.0.0.0"

	if WebSocketServer {
		// websocket 服务器开启---------------------------------
		wsServer = new(NetWork.WSServer)
		wsServer.Addr = ServerAddress + ":"+strconv.Itoa(WebSocketPort)
		fmt.Println("websocket 绑定："+ wsServer.Addr)
		wsServer.MaxConnNum = int(math.MaxInt32)
		wsServer.PendingWriteNum = 1000			// 发送区缓存
		wsServer.MaxMsgLen = 4096
		wsServer.HTTPTimeout = 10 * time.Second
		wsServer.CertFile = ""
		wsServer.KeyFile = ""
		wsServer.NewAgent = func(conn *NetWork.WSConn) NetWork.Agent {
			ServerId := Lua.GetServerUid()
			a := Lua.NewMyServer(conn,ServerId) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}

		wsServer.Start()
	}
	if SocketServer{
		// socket 服务器开启----------------------------------
		tcpServer = new(NetWork.TCPServer)
		tcpServer.Addr = ServerAddress +":"+strconv.Itoa(SocketPort)
		fmt.Println("socket 绑定："+ tcpServer.Addr)
		tcpServer.MaxConnNum = int(math.MaxInt32)
		tcpServer.PendingWriteNum = 1000 // 发送区缓存
		tcpServer.LenMsgLen = 4
		tcpServer.MaxMsgLen = math.MaxUint32
		tcpServer.NewAgent = func(conn *NetWork.TCPConn) NetWork.Agent {
			ServerId := Lua.GetServerUid()
			a := Lua.NewMyServer(conn,ServerId) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}
		tcpServer.Start()
	}
	if UDPServer{
		// socket udp 服务器开启----------------------------------
		udpServer = new(NetWork.UDPServer)
		udpServer.Addr = ServerAddress +":"+strconv.Itoa(UdpPort)
		fmt.Println("socket udp 绑定：", udpServer.Addr)
		udpServer.MaxConnNum = int(math.MaxInt32)
		udpServer.PendingWriteNum = 1000 // 发送区缓存
		udpServer.LenMsgLen = 4
		udpServer.MaxMsgLen = math.MaxUint32
		udpServer.NewAgent = func(conn *NetWork.UdpConn) NetWork.Agent {
			a := Lua.NewMyUdpServer(conn) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}
		udpServer.Start()
		Lua.MyUdpListen = udpServer.Listen
	}
}

//-----------------------------------游戏公共逻辑处理-------------------------------------------------------
func GameManagerInit() {

	GameManagerLua = Lua.NewMyLua()

	GameManagerLua.Init() // 绑定lua脚本
	//Lua.GoCallLuaTest(GameManagerLua.L,1)

	Lua.GameManagerLuaHandle = GameManagerLua // 把句柄传递给lua保存一份
	GameManagerLuaReloadTime = LuaReloadTime
	GameManagerLua.GoCallLuaSetStringVar("GlobalVar","ServerIP_Port", ServerAddress+ ":" + strconv.Itoa(SocketPort)) 	//把服务器地址传递给lua
	GameManagerLua.GoCallLuaSetIntVar("GlobalVar","GameRoomServerID", GameRoomServerID) 								//把服务器地址传递给lua
	GameManagerLua.GoCallLuaSetStringVar("GlobalVar","ServerTypeName", ServerTypeName) 								//把参数传递给lua
}

// 检查通用逻辑部分的lua是否需要更新
func GameManagerLuaReloadCheck() {
	if GameManagerLuaReloadTime == LuaReloadTime {
		//return
	}
	// 如果跟本地的lua时间戳不一致，就更新
	err = GameManagerLua.GoCallLuaReload()
	if err == nil{
		// 热更新成功
		GameManagerLuaReloadTime = LuaReloadTime
	}
}

func TimerCommonLogicStart() {
	// -------------------创建计时器，定期去run公共逻辑---------------------
	//ztimer.TimerCheckUpdate(func() {
	//	connectShow, successSendMsg, successRecMsg, WriteChan := GetAllConnectMsg()
	//	memoryShow, heapInUse := zLog.GetSysMemInfo()
	//
	//	zLog.PrintfLogger("[%s] 拼接成功%d    标识错误%d   %s   %s  处理消息平均时间：%d  ", ServerAddress+":"+strconv.Itoa(SocketPort),
	//		Lua.StaticDataPackagePasteSuccess, Lua.StaticDataPackageHeadFlagError, connectShow, memoryShow, Lua.StaticNetWorkReceiveToSendCostTime)
	//
	//	// 把服务器的状态信息，传递给lua
	//	GameManagerLua.GoCallLuaSetIntVar("ServerStateSendNum", successSendMsg)
	//	GameManagerLua.GoCallLuaSetIntVar("ServerStateReceiveNum", successRecMsg)
	//	GameManagerLua.GoCallLuaSetIntVar("ServerSendWriteChannelNum", WriteChan)
	//	GameManagerLua.GoCallLuaSetIntVar("ServerDataHeadErrorNum", Lua.StaticDataPackageHeadFlagError)
	//	GameManagerLua.GoCallLuaSetIntVar("ServerHeapInUse", heapInUse)
	//	GameManagerLua.GoCallLuaSetIntVar("ServerNetWorkDelay", Lua.StaticNetWorkReceiveToSendCostTime)
	//
	//	ztimer.CheckRunTimeCost(func() {
	//		GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogicRun") //公共逻辑处理循环
	//	}, "GoCallLuaCommonLogicRun")
	//
	//}, 10)

	// ---------------------创建计时器，定期去保存服务器状态---------------------
	//ztimer.TimerCheckUpdate(func() {
	//	ztimer.CheckRunTimeCost(func() {
	//		GameManagerLua.GoCallLuaLogic("GoCallLuaSaveServerState") // 保存服务器的状态
	//	}, "GoCallLuaSaveServerState")
	//
	//
	//	runtime.GC()
	//}, 60)	// 1 分钟

	//---------------------创建计时器，定期去更新lua脚本reload---------------------
	ztimer.TimerCheckUpdate(func() {
		// 定期更新后台的配置数据
		ztimer.CheckRunTimeCost(func() {
			UpdateLuaReload()
		}, "UpdateLuaReload")

	},  10 * 1)  // 10 秒

	//---------------------创建计时器，夜里12点触发---------------------
	//ztimer.TimerClock0(func() {
	//	ztimer.CheckRunTimeCost(func() {
	//		GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogic12clock") //公共逻辑处理循环
	//	}, "GoCallLuaCommonLogic12clock")
	//})


	//---------------------创建计时器，定期执行所有的接受消息队列---------------------
	ztimer.TimerMillisecondCheckUpdate(func() {
		ztimer.CheckRunTimeCost(func() {
			Lua.QueueRun()
		}, "Lua.QueueRun")
	},  20)  //

}



// 这是用来统计所有连接数量，及连接包不全的缓存大小
func GetAllConnectMsg() (string,int,int,int)  {
	connNum := 0		//所有包不全缓存大小
	successSendClients := 0
	successRecClients := 0
	successSendMsg := 0
	successRecMsg := 0
	WriteChan := 0
	AllConnect :=0

	//GlobalVar.RWMutex.RLock()
	Lua.ConnectMyTcpServer.Range(func(key, value interface{}) bool {
		v:= value.(*Lua.MyTcpServer)
		if v!=nil {
			AllConnect ++
			connNum += len(v.ReceiveBuf)
			if v.SendMsgNum>0{
				successSendClients++
				successSendMsg += v.SendMsgNum
				v.SendMsgNum = 0
			}
			if v.ReceiveMsgNum>0{
				successRecClients ++
				successRecMsg += v.ReceiveMsgNum
				v.ReceiveMsgNum = 0
			}
			WriteChan += v.Conn.GetWriteChanCap()
		}
		return true
	})
	//if AllConnect>0{
	//	WriteChan = WriteChan/AllConnect
	//}
	//GlobalVar.RWMutex.RUnlock()
	GameManagerLua.GoCallLuaSetIntVar("GlobalVar","ServerSendWriteChannelNum", WriteChan)                       // 发送缓冲区大小
	GameManagerLua.GoCallLuaSetIntVar("GlobalVar","ServerDataHeadErrorNum", Lua.StaticDataPackageHeadFlagError) // 把数据头尾错误发送给lua
	str:=fmt.Sprintf(" 发送连接数量 %d  接收连接数量  %d 每秒发送 %d  每秒接收 %d    发送缓存WriteChan %d  消息队列长度 %d ",   successSendClients, successRecClients, successSendMsg , successRecMsg, WriteChan, Lua.QueueGetLen())
	return "所有连接数量："+ strconv.Itoa(AllConnect) + str , successSendMsg , successRecMsg, WriteChan
}
