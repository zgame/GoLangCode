package Lua

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"sync"
)
//--------------------------------------------------------------------------------
// lua的接口，包含热更新
//--------------------------------------------------------------------------------

var GameManagerLuaHandle *MyLua      // 主线程的lua句柄
var ConnectMyTcpServer sync.Map      //[int]*MyTcpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调
var ConnectMyTcpServerByUid sync.Map //[int]*MyTcpServer // 将uid跟连接句柄进行哈希

//var ConnectMyUdpServer sync.Map   // [int]*MyUdpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调

type MyLua struct {
	L *lua.LState
}


func NewMyLua() *MyLua {
	l := lua.NewState()
	return &MyLua{L: l}
}

// --------------------全局变量初始化--------------------------
func InitGlobalVar() {

}


//----------------------对象个体初始化-----------------------
func (m *MyLua)Init()   {
	m.InitResister() // 这里是统一的lua函数注册入口

	if err := m.L.DoFile("Script/main.lua"); err != nil {
		fmt.Println("加载main.lua文件出错了！")
		fmt.Println(err.Error())
		panic("--------------error-----------------")
	}
}
