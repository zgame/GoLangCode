package Lua

import (
	"fmt"
	"net"
	"../NetWork"
	"../GlobalVar"
	"../Utils/log"
	"time"
	"math"
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
	TokenId int					// 玩家发包的id标识
	ReceiveBuf []byte			// 上次不完整的数据包
	//userData interface{}
}



// 分配一个玩家处理逻辑模块的内存
func NewMyServer(conn NetWork.Conn,GameManagerLua *MyLua)  *MyServer{
	//myLua := NewMyLua()
	myLua:= GameManagerLua		// 改为统一一个LState

	GlobalVar.Mutex.Lock()
	ServerId := MyServerUUID
	MyServerUUID ++
	if MyServerUUID > int(math.MaxInt32) {
		MyServerUUID = 0
	}
	GlobalVar.Mutex.Unlock()
	return &MyServer{conn:conn,myLua:myLua,ServerId:ServerId,ReceiveBuf:nil}
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
		//errNum :=0
		//GlobalVar.Mutex.Lock()
		for {
			if a.ReceiveBuf !=nil {		// 如果上次有没有接收完整的包，那么合并一下
				//str:= fmt.Sprintf("%d上次buf: %x ", this.Index,this.ReceiveBuf)
				//this.Zlog(str)
				//str= fmt.Sprintf("%d本次buf: %x ", this.Index,buf)
				//this.Zlog(str)

				buf2 := make([]byte,len(a.ReceiveBuf)+bufLen)		//缓存从新组合包
				copy(buf2, a.ReceiveBuf)
				copy(buf2[len(a.ReceiveBuf):],buf[:bufLen])
				//str= fmt.Sprintf("%d合并后buf2: %x ", this.Index,buf2)
				//this.Zlog(str)
				buf = buf2
				bufLen= len(buf2)
			}
			//fmt.Println(" buf ",buf)
			//fmt.Println(" bufsize ",bufLen)
			bufTemp := buf[bufHead:bufLen]         //要处理的buffer
			bufHeadTemp := a.HandlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
			bufHead += bufHeadTemp
			time.Sleep(time.Millisecond * 1)
			//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
			if bufHeadTemp == 0 {
					break 	// 数据不全， 继续接收数据
			}else if bufHeadTemp > 0 {				// 解析完成
				if a.ReceiveBuf != nil {			// 如果是拼接包，清理一下
					//str := fmt.Sprintf("%d 拼接后成功解析%x", this.Index, buf)
					//this.Zlog(str)
					a.ReceiveBuf = nil
				}
			}else if bufHeadTemp == -1 {
				a.ReceiveBuf = nil
				break 		//数据包不正确，放弃
			}
			if bufHead >= bufLen{
				break		// 处理完毕，继续接收
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
	//-----------------------------头部数据不完整----------------------------
	if len(buf)< NetWork.TCPHeaderSize {
		//fmt.Printf("数据包头部数据不全 : %x \n",buf)
		a.ReceiveBuf = buf			// 接受不全，那么缓存
		return 0
	}
	//-----------------------------解析头部信息----------------------------
	offset := NetWork.TCPHeaderSize
	headFlag,msgId, subMsgId, bufferSize, tokenId ,msgSize := NetWork.DealRecvTcpDeaderData(buf)

	//-----------------------------头部信息错误----------------------------
	if headFlag != uint8(254){
		str:= fmt.Sprintf("%d数据包头部判断不正确 %x",a.UserId, buf)
		log.PrintLogger(str)
		return -1 			// 数据包格式校验不正确
	}
	//-----------------------------数据包重复----------------------------
	if int(tokenId) > a.TokenId{
		a.TokenId = int(tokenId)		// 记录当前最后接收的数据包编号，防止重复
	}else{
		//log.PrintLogger( strconv.Itoa(a.UserId)+" 出现重复的数据包,包id："+ strconv.Itoa(int(tokenId)))
		return int(bufferSize) + offset + int(msgSize)  // 如果重复，那么跳过解析这个数据包
	}

	//fmt.Println("len(buf)",len(buf))
	//fmt.Println("offset",offset)
	//fmt.Println("bufferSize",bufferSize)

	//-----------------------------错误提示----------------------------
	if msgSize >0 {
		//fmt.Println("有错误提示了")
		//msgBuffer := buf[offset + int(bufferSize):offset + int(bufferSize)+ int(msgSize)]
		//fmt.Println(string(msgBuffer))
		return int(bufferSize) + offset + int(msgSize)
	}

	//-----------------------------proto buffer 内容不完整----------------------------
	if len(buf) < offset + int(bufferSize) + int(msgSize){
		//fmt.Printf("数据包格式不正确buflen%d,bufferSize%d,%x  \n",len(buf),int(bufferSize),buf)
		a.ReceiveBuf = buf			// 接受不全，那么缓存
		return  0 //int(bufferSize) + offset
	}

	//-----------------------------取出proto buffer的内容----------------------------
	finalBuffer := buf[offset:offset + int(bufferSize)]
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")

	a.myLua.GoCallLuaNetWorkReceive( a.ServerId,  a.UserId,int(msgId),int(subMsgId),string(finalBuffer))		// 把收到的数据传递给lua进行处理
	return int(bufferSize) + offset + int(msgSize)

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
		//fmt.Printf("玩家的网络中断，不能正常发送消息给该玩家%x \n",msg)
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

	GlobalVar.Mutex.Lock()
	if luaConnectMyServer[a.ServerId] != nil {
		fmt.Println("luaConnectMyServer  已经有了, map重复了", a.ServerId,  a.UserId)
	}
	luaConnectMyServer[a.ServerId] = a
	GlobalVar.Mutex.Unlock()

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
