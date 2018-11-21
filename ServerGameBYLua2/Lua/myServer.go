package Lua

import (
	"fmt"
	"net"
	"../NetWork"
	"../GlobalVar"

	"time"
)

// ----------------------------服务器处理的统一接口----------------------------------
// myServer其实是一个个连接单独处理的模块
//-----------------------------------------------------------------------------------

var MyServerUUID = 0


type MyServer struct {
	conn NetWork.Conn			// 对应的每个玩家的连接
	myLua     *MyLua			// 处理该玩家的lua脚本
	luaReloadTime	int			// 记录上次lua脚本更新的时间戳
	ServerId int
	UserId  int					// 玩家uid
	//userData interface{}
}



// 分配一个玩家处理逻辑模块的内存
func NewMyServer(conn NetWork.Conn,GameManagerLua *MyLua)  *MyServer{
	//myLua := NewMyLua()
	myLua:= GameManagerLua		// 改为统一一个LState

	GlobalVar.Mutex.Lock()
	ServerId := MyServerUUID
	MyServerUUID ++
	if MyServerUUID >= 9990000{
		MyServerUUID = 0
	}
	GlobalVar.Mutex.Unlock()
	return &MyServer{conn:conn,myLua:myLua,ServerId:ServerId}
}

//--------------------------各个玩家连接逻辑主循环------------------------------
func (a *MyServer) Run() {
	//fmt.Println("-------------各个玩家连接逻辑主循环---------")
	a.Init()
	for {
		buf,bufLen, err := a.conn.ReadMsg()
		if err != nil {
			//fmt.Println("跟对方的连接中断了")
			// 中断网络连接，关闭网络连接，关闭lua
			break
		}
		//fmt.Printf("收到消息------------%x \n", buf)
		bufHead := 0
		//num:=0

		//GlobalVar.Mutex.Lock()
		for {
			//fmt.Println(" buf ",buf)
			//fmt.Println(" bufsize ",bufLen)
			bufTemp := buf[bufHead:bufLen]         //要处理的buffer
			bufHead += a.HandlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
			time.Sleep(time.Millisecond * 100)
			//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
			//num++
			//fmt.Println("num",num)
			if bufHead >= bufLen{
				break
			}
		}
		//GlobalVar.Mutex.Unlock()

		//a.myLua.GoCallLuaNetWorkReceive(string(data))		// 把收到的数据传递给lua进行处理
		//a.WriteMsg([]byte("服务器收到你的消息-------" + string(data)))


		//a.CheckLuaReload()		// 检查lua脚本是否需要更新
	}
}

func (a * MyServer)HandlerRead(buf []byte) int {
	//fmt.Printf("buf......%x",buf)
	if len(buf)< 10 {
		fmt.Printf("error buf len < 10 : %x",buf)
		return 0
	}


	msgId, subMsgId, bufferSize, _ := NetWork.DealRecvTcpDeaderData(buf)

	offset := 10

	//fmt.Println("len(buf)",len(buf))
	//fmt.Println("offset",offset)
	//fmt.Println("bufferSize",bufferSize)

	if len(buf) < offset + int(bufferSize){
		fmt.Println("出现数据包异常")
		return  int(bufferSize) + offset
	}

	//fmt.Println("")
	//if ver > 0{
	//	offset = 12		// version == 1 的时候， 加了一个token
	//}
	finalBuffer := buf[offset:offset + int(bufferSize)]
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")


	a.myLua.GoCallLuaNetWorkReceive( a.ServerId,  a.UserId,int(msgId),int(subMsgId),string(finalBuffer))		// 把收到的数据传递给lua进行处理



	return int(bufferSize) + offset

}

// 在网络中断的时候会自动调用， 关闭lua脚本
func (a *MyServer) OnClose() {
	//log.PrintLogger("玩家中断了网络连接， 我们要关闭网络， 同时关闭玩家的lua文件")

//	a.myLua.L.DoString(`	// 关闭channel
//	GameManagerReceiveCh:close()
//    GameManagerSendCh:close()
//`)
	a.myLua.L.Close() // 关闭lua调用
}

// ---------------------发送数据到网络-------------------------
func (a *MyServer) WriteMsg(msg ... []byte) bool{

	//GlobalVar.Mutex.Lock()
	err := a.conn.WriteMsg(msg...)
	//GlobalVar.Mutex.Unlock()
	if err != nil {
		fmt.Printf("玩家的网络中断，不能正常发送消息给该玩家%x \n",msg)
		return  false   // 发送失败
	}
	return true    // 发送成功
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
func (a *MyServer) Init() {


	//a.myLua.Init() // 绑定lua脚本
	//a.luaReloadTime = GlobalVar.LuaReloadTime

	//GlobalVar.Mutex.Lock()
	luaConnectMyServer[a.ServerId] = a
	//GlobalVar.Mutex.Unlock()

	// 以后这里可以初始化玩家自己solo的游戏服务器


	// 以后如果有逻辑需要循环， 可以这里加一个协程，做逻辑的run
	//go func() {
	//	// 调用lua的逻辑run
	//}()
}

////---------------------------热更新检查-----------------------------
//func (a *MyServer) CheckLuaReload() {
//	// 检查一下lua更新的时间戳
//	if a.luaReloadTime == GlobalVar.LuaReloadTime{
//		return
//	}
//
//	// 如果跟本地的lua时间戳不一致，就更新
//	err := a.myLua.GoCallLuaReload()
//	if err == nil{
//		// 热更新成功
//		a.luaReloadTime = GlobalVar.LuaReloadTime
//	}
//}
