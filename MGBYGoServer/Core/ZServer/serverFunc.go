package ZServer

import (
	"../Utils/log"
	"../Utils/ztimer"
	"time"
	"strconv"
	"bytes"
	"encoding/binary"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 发消息给客户端， 通过 serverId
func NetWorkSendByServerId(serverId int, data string, msg string, mainCmd int, subCmd int)  bool{
	var result bool
	// 发送出去
	result = GetMyServerByServerId(serverId).SendMsg(data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
	return result
}

// 发消息给客户端，通过 uid
func NetWorkSendByUid(uid int, data string, msg string, mainCmd int, subCmd int)  bool{
	var result bool
	// 发送出去
	result = GetMyServerByUID(uid).SendMsg(data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
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
			log.PrintfLogger("玩家 %d ,连接并不存在" ,userId)
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


//-----------------------------游客密码的生成规则---------------------------------------------
func luaCallGoGetPWD(user_id int, mac string) string {
	pwd := mac + "               " + strconv.Itoa(user_id)

	buffertt := new(bytes.Buffer)
	for _,v := range pwd{
		binary.Write(buffertt, binary.LittleEndian, uint16(v))
	}

	tokenBuf := buffertt.Bytes()
	//fmt.Println("", tokenBuf)

	h:= md5.New()
	h.Write(tokenBuf)
	cips := h.Sum(nil)			// h.Sum(nil) 将h的hash转成[]byte格式
	pwdmd5 := hex.EncodeToString(cips)
	pwdmd5 = strings.ToUpper(pwdmd5)
	//fmt.Println("pwdmd5: ", pwdmd5)
	//L.Push(lua.LString(pwdmd5))
	return pwdmd5
}