package main

import (
	"./netWork"
	"reflect"
	"net"
	"fmt"
	"math"
	"time"
	"strconv"
)

var wsServer *NetWork.WSServer
var server *NetWork.TCPServer
var IsWebSocket bool

func main() {

	IsWebSocket = true // webscoket


	if IsWebSocket {
		// websocket 服务器开启---------------------------------
		wsServer = new(NetWork.WSServer)
		wsServer.Addr = "localhost:8089"
		wsServer.MaxConnNum = 2
		wsServer.PendingWriteNum = 100
		wsServer.MaxMsgLen = 4096
		wsServer.HTTPTimeout = 10 * time.Second
		wsServer.CertFile = ""
		wsServer.KeyFile = ""
		wsServer.NewAgent = func(conn *NetWork.WSConn) NetWork.Agent {
			a := &agentServer{conn: conn}
			return a
		}
		wsServer.Start()
	}
	if IsWebSocket{

		// socket 服务器开启----------------------------------
		server = new(NetWork.TCPServer)
		server.Addr = "172.16.140.131:8123"
		server.MaxConnNum = int(math.MaxInt32)
		server.PendingWriteNum = 100
		server.LenMsgLen = 4
		server.MaxMsgLen = math.MaxUint32
		server.NewAgent = func(conn *NetWork.TCPConn) NetWork.Agent {
			a := &agentServer{conn: conn}
			return a
		}
		server.Start()

	}

	for {
		select {}
	}
}

//type Agent interface {
//	WriteMsg(msg ... []byte)
//	LocalAddr() net.Addr
//	RemoteAddr() net.Addr
//	Close()
//	Destroy()
//	UserData() interface{}
//	SetUserData(data interface{})
//}

// wsServer.NewAgent 服务器连接的代理
type agentServer struct {
	conn NetWork.Conn
	//gate     *Gate
	userData interface{}
}

func (a *agentServer) Run() {
	//fmt.Println("run")
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("跟对方的连接中断了")
			break
		}
		fmt.Println("收到消息------------")
		fmt.Println("消息:  ", string(data))

		for i:=0;i<100 ;i++  {
			a.WriteMsg([]byte("服务器收到你的消息-------" + strconv.Itoa(i)+ string(data) ))
			fmt.Println("",i)
			time.Sleep(time.Microsecond * 1000)

		}
		a.WriteMsg([]byte("服务器收到你的消息-------" + string(data)))

		//if a.gate.Processor != nil {
		//	msg, err := a.gate.Processor.Unmarshal(data)
		//	if err != nil {
		//		fmt.Printf("unmarshal message error: %v", err)
		//		break
		//	}
		//	err = a.gate.Processor.Route(msg, a)
		//	if err != nil {
		//		fmt.Printf("route message error: %v", err)
		//		break
		//	}
		//}
	}
}

func (a *agentServer) OnClose() {
	//if a.gate.AgentChanRPC != nil {
	//	err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
	//	if err != nil {
	//		fmt.Printf("chanrpc error: %v", err)
	//	}
	//}
}

func (a *agentServer) WriteMsg(msg ... []byte) {
	//if a.gate.Processor != nil {
	//	data, err := a.gate.Processor.Marshal(msg)
	//	if err != nil {
	//		fmt.Printf("marshal message %v error: %v", reflect.TypeOf(msg), err)
	//		return
	//	}
	err := a.conn.WriteMsg(msg...)
	if err != nil {
		fmt.Printf("write message %v error: %v", reflect.TypeOf(msg), err)
	}
	//}
}

func (a *agentServer) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agentServer) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agentServer) Close() {
	a.conn.Close()
}

func (a *agentServer) Destroy() {
	a.conn.Destroy()
}

func (a *agentServer) UserData() interface{} {
	return a.userData
}

func (a *agentServer) SetUserData(data interface{}) {
	a.userData = data
}
