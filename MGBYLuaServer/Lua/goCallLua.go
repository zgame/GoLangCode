package Lua

import (
	"MGBYLuaServer/GlobalVar"
	"MGBYLuaServer/Utils/log"
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


// -------------------go通知lua网络连接成功-------------------
func (m *MyLua)GoCallLuaNetWorkInit(serverId int) {
	GlobalVar.GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("GoCallLuaNetWorkInit"),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}, lua.LNumber(serverId)); err != nil {		// 参数
		log.PrintfLogger("GoCallLuaNetWorkInit  error :  serverId:%d error:%s",serverId, err.Error())
	}
	GlobalVar.GlobalMutex.Unlock()
}

// -------------------go传递接收到的网络数据包给lua-------------------
func (m *MyLua)GoCallLuaNetWorkReceive(serverId int,userId int,msgId int , subMsgId int ,buf string, token int) {
	GlobalVar.GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("GoCallLuaNetWorkReceive"),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}, lua.LNumber(serverId),lua.LNumber(userId),lua.LNumber(msgId), lua.LNumber(subMsgId), lua.LString(buf),  lua.LNumber(token)); err != nil {		// 参数
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
//------------------------go 给lua传递 2个 int-----------------------------------------------
func (m *MyLua) GoCallLuaLogicInt2(funcName string,ii int, ii2 int) {
	GlobalVar.GlobalMutex.Lock()
	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	},lua.LNumber(ii),lua.LNumber(ii2)); err != nil {		// 参数
		log.PrintLogger("GoCallLuaLogicInt2 error :"+ funcName+"      "+err.Error())
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

// ----------------------Lua连接mysql----------------------------------------
func (m *MyLua)GoCallLuaConnectMysql(addr string, port string, db string ,user string ,pwd string) bool {
	GlobalVar.GlobalMutex.Lock()

	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("MysqlConnect"),
		NRet: 1,
		Protect: true,
	},lua.LString(addr),lua.LString(port),lua.LString(db),lua.LString(user),lua.LString(pwd)); err != nil {		// 参数
		log.PrintLogger("GoCallLuaConnectMysql error :"+err.Error())
	}
	ret := m.L.Get(1) // returned value
	//fmt.Println("ret",ret, reflect.TypeOf(ret))
	m.L.Pop(1)  // remove received value
	GlobalVar.GlobalMutex.Unlock()
	if ret == lua.LTrue {
		return true
	}
	return false
}

// ----------------------Lua连接sql server----------------------------------------
func (m *MyLua)GoCallLuaConnectSqlServer(addr string, port string, db string ,user string ,pwd string) bool {
	GlobalVar.GlobalMutex.Lock()

	if err := m.L.CallByParam(lua.P{
		Fn: m.L.GetGlobal("SqlServerConnect"),
		NRet: 1,
		Protect: true,
	},lua.LString(addr),lua.LString(port),lua.LString(db),lua.LString(user),lua.LString(pwd)); err != nil {		// 参数
		log.PrintLogger("GoCallLuaConnectSqlServer error :"+err.Error())
	}
	ret := m.L.Get(1) // returned value
	//fmt.Println("ret",ret, reflect.TypeOf(ret))
	m.L.Pop(1)  // remove received value
	GlobalVar.GlobalMutex.Unlock()
	if ret == lua.LTrue {
		return true
	}
	return false
}
// ----------------------将go的变量传递给lua， 用来改变lua的全局变量值，一般用于统计和监控----------------------
func (m *MyLua) GoCallLuaSetStringVar(name string, value string) {
	GlobalVar.GlobalMutex.Lock()
	m.L.SetGlobal(name, lua.LString(value))
	GlobalVar.GlobalMutex.Unlock()
}
func (m *MyLua) GoCallLuaSetIntVar(name string, value int) {
	GlobalVar.GlobalMutex.Lock()
	m.L.SetGlobal(name, lua.LNumber(value))
	GlobalVar.GlobalMutex.Unlock()
}