package Lua

import (
	"GoLuaServerV2.1/Utils/mongoDB"
	"GoLuaServerV2.1/Utils/mySql"
	"GoLuaServerV2.1/Utils/redis"
	"GoLuaServerV2.1/Utils/sqlServer"
	"GoLuaServerV2.1/Utils/zBit32"
	"GoLuaServerV2.1/Utils/zCrypto"
	"GoLuaServerV2.1/Utils/zJson"
	"GoLuaServerV2.1/Utils/zLog"
	"GoLuaServerV2.1/Utils/zProtocol"
	"GoLuaServerV2.1/Utils/zStrings"
	"GoLuaServerV2.1/Utils/zip"
	"GoLuaServerV2.1/Utils/ztimer"
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"github.com/yuin/gopher-lua"
	"time"
	//"github.com/cjoudrey/gluahttp"
	//mySql "github.com/tengattack/gluasql/mySql"
	"strconv"
	"strings"
)

//--------------------------------------------------------------------------------
// Lua调用的go函数
// 需要像下面一样，start的时候先注册进去，才可以正常调用
// L.SetGlobal("double", L.NewFunction(Double))
//--------------------------------------------------------------------------------

// 统一的go给lua调用的函数注册点
func (m *MyLua) InitResister() {
	// Lua调用go函数声明
	m.L.SetGlobal("luaCallGoNetWorkSendUdp", m.L.NewFunction(luaCallGoNetWorkSendUdp))                             //注册到lua 网络发送函数
	m.L.SetGlobal("luaCallGoNetWorkSend", m.L.NewFunction(luaCallGoNetWorkSend))                             //注册到lua 网络发送函数
	m.L.SetGlobal("luaCallGoNetWorkConnectOtherServer", m.L.NewFunction(luaCallGoNetWorkConnectOtherServer)) //注册到lua 网络 申请连接其他服务器
	m.L.SetGlobal("luaCallGoNetWorkClose", m.L.NewFunction(luaCallGoNetWorkClose))                           //注册到lua 网络关闭

	m.L.SetGlobal("luaCallGoPrintLogger", m.L.NewFunction(luaCallGoPrintLogger))                   //注册到lua 日志打印
	m.L.SetGlobal("luaCallGoGetOsTimeMillisecond", m.L.NewFunction(luaCallGoGetOsTimeMillisecond)) //注册到lua 获取毫秒时间
	m.L.SetGlobal("luaCallGoCreateNewTimer", m.L.NewFunction(luaCallGoCreateNewTimer))             //注册到lua 设置定时器
	m.L.SetGlobal("luaCallGoCreateNewClockTimer", m.L.NewFunction(luaCallGoCreateNewClockTimer))   //注册到lua 设置定时器，固定时间
	m.L.SetGlobal("luaCallGoResisterUID", m.L.NewFunction(luaCallGoResisterUID))                   //注册到lua 将uid注册到列表中


	//m.L.SetGlobal("luaCallGoCreateGoroutine", m.L.NewFunction(luaCallGoCreateGoroutine))		//注册到lua 创建go协程
	m.L.SetGlobal("luaCallGoGetPWD", m.L.NewFunction(luaCallGoGetPWD)) //注册到lua 生成用户密码


	zProtocol.LuaProtocolLoad(m.L) //加载protobuf的lua调用
	zBit32.LuaBit32Load(m.L)       // 加载bit32
	zJson.LuaJsonLoad(m.L)         // 加载json
	zCrypto.LuaCryptoLoad(m.L)
	zip.LuaZipLoad(m.L)
	zStrings.LuaStringsLoad(m.L)

	m.L.PreloadModule("mySql", mySql.Loader)         //加载mysql的lua调用 ，性能一般，写起来方便
	m.L.PreloadModule("sqlServer", sqlServer.Loader) //加载sql server 的lua调用
	m.L.PreloadModule("mongodb", mongoDB.Loader)     //加载mongodb 的lua调用
	m.L.PreloadModule("redis", redis.Loader)         //加载redis 的lua调用

	//m.L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)		//访问其他http地址
}

//------------------------------------------------------------------------------------------------------------------------
// 下面是lua 和 go 的交互函数
//------------------------------------------------------------------------------------------------------------------------

// lua发送网络数据
func luaCallGoNetWorkSend(L *lua.LState) int {
	userId := L.ToInt(1)
	serverId := L.ToInt(2)
	mainCmd := L.ToInt(3)
	subCmd := L.ToInt(4)
	data := L.ToString(5)
	msg := L.ToString(6)
	//token := L.ToInt(7)

	// lua传递过来之后， 立即开启一个新的协程去专门做发送工作
	//go func() {
	//bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0) // token始终是0，服务器不用发token
	//_, err := Conn.Write(bufferEnd)
	//zLog.CheckError(err)

	var result bool
	// 发送出去
	if userId == 0 {
		// 给玩家自己回复消息
		result = GetMyServerByServerId(serverId).SendMsg(data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
		//result = GetMyServerByServerId(serverId).WriteMsg(bufferEnd)
	} else {
		// 给其他玩家发送消息
		result = GetMyServerByUID(userId).SendMsg(data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
		//result = GetMyServerByUID(userId).WriteMsg(bufferEnd)
	}
	//}()

	L.Push(lua.LBool(result)) /* push result */
	//fmt.Println("lua send :" + str)
	return 1 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}

// user id 要注册，方便以后查询
func luaCallGoResisterUID(L *lua.LState) int {
	uid := L.ToNumber(1)                           // 玩家uid
	serverId := L.ToNumber(2)                      //
	server := GetMyServerByServerId(int(serverId)) // my server

	//GlobalVar.RWMutex.Lock()
	//ConnectMyTcpServerByUid[int(uid)] = server // 进行关联 ,  因为lua是单线程跑， 所以不存在线程安全问题， 如果是go，需要加锁
	//GlobalVar.RWMutex.Unlock()
	ConnectMyTcpServerByUid.Store(int(uid),server)

	server.UserId = int(uid) // 保存uid
	return 0
}

//-------------------------------------建立其他服务器的连接----------------------------------------------------------
// lua申请连接另外的服务器地址
func luaCallGoNetWorkConnectOtherServer(L *lua.LState) int {
	//serverId := L.ToInt(1)
	serverAddressAndPort := L.ToString(1)
	serverId := ConnectOtherServer(serverAddressAndPort)

	//if userId == 0 {
	//	result = GetMyServerByServerId(serverId).SendMsg(data, msg, mainCmd, subCmd)
	//} else {
	//	result = GetMyServerByUID(userId).SendMsg(data, msg, mainCmd, subCmd)
	//}
	//
	L.Push(lua.LNumber(serverId)) /* push result */
	//fmt.Println("lua send :" , serverId)
	return 1 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}

// lua 请求关闭网络连接
func luaCallGoNetWorkClose(L *lua.LState) int {
	userId := L.ToInt(1)
	serverId := L.ToInt(2)
	if userId > 0 {
		if GetMyServerByUID(userId) != nil {
			GetMyServerByUID(userId).LuaCallClose = true
		} else {
			zLog.PrintfLogger("玩家 %d ,连接并不存在", userId)
		}
	} else {
		GetMyServerByServerId(serverId).LuaCallClose = true
	}
	return 0 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}

//--------------------------------Utils-------------------------------------
// lua的日志处理
func luaCallGoPrintLogger(L *lua.LState) int {
	str := L.ToString(1)
	zLog.PrintLogger(str)
	return 0
}

// lua 创建一个计时器
func luaCallGoCreateNewTimer(L *lua.LState) int {
	module := L.ToString(1) //
	funcName := L.ToString(2) // 定期调用函数名字
	time1 := L.ToInt(3)       // 时间，秒

	ztimer.TimerMillisecondCheckUpdate(func() {
		GameManagerLuaHandle.GoCallLuaLogic(module,funcName) //定时调用函数
	}, time.Duration(time1))

	return 0
}

// lua 创建一个到固定时间触发器
func luaCallGoCreateNewClockTimer(L *lua.LState) int {
	module := L.ToString(1) //
	funcName := L.ToString(2) // 定期调用函数名字
	clock := L.ToInt(3)       // 时间，几点，24小时制

	ztimer.TimerClock(func() {
		GameManagerLuaHandle.GoCallLuaLogic(module,funcName) //定时调用函数
	}, clock)

	return 0
}

// 获取毫秒级系统时间
func luaCallGoGetOsTimeMillisecond(L *lua.LState) int {
	L.Push(lua.LNumber(ztimer.GetOsTimeMillisecond()))
	return 1
}




//-----------------------------游客密码的生成规则---------------------------------------------
func luaCallGoGetPWD(L *lua.LState) int {
	user_id := L.ToInt(1)
	mac := L.ToString(2)

	pwd := mac + "               " + strconv.Itoa(user_id)

	buffertt := new(bytes.Buffer)
	for _, v := range pwd {
		binary.Write(buffertt, binary.LittleEndian, uint16(v))
	}

	tokenBuf := buffertt.Bytes()
	//fmt.Println("", tokenBuf)

	h := md5.New()
	h.Write(tokenBuf)
	cips := h.Sum(nil) // h.Sum(nil) 将h的hash转成[]byte格式
	pwdmd5 := hex.EncodeToString(cips)
	pwdmd5 = strings.ToUpper(pwdmd5)
	//fmt.Println("pwdmd5: ", pwdmd5)
	L.Push(lua.LString(pwdmd5))
	return 1
}

//--------------------------------udp-----------------------------------------

// lua发送网络数据udp
func luaCallGoNetWorkSendUdp(L *lua.LState) int {
	//userId := L.ToInt(1)
	serverAddr := L.ToString(2)		// udp address
	mainCmd := L.ToInt(3)
	subCmd := L.ToInt(4)
	data := L.ToString(5)
	msg := L.ToString(6)


	//GetMyUdpServerByLSate(serverAddr).SendMsg(data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
	UdpSendMsg(serverAddr, data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回

	return 0 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}