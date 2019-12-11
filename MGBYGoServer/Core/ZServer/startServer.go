package ZServer
//-----------------------------------------------------------------------------------------------------------------
//  启动连接服务
//-----------------------------------------------------------------------------------------------------------------
import (
	"time"
	"math"
	"fmt"
	"../NetWork"
	"strconv"
)

var wsServer *NetWork.WSServer
var server *NetWork.TCPServer

var WebSocketServer = true	// websocket 开启
var SocketServer = true		// socket 开启

//-----------------------------------建立服务器的网络功能---------------------------------------------------------------
func NetWorkServerStart(ServerAddress string, SocketPort int, WebSocketPort int)  {

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
			ServerId := GetServerUid()
			a := NewMyServer(conn,ServerId) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}

		wsServer.Start()
	}
	if SocketServer{
		// socket 服务器开启----------------------------------
		server = new(NetWork.TCPServer)
		server.Addr = ServerAddress +":"+strconv.Itoa(SocketPort)
		fmt.Println("socket 绑定："+ server.Addr)
		server.MaxConnNum = int(math.MaxInt32)
		server.PendingWriteNum = 1000		// 发送区缓存
		server.LenMsgLen = 4
		server.MaxMsgLen = math.MaxUint32
		server.NewAgent = func(conn *NetWork.TCPConn) NetWork.Agent {
			ServerId := GetServerUid()
			a := NewMyServer(conn,ServerId) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}
		server.Start()
	}
}


// ----------------------------要跟其他服务器做一个 tcp 连接-------------------------------------------
func ConnectOtherServer(ServerAddressAndPort string) int{
	client := new(NetWork.TCPClient)
	client.Addr = ServerAddressAndPort
	client.ConnNum = 1  //废了
	client.ConnectInterval = 3 * time.Second	// 客户端自动重连
	client.PendingWriteNum = 1000	// 发送缓冲区
	client.LenMsgLen = 4
	client.MaxMsgLen = math.MaxUint32
	client.AutoReconnect = true		// 支持断线重联

	fmt.Println("0")
	//serverId :=  MyServerUUID 		// serverId
	serverId :=  GetServerUid() 		// serverId

	client.NewAgent = func(conn *NetWork.TCPConn,index int) NetWork.Agent {
		b := NewMyServer(conn,serverId)				// 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
		return b
	}

	fmt.Println("开始连接 -- 服务器 -- ", client.Addr, serverId)
	client.Start(serverId, serverId)
	return serverId
}