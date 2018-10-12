package Lua

import (
	"github.com/yuin/gopher-lua"
	"fmt"
)

// go调用lua函数--------------------------------------------------
func GoCallLua(L *lua.LState, num int)  {
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

// go触发lua函数，不带参数和返回值
func GoCallLuaCommonLogic(L *lua.LState, funcName string)  {
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}); err != nil {		// 参数
		fmt.Println("",err.Error())
	}
}

