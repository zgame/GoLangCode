package main

import (
	"./netWork"
	"time"
	"math"
	"net"
	"fmt"
	"reflect"
)

//var origin = "http://localhost/"
//var url = "ws://localhost:8080/echo"

func main() {
	// socket client
	//var clients []*network.TCPClient
	//
	//client := new(network.TCPClient)
	//client.Addr = "127.0.0.1:8089"
	//client.ConnNum = 1
	//client.ConnectInterval = 3 * time.Second
	//client.PendingWriteNum = 100
	//client.LenMsgLen = 4
	//client.MaxMsgLen = math.MaxUint32
	//client.NewAgent = func(conn *network.TCPConn) network.Agent  {
	//	a := &agentClient{conn: conn}
	//	return a
	//}
	//
	//client.Start()
	//clients = append(clients, client)

	// websocket client
	var wsclients []*network.WSClient

	wsclient := new(network.WSClient)
	wsclient.Addr = "ws://localhost:8089/"
	wsclient.ConnNum = 1
	wsclient.ConnectInterval = 3 * time.Second
	wsclient.PendingWriteNum = 100
	wsclient.HandshakeTimeout = 10 * time.Second
	wsclient.MaxMsgLen = math.MaxUint32
	wsclient.NewAgent = func(conn *network.WSConn) network.Agent {
		a := &agentClient{conn: conn}
		return a
	}

	wsclient.Start()
	wsclients = append(wsclients, wsclient)

	wsclient.NewAgent.WriteMsg([]byte("服务器收到你的消息-------"+ string(data)))
wsclient.NewAgent.WriteMsg()

}



// wsServer.NewAgent 服务器连接的代理
type agentClient struct {
	conn     network.Conn
	//gate     *Gate
	userData interface{}
}

func (a *agentClient) Run() {
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
