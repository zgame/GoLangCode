package main

import (
	"fmt"
	"net"
	"./NetWork"
	"github.com/yuin/gopher-lua"
	"./Lua"
)

// ----------------------------服务器处理的统一接口----------------------------------
// myServer其实是一个个连接单独处理的模块
//-----------------------------------------------------------------------------------

//type Agent interface {
//	WriteMsg(msg ... []byte)
//	LocalAddr() net.Addr
//	RemoteAddr() net.Addr
//	Close()
//	Destroy()
//	UserData() interface{}
//	SetUserData(data interface{})
//}

// wsServer.NewAgent 服务器连接的代理
type myServer struct {
	conn NetWork.Conn			// 对应的每个玩家的连接
	L     *lua.LState			// 处理该玩家的lua脚本
	luaReloadTime	int			// 记录上次lua脚本更新的时间戳
	userData interface{}
}

func (a *myServer) Run() {
	a.Init()
	//fmt.Println("run")
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("跟对方的连接中断了")
			break
		}
		fmt.Println("收到消息------------", string(data))

		a.WriteMsg([]byte("服务器收到你的消息-------" + string(data)))


		a.CheckLuaReload()		// 检查lua脚本是否需要更新
	}
}

func (a *myServer) OnClose() {
	a.L.Close()		// 关闭lua调用
}

func (a *myServer) WriteMsg(msg ... []byte) {

	err := a.conn.WriteMsg(msg...)
	if err != nil {
		fmt.Printf("发送消息：%v， 出错了！", msg)
	}
	//}
}

func (a *myServer) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *myServer) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *myServer) Close() {
	a.conn.Close()
}

func (a *myServer) Destroy() {
	a.conn.Destroy()
}

func (a *myServer) UserData() interface{} {
	return a.userData
}

func (a *myServer) SetUserData(data interface{}) {
	a.userData = data
}

//--------------------------lua 启动-------------------------------
func (a *myServer) Init()  {
	a.L = Lua.Init(luaCodeToShare) // 绑定lua脚本
	a.luaReloadTime = LuaReloadTime
}

//---------------------------热更新检查-----------------------------
func (a *myServer) CheckLuaReload()  {
	// 检查一下lua更新的时间戳
	if a.luaReloadTime == LuaReloadTime{
		return
	}

	// 如果跟本地的lua时间戳不一致，就更新
	err = Lua.GoCallLuaReload(a.L)
	if err == nil{
		// 热更新成功
		a.luaReloadTime = LuaReloadTime
	}
}
