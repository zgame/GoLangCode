package Lua

import (
	"github.com/oleiade/lane"
)

// 接受到的网络消息
type NetWorkMessage struct {
	//MyServerHandler *MyTcpServer
	ServerId int
	UserId int
	MsgId int
	SubMsgId int
	Buffer string
	Token int
	UDPAddr string
}

var queueTcp *lane.Queue // 网络消息队列
var queueUdp *lane.Queue // 网络消息队列

// 队列初始化
func QueueInit()  {
	queueTcp = lane.NewQueue()
	queueUdp = lane.NewQueue()
}

// 把网络消息保存到队列中
func QueueAddTcp( serverId int ,userId int, msgId int ,subMsgId int ,buffer string ,token int){
	message := &NetWorkMessage{ServerId:serverId,UserId:userId,MsgId:msgId, SubMsgId:subMsgId, Buffer:buffer, Token:token}
	queueTcp.Enqueue(message)
}
// 把网络消息保存到队列中
func QueueAddUdp( serverAddr string ,userId int, msgId int ,subMsgId int ,buffer string ,token int){
	message := &NetWorkMessage{UDPAddr:serverAddr,UserId:userId,MsgId:msgId, SubMsgId:subMsgId, Buffer:buffer, Token:token}
	queueUdp.Enqueue(message)
}

// 把队列中的网络消息依次传递给lua进行处理
func QueueRun() {
	//fmt.Println("处理消息队列",queueTcp.Size())
	for queueTcp.Head() != nil{
		item := queueTcp.Dequeue().(*NetWorkMessage)
		GameManagerLuaHandle.GoCallLuaNetWorkReceive( item.ServerId, item.UserId, item.MsgId, item.SubMsgId, item.Buffer, item.Token)
	}
	for queueUdp.Head() != nil{
		item := queueUdp.Dequeue().(*NetWorkMessage)
		GameManagerLuaHandle.GoCallLuaNetWorkReceiveUdp( item.UDPAddr, item.MsgId, item.SubMsgId, item.Buffer)
	}


}

// 获取网络消息队列的长度
func QueueGetLen() int {
	return queueTcp.Size()
}