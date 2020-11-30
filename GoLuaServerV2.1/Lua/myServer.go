package Lua

import (
	"GoLuaServerV2.1/GlobalVar"
	"GoLuaServerV2.1/NetWork"
	"GoLuaServerV2.1/Utils/log"
	"GoLuaServerV2.1/Utils/ztimer"
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// ----------------------------服务器处理的统一接口----------------------------------
// myServer其实是一个个连接单独处理的模块
//-----------------------------------------------------------------------------------

var MyServerUUID int32 = 1		// 自定义玩家连接的临时编号，用来传给lua，这样lua就知道消息给谁返回
var StaticDataPackageHeadLess = 0  // 统计信息，数据包 头部数据不全
var StaticDataPackageProtoDataLess = 0  // 统计信息，数据包 pb数据不全
var StaticDataPackagePasteNum = 0   // 统计信息，拼接次数
var StaticDataPackagePasteSuccess = 0   // 统计信息，成功拼接后，解析成功
var StaticDataPackageHeadFlagError = 0   // 统计信息，数据包头部标识不正确

var StaticNetWorkReceiveToSendCostTime = 0   // 统计信息，接收客户端消息到发送回该消息所消耗的时间
var StaticNetWorkReceiveToSendCostTimeAll = 0   // 统计信息，接收客户端消息到发送回该消息所消耗的时间
var StaticNetWorkReceiveToSendCostTimeNum = 0   // 统计信息，接收客户端消息到发送回该消息所消耗的时间

var LuaConnectMyServer map[int]*MyServer    // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调
var luaUIDConnectMyServer map[int]*MyServer // 将uid跟连接句柄进行哈希


// MyServer是每个客户端的连接
type MyServer struct {
	Conn  NetWork.Conn // 对应的每个玩家的连接
	myLua *MyLua       // 处理该玩家的lua脚本
	//luaReloadTime	int			// 记录上次lua脚本更新的时间戳，后面统一到一个Lstate之后，这个作废了
	ServerId int				// 自己分配的连接编号
	UserId  int					// 玩家uid
	TokenId int					// 玩家发包的id标识
	TokenTime int64				// 接收token的时间


	MyServerMutex sync.Mutex // 连接自己的锁主要用于防止发送时候产生的线程不安全

	ReceiveBuf []byte			// 上次不完整的数据包
	SuccessBuf []byte			// 上次成功的数据包
	LastBuf []byte				// 最后一次接收数据包
	ReceiveMsgNum int		// 接收包数量
	SendMsgNum int			// 发送包的数量


	LuaCallClose bool		// lua申请关闭连接
}

// 获取唯一的ServerId
func GetServerUid() int {


retry:
	MyServerUUID = atomic.AddInt32(&MyServerUUID, 1)

	if MyServerUUID > math.MaxInt32 {
		MyServerUUID = 0		// 如果越界了， 那么重头来过
	}
	ServerId := int(MyServerUUID)
	if  GetMyServerByServerId(int(MyServerUUID)) != nil {
		// 如果被占用了， 那么尝试下一个
		goto retry
		fmt.Printf("serverId  %d 被占用", MyServerUUID)
	}


	fmt.Println("连接创建了ServerId ： ",ServerId)

	return ServerId
}



// 分配一个玩家处理逻辑模块的内存
func NewMyServer(conn NetWork.Conn,ServerId int)  *MyServer{
	//myLua := NewMyLua()
	//myLua:= GameManagerLua		// 改为统一一个LState
	return &MyServer{Conn:conn,myLua:GameManagerLuaHandle,ServerId:ServerId,ReceiveBuf:nil}
}

//--------------------------各个玩家连接逻辑主循环------------------------------
func (a *MyServer) Run() {



	a.Init()
	for {
		//fmt.Println("-------------各个玩家连接逻辑主循环---------")
		if a.LuaCallClose {
			// lua 申请关闭网络连接
			return    // 那么主动关闭吧
		}


		//a.GlobalMutex.Lock()
		buf,bufLen, err := a.Conn.ReadMsg()
		//a.GlobalMutex.Unlock()
		//panic(strconv.Itoa(a.UserId)+"出现严重错误"+strconv.Itoa(a.ServerId))
		//if r := recover(); r != nil {
		//	fmt.Printf("panic的内容%v\n", r)
		//}

		if err != nil {
			log.PrintfLogger("跟对方的连接中断了")
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
			//log.PrintLogger(str)
			//str= fmt.Sprintf("%d本次buf: %x ", a.UserId,buf)
			//log.PrintLogger(str)

			buf2 := make([]byte,len(a.ReceiveBuf)+bufLen)		//缓存从新组合包
			copy(buf2, a.ReceiveBuf)
			copy(buf2[len(a.ReceiveBuf):],buf[:bufLen])
			//str= fmt.Sprintf("%d合并后buf2: %x ", a.UserId,buf2)
			//log.PrintLogger(str)
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
			bufTemp := buf[bufHead:bufLen]         //要处理的buffer
			bufHeadTemp := a.HandlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
			bufHead += bufHeadTemp
			//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
			if bufHeadTemp == 0 {
				a.ReceiveBuf = bufTemp			// 接收不全，那么缓存
				break 	// 数据不全， 继续接收数据
			}else if bufHeadTemp > 0 {				// 解析完成
				if a.ReceiveBuf != nil {			// 如果是拼接包，只要成功解析，就可以清理了
					//str := fmt.Sprintf("%d 拼接后成功解析%x", a.UserId, buf)
					//log.PrintLogger(str)
					a.ReceiveBuf = nil

					//GlobalVar.GlobalMutex.Lock()
					StaticDataPackagePasteSuccess ++
					//GlobalVar.GlobalMutex.Unlock()
				}
			}else if bufHeadTemp == -1 {
				a.ReceiveBuf = nil
				log.PrintfLogger("最后一次成功的buf：%x  bufHeadTemp%d  bufHead %d",a.SuccessBuf , bufHeadTemp, bufHead)
				log.PrintfLogger("最后一次接收的buf：%x  len:%d",a.LastBuf, len(a.LastBuf))
				log.PrintfLogger("最后一次保存的不完整buf：%x",a.ReceiveBuf)
				log.PrintfLogger("当前buf：%x   bufLen %d",buf , bufLen)
				return  		//数据包不正确，放弃连接
			}

			if bufHead >= bufLen{
				a.LastBuf = buf[:bufLen]		//记录上次接收buf
				if bufHead > bufLen{
					log.PrintfLogger(" %d bufHead  %d > bufLen %d  bufHeadTemp %d  buf：%x", a.UserId, bufHead ,  bufHeadTemp ,bufLen,buf[:bufLen])
				}
				break		// 处理完毕，继续接收
			}
			//time.Sleep(time.Millisecond * 50)			//服务器在接收客户端消息的时候， 1秒最多接收20个消息， 防止客户端狂发消息给服务器
			time.Sleep(time.Millisecond * 50)			//服务器在接收客户端消息的时候， 1秒最多接收20个消息， 防止客户端狂发消息给服务器
		}
		//GlobalVar.GlobalMutex.Unlock()

		//a.myLua.GoCallLuaNetWorkReceive(string(data))		// 把收到的数据传递给lua进行处理
		//a.WriteMsg([]byte("服务器收到你的消息-------" + string(data)))


		//a.CheckLuaReload()		// 检查lua脚本是否需要更新
	}
}



func (a * MyServer)HandlerRead(buf []byte) int {
	//fmt.Printf("buf......%x",buf)
	//-----------------------------头部数据不完整----------------------------
	if len(buf)< NetWork.TCPHeaderSize {
		//str:= fmt.Sprintf("%d数据包头部数据不全 : %x \n",a.UserId,buf)
		//log.PrintLogger(str)
		//GlobalVar.GlobalMutex.Lock()
		StaticDataPackageHeadLess ++
		//GlobalVar.GlobalMutex.Unlock()

		return 0
	}
	//-----------------------------解析头部信息----------------------------

	msgId, subMsgId, bufferSize, ver := NetWork.DealReceiveTcpDeaderData(buf)
	BufAllSize := NetWork.TCPHeaderSize + int(bufferSize)

	////-----------------------------头部信息错误----------------------------
	//if headFlag != uint8(254){
	//	log.PrintfLogger("%d 数据包头部标识不正确 %x",a.UserId, buf)
	//
	//	//GlobalVar.GlobalMutex.Lock()
	//	StaticDataPackageHeadFlagError ++
	//	//GlobalVar.GlobalMutex.Unlock()
	//	return -1 			// 数据包格式校验不正确
	//}


	//fmt.Println("len(buf)",len(buf))
	//fmt.Println("offset",offset)
	//fmt.Println("bufferSize",bufferSize)

	//-----------------------------错误提示----------------------------
	//if msgSize >0 {
	//	//fmt.Println("有错误提示了")
	//	//msgBuffer := buf[offset + int(bufferSize):offset + int(bufferSize)+ int(msgSize)]
	//	//fmt.Println(string(msgBuffer))
	//	return BufAllSize
	//}

	//-----------------------------proto buffer 内容不完整----------------------------
	if len(buf) < BufAllSize{
		//str:= fmt.Sprintf("%d数据包格式不正确buflen%d,bufferSize%d,%x  \n", a.UserId,len(buf),int(bufferSize),buf)
		//log.PrintLogger(str)
		//GlobalVar.GlobalMutex.Lock()
		StaticDataPackageProtoDataLess ++
		//GlobalVar.GlobalMutex.Unlock()
		//a.ReceiveBuf = buf			// 接受不全，那么缓存
		return  0 //int(bufferSize) + offset
	}

	//// ------------------------数据包尾部的判断----------------------
	//endData := NetWork.DealRecvTcpEndData(buf[BufAllSize -1 :BufAllSize])
	//if endData!= uint8(NetWork.TCPEnd){		// EE
	//	log.PrintfLogger("%d数据包尾部判断不正确 %x ",a.UserId, buf)
	//	return -1
	//}

	//-----------------------------数据包重复----------------------------
	//if int(ver) == a.TokenId{
	//	log.PrintLogger( strconv.Itoa(a.UserId)+" 出现重复的数据包,包id："+ strconv.Itoa(int(ver)))
	//	//return BufAllSize  // 如果重复，那么跳过解析这个数据包
	//}

	//-----------------------------取出proto buffer的内容----------------------------
	//var finalBuffer []byte

	offset := NetWork.TCPHeaderSize
	token := 0
	if ver > 0{
		offset = 12		// version == 1 的时候， 加了一个token

		buf1 := bytes.NewBuffer( buf[NetWork.TCPHeaderSize:offset])
		binary.Read(buf1,binary.LittleEndian,&token)
	}
	//fmt.Println("token: ",token)
	finalBuffer := buf[offset: NetWork.TCPHeaderSize + int(bufferSize)]
	// 解密
	if ver > 0 {
		//fmt.Printf("buffer: %x\n", finalBuffer)
		//fmt.Println("开始解密")
		finalBuffer = NetWork.Decryp(finalBuffer)
		//fmt.Printf("解密后buffer: %x\n", finalBuffer)
	}

	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")

	a.TokenId = int(ver) // 记录当前最后接收的数据包编号，防止重复
	a.TokenTime = ztimer.GetOsTimeMillisecond()

	QueueAdd(a, a.ServerId, a.UserId, int(msgId), int(subMsgId), string(finalBuffer), int(ver)) // 把收到的数据传递给队列， 后期进行lua进行处理

	a.ReceiveMsgNum++
	a.SuccessBuf = buf 	// 记录最后一次成功的buf

	return BufAllSize

}

// 在网络中断的时候会自动调用， 关闭lua脚本
func (a *MyServer) OnClose() {
	//log.PrintLogger("玩家中断了网络连接， 我们要关闭网络")
	//	a.myLua.L.DoString(`	// 关闭channel
	//	GameManagerReceiveCh:close()
	//    GameManagerSendCh:close()
	//`)
	//	a.myLua.L.Close() // 关闭lua调用



	//if a.UserId > 0 {
		// 连接关闭了， 通知lua， 这个玩家网络中断了
	a.myLua.GoCallLuaLogicInt2("GoCallLuaPlayerNetworkBroken", a.UserId, a.ServerId)
	//}

	// 清理掉一些调用关系
	GlobalVar.RWMutex.Lock()
	delete(LuaConnectMyServer, a.ServerId)
	delete(luaUIDConnectMyServer, a.UserId)
	//LuaConnectMyServer[a.ServerId] = nil
	//luaUIDConnectMyServer[a.UserId] = nil
	GlobalVar.RWMutex.Unlock()

	//runtime.GC()

}

// ---------------------发送数据到网络-------------------------

func (a *MyServer) SendMsg(data string, msg string, mainCmd int, subCmd int ) bool{
	bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0) // token始终是0，服务器不用发token

	//if token!=0 {
	//	fmt.Println("token", token)
	//	fmt.Println("a.TokenId ", a.TokenId)
	//}

	// 计算一下从消息的接收 --  消息的处理  --- 消息的发送  所消耗的时间
	//if token==a.TokenId {
	//
	//	now := ztimer.GetOsTimeMillisecond()
	//	cost := int(now - a.TokenTime)
	//
	//	if cost > GlobalVar.WarningTimeCost {
	//		log.PrintfLogger("UID: %d  处理消息花费时间 %d  mainCmd   %d  subCmd  %d", a.UserId, int(cost), mainCmd, subCmd)
	//	}
	//	if StaticNetWorkReceiveToSendCostTimeAll> 99999999 {
	//		StaticNetWorkReceiveToSendCostTimeAll = 0	// 定期清理，防止数字过大
	//		StaticNetWorkReceiveToSendCostTimeNum = 0
	//	}
	//
	//	StaticNetWorkReceiveToSendCostTimeNum++
	//	StaticNetWorkReceiveToSendCostTimeAll+= cost
	//	StaticNetWorkReceiveToSendCostTime = StaticNetWorkReceiveToSendCostTimeAll/StaticNetWorkReceiveToSendCostTimeNum
	//	//GlobalVar.RWMutex.Lock()
	//	//if StaticNetWorkReceiveToSendCostTime == 0{
	//	//	StaticNetWorkReceiveToSendCostTime = cost
	//	//}else {
	//	//	StaticNetWorkReceiveToSendCostTime = (StaticNetWorkReceiveToSendCostTime+cost)/2	//求平均值
	//	//}
	//	//GlobalVar.RWMutex.Unlock()
	//	//log.PrintfLogger("UID: %d  处理消息花费时间 %d", a.UserId, int(cost))
	//}
	return a.WriteMsg(bufferEnd)
}

func (a *MyServer) WriteMsg(msg ... []byte) bool{
	if a == nil ||  a.Conn == nil{
		log.PrintLogger("当前连接已经关闭, 不发送了")
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

func (a *MyServer) LocalAddr() net.Addr {
	return a.Conn.LocalAddr()
}

func (a *MyServer) RemoteAddr() net.Addr {
	return a.Conn.RemoteAddr()
}

//func (a *MyServer) Close() {
//	a.Conn.Close()
//}
//
//func (a *MyServer) Destroy() {
//	a.Conn.Destroy()
//}
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

	GlobalVar.RWMutex.Lock()
	if LuaConnectMyServer[a.ServerId] != nil {
		log.PrintfLogger("LuaConnectMyServer  已经有了, map重复了", a.ServerId,  a.UserId)
	}
	LuaConnectMyServer[a.ServerId] = a
	GlobalVar.RWMutex.Unlock()

	a.myLua.GoCallLuaNetWorkInit(a.ServerId)
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


// ----------------------------要跟其他服务器做一个连接-------------------------------------------
func ConnectOtherServer(ServerAddressAndPort string) int{
	client := new(NetWork.TCPClient)
	client.Addr = ServerAddressAndPort
	client.ConnNum = 1  //废了
	client.ConnectInterval = 3 * time.Second	// 客户端自动重连
	client.PendingWriteNum = 1000	// 发送缓冲区
	client.LenMsgLen = 4
	client.MaxMsgLen = math.MaxUint32
	client.AutoReconnect = true		// 支持断线重联

	fmt.Println("0")
	//serverId :=  MyServerUUID 		// serverId
	serverId :=  GetServerUid() 		// serverId

	client.NewAgent = func(conn *NetWork.TCPConn,index int) NetWork.Agent {
		b := NewMyServer(conn,serverId)				// 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
		return b
	}

	fmt.Println("开始连接 -- 服务器 -- ", client.Addr, serverId)
	client.Start(serverId, serverId)
	return serverId
}