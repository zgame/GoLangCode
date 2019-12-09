package zBit32

import (
	"github.com/yuin/gopher-lua"
)

// ----------------------------------------------------------------------------
// lua 位运算
// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"band":   bandFn,	//与
	"bnot":   bnotFn,	//非
	"bor":    borFn,	//或
	"bxor":   bxorFn,	//异或
	"lshift": lshiftFn,	//左移
	"rshift": rshiftFn,	//右移
}

// ----------------------------------------------------------------------------

func Bit32Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

// ----------------------------------------------------------------------------

func LuaBit32Load(L *lua.LState) {
	L.PreloadModule("bit32", Bit32Loader)
}


// 与
func bandFn(L *lua.LState) int {
	a := L.CheckInt64(1)
	b := L.CheckInt64(2)
	result := a & b
	L.Push(lua.LNumber(result))
	return 1
}

// 非
func bnotFn(L *lua.LState) int {
	a := L.CheckInt64(1)
	result := ^a
	L.Push(lua.LNumber(result))
	return 1
}

// 或
func borFn(L *lua.LState) int {
	a := L.CheckInt64(1)
	b := L.CheckInt64(2)
	result := a | b
	L.Push(lua.LNumber(result))
	return 1
}

// 异或
func bxorFn(L *lua.LState) int {
	a := L.CheckInt64(1)
	b := L.CheckInt64(2)
	result := a ^ b
	L.Push(lua.LNumber(result))
	return 1
}
// 左移
func lshiftFn(L *lua.LState) int {
	a := L.CheckInt64(1)
	b := L.CheckInt(2)
	result := a << uint(b)
	L.Push(lua.LNumber(result))
	return 1
}
// 右移
func rshiftFn(L *lua.LState) int {
	n := L.CheckInt64(1)
	n2 := L.CheckInt(2)
	result := n >> uint8(n2)
	L.Push(lua.LNumber(result))
	return 1
}
