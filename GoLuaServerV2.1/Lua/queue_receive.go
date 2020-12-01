package Lua

import (
	"github.com/oleiade/lane"
)

// 接受到的网络消息
type NetWorkMessage struct {
	MyServerHandler *MyServer
	ServerId int
	UserId int
	MsgId int
	SubMsgId int
	Buffer string
	Token int
}

var queue *lane.Queue	// 网络消息队列

// 队列初始化
func QueueInit()  {
	queue = lane.NewQueue()
}

// 把网络消息保存到队列中
func QueueAdd(myServer *MyServer, serverId int ,userId int, msgId int ,subMsgId int ,buffer string ,token int)  {
	message := &NetWorkMessage{MyServerHandler:myServer,ServerId:serverId,UserId:userId,MsgId:msgId, SubMsgId:subMsgId, Buffer:buffer, Token:token}
	queue.Enqueue(message)
}

// 把队列中的网络消息依次传递给lua进行处理
func QueueRun()  {
	//fmt.Println("处理消息队列",queue.Size())
	for queue.Head() != nil{
		item := queue.Dequeue().(*NetWorkMessage)
		item.MyServerHandler.myLua.GoCallLuaNetWorkReceive( item.ServerId, item.UserId, item.MsgId, item.SubMsgId, item.Buffer, item.Token)
	}


}

// 获取网络消息队列的长度
func QueueGetLen() int {
	return queue.Size()
}