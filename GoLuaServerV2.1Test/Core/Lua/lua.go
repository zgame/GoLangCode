package Lua

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"sync"
)
//--------------------------------------------------------------------------------
// lua的接口，包含热更新
//--------------------------------------------------------------------------------


var (
	GameManagerLuaHandle *MyLua // 主线程的lua句柄
	GlobalMutex sync.Mutex      // 主要用于lua逻辑调用时候的加锁
	RWMutex sync.RWMutex        // 主要用于针对map进行读写时候的锁
)


type MyLua struct {
	L *lua.LState
}


func NewMyLua() *MyLua {
	l := lua.NewState()
	return &MyLua{L: l}
}

// --------------------全局变量初始化--------------------------
func InitGlobalVar() {
	ConnectMyUdpServer = make(map[int]*MyUdpServer)
	ConnectMyTcpServer = make(map[int]*MyTcpServer)
	ConnectMyTcpServerByUID = make(map[int]*MyTcpServer)
	//GameManagerReceiveCh = make(chan lua.LValue)// 这是每个玩家线程跟主线程之间的通信用channel
	//GameManagerSendCh = make(chan lua.LValue)

}

// 通过lua堆栈找到对应的是哪个myServer
func GetMyTcpServerByLSate(id int) *MyTcpServer {
	RWMutex.RLock()
	re := ConnectMyTcpServer[id] // 这是全局变量，所以要加锁， 读写都要加
	RWMutex.RUnlock()
	return re
}
// 通过 user id 找到对应的是哪个myServer
func GetMyTcpServerByUID(uid int) *MyTcpServer {
	RWMutex.RLock()
	re:= ConnectMyTcpServerByUID[uid] // 这是全局变量，所以要加锁， 读写都要加
	RWMutex.RUnlock()
	return re
}


// 通过lua堆栈找到对应的是哪个myServer
func GetMyUdpServerByLSate(id int) *MyUdpServer {
	RWMutex.RLock()
	re := ConnectMyUdpServer[id] // 这是全局变量，所以要加锁， 读写都要加
	RWMutex.RUnlock()
	return re
}

//----------------------对象个体初始化-----------------------
func (m *MyLua)Init()   {
	//L := luaPool.Get()		// 这是用池的方式， 但是玩家数据需要清理重置，以后再考虑吧
	//defer luaPool.Put(L)
	//m.L.SetGlobal("GameManagerReceiveCh", lua.LChannel(GameManagerReceiveCh))// 这是每个玩家线程跟主线程之间的通信用channel
	//m.L.SetGlobal("GameManagerSendCh", lua.LChannel(GameManagerSendCh))

	m.InitResister() // 这里是统一的lua函数注册入口

	if err := m.L.DoFile("Script/main.lua"); err != nil {
		fmt.Println("加载main.lua文件出错了！")
		fmt.Println(err.Error())
		panic("--------------error-----------------")
	}
	//DoCompiledFile(m.L, GlobalVar.LuaCodeToShare)
	//return m.L
	//fmt.Println("--------lua 脚本 加载完成！---------------")
}
