package main

import (
	"./NetWork"
	"net"
	"fmt"
	"reflect"
	"time"
	"math"
)

//var origin = "http://localhost/"
//var url = "ws://localhost:8080/echo"


func main() {
	IsWebSocket := true

	if false {
		// socket client----------------------------------------------------------
		var clients []*NetWork.TCPClient

		client := new(NetWork.TCPClient)
		client.Addr = "127.0.0.1:8088"
		client.ConnNum = 1
		client.ConnectInterval = 3 * time.Second
		client.PendingWriteNum = 100
		client.LenMsgLen = 4
		client.MaxMsgLen = math.MaxUint32
		client.NewAgent = func(conn *NetWork.TCPConn) NetWork.Agent {
			a := &agentClient{conn: conn}
			return a
		}

		client.Start()
		clients = append(clients, client)
	}
	if IsWebSocket{
		// websocket client------------------------------------------------------------------
		var wsclients []*NetWork.WSClient

		wsclient := new(NetWork.WSClient)
		wsclient.Addr = "ws://localhost:8089/"
		wsclient.ConnNum = 1
		wsclient.ConnectInterval = 3 * time.Second
		wsclient.PendingWriteNum = 100
		wsclient.HandshakeTimeout = 10 * time.Second
		wsclient.MaxMsgLen = math.MaxUint32
		wsclient.NewAgent = func(conn *NetWork.WSConn) NetWork.Agent {
			a := &agentClient{conn: conn}
			return a
		}

		fmt.Println("开始客户端")
		wsclient.Start()
		wsclients = append(wsclients, wsclient)
	}

	for {
		select {

		}
	}

}



// wsServer.NewAgent 服务器连接的代理
type agentClient struct {
	conn NetWork.Conn
	//gate     *Gate
	userData interface{}
}

func (a *agentClient) Run() {
	a.WriteMsg([]byte("我是客户端哟"))

	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("跟对方的连接中断了")
			break
		}
		fmt.Println("收到消息------------")
		fmt.Println("消息:  ",string(data))


		//a.WriteMsg([]byte("服务器收到你的消息-------"+ string(data)))

	}
}

func (a *agentClient) OnClose() {

}

func (a *agentClient) WriteMsg(msg... []byte) {
	err := a.conn.WriteMsg(msg...)
	if err != nil {
		fmt.Printf("write message %v error: %v", reflect.TypeOf(msg), err)
	}
	//}
}

func (a *agentClient) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agentClient) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agentClient) Close() {
	a.conn.Close()
}

func (a *agentClient) Destroy() {
	a.conn.Destroy()
}

func (a *agentClient) UserData() interface{} {
	return a.userData
}

func (a *agentClient) SetUserData(data interface{}) {
	a.userData = data
}
