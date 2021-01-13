package zRandom

import (
	"GoLuaServerV2.1/Core/Utils/zLua"
	lua "github.com/yuin/gopher-lua"
	"math/rand"
	"time"
)


var exports = map[string]lua.LGFunction{
	"intN": IntN,
	"random":  Random,
	"float":  RandomFloatTo,
	"percent":  PercentRate,
	"normal":  Normal,
	"exp":  Exp,
	"perm":  Perm,
}

// ----------------------------------------------------------------------------

func zRandomLoader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

// ----------------------------------------------------------------------------

func LuaRandomLoad(L *lua.LState) {
	L.PreloadModule("zRandom", zRandomLoader)
}


// 随机 [0,num)
func IntN(L *lua.LState) int {
	num := L.CheckInt(1)
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	x := rand.Intn(num)
	L.Push(lua.LNumber(x))
	return 1
}

// 随机[ min, max)
func Random(L *lua.LState) int {
	min := L.CheckInt(1)
	max := L.CheckInt(2)
	if min >= max || max == 0 {
		//fmt.Println("随机数格式不正确")
		return max
	}
	rand.Seed(time.Now().UnixNano())
	L.Push(lua.LNumber(rand.Intn(max-min) + min))
	return 1
}

// 随机[ min, max) float32
func RandomFloatTo(L *lua.LState) int {
	min := L.CheckNumber(1)
	max := L.CheckNumber(2)
	if min >= max || max <= 0 {
		//fmt.Println("随机数格式不正确")
		L.Push(max)
		return 1
	}
	rand.Seed(time.Now().UnixNano())
	ran := lua.LNumber(rand.Float64())
	L.Push(min + ran * (max-min))
	return 1
}
// 随机[ min, max)
func RandomTo(min int,max int) int {
	if min >= max || max == 0 {
		//fmt.Println("随机数格式不正确")
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// 获取百分比方法， 比如10几率， 那么小于等于10，返回true
func PercentRate(L *lua.LState) int {
	rate := L.CheckNumber(1)
	rr := lua.LNumber(RandomTo(1,101))
	if rr <= rate{
		L.Push(lua.LTrue)
	}else{
		L.Push(lua.LFalse)
	}

	return 1
}

// 正态分布
func Normal(L * lua.LState)  int{
	rand.Seed(time.Now().UnixNano())
	L.Push(lua.LNumber(rand.NormFloat64()))
	return 1
}

// 指数分布
func Exp(L * lua.LState)  int{
	rand.Seed(time.Now().UnixNano())
	L.Push(lua.LNumber(rand.ExpFloat64()))
	return 1
}

// 随机数组切片
func Perm(L * lua.LState)  int{
	len := L.CheckInt(1)
	rand.Seed(time.Now().UnixNano())
	L.Push(zLua.LuaSetValue(L,rand.Perm(len)))
	return 1
}