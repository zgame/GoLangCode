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
		buf,bufLen, err := a.conn.ReadMsg()
		if err != nil {
			fmt.Println("跟对方的连接中断了")
			break
		}
		//fmt.Println("收到消息------------", string(buf))
		bufHead := 0
		num:=0
		for {
			//fmt.Println(" buf ",buf)
			//fmt.Println(" bufsize ",bufLen)
			bufTemp := buf[bufHead:bufLen]         //要处理的buffer
			bufHead += a.HandlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
			//time.Sleep(time.Millisecond * 100)
			//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
			num++
			//fmt.Println("num",num)
			if bufHead >= bufLen{
				break
			}
		}


		//a.myLua.GoCallLuaNetWorkReceive(string(data))		// 把收到的数据传递给lua进行处理
		//a.WriteMsg([]byte("服务器收到你的消息-------" + string(data)))


		a.CheckLuaReload()		// 检查lua脚本是否需要更新
	}
}

func (a * MyServer)HandlerRead(buf []byte) int {
	msgId, subMsgId, bufferSize, ver := NetWork.DealRecvTcpDeaderData(buf)

	offset := 10

	//fmt.Println("len(buf)",len(buf))
	//fmt.Println("offset",offset)
	//fmt.Println("bufferSize",bufferSize)

	if len(buf) < offset + int(bufferSize){
		fmt.Println("出现数据包异常")
		return  int(bufferSize) + offset
	}

	if ver > 0{
		offset = 12		// version == 1 的时候， 加了一个token
	}
	finalBuffer := buf[offset:offset + int(bufferSize)]
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")
	a.myLua.GoCallLuaNetWorkReceive(int(msgId),int(subMsgId),string(finalBuffer))		// 把收到的数据传递给lua进行处理


	return int(bufferSize) + offset

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
