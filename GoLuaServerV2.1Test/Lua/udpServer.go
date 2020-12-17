package Lua

import (
	"GoLuaServerV2.1Test/NetWork"
	"GoLuaServerV2.1Test/Utils/zLog"
	"fmt"
	"math"
	"time"
)

var MyUdpServerUUID = 0                // 自定义玩家连接的临时编号，用来传给lua，这样lua就知道消息给谁返回
type MyUdpServer struct {
	Conn  NetWork.Conn // 对应的每个玩家的连接
	myLua *MyLua       // 处理该玩家的lua脚本
	ServerId int
}

func NewMyUdpServer(conn NetWork.Conn, GameManagerLua *MyLua) *MyUdpServer {
	//myLua := NewMyLua()
	myLua:= GameManagerLua		// 改为统一一个LState
	GlobalMutex.Lock()

	if MyUdpServerUUID == 0 {
		MyUdpServerUUID = ClientStart
	}

	ServerId := MyUdpServerUUID
	MyUdpServerUUID++
	if MyUdpServerUUID > int(math.MaxInt32) {
		MyUdpServerUUID = 0
	}
	GlobalMutex.Unlock()
	return &MyUdpServer{Conn: conn,myLua:myLua,ServerId: ServerId}
}

func (a *MyUdpServer) init()  {
	RWMutex.Lock()
	if ConnectMyUdpServer[a.ServerId] != nil {
		zLog.PrintfLogger("ConnectMyTcpServer  已经有了, map重复了", a.ServerId)
	}
	ConnectMyUdpServer[a.ServerId] = a
	RWMutex.Unlock()
	a.myLua.GoCallLuaLogicInt("GoCallLuaStartGamesServers",a.ServerId)
}

func (a *MyUdpServer) Run() {
	a.init()

	for{
		msgData,_, err := a.Conn.ReadMsg()
		if err!=nil {
			fmt.Println(err.Error())
			continue
		}
		//fmt.Printf("接收消息： %s \n",string(msgData[:Len]))
		bufHeadTemp,msgId,subMsgId,finalBuffer := HandlerRead(msgData,-1) //处理结束之后返回，接下来要开始的范围
		if bufHeadTemp>0 {
			GameManagerLuaHandle.GoCallLuaNetWorkReceiveUdp( "",  msgId,subMsgId,finalBuffer)		// 把收到的数据传递给lua进行处理
		}else {
			zLog.PrintLogger("数据包格式不合法")
		}

		time.Sleep(time.Microsecond * 20)
	}
}

func (a *MyUdpServer) OnClose() {
	// 清理掉一些调用关系
	RWMutex.Lock()
	delete(ConnectMyUdpServer, a.ServerId)
	RWMutex.Unlock()
}

func (a *MyUdpServer) SendMsg(data string, msg string, mainCmd int, subCmd int ) bool {
	bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0)
	//println("send msg")
	a.Conn.WriteMsg(bufferEnd)

	return true
}