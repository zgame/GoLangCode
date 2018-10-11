package main

import (
	"fmt"
	"net"
	"./NetWork"
)

// ----------------------------服务器处理的统一接口----------------------------------


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
type myServer struct {
	conn NetWork.Conn
	//gate     *Gate
	userData interface{}
}

func (a *myServer) Run() {
	//fmt.Println("run")
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("跟对方的连接中断了")
			break
		}
		fmt.Println("收到消息------------", string(data))

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

func (a *myServer) OnClose() {

}

func (a *myServer) WriteMsg(msg ... []byte) {

	err := a.conn.WriteMsg(msg...)
	if err != nil {
		fmt.Printf("发送消息：%v， 出错了！", msg)
	}
	//}
}

func (a *myServer) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *myServer) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *myServer) Close() {
	a.conn.Close()
}

func (a *myServer) Destroy() {
	a.conn.Destroy()
}

func (a *myServer) UserData() interface{} {
	return a.userData
}

func (a *myServer) SetUserData(data interface{}) {
	a.userData = data
}

