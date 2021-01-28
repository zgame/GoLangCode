package Lua

import (
	"GoLuaServerV2.1/Core/NetWork"
	"GoLuaServerV2.1/Core/Utils/zLog"
	"fmt"
	"math"
	"net"
	"sync"
	"time"
)

// ----------------------------服务器处理的统一接口----------------------------------
// TcpServer其实是一个个连接单独处理的模块
//-----------------------------------------------------------------------------------

var ConnectMyTcpServer sync.Map      //[int]*MyTcpServer      // 将lua的句柄跟对应的服务器句柄进行一个哈希，方便以后的lua发送时候回调
var ConnectMyTcpServerByUid sync.Map //[int]*MyTcpServer // 将uid跟连接句柄进行哈希
//var TcpServerUUID int32 = 0            // 自定义玩家连接的临时编号，用来传给lua，这样lua就知道消息给谁返回
//var StaticDataPackageHeadLess = 0      // 统计信息，数据包 头部数据不全
//var StaticDataPackageProtoDataLess = 0 // 统计信息，数据包 pb数据不全
//var StaticDataPackagePasteNum = 0      // 统计信息，拼接次数
//var StaticDataPackagePasteSuccess = 0  // 统计信息，成功拼接后，解析成功
//var StaticDataPackageHeadFlagError = 0 // 统计信息，数据包头部标识不正确
//var StaticNetWorkReceiveToSendCostTime = 0   // 统计信息，接收客户端消息到发送回该消息所消耗的时间
//var StaticNetWorkReceiveToSendCostTimeAll = 0   // 统计信息，接收客户端消息到发送回该消息所消耗的时间
//var StaticNetWorkReceiveToSendCostTimeNum = 0   // 统计信息，接收客户端消息到发送回该消息所消耗的时间



// MyServer是每个客户端的连接
type MyTcpServer struct {
	Conn  NetWork.Conn // 对应的每个玩家的连接
	myLua *MyLua       // 处理该玩家的lua脚本
	//luaReloadTime	int			// 记录上次lua脚本更新的时间戳，后面统一到一个Lstate之后，这个作废了
	ServerId int				// 自己分配的连接编号
	UserId  int					// 玩家uid
	TokenId int					// 玩家发包的id标识
	TokenTime int64				// 接收token的时间

	ReceiveBuf []byte			// 上次不完整的数据包
	SuccessBuf []byte			// 上次成功的数据包
	LastBuf []byte				// 最后一次接收数据包
	ReceiveMsgNum int		// 接收包数量
	SendMsgNum int			// 发送包的数量

	LuaCallClose bool		// lua申请关闭连接
}


// 通过lua堆栈找到对应的是哪个myServer
func GetMyServerByServerId(serverId int) *MyTcpServer {
	if re,ok := ConnectMyTcpServer.Load(serverId); ok{
		return re.(*MyTcpServer)
	}
	return nil
}
// 通过 user id 找到对应的是哪个myServer
func GetMyServerByUID(uid int) *MyTcpServer {
	if re,ok:= ConnectMyTcpServerByUid.Load(uid);ok{
		return re.(*MyTcpServer)
	}
	return nil
}


// 获取唯一的ServerId
func GetServerUid() int {
	for i := 1; i < math.MaxInt32; i++ {
		if  GetMyServerByServerId(i) == nil{
			return i
		}
	}
	return 0
//
//retry:
//	TcpServerUUID = atomic.AddInt32(&TcpServerUUID, 1)
//
//	if TcpServerUUID > math.MaxInt32/2 {
//		TcpServerUUID = 1 // 如果越界了， 那么重头来过
//	}
//	ServerId := int(TcpServerUUID)
//	if  GetMyServerByServerId(ServerId) != nil {
//		// 如果被占用了， 那么尝试下一个
//		goto retry
//		fmt.Printf("serverId  %d 被占用", TcpServerUUID)
//	}
//	//fmt.Println("连接创建了ServerId ： ",ServerId)
//
//	return ServerId
}

// 分配一个玩家处理逻辑模块的内存
func NewMyServer(conn NetWork.Conn,ServerId int)  *MyTcpServer {
	return &MyTcpServer{Conn: conn,myLua: GameManagerLuaHandle,ServerId:ServerId,ReceiveBuf:nil}
}

//--------------------------各个玩家连接逻辑主循环------------------------------
func (a *MyTcpServer) Run() {

	a.Init()
	for {
		//fmt.Println("-------------各个玩家连接逻辑主循环---------")
		if a.LuaCallClose {
			// lua 申请关闭网络连接
			return    // 那么主动关闭吧
		}

		buf,bufLen, err := a.Conn.ReadMsg()

		if err != nil {
			//zLog.PrintfLogger("跟对方的连接中断了")
			// 中断网络连接，关闭网络连接，关闭lua
			break
		}
		//fmt.Printf("收到消息------------%x \n", buf)
		bufStartP := 0
		//----------------粘包-------------------------
		if a.ReceiveBuf !=nil {		// 如果上次有没有接收完整的包，那么合并一下
			bufAddOld := make([]byte,len(a.ReceiveBuf)+bufLen) //缓存从新组合包
			copy(bufAddOld, a.ReceiveBuf)                      // copy上次缓存的数据
			copy(bufAddOld[len(a.ReceiveBuf):],buf[:bufLen])   // copy本次收到的
			buf = bufAddOld
			bufLen= len(bufAddOld)
		}
		for {
			// ----------------分析和拆包--------------------------------
			bufPackage := buf[bufStartP:bufLen]                                                  //要处理的buffer
			bufPackageSize,msgId, subMsgId, finalBuffer := ReadDataPackage(bufPackage, a.UserId) //处理结束之后返回，接下来要开始的范围

			if bufPackageSize == 0 {
				a.ReceiveBuf = bufPackage // 接收不全，那么缓存
				break                     // 数据不全， 继续接收数据
			}else if bufPackageSize > 0 { // 解析完成
				if a.ReceiveBuf != nil {			// 如果是拼接包，只要成功解析，就可以清理了
					a.ReceiveBuf = nil
				}
				QueueAddTcp( a.ServerId,a.UserId, msgId, subMsgId, finalBuffer, 0) // 把收到的数据传递给队列， 后期进行lua进行处理
			}else if bufPackageSize == -1 {
				a.ReceiveBuf = nil
				return  		//数据包不正确，放弃连接
			}

			bufStartP += bufPackageSize
			if bufStartP >= bufLen{
				//a.LastBuf = buf[:bufLen]		//记录上次接收buf
				if bufStartP > bufLen{
					zLog.PrintfLogger(" %d bufStartP  %d > bufLen %d  bufPackageSize %d  buf：%x", a.UserId, bufStartP, bufPackageSize,bufLen,buf[:bufLen])
				}
				break		// 处理完毕，继续接收
			}
			time.Sleep(time.Millisecond * 50)			//服务器在接收客户端消息的时候， 1秒最多接收20个消息， 防止客户端狂发消息给服务器
		}
	}
}





// 在网络中断的时候会自动调用， 关闭lua脚本
func (a *MyTcpServer) OnClose() {
	// 连接关闭了， 通知lua， 这个玩家网络中断了
	a.myLua.GoCallLuaLogicInt2("ServerNetwork","PlayerNetworkBroken", a.UserId, a.ServerId)

	// 清理掉一些调用关系
	ConnectMyTcpServer.Delete(a.ServerId)
	ConnectMyTcpServerByUid.Delete(a.UserId)
}

// ---------------------发送数据到网络-------------------------

func (a *MyTcpServer) SendMsg(data string, msg string, mainCmd int, subCmd int ) bool{
	bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0) // token始终是0，服务器不用发token
	return a.WriteMsg(bufferEnd)
}

func (a *MyTcpServer) WriteMsg(msg ... []byte) bool{
	if a == nil ||  a.Conn == nil{
		zLog.PrintLogger("当前连接已经关闭, 不发送了")
		return false
	}
	err := a.Conn.WriteMsg(msg...)

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

	if _,ok:= ConnectMyTcpServer.Load(a.ServerId) ;ok {
		zLog.PrintfLogger("ConnectMyTcpServer  已经有了, map重复了", a.ServerId,  a.UserId)
	}
	ConnectMyTcpServer.Store(a.ServerId,a)

	a.myLua.GoCallLuaNetWorkInit(a.ServerId)

}


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

	serverId :=  GetServerUid() // serverId

	client.NewAgent = func(conn *NetWork.TCPConn,index int) NetWork.Agent {
		b := NewMyServer(conn,serverId) // 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
		return b
	}

	fmt.Println("开始连接 -- 服务器 -- ", client.Addr, serverId)
	client.Start(serverId, serverId)
	return serverId
}