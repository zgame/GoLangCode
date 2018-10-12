package Lua

import (
	"github.com/yuin/gopher-lua"
	"strconv"
)

// 统一的go给lua调用的函数注册点
func InitResister(L *lua.LState) {
	// Lua调用go函数声明
	L.SetGlobal("double", L.NewFunction(Double))
}





// Lua调用的go函数-----------------------------------------------------
// 需要像下面一样，start的时候先注册进去，才可以正常调用
//L.SetGlobal("double", L.NewFunction(Double))
func Double(L *lua.LState) int {
	lv := L.ToInt(1)             //第一个参数
	lv2 :=  L.ToInt(2)			 //第二个参数
	str := L.ToString(3)

	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */

	return 1                     /* number of results */
}

func NetWorkSend(L *lua.LState) int {
	lv := L.ToInt(1)             //第一个参数
	lv2 :=  L.ToInt(2)			 //第一个参数
	str := L.ToString(3)

	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */

	return 1			// 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}