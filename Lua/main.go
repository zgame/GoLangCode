package main

import (
	"github.com/yuin/gopher-lua"
	"fmt"
	"strconv"
)


func main() {
	L := lua.NewState()
	defer L.Close()

	// Lua调用go函数声明
	// 声明double函数为Lua的全局函数，绑定go函数Double
	L.SetGlobal("double", L.NewFunction(Double))


	// 执行lua文件
	if err := L.DoFile("hello.lua"); err != nil {
		fmt.Println("---------------")
		fmt.Println(err.Error())
		fmt.Println("---------------")
	}

	fmt.Println("----------go call lua--------------")
	// 这里是go调用lua的函数
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("Zsw2"),
		NRet: 2,
		Protect: true,
	}, lua.LNumber(2),lua.LNumber(2)); err != nil {
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

//-----------------------------------------------------
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


// Lua调用的go函数
func Double(L *lua.LState) int {
	lv := L.ToInt(1)             //第一个参数
	lv2 :=  L.ToInt(2)			 //第一个参数
	str := L.ToString(3)

	L.Push(lua.LNumber(22))
	L.Push(lua.LNumber(22))
	L.Push(lua.LNumber(22))
	L.Push(lua.LNumber(22))

	L.Push(lua.LNumber(lv * lv2)) /* push result */

	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */
	L.Push(lua.LNumber(22))


	return 1                     /* number of results */
}


