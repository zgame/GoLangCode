package Lua

import (
	"GoLuaServerV2.1/Utils/zLog"
	"fmt"
	"github.com/yuin/gopher-lua"
	"sync"
)

//--------------------------------------------------------------------------------
// go 调用 lua的函数
//--------------------------------------------------------------------------------
var MyLuaMutex sync.Mutex // 主要用于lua逻辑调用时候的加锁

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
func (m *MyLua) GoCallLuaLogic(module string,funcName string) {
	MyLuaMutex.Lock()
	table:= m.L.GetGlobal(module)
	value := m.L.GetField(table,funcName)
	if err := m.L.CallByParam(lua.P{
		Fn: value,		// lua的函数名字
		NRet: 0,
		Protect: true,
	}); err != nil {		// 参数
		zLog.PrintLogger("GoCallLuaLogic error :"+funcName+"      "+err.Error())
	}
	MyLuaMutex.Unlock()
}


// -------------------go通知lua网络连接成功-------------------
func (m *MyLua)GoCallLuaNetWorkInit(serverId int) {
	MyLuaMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("GoCallLuaNetWorkInit"),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}, lua.LNumber(serverId)); err != nil {		// 参数
		zLog.PrintfLogger("GoCallLuaNetWorkInit  error :  serverId:%d error:%s",serverId, err.Error())
	}
	MyLuaMutex.Unlock()
}

// -------------------go传递接收到的网络数据包给lua-------------------
func (m *MyLua)GoCallLuaNetWorkReceive(serverId int,userId int,msgId int , subMsgId int ,buf string, token int) {
	MyLuaMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("GoCallLuaNetWorkReceive"),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}, lua.LNumber(serverId),lua.LNumber(userId),lua.LNumber(msgId), lua.LNumber(subMsgId), lua.LString(buf),  lua.LNumber(token)); err != nil {		// 参数
		zLog.PrintfLogger("GoCallLuaNetWorkReceive  error :  msgId:%d  subMsgId %d  buf:%s   error:%s",msgId , subMsgId, buf, err.Error())
	}
	MyLuaMutex.Unlock()
}

// -------------------go传递接收到的网络数据包给lua-------------------
func (m *MyLua)GoCallLuaNetWorkReceiveUdp(serverAddr string,msgId int , subMsgId int ,buf string) {
	MyLuaMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("GoCallLuaNetWorkUdpReceive"),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}, lua.LString(serverAddr),lua.LNumber(msgId), lua.LNumber(subMsgId), lua.LString(buf)); err != nil {		// 参数
		zLog.PrintfLogger("GoCallLuaNetWorkReceive  error :  msgId:%d  subMsgId %d  buf:%s   error:%s",msgId , subMsgId, buf, err.Error())
	}
	MyLuaMutex.Unlock()
}

//------------------------go 给lua传递 1个 int-----------------------------------------------
func (m *MyLua) GoCallLuaLogicInt(funcName string,ii int) {
	MyLuaMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	},lua.LNumber(ii)); err != nil {		// 参数
		zLog.PrintLogger("GoCallLuaLogicInt error :"+ funcName+"      "+err.Error())
	}
	MyLuaMutex.Unlock()
}
//------------------------go 给lua传递 2个 int-----------------------------------------------
func (m *MyLua) GoCallLuaLogicInt2(funcName string,ii int, ii2 int) {
	MyLuaMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	},lua.LNumber(ii),lua.LNumber(ii2)); err != nil {		// 参数
		zLog.PrintLogger("GoCallLuaLogicInt2 error :"+ funcName+"      "+err.Error())
	}
	MyLuaMutex.Unlock()
}

// ----------------------Lua重新加载，Lua的热更新按钮----------------------------------------
func (m *MyLua)GoCallLuaReload() error {
	//fmt.Println("----------lua reload--------------")
	MyLuaMutex.Lock()
	var err error
	err = m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("ReloadAll"), //reloadUp  ReloadAll
		NRet: 0,
		Protect: true,
	})

	MyLuaMutex.Unlock()
	if err != nil {
		zLog.PrintLogger("热更新出错 "+err.Error())
	}
	return err
}


// ----------------------将go的变量传递给lua， 用来改变lua的全局变量值，一般用于统计和监控----------------------
func (m *MyLua) GoCallLuaSetStringVar(module string,key string, value string) {
	MyLuaMutex.Lock()
	table := m.L.GetGlobal(module)
	m.L.SetField(table,key, lua.LString(value))
	MyLuaMutex.Unlock()
}
func (m *MyLua) GoCallLuaSetIntVar(module string, key string, value int) {
	MyLuaMutex.Lock()
	//m.L.SetGlobal(name, lua.LNumber(value))
	table := m.L.GetGlobal(module)
	m.L.SetField(table,key, lua.LNumber(value))
	MyLuaMutex.Unlock()
}


