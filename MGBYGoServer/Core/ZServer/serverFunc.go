package ZServer

//-----------------------------------------------------------------------------------------------------------------
//  给业务逻辑调用的接口
//-----------------------------------------------------------------------------------------------------------------


import (
	"../Utils/zLog"
	"../Utils/ztimer"
	"time"
	"github.com/golang/protobuf/proto"
)

// 发消息给客户端， 通过 serverId
func NetWorkSendByServerId(serverId int, data proto.Message, mainCmd int, subCmd int)  bool{
	var result bool
	// 发送出去
	result = GetMyServerByServerId(serverId).SendMsg(data, "", mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
	return result
}

// 发消息给客户端，通过 uid
func NetWorkSendByUid(uid int, data proto.Message,mainCmd int, subCmd int)  bool{
	var result bool
	// 发送出去
	result = GetMyServerByUID(uid).SendMsg(data, "", mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
	return result
}


// uid 注册 myServer
func ResisterUID(serverId int, uid int)  {
	server := GetMyServerByServerId(int(serverId)) // my server

	RWMutex.Lock()
	UidConnectMyServer[int(uid)] = server // 进行关联 ,  因为lua是单线程跑， 所以不存在线程安全问题， 如果是go，需要加锁
	RWMutex.Unlock()

	server.UserId = int(uid)                       // 保存uid
}


// 关闭连接
func NetWorkClose(serverId int ,userId int)  {
	if userId > 0 {
		if GetMyServerByUID(userId) != nil {
			GetMyServerByUID(userId).LuaCallClose = true
		}else {
			zLog.PrintfLogger("玩家 %d ,连接并不存在" ,userId)
		}
	}else{
		GetMyServerByServerId(serverId).LuaCallClose = true
	}
}

// 创建毫秒计时器
func CreateNewTimer(f func(), time1 time.Duration)  {
	ztimer.TimerMillisecondCheckUpdate(func() {
		f()
	},  time.Duration(time1) )
}

// 创建定时器
func CreateOclockTimer(f func(), clock int )  {
	ztimer.TimerClock(func() {
		f()
	},  clock )

}
