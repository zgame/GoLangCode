package Lua

import (
	"fmt"
	"net"
	"../NetWork"
	"../GlobalVar"

)

// ----------------------------服务器处理的统一接口----------------------------------
// myServer其实是一个个连接单独处理的模块
//-----------------------------------------------------------------------------------



type MyServer struct {
	conn NetWork.Conn			// 对应的每个玩家的连接
	myLua     *MyLua			// 处理该玩家的lua脚本
	luaReloadTime	int			// 记录上次lua脚本更新的时间戳

	//userData interface{}
}



// 分配一个玩家处理逻辑模块的内存
func NewMyServer(conn NetWork.Conn)  *MyServer{
	myLua := NewMyLua()
	return &MyServer{conn:conn,myLua:myLua}
}

//--------------------------各个玩家连接逻辑主循环------------------------------
func (a *MyServer) Run() {
	a.Init()
	//fmt.Println("----logic start---")
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("跟对方的连接中断了")
			break
		}
		fmt.Println("收到消息------------", string(data))
		a.myLua.GoCallLuaNetWorkReceive(string(data))		// 把收到的数据传递给lua进行处理

		a.WriteMsg([]byte("服务器收到你的消息-------" + string(data)))


		a.CheckLuaReload()		// 检查lua脚本是否需要更新
	}
}

func (a *MyServer) OnClose() {
	a.myLua.L.Close()		// 关闭lua调用
}

// ---------------------发送数据到网络-------------------------
func (a *MyServer) WriteMsg(msg ... []byte) {

	err := a.conn.WriteMsg(msg...)
	if err != nil {
		fmt.Printf("发送消息：%v， 出错了！", msg)
	}
	//}
}

func (a *MyServer) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *MyServer) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *MyServer) Close() {
	a.conn.Close()
}

func (a *MyServer) Destroy() {
	a.conn.Destroy()
}
//
//func (a *MyServer) UserData() interface{} {
//	return a.userData
//}
//
//func (a *MyServer) SetUserData(data interface{}) {
//	a.userData = data
//}

//--------------------------lua 启动-------------------------------
func (a *MyServer) Init()  {
	a.myLua.Init() // 绑定lua脚本
	a.luaReloadTime = GlobalVar.LuaReloadTime
	LuaConnectMyServer[a.myLua.L] = a
}

//---------------------------热更新检查-----------------------------
func (a *MyServer) CheckLuaReload()  {
	// 检查一下lua更新的时间戳
	if a.luaReloadTime == GlobalVar.LuaReloadTime{
		return
	}

	// 如果跟本地的lua时间戳不一致，就更新
	err := a.myLua.GoCallLuaReload()
	if err == nil{
		// 热更新成功
		a.luaReloadTime = GlobalVar.LuaReloadTime
	}
}
