package Lua

import (
	"GoLuaServerV2.1Test/Core/Utils/zLog"
	"GoLuaServerV2.1Test/Core/Utils/zPbc"
	"GoLuaServerV2.1Test/Core/Utils/ztimer"
	"fmt"
	"github.com/yuin/gopher-lua"
	"time"
	//mysql "github.com/tengattack/gluasql/mysql"
)

//--------------------------------------------------------------------------------
// Lua调用的go函数
// 需要像下面一样，start的时候先注册进去，才可以正常调用
// L.SetGlobal("double", L.NewFunction(Double))
//--------------------------------------------------------------------------------

// 统一的go给lua调用的函数注册点
func (m *MyLua)InitResister() {
	// Lua调用go函数声明
	//m.L.SetGlobal("double", m.L.NewFunction(Double))
	m.L.SetGlobal("luaCallGoNetWorkSendUdp", m.L.NewFunction(luaCallGoNetWorkSendUdp))             //注册到lua 网络发送函数
	m.L.SetGlobal("luaCallGoNetWorkSend", m.L.NewFunction(luaCallGoNetWorkSend))                   //注册到lua 网络发送函数
	m.L.SetGlobal("luaCallGoPrintLogger", m.L.NewFunction(luaCallGoPrintLogger))                   //注册到lua 日志打印
	m.L.SetGlobal("luaCallGoGetOsTimeMillisecond", m.L.NewFunction(luaCallGoGetOsTimeMillisecond)) //注册到lua 获取毫秒时间
	m.L.SetGlobal("luaCallGoCreateNewTimer", m.L.NewFunction(luaCallGoCreateNewTimer))             //注册到lua 设置定时器
	m.L.SetGlobal("luaCallGoCreateNewClockTimer", m.L.NewFunction(luaCallGoCreateNewClockTimer))   //注册到lua 设置定时器，固定时间
	m.L.SetGlobal("luaCallGoResisterUID", m.L.NewFunction(luaCallGoResisterUID))                   //注册到lua 将uid注册到列表中
	//m.L.SetGlobal("luaCallGoRedisSaveString", m.L.NewFunction(luaCallGoRedisSaveString))		//注册到lua redis save
	//m.L.SetGlobal("luaCallGoRedisGetString", m.L.NewFunction(luaCallGoRedisGetString))		//注册到lua redis load
	//m.L.SetGlobal("luaCallGoRedisDelKey", m.L.NewFunction(luaCallGoRedisDelKey))		//注册到lua redis del key
	//m.L.SetGlobal("luaCallGoAddNumberToRedis", m.L.NewFunction(luaCallGoAddNumberToRedis))		//注册到lua redis add number

	//m.L.SetGlobal("luaCallGoCreateGoroutine", m.L.NewFunction(luaCallGoCreateGoroutine))		//注册到lua 创建go协程

	//加载protobuf
	zPbc.LuaOpenPb(m.L)
	//m.L.PreloadModule("mysql", mysql.Loader)
}



//------------------------------------------------------------------------------------------------------------------------
// 下面是lua 和 go 的交互函数
//------------------------------------------------------------------------------------------------------------------------

//// test
//func Double(L *lua.LState) int {
//	lv := L.ToInt(1)             //第一个参数
//	lv2 :=  L.ToInt(2)			 //第二个参数
//	str := L.ToString(3)
//
//	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */
//
//	return 1                     /* number of results */
//}

// lua发送网络数据
func luaCallGoNetWorkSend(L *lua.LState) int {
	userId := L.ToInt(1)
	serverId := L.ToInt(2)
	mainCmd := L.ToInt(3)
	subCmd := L.ToInt(4)
	data := L.ToString(5)
	msg := L.ToString(6)

	// lua传递过来之后， 立即开启一个新的协程去专门做发送工作
	//go func() {
	//bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0) // token始终是0，服务器不用发token
	//_, err := Conn.Write(bufferEnd)
	//zLog.CheckError(err)

	fmt.Println("tcp")

	var result bool
	// 发送出去
	if userId == 0 {
		// 给玩家自己回复消息
		result = GetMyTcpServerByLSate(serverId).SendMsg(data, msg, mainCmd, subCmd)
		//result = GetMyTcpServerByLSate(serverId).WriteMsg(bufferEnd)
	} else {
		// 给其他玩家发送消息
		result = GetMyTcpServerByUID(userId).SendMsg(data, msg, mainCmd, subCmd)
		//result = GetMyTcpServerByUID(userId).WriteMsg(bufferEnd)
	}
	//}()

	L.Push(lua.LBool(result))		 /* push result */
	//fmt.Println("lua send :" + str)
	return 1 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}


// user id 要注册，方便以后查询
func luaCallGoResisterUID(L * lua.LState) int  {
	uid := L.ToNumber(1)                           // 玩家uid
	serverId := L.ToNumber(2)                      //
	server := GetMyTcpServerByLSate(int(serverId)) // my server
	ConnectMyTcpServerByUID[int(uid)] = server     // 进行关联 ,  因为lua是单线程跑， 所以不存在线程安全问题， 如果是go，需要加锁
	server.UserId = int(uid)                       // 保存uid
	return 0
}




//--------------------------------Utils-------------------------------------
// lua的日志处理
func luaCallGoPrintLogger(L * lua.LState) int  {
	str := L.ToString(1)
	zLog.PrintLogger(str)
	return 0
}

// lua 创建一个计时器
func luaCallGoCreateNewTimer(L * lua.LState) int  {
	module := L.ToString(1) //
	funcName := L.ToString(2) // 定期调用函数名字
	time1 := L.ToInt(3)       // 时间，秒

	ztimer.TimerMillisecondCheckUpdate(func() {
		GameManagerLuaHandle.GoCallLuaLogic(module,funcName) //定时调用函数
	}, time.Duration(time1))

	return 0
}

// lua 创建一个到固定时间触发器
func luaCallGoCreateNewClockTimer(L * lua.LState) int  {
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



//--------------------------------udp-----------------------------------------

// lua发送网络数据udp
func luaCallGoNetWorkSendUdp(L *lua.LState) int {
	//userId := L.ToInt(1)
	serverId := L.ToInt(2)		// udp address
	mainCmd := L.ToInt(3)
	subCmd := L.ToInt(4)
	data := L.ToString(5)
	msg := L.ToString(6)

println("udp send")
	GetMyUdpServerByLSate(serverId).SendMsg(data, msg, mainCmd, subCmd) // 把客户端发来的token返回给客户端，标记出这是哪个消息的返回

	return 0 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}