package Lua

import (
	"GoLuaServerV2.1/GlobalVar"
	"fmt"
	"github.com/yuin/gopher-lua"
)
//--------------------------------------------------------------------------------
// lua的接口，包含热更新
//--------------------------------------------------------------------------------

//var GameManagerReceiveCh chan lua.LValue		// 这是每个玩家线程跟主线程之间的通信用channel
//var GameManagerSendCh chan lua.LValue			// 这是主线程给每个玩家线程跟之间的通信用channel
var GameManagerLuaHandle *MyLua                  // 主线程的lua句柄
var ConnectMyTcpServer map[int]*MyTcpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调
var ConnectMyTcpServerByUid map[int]*MyTcpServer // 将uid跟连接句柄进行哈希
var ConnectMyUdpServer map[int]*MyUdpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调

type MyLua struct {
	L *lua.LState
}


func NewMyLua() *MyLua {
	l := lua.NewState()
	return &MyLua{L:l}
}

// --------------------全局变量初始化--------------------------
func InitGlobalVar() {
	ConnectMyTcpServer = make(map[int]*MyTcpServer)
	ConnectMyTcpServerByUid = make(map[int]*MyTcpServer)
	ConnectMyUdpServer = make(map[int]*MyUdpServer)
	//GameManagerReceiveCh = make(chan lua.LValue)// 这是每个玩家线程跟主线程之间的通信用channel
	//GameManagerSendCh = make(chan lua.LValue)

}

// 通过lua堆栈找到对应的是哪个myServer
func GetMyServerByServerId(serverId int) *MyTcpServer {
	GlobalVar.RWMutex.RLock()
	re := ConnectMyTcpServer[serverId] // 这是全局变量，所以要加锁， 读写都要加
	GlobalVar.RWMutex.RUnlock()
	return re
}
// 通过 user id 找到对应的是哪个myServer
func GetMyServerByUID(uid int) *MyTcpServer {
	GlobalVar.RWMutex.RLock()
	re:= ConnectMyTcpServerByUid[uid] // 这是全局变量，所以要加锁， 读写都要加
	GlobalVar.RWMutex.RUnlock()
	return re
}

// 通过lua堆栈找到对应的是哪个myServer
func GetMyUdpServerByLSate(id int) *MyUdpServer {
	GlobalVar.RWMutex.RLock()
	re := ConnectMyUdpServer[id] // 这是全局变量，所以要加锁， 读写都要加
	GlobalVar.RWMutex.RUnlock()
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


//------------------编译lua文件------------------------------
// CompileLua reads the passed lua file from disk and compiles it.
//func CompileLua(filePath string) (*lua.FunctionProto, error) {
//	file, err := os.Open(filePath)
//	defer file.Close()
//	if err != nil {
//		return nil, err
//	}
//	reader := bufio.NewReader(file)
//	chunk, err := parse.Parse(reader, filePath)
//	if err != nil {
//		return nil, err
//	}
//	proto, err := lua.Compile(chunk, filePath)
//	if err != nil {
//		return nil, err
//	}
//	return proto, nil
//}
//
//// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
//// to calling DoFile on the LState with the original source file.
//func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
//	lFunc := L.NewFunctionFromProto(proto)
//	L.Push(lFunc)
//	return L.PCall(0, lua.MultRet, nil)
//}



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


