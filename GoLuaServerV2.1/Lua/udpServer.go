package Lua

import (
	"GoLuaServerV2.1/GlobalVar"
	"GoLuaServerV2.1/NetWork"
	"GoLuaServerV2.1/Utils/zLog"
	"fmt"
	"time"
)



var MyUdpServerUUID = 0                // 自定义玩家连接的临时编号，用来传给lua，这样lua就知道消息给谁返回
type MyUdpServer struct {
	Conn  NetWork.Conn // 对应的每个玩家的连接
	myLua *MyLua       // 处理该玩家的lua脚本
	ServerId int
}

func NewMyUdpServer(conn NetWork.Conn, ServerId int) *MyUdpServer {
	return &MyUdpServer{Conn: conn,myLua:GameManagerLuaHandle,ServerId: ServerId}
}

func (a *MyUdpServer) init()  {
	GlobalVar.RWMutex.Lock()
	if ConnectMyUdpServer[a.ServerId] != nil {
		zLog.PrintfLogger("ConnectMyTcpServer  已经有了, map重复了", a.ServerId)
	}
	ConnectMyUdpServer[a.ServerId] = a
	GlobalVar.RWMutex.Unlock()
	a.myLua.GoCallLuaLogicInt("GoCallLuaStartGamesServers",a.ServerId)
}

func (a *MyUdpServer) Run() {
	a.init()

	for{
		msgData,Len, err := a.Conn.ReadMsg()
		if err!=nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("接收消息： %s \n",string(msgData[:Len]))
		time.Sleep(time.Microsecond * 20)
	}
}

func (a *MyUdpServer) OnClose() {
	// 清理掉一些调用关系
	GlobalVar.RWMutex.Lock()
	delete(ConnectMyUdpServer, a.ServerId)
	GlobalVar.RWMutex.Unlock()
}

func (a *MyUdpServer) SendMsg(data string, msg string, mainCmd int, subCmd int ) bool {
	bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0)
	//println("send msg")
	a.Conn.WriteMsg(bufferEnd)

	return true
}