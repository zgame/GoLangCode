//------------------------------------------------------------------
// game server入口文件
//------------------------------------------------------------------

package main

import (
	"ServerZgameTB/Utils/log"
	"ServerZgameTB/Utils/zMySql"
	"ServerZgameTB/Utils/zRedis"
	"fmt"
	"github.com/go-ini/ini"
	"strconv"
	"ServerZgameTB/Lua"
	"ServerZgameTB/NetWork"
	"ServerZgameTB/Utils/ztimer"
	"math"
	//"./Games"
	//"./Logic/Player"
	//"./CSV"
	"time"
	//"github.com/yuin/gopher-lua"
	"ServerZgameTB/GlobalVar"
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


// ---------------------------程序入口-----------------------------------
var wsServer *NetWork.WSServer
var server *NetWork.TCPServer

var WebSocketServer bool	// websocket 开启
var SocketServer bool		// socket 开启

var WebSocketPort int
var SocketPort int
var ServerAddress string    // ServerAddress 服务器地址
//var WebSocketAddress string // WebSocketAddress 服务器地址


var RedisAddress string		// redis 服务器地址
var RedisPass string		// redis pwd
var err error

var MySqlServerIP string		// mysql
var MySqlServerPort string		// mysql port
var MySqlDatabase string
var MySqlUid string
var MySqlPwd string



var GameManagerLua *Lua.MyLua    // 公共部分lua脚本
var GameManagerLuaReloadTime int // 公共逻辑处理的lua更新时间



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
	flag.Parse()
	WebSocketPort = *wsPort
	SocketPort = *sPort
	fmt.Println("WebSocketPort=",WebSocketPort,"SocketPort=",SocketPort)
	if WebSocketPort==0 || SocketPort==0{
		for{
			fmt.Println("缺少命令行参数！ 参数要设置类似 -WebSocketPort=8089 -SocketPort=8123")
			time.Sleep(time.Second)
		}
	}
	log.ServerPort = SocketPort		// 传递给log日志，让日志记录的时候区分服务器端口
	fmt.Println("-------------------读取本地配置文件---------------------------")
	initSetting()

	fmt.Println("------------------------检查日志目录---------------------------------")
	log.CheckLogDir()

	//fmt.Println("-------------------lua编译---------------------------")
	//GlobalVar.LuaCodeToShare,err = Lua.CompileLua("Script/main.lua")
	//if err!=nil{
	//	fmt.Println("加载main.lua文件出错了！")
	//}

	fmt.Println("-------------------Redis 数据库连接---------------------------")
	if zRedis.InitRedis(RedisAddress,RedisPass) == false{
		return
	}
	fmt.Println("-------------------MySql 数据库连接---------------------------")
	if zMySql.ConnectDB(MySqlServerIP, MySqlServerPort , MySqlDatabase,MySqlUid,MySqlPwd) == false{
		return
	}


	fmt.Println("-------------------服务器初始化---------------------------")
	initVar()


	//Player.GetALLUserUUID()			// 获取玩家的总体分配UUID

	////-------------------------------------创建各个游戏，以后新增游戏，要在这里增加进去即可-----------------------------------
	//Games.AddGame("满贯捕鱼", Games.GameTypeBY)
	//Client.AddGame("满贯捕鱼2", Client.GameTypeBY2)
	//Client.AddGame("满贯捕鱼3", Client.GameTypeBY3)
	// 后续更多游戏可添加到此处...
	// ...
	// ...
	// ...





	fmt.Println("-------------------Lua逻辑处理器---------------------------")
	GameManagerInit()
	fmt.Println("Lua 代码初始化完成")


	fmt.Println("-------------------启动 Lua 访问 MySql ---------------------------")
	if GameManagerLua.GoCallLuaConnectMysql(MySqlServerIP, MySqlServerPort , MySqlDatabase,MySqlUid,MySqlPwd) == false{
		fmt.Println("lua mysql 数据库没有连接成功")
		return
	}

	fmt.Println("-------------------启动gameManager---------------------------")
	GameManagerLua.GoCallLuaLogic("GoCallLuaStartGamesServers")

	// 服务器状态的记录
	TimerCommonLogicStart()

	fmt.Println("-------------------读取数据库设置---------------------------")
	//UpdateLuaReload()

	fmt.Println("-------------------游戏服务器开始建立连接---------------------------")
	NetWorkServerStart()



	for {

		ztimer.CheckRunTimeCost(func() {
				GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 桌子的run
			}, "桌子循环GoCallLuaGoRoutineForLuaGameTable"		)


		runtime.GC()
		time.Sleep(time.Millisecond * 1000)                                //给其他协程让出1秒的时间， 这个可以后期调整

	}


}




// -WebSocketPort=8089 -SocketPort=8124
//-----------------------------本地配置文件---------------------------------------------------
func initSetting()  {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("读取配置文件出错")
		return
	}

	log.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()
	WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()
	RedisAddress = f.Section("Server").Key("RedisAddress").String()
	RedisPass = f.Section("Server").Key("RedisPass").String()
	ServerAddress = f.Section("Server").Key("ServerAddress").String()
	//WebSocketAddress = f.Section("Server").Key("WebSocketAddress").String()
	//GoroutineMax ,err  = f.Section("Server").Key("GoroutineMax").Int()

	MySqlServerIP = f.Section("Server").Key("MySqlServerIP").Value()
	MySqlServerPort = f.Section("Server").Key("MySqlServerPort").Value()
	MySqlDatabase = f.Section("Server").Key("MySqlDatabase").Value()
	MySqlUid = f.Section("Server").Key("uid").Value()
	MySqlPwd = f.Section("Server").Key("pwd").Value()

	log.CheckError(err)
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
			a := Lua.NewMyServer(conn,GameManagerLua)				// 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
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
			a := Lua.NewMyServer(conn,GameManagerLua)		// 每个新连接进来的时候创建一个对应的网络处理的MyServer对象
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
	GameManagerLua.GoCallLuaSetStringVar("ServerIP_Port", ServerAddress+ ":" + strconv.Itoa(SocketPort)) //把服务器地址传递给lua

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

	}, 60 * 1)   // 30秒

	// ---------------------创建计时器，定期去保存服务器状态---------------------
	ztimer.TimerCheckUpdate(func() {
		ztimer.CheckRunTimeCost(func() {
			GameManagerLua.GoCallLuaLogic("GoCallLuaSaveServerState") // 保存服务器的状态
		}, "GoCallLuaSaveServerState")


		runtime.GC()
	}, 60 * 1)	// 1 分钟

	// ---------------------创建计时器，定期去更新lua脚本reload---------------------
	ztimer.TimerCheckUpdate(func() {
		// 定期更新后台的配置数据
		ztimer.CheckRunTimeCost(func() {
			UpdateLuaReload()
		}, "UpdateLuaReload")
	},  10 * 1)  // 1 分钟


	//---------------------创建计时器，夜里12点触发---------------------lua也可以创建固定时间计时器
	ztimer.TimerClock0(func() {
		ztimer.CheckRunTimeCost(func() {
			GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogic12clock") //公共逻辑处理循环
		}, "GoCallLuaCommonLogic12clock")

	})
}



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
	str:=fmt.Sprintf(" 发送连接数量 %d  接收连接数量  %d 每秒发送 %d  每秒接收 %d    发送缓存WriteChan %d",   successSendClients, successRecClients, successSendMsg , successRecMsg, WriteChan)
	return "所有连接数量："+ strconv.Itoa(AllConnect) + str , successSendMsg , successRecMsg, WriteChan
}
