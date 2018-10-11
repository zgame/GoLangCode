package main

import (
	"github.com/yuin/gopher-lua"
	"os"
	"bufio"
	"github.com/yuin/gopher-lua/parse"
	"fmt"
	"strconv"
)

//------------------编译lua文件------------------------------

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file.
func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}




// Lua重新加载，Lua的热更新按钮
func goCallLuaReload(L *lua.LState)  {
	//fmt.Println("----------lua reload--------------")
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("ReloadAll"), //reloadUp  ReloadAll
		NRet: 0,
		Protect: true,
	}); err != nil {
		fmt.Println("",err.Error())
	}
}

// go调用lua函数
func goCallLua(L *lua.LState, num int)  {
	fmt.Println("----------go call lua--------------")
	// 这里是go调用lua的函数
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("Zsw2"),
		NRet: 2,
		Protect: true,
	}, lua.LNumber(num),lua.LNumber(num)); err != nil {
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

//-----------------------lua 对应的类型列表------------------------------
//Type name	Go type	Type() value	Constants
//LNilType	(constants)	LTNil	LNil
//LBool	(constants)	LTBool	LTrue, LFalse
//LNumber	float64	LTNumber	-
//LString	string	LTString	-
//LFunction	struct pointer	LTFunction	-
//LUserData	struct pointer	LTUserData	-
//LState	struct pointer	LTThread	-
//LTable	struct pointer	LTTable	-
//LChannel	chan LValue	LTChannel	-
//-----------------------------------------------------




// Lua调用的go函数，需要像下面一样，start的时候先注册进去，才可以正常调用
//L.SetGlobal("double", L.NewFunction(Double))
func Double(L *lua.LState) int {
	lv := L.ToInt(1)             //第一个参数
	lv2 :=  L.ToInt(2)			 //第一个参数
	str := L.ToString(3)

	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */

	return 1                     /* number of results */
}
