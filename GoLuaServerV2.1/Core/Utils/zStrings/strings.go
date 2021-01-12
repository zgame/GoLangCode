package zStrings

//----------------------------------------------------------------------------
// 字符串处理
//----------------------------------------------------------------------------

import (
	"GoLuaServerV2.1/Core/Utils/zLua"
	"github.com/yuin/gopher-lua"
	"strings"
)

var exports = map[string]lua.LGFunction{
	"split": split,
	"join":  join,
	"trim":  Trim,
}

// ----------------------------------------------------------------------------

func zStringsLoader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

// ----------------------------------------------------------------------------

func LuaStringsLoad(L *lua.LState) {
	L.PreloadModule("zStrings", zStringsLoader)
}

// 字符串分割
func split(L *lua.LState) int {
	src := L.CheckString(1)
	sep := L.CheckString(2)
	converted := strings.Split(src,sep)
	arr := L.NewTable()
	for _, item := range converted {
		arr.Append(lua.LString(item))
	}
	L.Push(arr)
	return 1
}
// 字符串衔接
func join(L *lua.LState) int {
	list:= zLua.LuaGetValue(L,1)
	sep := L.CheckString(2)

	str:= strings.Join(list.([]string),sep)
	//fmt.Println(string(out))
	L.Push(lua.LString(str))

	return 1
}

// 字符串修剪
func Trim(L *lua.LState) int {
	str := L.CheckString(1)
	cut := L.CheckString(2)
	re:= strings.Trim(str, cut)
	L.Push(lua.LString(re))
	return 1
}