package main

import (
	"github.com/yuin/gopher-lua"
	"fmt"
	"strconv"
	"time"
)


func main() {

	go start(1)
	go start(2)
	//goCallLua(L)


	for{
		select {

		}
	}


}

func start(timer time.Duration) {
	L := lua.NewState()
	defer L.Close()

	// Lua调用go函数声明
	// 声明double函数为Lua的全局函数，绑定go函数Double
	L.SetGlobal("double", L.NewFunction(Double))


	// 执行lua文件
	if err := L.DoFile("main.lua"); err != nil {
		fmt.Println("加载main.lua文件出错了！")
		fmt.Println(err.Error())
	}


	tickerCheckUpdateData := time.NewTicker(time.Second * timer)
	defer tickerCheckUpdateData.Stop()

	for{
		select {
		case <-tickerCheckUpdateData.C:
			timerFunc(L,timer)
		}
	}

}


//-------------计时器------------------------
func timerFunc(L *lua.LState,timer time.Duration)  {
	//fmt.Println("timer--------")
	goCallLuaReload(L)
	goCallLua(L,int(timer))
	//goCallLua(L)
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


	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */

	return 1                     /* number of results */
}


// 计时器，用来定期检查配置的更新，包括后台控制的活动，开关，配置文件更新，用数据版本号来控制
func TimerCheckUpdate(f func())  {
	go func() {
		tickerCheckUpdateData := time.NewTicker(time.Second * 2)
		defer tickerCheckUpdateData.Stop()

		for {
			select {
			case <-tickerCheckUpdateData.C:
				f()
			}
		}
	}()
}
