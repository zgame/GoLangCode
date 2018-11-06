package Lua

import (
	"github.com/yuin/gopher-lua"
	"../Utils/log"
	"../Utils/ztimer"
	"../NetWork"
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
	m.L.SetGlobal("luaCallGoNetWorkSend", m.L.NewFunction(luaCallGoNetWorkSend))		//注册到lua 网络发送函数
	m.L.SetGlobal("luaCallGoPrintLogger", m.L.NewFunction(luaCallGoPrintLogger))		//注册到lua 日志打印
	m.L.SetGlobal("luaCallGoGetOsTimeMillisecond", m.L.NewFunction(luaCallGoGetOsTimeMillisecond))		//注册到lua 获取毫秒时间
	//m.L.SetGlobal("luaCallGoCreateGoroutine", m.L.NewFunction(luaCallGoCreateGoroutine))		//注册到lua 创建go协程

	//加载protobuf
	luaopen_pb(m.L)
}

// 通过lua堆栈找到对应的是哪个myServer
func GetMyServerByLSate(L *lua.LState) *MyServer {
	return LuaConnectMyServer[L]
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
	mainCmd := L.ToInt(1)
	subCmd := L.ToInt(2)
	data := L.ToString(3)
	msg := L.ToString(4)

	// 组合一下str
	protoData := []byte(data)
	protoDataSize := len(protoData)

	// 增加服务器的错误提示msg
	msgData := []byte(msg)
	msgSize :=  len(msgData)

	// 生成数据包头部信息
	bufferHead := NetWork.GetSendTcpHeaderData(uint16(mainCmd),uint16( subCmd), uint16(protoDataSize), uint8(msgSize))
	headSize :=  len(bufferHead)


	// 发送最后数据包
	bufferEnd := make([]byte, protoDataSize + headSize + msgSize)
	copy(bufferEnd, bufferHead)		// copy 数据包头部
	copy(bufferEnd[headSize:protoDataSize + headSize], protoData)	// copy protobuffer 数据
	copy(bufferEnd[protoDataSize + headSize:], msgData)		// copy msg 数据
	//_, err := Conn.Write(bufferEnd)
	//log.CheckError(err)

	// 发送出去
	GetMyServerByLSate(L).WriteMsg(bufferEnd)

	//fmt.Println("lua send :" + str)
	return 0			// 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}

// lua的日志处理
func luaCallGoPrintLogger(L * lua.LState) int  {
	str := L.ToString(1)
	log.PrintLogger(str)
	return 0
}
//
//// lua 创建一个go协程
//func luaCallGoCreateGoroutine(L * lua.LState) int  {
//	funcName := L.ToString(1)
//	go func() {
//		if err := L.CallByParam(lua.P{
//			Fn: L.GetGlobal(funcName),		// lua的函数名字
//			NRet: 0,
//			Protect: true,
//		}); err != nil {		// 参数
//			fmt.Println("luaCallGoCreateGoroutine error :",err.Error())
//		}
//	}()
//	return 0
//}

// 获取毫秒级系统时间
func luaCallGoGetOsTimeMillisecond(L *lua.LState) int {
	L.Push(lua.LNumber(ztimer.GetOsTimeMillisecond()))
	return 1
}