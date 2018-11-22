//------------------------------------------------------------------
// game server入口文件
//------------------------------------------------------------------

package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"strconv"
	"./Utils/log"
	"./Utils/zRedis"
	//"./Games"
	//"./Logic/Player"
	//"./CSV"
	"time"
	"math"
	"./NetWork"
	"./Lua"
	"./Utils/ztimer"
	//"github.com/yuin/gopher-lua"
	"./GlobalVar"
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

var WebSocketServer bool	// websocket 开启
var SocketServer bool		// socket 开启
var WebSockewtPort int
var SockewtPort int
var SocketAddress string		// SocketAddress 服务器地址
var WebSocketAddress string		// WebSocketAddress 服务器地址


var RedisAddress string		// redis 服务器地址
var err error


var GameManagerLua *Lua.MyLua    // 公共部分lua脚本
var GameManagerLuaReloadTime int // 公共逻辑处理的lua更新时间



//var GoroutineMax int 			// 给lua的游戏桌子使用的协程数量		暂时没用
//var GoroutineTableLua *Lua.MyLua		// 桌子lua脚本
//var GoroutineTableLuaLuaReloadTime int  // 公共逻辑处理的lua更新时间

func main() {

	runtime.GOMAXPROCS(4)

	fmt.Println("-------------------读取本地配置文件---------------------------")
	initSetting()

	fmt.Println("-------------------lua编译---------------------------")
	GlobalVar.LuaCodeToShare,err = Lua.CompileLua("Script/main.lua")
	if err!=nil{
		fmt.Println("加载main.lua文件出错了！")
	}

	fmt.Println("-------------------数据库连接---------------------------")
	zRedis.InitRedis(RedisAddress)

	//fmt.Println("-------------------读取CVS数据文件---------------------------")
	//CSV.LoadFishServerExcel()

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

	//-------------------------------------创建计时器-----------------------------------
	//ztimer.TimerCheckUpdate(func() {
	//	fmt.Println("我是用来定时检查是否有数据更新的")		// 开启定时器检查数据更新
	//})
	//ztimer.TimerClock12(func() {
	//	fmt.Println("我是用来定时刷新数据的")		// 开启定时器每天夜里刷新
	//})



	fmt.Println("-------------------游戏公共逻辑处理器---------------------------")
	GameManagerInit()
	CommonLogicStart()

	//fmt.Println("-------------------多核桌子逻辑处理器---------------------------")
	//CreateGoroutineForLuaGameTable()

	fmt.Println("-------------------启动gameManager---------------------------")
	GameManagerLua.GoCallLuaLogic("GoCallLuaStartGamesServers")
	//StartMultiThreadChannelPlayerToGameManager()

	fmt.Println("-------------------读取数据库设置---------------------------")
	UpdateDBSetting()

	fmt.Println("-------------------游戏服务器开始建立连接---------------------------")
	NetWorkServerStart()


	//service := "127.0.0.1:"+strconv.Itoa(SockewtPort)
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
		//	UpdateDBSetting()
		//}
		start:= ztimer.GetOsTimeMillisecond()
		//GameManagerLua.GoCallLuaLogic("MultiThreadChannelPlayerToGameManager") //公共逻辑处理循环
		GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 给lua的桌子用的 n个协程函数
		time.Sleep(time.Millisecond * 1000)                                //给其他协程让出1秒的时间， 这个可以后期调整
		end:= ztimer.GetOsTimeMillisecond()
		if end - start > 120 {
			//fmt.Println("一个循环用时", end-start)
		}
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





//-----------------------------本地配置文件---------------------------------------------------
func initSetting()  {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("读取配置文件出错")
		return
	}
	WebSockewtPort ,err  = f.Section("Server").Key("WebSocketPort").Int()
	SockewtPort ,err   = f.Section("Server").Key("SocketPort").Int()
	log.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()
	WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()
	RedisAddress = f.Section("Server").Key("SocketServer").String()
	SocketAddress = f.Section("Server").Key("SocketAddress").String()
	WebSocketAddress = f.Section("Server").Key("WebSocketAddress").String()
	//GoroutineMax ,err  = f.Section("Server").Key("GoroutineMax").Int()
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
func UpdateDBSetting() {
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
		wsServer.Addr = WebSocketAddress + ":"+strconv.Itoa(WebSockewtPort)
		//fmt.Println("websocket 绑定："+ wsServer.Addr)
		wsServer.MaxConnNum = 2000
		wsServer.PendingWriteNum = 100
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
		server.Addr = SocketAddress +":"+strconv.Itoa(SockewtPort)
		//fmt.Println("socket 绑定："+ server.Addr)
		server.MaxConnNum = int(math.MaxInt32)
		server.PendingWriteNum = 100
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
	GameManagerLuaReloadTime = GlobalVar.LuaReloadTime


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

func CommonLogicStart() {
	// 创建计时器，定期去run公共逻辑
	ztimer.TimerCheckUpdate(func() {
		GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogicRun") //公共逻辑处理循环
	}, 5)

	// 创建计时器，定期去run公共逻辑
	ztimer.TimerCheckUpdate(func() {
		// 定期更新后台的配置数据
		UpdateDBSetting()
	}, 10)  // 60秒

	ztimer.TimerClock12(func() { // 创建计时器，夜里12点触发
		GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogic12clock") //公共逻辑处理循环
	})

}
//func StartMultiThreadChannelPlayerToGameManager()  {
//
//	//开始开启监听线程，监听玩家消息，进行线程间通信
//	go func() {
//		for {
//			GameManagerLua.GoCallLuaLogic("MultiThreadChannelPlayerToGameManager") //公共逻辑处理循环
//		}
//	}()
//
//}

////----------------------------------------------------桌子逻辑部分-------------------------------------------------------------------
//
//// 一次性创建好多个协程给lua的游戏使用，GoroutineMax的数量跟cpu有几个核数量一样效率比较高
//func CreateGoroutineForLuaGameTable() {
//
//	GoroutineTableLua = make([]*Lua.MyLua,GoroutineMax+1)
//	GoroutineTableLuaLuaReloadTime = make([]int,GoroutineMax+1)
//
//	for i:=1;i<= GoroutineMax;i++{
//		GoroutineTableLua[i] = Lua.NewMyLua()
//		GoroutineTableLua[i].Init() // 绑定lua脚本
//		//Lua.GoCallLuaTest(GameManagerLua.L,1)
//		GoroutineTableLuaLuaReloadTime[i] = GlobalVar.LuaReloadTime
//
//		GoroutineTableLua[i].GoCallLuaLogicInt("GoCallLuaSetGoRoutineMax", GoroutineMax)	// 把上限传递给lua
//
//		// run tables
//		go func(index int) {
//			for {
//				functionName := "GoCallLuaGoRoutineForLuaGameTable"+strconv.Itoa(index)
//				//fmt.Println("",functionName)
//				GoroutineTableLua[i].GoCallLuaLogic(functionName) 		// 给lua的桌子用的 n个协程函数
//				time.Sleep(time.Millisecond * 1000)		//给其他协程让出1秒的时间， 这个可以后期调整
//			}
//		}(i)
//	}
//}
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