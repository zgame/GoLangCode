//------------------------------------------------------------------
// game server入口文件
//------------------------------------------------------------------

package main

import (
	"GoLuaServerV2.1/Lua"
	"GoLuaServerV2.1/NetWork"
	"GoLuaServerV2.1/Utils/log"
	"GoLuaServerV2.1/Utils/ip"
	"GoLuaServerV2.1/Utils/ztimer"
	"fmt"
	"github.com/go-ini/ini"
	"math"
	"strconv"
	//"./Games"
	//"./Logic/Player"
	//"./CSV"
	"time"
	//"github.com/yuin/gopher-lua"
	"GoLuaServerV2.1/GlobalVar"
	"flag"
	oldLog "log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)


//func checkPanic(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//var ServerOpen bool		//服务器开启完毕

// ---------------------------程序入口-----------------------------------
var wsServer *NetWork.WSServer
var server *NetWork.TCPServer

var WebSocketServer = true	// websocket 开启
var SocketServer = true		// socket 开启

var WebSocketPort int
var SocketPort int
var GameRoomServerID int	// 游戏原有的ServerID(来源于GameRoomInfo中对应的)
var ServerAddress string    // ServerAddress 服务器地址
//var WebSocketAddress string // WebSocketAddress 服务器地址


//var RedisAddress string		// redis 服务器地址
//var RedisPass string		// redis pwd
var err error

//var MySqlServerIP string		// mySql
//var MySqlServerPort string		// mySql port
//var MySqlDatabase string
//var MySqlUid string
//var MySqlPwd string

//var SqlServerIP string		// sql server
//var SqlServerPort string		// sql  server port
//var SqlServerDatabase string
//var SqlServerUid string
//var SqlServerPwd string



var GameManagerLua *Lua.MyLua    // 公共部分lua脚本
var GameManagerLuaReloadTime int // 公共逻辑处理的lua更新时间



//var GoroutineMax int 			// 给lua的游戏桌子使用的协程数量		暂时没用
//var GoroutineTableLua *Lua.MyLua		// 桌子lua脚本
//var GoroutineTableLuaLuaReloadTime int  // 公共逻辑处理的lua更新时间

func main() {

	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	//远程获取pprof数据打开浏览器http://localhost:8081/debug/pprof/
	go func() {
		oldLog.Println(http.ListenAndServe("localhost:8081", nil))
	}()

	fmt.Println("------------------首先读取命令行参数---------------------------")
	wsPort := flag.Int("WebSocketPort", 0, "")
	sPort := flag.Int("SocketPort", 0, "")
	iGameServerID := flag.Int("ServerID", 0, "")
	flag.Parse()
	WebSocketPort = *wsPort
	SocketPort = *sPort
	GameRoomServerID = *iGameServerID

	fmt.Println("WebSocketPort=",WebSocketPort,"SocketPort=",SocketPort,"GameServerID=",GameRoomServerID)
	if WebSocketPort==0 || SocketPort==0{
		for{
			fmt.Println("缺少命令行参数！ 参数要设置类似 -WebSocketPort=8089 -SocketPort=8123")
			time.Sleep(time.Second)
		}
	}
	log.ServerPort = SocketPort		// 传递给log日志，让日志记录的时候区分服务器端口
	fmt.Println("-------------------读取本地配置文件---------------------------")
	initSetting()

	//fmt.Println("-------------------lua编译---------------------------")
	//GlobalVar.LuaCodeToShare,err = Lua.CompileLua("Script/main.lua")
	//if err!=nil{
	//	fmt.Println("加载main.lua文件出错了！")
	//}

	//fmt.Println("-------------------Redis 数据库连接---------------------------")
	//if redis.InitRedis(RedisAddress,RedisPass) == false{
	//	return
	//}
	//fmt.Println("-------------------MySql 数据库连接---------------------------")
	//if zMySql.ConnectDB(MySqlServerIP, MySqlServerPort , MySqlDatabase,MySqlUid,MySqlPwd) == false{
	//	return
	//}

	//fmt.Println("-------------------读取CVS数据文件---------------------------")
	//CSV.LoadFishServerExcel()

	fmt.Println("-------------------服务器初始化---------------------------")
	initVar()
	Lua.QueueInit()

	//Player.GetALLUserUUID()			// 获取玩家的总体分配UUID

	////-------------------------------------创建各个游戏，以后新增游戏，要在这里增加进去即可-----------------------------------
	//Games.AddGame("满贯捕鱼", Games.GameTypeBY)
	//Client.AddGame("满贯捕鱼2", Client.GameTypeBY2)
	//Client.AddGame("满贯捕鱼3", Client.GameTypeBY3)
	// 后续更多游戏可添加到此处...
	// ...
	// ...
	// ...

	//-------------------------------------创建计时器-----------------------------------
	//ztimer.TimerCheckUpdate(func() {
	//	fmt.Println("我是用来定时检查是否有数据更新的")		// 开启定时器检查数据更新
	//})
	//ztimer.TimerClock12(func() {
	//	fmt.Println("我是用来定时刷新数据的")		// 开启定时器每天夜里刷新
	//})



	fmt.Println("-------------------Lua逻辑处理器---------------------------")
	GameManagerInit()
	fmt.Println("Lua 代码初始化完成")

	//fmt.Println("-------------------启动 Lua 访问 MySql ---------------------------")
	//if GameManagerLua.GoCallLuaConnectMysql(MySqlServerIP, MySqlServerPort , MySqlDatabase,MySqlUid,MySqlPwd) == false{
	//	fmt.Println("lua mySql 数据库没有连接成功")
	//	return
	//}
	//fmt.Println("-------------------启动 Lua 访问 Sql Server---------------------------")
	//if GameManagerLua.GoCallLuaConnectSqlServer(SqlServerIP, "" , SqlServerDatabase, SqlServerUid,SqlServerPwd) == false{
	//	fmt.Println("lua sql  server 数据库没有连接成功")
	//	return
	//}

	//fmt.Println("-------------------多核桌子逻辑处理器---------------------------")
	//CreateGoroutineForLuaGameTable()

	fmt.Println("-------------------启动gameManager---------------------------")
	//if GameManagerLua.GoCallLuaConnectMysql(MySqlServerIP,Database,MySqlUid,MySqlPwd) == false{
	//	fmt.Println("lua mySql 数据库没有连接成功")
	//	return
	//}
	GameManagerLua.GoCallLuaLogic("GoCallLuaStartGamesServers")
	//StartMultiThreadChannelPlayerToGameManager()

	TimerCommonLogicStart()

	fmt.Println("-------------------读取数据库设置---------------------------")
	//UpdateLuaReload()

	fmt.Println("-------------------游戏服务器开始建立连接---------------------------")
	NetWorkServerStart()


	//service := "127.0.0.1:"+strconv.Itoa(SocketPort)
	//listener, err := net.Listen("tcp", service)
	//log.CheckError(err)
	//for {
	//	conn, err := listener.Accept()
	//	if log.CheckError(err){
	//		continue
	//	}
	//
	//	//clientNew := &Client.Client{conn, 0, nil,nil , nil, 0, false, time.Now(), time.Now(),  time.Now(),0 ,0,0}
	//	var clientNew *Client.Client
	//	clientNew = clientNew.NewClient(conn)
	//	Client.AllClientsList[clientNew] = struct{}{} //把新客户端地址增加到哈希表中保存，方便以后遍历
	//
	//	go startClient(clientNew)
	//}

	// ----------------------主循环计时器----------------------------------------
	//tickerCheckUpdateData := time.NewTicker(time.Second * 5)		// 每5秒触发一次计时器，用于定期更新数据库的通用配置信息
	//defer tickerCheckUpdateData.Stop()

	for {
		//select {
		//case <-tickerCheckUpdateData.C:
		//	// 定期更新后台的配置数据
		//	UpdateLuaReload()
		//}
		//start:= ztimer.GetOsTimeMillisecond()
		//GameManagerLua.GoCallLuaLogic("MultiThreadChannelPlayerToGameManager") //公共逻辑处理循环

		//UpdateDBSetting()
		ztimer.CheckRunTimeCost(func() {
				GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 桌子的run
			}, "桌子循环GoCallLuaGoRoutineForLuaGameTable"		)
		//startTime := ztimer.GetOsTimeMillisecond()
		//GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 桌子的run
		//if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		//	log.PrintfLogger("--------!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!![ 警告 ] 桌子循环 消耗时间过长: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
		//}
		runtime.GC()
		time.Sleep(time.Millisecond * 1000)                                //给其他协程让出1秒的时间， 这个可以后期调整
		//end:= ztimer.GetOsTimeMillisecond()
		//if end - start > 120 {
			//fmt.Println("一个循环用时", end-start)
		//}
	}


	// 主逻辑退出时候
//	defer func() {
//		GameManagerLua.L.DoString(`	// 关闭channel
//	GameManagerReceiveCh:close()
//    GameManagerSendCh:close()
//`)
//		GameManagerLua.L.Close()
//	}()

}




// -WebSocketPort=8089 -SocketPort=8124
//-----------------------------本地配置文件---------------------------------------------------
func initSetting()  {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("读取配置文件出错")
		return
	}

	//-------------------------------------------------------------------
	//if WebSocketPort == 0 {
	//	WebSocketPort, err = f.Section("Server").Key("WebSocketPort").Int()
	//}
	//if SocketPort == 0 {
	//	fmt.Println("Warning!!!! You sould write arguments like : -WebSocketPort=8089 -SocketPort=8124")
	//	SocketPort, err = f.Section("Server").Key("SocketPort").Int()
	//}

	log.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()
	//WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	//SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()
	//RedisAddress = f.Section("Server").Key("RedisAddress").String()
	//RedisPass = f.Section("Server").Key("RedisPass").String()
	ServerAddress = f.Section("Server").Key("ServerAddress").String()
	//WebSocketAddress = f.Section("Server").Key("WebSocketAddress").String()
	//GoroutineMax ,err  = f.Section("Server").Key("GoroutineMax").Int()

	//MySqlServerIP = f.Section("Server").Key("MySqlServerIP").Value()
	//MySqlServerPort = f.Section("Server").Key("MySqlServerPort").Value()
	//MySqlDatabase = f.Section("Server").Key("MySqlDatabase").Value()
	//MySqlUid = f.Section("Server").Key("MySqlUid").Value()
	//MySqlPwd = f.Section("Server").Key("MySqlPwd").Value()

	//SqlServerIP = f.Section("Server").Key("SqlServerIP").Value()
	//SqlServerDatabase = f.Section("Server").Key("SqlServerDatabase").Value()
	//SqlServerUid = f.Section("Server").Key("SqlServerUid").Value()
	//SqlServerPwd = f.Section("Server").Key("SqlServerPwd").Value()

	log.CheckError(err)

	ServerAddress = string(ip.GetInternal(0)) // 获取本机内网ip
	fmt.Println("本机内网ip :",ServerAddress)
}

//----------------------------变量的初始化---------------------------------------------------------------
func initVar()  {
	//Client.AllClientsList = make(map[*Client.Client]struct{})
	//Client.AllUserClientList = make(map[uint32]*Client.Client)
	//Games.AllGamesList = make(map[int]*Games.Games)
	Lua.InitGlobalVar()
}


//-------------------------定期更新后台的配置数据---------------------------------------------------------
func UpdateLuaReload() {
	// 以后增加读取数据库的代码
	//...

	// 从服务器更新lua热更新的时间戳
	GlobalVar.LuaReloadTime = 1111
	GameManagerLuaReloadCheck() //共有逻辑检查一下是否需要更新, 玩家部分每个连接自己检查
	//GoroutineTableLuaReloadCheck()
	


}

//-----------------------------------建立服务器的网络功能---------------------------------------------------------------
func NetWorkServerStart()  {



	if WebSocketServer {
		// websocket 服务器开启---------------------------------
		wsServer = new(NetWork.WSServer)
		wsServer.Addr = ServerAddress + ":"+strconv.Itoa(WebSocketPort)
		fmt.Println("websocket 绑定："+ wsServer.Addr)
		wsServer.MaxConnNum = int(math.MaxInt32)
		wsServer.PendingWriteNum = 1000			// 发送区缓存
		wsServer.MaxMsgLen = 4096
		wsServer.HTTPTimeout = 10 * time.Second
		wsServer.CertFile = ""
		wsServer.KeyFile = ""
		wsServer.NewAgent = func(conn *NetWork.WSConn) NetWork.Agent {
			ServerId := Lua.GetServerUid()
			a := Lua.NewMyServer(conn,ServerId)				// 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}

		wsServer.Start()
	}
	if SocketServer{
		// socket 服务器开启----------------------------------
		server = new(NetWork.TCPServer)
		server.Addr = ServerAddress +":"+strconv.Itoa(SocketPort)
		fmt.Println("socket 绑定："+ server.Addr)
		server.MaxConnNum = int(math.MaxInt32)
		server.PendingWriteNum = 1000		// 发送区缓存
		server.LenMsgLen = 4
		server.MaxMsgLen = math.MaxUint32
		server.NewAgent = func(conn *NetWork.TCPConn) NetWork.Agent {
			ServerId := Lua.GetServerUid()
			a := Lua.NewMyServer(conn,ServerId)		// 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
			return a
		}
		server.Start()
	}
}

//-----------------------------------游戏公共逻辑处理-------------------------------------------------------
func GameManagerInit() {

	GameManagerLua = Lua.NewMyLua()

	GameManagerLua.Init() // 绑定lua脚本
	//Lua.GoCallLuaTest(GameManagerLua.L,1)

	Lua.GameManagerLuaHandle = GameManagerLua  // 把句柄传递给lua保存一份
	GameManagerLuaReloadTime = GlobalVar.LuaReloadTime
	GameManagerLua.GoCallLuaSetStringVar("ServerIP_Port", ServerAddress+ ":" + strconv.Itoa(SocketPort)) 	//把服务器地址传递给lua
	GameManagerLua.GoCallLuaSetIntVar("GameRoomServerID", GameRoomServerID) 								//把服务器地址传递给lua
}

// 检查通用逻辑部分的lua是否需要更新
func GameManagerLuaReloadCheck() {
	if GameManagerLuaReloadTime == GlobalVar.LuaReloadTime {
		//return
	}
	// 如果跟本地的lua时间戳不一致，就更新
	err = GameManagerLua.GoCallLuaReload()
	if err == nil{
		// 热更新成功
		GameManagerLuaReloadTime = GlobalVar.LuaReloadTime
	}
}

func TimerCommonLogicStart() {
	// -------------------创建计时器，定期去run公共逻辑---------------------
	ztimer.TimerCheckUpdate(func() {
		connectShow, successSendMsg, successRecMsg, WriteChan := GetAllConnectMsg()
		memoryShow, heapInUse := log.GetSysMemInfo()

		log.PrintfLogger("[%s] 拼接成功%d    标识错误%d   %s   %s  处理消息平均时间：%d  ", ServerAddress+":"+strconv.Itoa(SocketPort),
			Lua.StaticDataPackagePasteSuccess, Lua.StaticDataPackageHeadFlagError, connectShow, memoryShow, Lua.StaticNetWorkReceiveToSendCostTime)

		// 把服务器的状态信息，传递给lua
		GameManagerLua.GoCallLuaSetIntVar("ServerStateSendNum", successSendMsg)
		GameManagerLua.GoCallLuaSetIntVar("ServerStateReceiveNum", successRecMsg)
		GameManagerLua.GoCallLuaSetIntVar("ServerSendWriteChannelNum", WriteChan)
		GameManagerLua.GoCallLuaSetIntVar("ServerDataHeadErrorNum", Lua.StaticDataPackageHeadFlagError)
		GameManagerLua.GoCallLuaSetIntVar("ServerHeapInUse", heapInUse)
		GameManagerLua.GoCallLuaSetIntVar("ServerNetWorkDelay", Lua.StaticNetWorkReceiveToSendCostTime)

		ztimer.CheckRunTimeCost(func() {
			GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogicRun") //公共逻辑处理循环
		}, "GoCallLuaCommonLogicRun")

	}, 10)

	// ---------------------创建计时器，定期去保存服务器状态---------------------
	ztimer.TimerCheckUpdate(func() {
		ztimer.CheckRunTimeCost(func() {
			GameManagerLua.GoCallLuaLogic("GoCallLuaSaveServerState") // 保存服务器的状态
		}, "GoCallLuaSaveServerState")


		runtime.GC()
	}, 60)	// 1 分钟

	// ---------------------创建计时器，定期去更新lua脚本reload---------------------
	ztimer.TimerCheckUpdate(func() {
		// 定期更新后台的配置数据
		ztimer.CheckRunTimeCost(func() {
			UpdateLuaReload()
		}, "UpdateLuaReload")


	},  20 * 1)  // 1 分钟

	//---------------------创建计时器，夜里12点触发---------------------
	ztimer.TimerClock0(func() {
		ztimer.CheckRunTimeCost(func() {
			GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogic12clock") //公共逻辑处理循环
		}, "GoCallLuaCommonLogic12clock")

	})


	//---------------------创建计时器，定期执行所有的接受消息队列---------------------
	ztimer.TimerMillisecondCheckUpdate(func() {
		ztimer.CheckRunTimeCost(func() {
			Lua.QueueRun()
		}, "Lua.QueueRun")
	},  20)  //

}


////----------------------------------------------------桌子逻辑部分-------------------------------------------------------------------
//
//// 检查通用逻辑部分的lua是否需要更新
//func GoroutineTableLuaReloadCheck() {
//	for i:=1;i<= GoroutineMax;i++ {
//		if GoroutineTableLuaLuaReloadTime[i] == GlobalVar.LuaReloadTime {
//			return
//		}
//		// 如果跟本地的lua时间戳不一致，就更新
//		err = GoroutineTableLua[i].GoCallLuaReload()
//		if err == nil {
//			// 热更新成功
//			GoroutineTableLuaLuaReloadTime[i] = GlobalVar.LuaReloadTime
//		}
//	}
//}



// 这是用来统计所有连接数量，及连接包不全的缓存大小
func GetAllConnectMsg() (string,int,int,int)  {
	connNum := 0		//所有包不全缓存大小
	successSendClients := 0
	successRecClients := 0
	successSendMsg := 0
	successRecMsg := 0
	WriteChan := 0
	AllConnect :=0

	GlobalVar.RWMutex.RLock()
	for _,v := range Lua.LuaConnectMyServer{
		if v!=nil {
			AllConnect ++
			connNum += len(v.ReceiveBuf)
			if v.SendMsgNum>0{
				successSendClients++
				successSendMsg += v.SendMsgNum
				v.SendMsgNum = 0
			}
			if v.ReceiveMsgNum>0{
				successRecClients ++
				successRecMsg += v.ReceiveMsgNum
				v.ReceiveMsgNum = 0
			}
			WriteChan += v.Conn.GetWriteChanCap()
		}
	}
	//if AllConnect>0{
	//	WriteChan = WriteChan/AllConnect
	//}
	GlobalVar.RWMutex.RUnlock()
	GameManagerLua.GoCallLuaSetIntVar("ServerSendWriteChannelNum", WriteChan)		// 发送缓冲区大小
	GameManagerLua.GoCallLuaSetIntVar("ServerDataHeadErrorNum", Lua.StaticDataPackageHeadFlagError)  // 把数据头尾错误发送给lua
	str:=fmt.Sprintf(" 发送连接数量 %d  接收连接数量  %d 每秒发送 %d  每秒接收 %d    发送缓存WriteChan %d  消息队列长度 %d ",   successSendClients, successRecClients, successSendMsg , successRecMsg, WriteChan, Lua.QueueGetLen())
	return "所有连接数量："+ strconv.Itoa(AllConnect) + str , successSendMsg , successRecMsg, WriteChan
}
