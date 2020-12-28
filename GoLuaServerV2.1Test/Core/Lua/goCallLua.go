package Lua

import (
	"GoLuaServerV2.1Test/Core/Utils/zLog"
	"fmt"
	"github.com/yuin/gopher-lua"
)

//--------------------------------------------------------------------------------
// go 调用 lua的函数
//--------------------------------------------------------------------------------

// ----------------测试用例--------------------------------
func GoCallLuaTest(L *lua.LState, num int)  {
	// 这里是go调用lua的函数
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("Zsw2"),		// lua的函数名字
		NRet: 2,
		Protect: true,
	}, lua.LNumber(num),lua.LNumber(num)); err != nil {		// 参数
		fmt.Println("---------------")
		fmt.Println("",err.Error())
		fmt.Println("----------------")
	}

	ret := L.Get(1) // returned value
	fmt.Println("lua return: ",ret)
	ret = L.Get(2) // returned value
	fmt.Println("lua return: ",ret)
	L.Pop(1)  // remove received value
	L.Pop(1)  // remove received value
}

// -------------------go触发lua函数，不带参数和返回值-------------------
func (m *MyLua) GoCallLuaLogic(module string,funcName string,param int) {
	GlobalMutex.Lock()
	table:= m.L.GetGlobal(module)
	value := m.L.GetField(table,funcName)
	if err := m.L.CallByParam(lua.P{
		Fn: value,		// lua的函数名字
		NRet: 0,
		Protect: true,
	},lua.LNumber(param)); err != nil {		// 参数
		zLog.PrintLogger("GoCallLuaLogic error :"+funcName+"      "+err.Error())
	}
	GlobalMutex.Unlock()
}

// -------------------go传递接收到的网络数据包给lua-------------------
func (m *MyLua)GoCallLuaNetWorkReceive(serverId int,userId int,msgId int , subMsgId int ,buf string) {
	table:= m.L.GetGlobal("Network")
	value := m.L.GetField(table,"Receive")
	GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: value,
		NRet: 0,
		Protect: true,
	}, lua.LNumber(serverId),lua.LNumber(userId),lua.LNumber(msgId), lua.LNumber(subMsgId), lua.LString(buf)); err != nil {		// 参数
		zLog.PrintfLogger("GoCallLuaNetWorkReceive  error :  msgId:%d  subMsgId %d  buf:%s   error:%s",msgId , subMsgId, buf, err.Error())
	}
	GlobalMutex.Unlock()
}

// -------------------go传递接收到的网络数据包给lua-------------------
func (m *MyLua)GoCallLuaNetWorkReceiveUdp(serverAddr string,msgId int , subMsgId int ,buf string) {
	table:= m.L.GetGlobal("Network")
	value := m.L.GetField(table,"UdpReceive")
	GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: value,
		NRet: 0,
		Protect: true,
	}, lua.LString(serverAddr),lua.LNumber(msgId), lua.LNumber(subMsgId), lua.LString(buf)); err != nil {		// 参数
		zLog.PrintfLogger("GoCallLuaNetWorkUdpReceive  error :  msgId:%d  subMsgId %d  buf:%s   error:%s",msgId , subMsgId, buf, err.Error())
	}
	GlobalMutex.Unlock()
}

//------------------------go 给lua传递 1个 int-----------------------------------------------
func (m *MyLua) GoCallLuaLogicInt(module string,funcName string,ii int) {
	table:= m.L.GetGlobal(module)
	value := m.L.GetField(table,funcName)
	GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: value,
		NRet: 0,
		Protect: true,
	},lua.LNumber(ii)); err != nil {		// 参数
		zLog.PrintLogger("GoCallLuaLogicInt error :"+ funcName+"      "+err.Error())
	}
	GlobalMutex.Unlock()
}

// ----------------------Lua重新加载，Lua的热更新按钮----------------------------------------
func (m *MyLua)GoCallLuaReload() error {
	//fmt.Println("----------lua reload--------------")
	GlobalMutex.Lock()
	var err error
	err = m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("ReloadAll"), //reloadUp  ReloadAll
		NRet: 0,
		Protect: true,
	})

	GlobalMutex.Unlock()
	if err != nil {
		zLog.PrintLogger("热更新出错 "+err.Error())
	}
	return err
}
//
//// ----------------------Lua连接mysql----------------------------------------
//func (m *MyLua)GoCallLuaConnectMysql(addr string,db string ,user string ,pwd string) bool {
//	GlobalMutex.Lock()
//
//	if err := m.L.CallByParam(lua.P{
//		Fn: m.L.GetGlobal("MysqlConnect"),
//		NRet: 1,
//		Protect: true,
//	},lua.LString(addr),lua.LString(db),lua.LString(user),lua.LString(pwd)); err != nil {		// 参数
//		zLog.PrintLogger("GoCallLuaConnectMysql error :"+err.Error())
//	}
//	ret := m.L.Get(1) // returned value
//	//fmt.Println("ret",ret, reflect.TypeOf(ret))
//	m.L.Pop(1)  // remove received value
//	GlobalMutex.Unlock()
//	if ret == lua.LTrue {
//		return true
//	}
//	return false
//}

// ----------------------将go的变量传递给lua， 用来改变lua的全局变量值，一般用于统计和监控----------------------
func (m *MyLua) GoCallLuaSetStringVar(name string, value string) {
	GlobalMutex.Lock()
	m.L.SetGlobal(name, lua.LString(value))
	GlobalMutex.Unlock()
}
func (m *MyLua) GoCallLuaSetIntVar(name string, value int) {
	GlobalMutex.Lock()
	m.L.SetGlobal(name, lua.LNumber(value))
	GlobalMutex.Unlock()
}