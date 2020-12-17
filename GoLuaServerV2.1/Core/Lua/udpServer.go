package Lua

import (
	//"GoLuaServerV2.1/GlobalVar"
	"GoLuaServerV2.1/Core/NetWork"
	"GoLuaServerV2.1/Core/Utils/zLog"
	"net"
)

var MyUdpListen *net.UDPConn

//var MyUdpServerUUID = 0                // 自定义玩家连接的临时编号，用来传给lua，这样lua就知道消息给谁返回
type MyUdpServer struct {
	Conn  *NetWork.UdpConn // 对应的每个玩家的连接
	myLua *MyLua           // 处理该玩家的lua脚本
	ServerId int           // 自己分配的连接编号
	UserId  int            // 玩家uid
}

// 通过lua堆栈找到对应的是哪个myServer
func GetMyUdpServerByLSate(clientAddr string) *MyUdpServer {
	re, _ := ConnectMyUdpServer.Load(clientAddr) // 这是全局变量，所以要加锁， 读写都要加
	return re.(*MyUdpServer)
}

func NewMyUdpServer(conn *NetWork.UdpConn) *MyUdpServer {
	return &MyUdpServer{Conn: conn, myLua: GameManagerLuaHandle}
}

func (a *MyUdpServer) init() {
	//if _,ok:=ConnectMyUdpServer.Load(a.Conn.UDPAddr.String());ok {
	//	zLog.PrintfLogger("ConnectMyTcpServer  已经有了, map重复了", a.ServerId)
	//}else{
	//	ConnectMyUdpServer.Store(a.Conn.UDPAddr.String(),a)
	//}
}

func (a *MyUdpServer) Run() {
	a.init()
	buf := a.Conn.Buffer.Bytes()
	bufPackageSize, msgId, subMsgId, finalBuffer := ReadDataPackage(buf,-1)
	if bufPackageSize >0 {
		QueueAddUdp(a.Conn.UDPAddr.String(), a.UserId, msgId, subMsgId, finalBuffer, 0) // 把收到的数据传递给队列， 后期进行lua进行处理
	}else {
		zLog.PrintLogger("udp 数据包格式不合法")
	}
}

func (a *MyUdpServer) OnClose() {
	// 清理掉一些调用关系

	//ConnectMyUdpServer.Delete(a.Conn.UDPAddr.String())
}

func UdpSendMsg(serverAddr string, data string, msg string, mainCmd int, subCmd int) bool {
	bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0)
	//println("send msg")

	UdpAddr, _ := net.ResolveUDPAddr("udp", serverAddr)
	if _, err := MyUdpListen.WriteToUDP(bufferEnd, UdpAddr); err != nil {
		zLog.PrintfLogger("udp 发送 error %s \n", err.Error())
	}

	return true
}
