package main

import (
	"github.com/yuin/gopher-lua"
	"fmt"
	"strconv"
	"time"
	"os"
	"bufio"
	"github.com/yuin/gopher-lua/parse"
	"unsafe"
	"runtime"
	"sync"
)

var num = 0
var codeToShare *lua.FunctionProto

var ch chan lua.LValue
var quit chan lua.LValue


var mutex sync.Mutex
//const FileLua  = "Script/main.lua"
//const FileLua  = "goroutine.lua"
//const FileLua  = "user.lua"
const FileLua  = "main.lua"

func main() {
	ch = make(chan lua.LValue)
	quit = make(chan lua.LValue)

	runtime.GOMAXPROCS(4)

	//defer luaPool.Shutdown()
	var err error
	codeToShare,err = CompileLua(FileLua)
	if err!=nil{
		fmt.Println("加载main.lua文件出错了！")
	}

	////支线程
	//for i:= 1;i<3;i++ {
	//	go start(1, i) // 间隔时间， 编号
	//
	//}


	//主线程
	L := lua.NewState()
	//L.SetGlobal("ch", lua.LChannel(ch))
	//L.SetGlobal("quit", lua.LChannel(quit))
	// 声明double函数为Lua的全局函数，绑定go函数Double
	//L.SetGlobal("zClose", L.NewFunction(zClose))

	// 直接调用luaopen_pb
	luaopen_pb(L)



	//DoCompiledFile(L, codeToShare)

	if err := L.DoFile(FileLua); err != nil {
		fmt.Println("加载main.lua文件出错了！")
		fmt.Println(err.Error())
	}

	//go func() {
	//	for {
	//		goCallLuaSelect(L)
	//	}
	//}()



	//for i:=0;i<10;i++{
	//	go func() {
	//		for{
	//			mutex.Lock()
	//			GoCallLuaLogic(L,"add_uid")
	//			GoCallLuaLogic(L,"show_uid")
	//			mutex.Unlock()
	//			//time.Sleep(time.Second*1)
	//		}
	//	}()
	//
	//}


	for{
		//fmt.Println("主循环")
		//GoCallLuaLogic(L,"test")
		mutex.Lock()
		GoCallLuaLogic(L,"Run")
		mutex.Unlock()
		//goCallLuaSelect(L)			// 主线程监听
		time.Sleep(time.Millisecond * 1000 * 1)
		//select {
		//
		//}
	}


}

// 把指针传递过去给dll
func IntPtr(L *lua.LState) uintptr {
	return uintptr(unsafe.Pointer(L))
}
func start(timer time.Duration, index int) {


	L := lua.NewState()
	//L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ch", lua.LChannel(ch))
	L.SetGlobal("quit", lua.LChannel(quit))

	//L := luaPool.Get()
	//defer luaPool.Put(L)


	// 直接调用luaopen_pb
	luaopen_pb(L)

	// Lua调用go函数声明
	// 声明double函数为Lua的全局函数，绑定go函数Double
	L.SetGlobal("zClose", L.NewFunction(zClose))
	//L.Register("double", Double)
	//lua.RegistrySize = 10244 * 20
	//lua.CallStackSize = 10244
	//DoCompiledFile(L, codeToShare)

	// 执行lua文件
	if err := L.DoFile(FileLua); err != nil {
		fmt.Println("加载main.lua文件出错了！")
		fmt.Println(err.Error())
	}

	// 通过dll加载luaopen_pb
	//DllTestDef := syscall.MustLoadDLL("libpb.dll")
	//add := DllTestDef.MustFindProc("luaopen_pb")
	//ret, _, err := add.Call(IntPtr(L))
	//if err!=nil{
	//	fmt.Println("返回",ret)
	//}





	tickerCheckUpdateData := time.NewTicker(time.Second * timer)
	defer tickerCheckUpdateData.Stop()

	for{
		//goCallLuaSelect(L)
		select {
		case <-tickerCheckUpdateData.C:
			timerFunc(L,index)

		}
	}

}



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


//// ---------------------lua 文件池-------------------------------------------
//type lStatePool struct {
//	m     sync.Mutex
//	saved []*lua.LState
//}
//
//func (pl *lStatePool) Get() *lua.LState {
//	pl.m.Lock()
//	defer pl.m.Unlock()
//	n := len(pl.saved)
//	if n == 0 {
//		return pl.New()
//	}
//	x := pl.saved[n-1]
//	pl.saved = pl.saved[0 : n-1]
//	return x
//}
//
//func (pl *lStatePool) New() *lua.LState {
//	L := lua.NewState()
//
//	// 执行lua文件
//	if err := L.DoFile("main.lua"); err != nil {
//		fmt.Println("加载main.lua文件出错了！")
//		fmt.Println(err.Error())
//	}
//	// setting the L up here.
//	// load scripts, set global variables, share channels, etc...
//	return L
//}
//
//func (pl *lStatePool) Put(L *lua.LState) {
//	pl.m.Lock()
//	defer pl.m.Unlock()
//	pl.saved = append(pl.saved, L)
//}
//
//func (pl *lStatePool) Shutdown() {
//	for _, L := range pl.saved {
//		L.Close()
//	}
//}
//
//// Global LState pool
//var luaPool = &lStatePool{
//	saved: make([]*lua.LState, 0, 4),
//}
//





//-------------计时器------------------------
func timerFunc(L *lua.LState,index int)  {
	//fmt.Println("timer--------")
	//goCallLuaReload(L)
	//goCallLua(L,int(timer))

	num++
	//goCallLuaSend(L, strconv.Itoa(index))			// 分支线程发消息
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
	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */

	return 2                    /* number of results */
}

// Lua调用的go函数
func zClose(L *lua.LState) int {
	L.Close()
	return 0                    /* number of results */
}


// go调用lua函数
func goCallLuaSelect(L *lua.LState)  {
	// 这里是go调用lua的函数
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("receivezz"),
		NRet: 0,
		Protect: true,
	}); err != nil {
		fmt.Println("",err.Error())
	}

	//ret := L.Get(1) // returned value
	//fmt.Println("lua return: ",ret)
	//ret = L.Get(2) // returned value
	//fmt.Println("lua return: ",ret)
	//L.Pop(1)  // remove received value
	//L.Pop(1)  // remove received value
}

// go调用lua函数
func goCallLuaSend(L *lua.LState,myName string)  {
	// 这里是go调用lua的函数
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("sendzz"),
		NRet: 0,
		Protect: true,
	},lua.LString(myName)); err != nil {
		fmt.Println("",err.Error())
	}

	//ret := L.Get(1) // returned value
	//fmt.Println("lua return: ",ret)
	//ret = L.Get(2) // returned value
	//fmt.Println("lua return: ",ret)
	//L.Pop(1)  // remove received value
	//L.Pop(1)  // remove received value
}

func GoCallLuaLogic(L *lua.LState,funcName string) {
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal(funcName),		// lua的函数名字
		NRet: 0,
		Protect: true,
	}); err != nil {		// 参数
		fmt.Println("GoCallLuaLogic error :",funcName, "      ",err.Error())
	}
}