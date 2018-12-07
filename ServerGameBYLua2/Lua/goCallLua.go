package Lua

import (
	"github.com/yuin/gopher-lua"
	"fmt"
	"../GlobalVar"
	"../Utils/log"
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
func (m *MyLua) GoCallLuaLogic(funcName string) {
	GlobalVar.GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}); err != nil {		// 参数
		log.PrintLogger("GoCallLuaLogic error :"+funcName+"      "+err.Error())
	}
	GlobalVar.GlobalMutex.Unlock()
}

// -------------------go传递接收到的网络数据包给lua-------------------
func (m *MyLua)GoCallLuaNetWorkReceive(serverId int,userId int,msgId int , subMsgId int ,buf string) {
	GlobalVar.GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("GoCallLuaNetWorkReceive"),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}, lua.LNumber(serverId),lua.LNumber(userId),lua.LNumber(msgId), lua.LNumber(subMsgId), lua.LString(buf)); err != nil {		// 参数
		log.PrintfLogger("GoCallLuaNetWorkReceive  error :  msgId:%d  subMsgId %d  buf:%s   error:%s",msgId , subMsgId, buf, err.Error())
	}
	GlobalVar.GlobalMutex.Unlock()
}

//------------------------go 给lua传递 1个 int-----------------------------------------------
func (m *MyLua) GoCallLuaLogicInt(funcName string,ii int) {
	GlobalVar.GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	},lua.LNumber(ii)); err != nil {		// 参数
		log.PrintLogger("GoCallLuaLogicInt error :"+ funcName+"      "+err.Error())
	}
	GlobalVar.GlobalMutex.Unlock()
}

// ----------------------Lua重新加载，Lua的热更新按钮----------------------------------------
func (m *MyLua)GoCallLuaReload() error {
	//fmt.Println("----------lua reload--------------")
	GlobalVar.GlobalMutex.Lock()
	var err error
	err = m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("ReloadAll"), //reloadUp  ReloadAll
		NRet: 0,
		Protect: true,
	})

	GlobalVar.GlobalMutex.Unlock()
	if err != nil {
		log.PrintLogger("热更新出错 "+err.Error())
	}
	return err
}

// 将服务器地址和端口传递给lua， 记录用
func (m *MyLua)GoCallLuaSetVar(name string, address string) {
	m.L.SetGlobal(name, lua.LString(address))
}