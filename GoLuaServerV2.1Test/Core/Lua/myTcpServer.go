package Lua

import (
	"GoLuaServerV2.1Test/Core/NetWork"
	"GoLuaServerV2.1Test/Core/Utils/zLog"
	"fmt"
	"math"
	"net"
	"sync"
)

// ----------------------------服务器处理的统一接口----------------------------------
// myServer其实是一个个连接单独处理的模块
//-----------------------------------------------------------------------------------

var MyTcpServerUUID = 0                // 自定义玩家连接的临时编号，用来传给lua，这样lua就知道消息给谁返回
var StaticDataPackageHeadLess = 0      // 统计信息，数据包 头部数据不全
var StaticDataPackageProtoDataLess = 0 // 统计信息，数据包 pb数据不全
var StaticDataPackagePasteNum = 0      // 统计信息，拼接次数
var StaticDataPackagePasteSuccess = 0  // 统计信息，成功拼接后，解析成功
var StaticDataPackageHeadFlagError = 0 // 统计信息，数据包头部标识不正确

var ConnectMyTcpServer map[int]*MyTcpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调
var ConnectMyUdpServer map[int]*MyUdpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调
var ConnectMyTcpServerByUID map[int]*MyTcpServer // 将uid跟连接句柄进行哈希
var ClientStart int


// MyServer是每个客户端的连接
type MyTcpServer struct {
	Conn  NetWork.Conn // 对应的每个玩家的连接
	myLua *MyLua       // 处理该玩家的lua脚本
	//luaReloadTime	int			// 记录上次lua脚本更新的时间戳，后面统一到一个Lstate之后，这个作废了
	ServerId int				// 自己分配的连接编号
	UserId  int					// 玩家uid
	TokenId int					// 玩家发包的id标识
	ChairID int					// 椅子


	MyServerMutex sync.Mutex // 连接自己的锁主要用于防止发送时候产生的线程不安全

	ReceiveBuf []byte			// 上次不完整的数据包
	SuccessBuf []byte			// 上次成功的数据包
	LastBuf []byte				// 最后一次接收数据包
	ReceiveMsgNum int		// 接收包数量
	SendMsgNum int			// 发送包的数量
}





// 分配一个玩家处理逻辑模块的内存
func NewMyTcpServer(conn NetWork.Conn,GameManagerLua *MyLua)  *MyTcpServer {
	//myLua := NewMyLua()
	myLua:= GameManagerLua		// 改为统一一个LState

	GlobalMutex.Lock()

	if MyTcpServerUUID == 0 {
		MyTcpServerUUID = ClientStart
	}

	ServerId := MyTcpServerUUID
	MyTcpServerUUID++
	if MyTcpServerUUID > int(math.MaxInt32) {
		MyTcpServerUUID = 0
	}
	GlobalMutex.Unlock()
	return &MyTcpServer{Conn: conn,myLua:myLua,ServerId:ServerId,ReceiveBuf:nil}
}

//--------------------------各个玩家连接逻辑主循环------------------------------
func (a *MyTcpServer) Run() {
	//fmt.Println("-------------各个玩家连接逻辑主循环---------")


	a.Init()
	for {

		//a.GlobalMutex.Lock()
		buf,bufLen, err := a.Conn.ReadMsg()
		//a.GlobalMutex.Unlock()
		//panic(strconv.Itoa(a.UserId)+"出现严重错误"+strconv.Itoa(a.ServerId))
		//if r := recover(); r != nil {
		//	fmt.Printf("panic的内容%v\n", r)
		//}

		if err != nil {
			//zLog.PrintfLogger("跟对方的连接中断了")
			// 中断网络连接，关闭网络连接，关闭lua
			break
		}
		//fmt.Printf("收到消息------------%x \n", buf)
		bufHead := 0
		//errNum :=0
		//GlobalVar.GlobalMutex.Lock()

		//----------------粘包-------------------------
		if a.ReceiveBuf !=nil {		// 如果上次有没有接收完整的包，那么合并一下
			//str:= fmt.Sprintf("%d上次buf: %x ", a.UserId,a.ReceiveBuf)
			//zLog.PrintLogger(str)
			//str= fmt.Sprintf("%d本次buf: %x ", a.UserId,buf)
			//zLog.PrintLogger(str)

			buf2 := make([]byte,len(a.ReceiveBuf)+bufLen)		//缓存从新组合包
			copy(buf2, a.ReceiveBuf)
			copy(buf2[len(a.ReceiveBuf):],buf[:bufLen])
			//str= fmt.Sprintf("%d合并后buf2: %x ", a.UserId,buf2)
			//zLog.PrintLogger(str)
			//GlobalVar.GlobalMutex.Lock()
			StaticDataPackagePasteNum++
			//GlobalVar.GlobalMutex.Unlock()

			buf = buf2
			bufLen= len(buf2)
		}
		for {
			// ----------------分析和拆包--------------------------------
			//fmt.Println(" buf ",buf)
			//fmt.Println(" bufsize ",bufLen)
			bufTemp := buf[bufHead:bufLen]                                          //要处理的buffer
			bufHeadTemp,msgId,subMsgId,finalBuffer := HandlerRead(bufTemp,a.UserId) //处理结束之后返回，接下来要开始的范围

			//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
			if bufHeadTemp == 0 {
				a.ReceiveBuf = bufTemp			// 接收不全，那么缓存
				break 	// 数据不全， 继续接收数据
			}else if bufHeadTemp > 0 {				// 解析完成
				if a.ReceiveBuf != nil {			// 如果是拼接包，只要成功解析，就可以清理了
					a.ReceiveBuf = nil
					a.myLua.GoCallLuaNetWorkReceive( a.ServerId,  a.UserId,int(msgId),int(subMsgId),string(finalBuffer))		// 把收到的数据传递给lua进行处理
				}
			}else if bufHeadTemp == -1 {
				a.ReceiveBuf = nil
				return  		//数据包不正确，放弃连接
			}
			bufHead += bufHeadTemp
			if bufHead >= bufLen{
				a.LastBuf = buf[:bufLen]		//记录上次接收buf
				if bufHead > bufLen{
					zLog.WritefLogger(" %d bufHead  %d > bufLen %d  bufHeadTemp %d  buf：%x", a.UserId, bufHead ,  bufHeadTemp ,bufLen,buf[:bufLen])
				}
				break		// 处理完毕，继续接收
			}
			//time.Sleep(time.Millisecond * 20)			// 客户端必须疯狂的接收服务器的消息，有多少收多少，防止服务器发太多，堵塞
		}
	}
}



func HandlerRead(buf []byte, uid int) (int,int,int,string) {
	//fmt.Printf("buf......%x",buf)
	//-----------------------------头部数据不完整----------------------------
	if len(buf)< NetWork.TCPHeaderSize {
		//str:= fmt.Sprintf("%d数据包头部数据不全 : %x \n",a.UserId,buf)
		return 0,0,0,""
	}
	//-----------------------------解析头部信息----------------------------

	headFlag,msgId, subMsgId, bufferSize, _ ,msgSize := NetWork.DealRecvTcpHeaderData(buf)
	BufAllSize := NetWork.TCPHeaderSize + int(bufferSize)+ int(msgSize) + 1 // 整个数据包长度，末尾有标示位

	//-----------------------------头部信息错误----------------------------
	if headFlag != uint8(254){
		zLog.WritefLogger("%d 数据包头部标识不正确 %x",uid, buf)
		return -1 ,0,0,""			// 数据包格式校验不正确
	}
	//-----------------------------proto buffer 内容不完整----------------------------
	if len(buf) < BufAllSize{
		//str:= fmt.Sprintf("%d数据包格式不正确buflen%d,bufferSize%d,%x  \n", a.UserId,len(buf),int(bufferSize),buf)
		return  0 ,0,0,""//int(bufferSize) + offset
	}

	// ------------------------数据包尾部的判断----------------------
	endData := NetWork.DealRecvTcpEndData(buf[BufAllSize -1 :BufAllSize])
	if endData!= uint8(NetWork.TCPEnd){ // EE
		zLog.WritefLogger("%d数据包尾部判断不正确 %x ",uid, buf)
		return -1,0,0,""
	}

	//-----------------------------错误提示----------------------------
	if msgSize >0 {
		fmt.Println("有错误提示了")
		msgBuffer := buf[NetWork.TCPHeaderSize+ int(bufferSize): NetWork.TCPHeaderSize+ int(bufferSize)+ int(msgSize)]
		zLog.WritefLogger(string(msgBuffer))
		return BufAllSize,int(msgId),int(subMsgId),""
	}

	//-----------------------------取出proto buffer的内容----------------------------
	finalBuffer := buf[NetWork.TCPHeaderSize : NetWork.TCPHeaderSize+ int(bufferSize)]


	return BufAllSize,int(msgId),int(subMsgId),string(finalBuffer)

}

// 在网络中断的时候会自动调用， 关闭lua脚本
func (a *MyTcpServer) OnClose() {
	//zLog.PrintLogger("玩家中断了网络连接， 我们要关闭网络")
	//	a.myLua.L.DoString(`	// 关闭channel
	//	GameManagerReceiveCh:close()
	//    GameManagerSendCh:close()
	//`)
	//	a.myLua.L.Close() // 关闭lua调用



	if a.UserId > 0 {
		// 连接关闭了， 通知lua， 这个玩家网络中断了
		a.myLua.GoCallLuaLogicInt("Network","Broken", a.UserId)
	}

	// 清理掉一些调用关系
	RWMutex.Lock()
	delete(ConnectMyTcpServer, a.ServerId)
	delete(ConnectMyTcpServerByUID, a.UserId)
	//ConnectMyTcpServer[a.ServerId] = nil
	//ConnectMyTcpServerByUID[a.UserId] = nil
	RWMutex.Unlock()

	//runtime.GC()

}

// ---------------------发送数据到网络-------------------------

func (a *MyTcpServer) SendMsg(data string, msg string, mainCmd int, subCmd int) bool{
	a.TokenId++
	bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, a.TokenId)
	return a.WriteMsg(bufferEnd)
}

func (a *MyTcpServer) WriteMsg(msg ... []byte) bool{
	if a == nil ||  a.Conn == nil{
		zLog.PrintLogger("当前连接已经关闭, 不发送了")
		return false
	}

	//a.MyServerMutex.Lock()
	err := a.Conn.WriteMsg(msg...)
	//a.MyServerMutex.Unlock()

	if err != nil {
		//fmt.Printf("玩家的网络中断，不能正常发送消息给该玩家%x \n",msg)
		return  false   // 发送失败
	}

	a.SendMsgNum++
	return true    // 发送成功
}

func (a *MyTcpServer) LocalAddr() net.Addr {
	return a.Conn.LocalAddr()
}

func (a *MyTcpServer) RemoteAddr() net.Addr {
	return a.Conn.RemoteAddr()
}


//--------------------------lua 启动-------------------------------
func (a *MyTcpServer) Init() {


	//a.myLua.Init() // 绑定lua脚本
	//a.luaReloadTime = GlobalVar.LuaReloadTime

	RWMutex.Lock()
	if ConnectMyTcpServer[a.ServerId] != nil {
		zLog.PrintfLogger("ConnectMyTcpServer  已经有了, map重复了", a.ServerId,  a.UserId)
	}
	ConnectMyTcpServer[a.ServerId] = a
	RWMutex.Unlock()

	// 以后这里可以初始化玩家自己solo的游戏服务器

	a.myLua.GoCallLuaLogicInt("GameClient","Start",a.ServerId)

	// 以后如果有逻辑需要循环， 可以这里加一个协程，做逻辑的run
	//go func() {
	//	// 调用lua的逻辑run
	//}()
}
