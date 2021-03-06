//------------------------------------------------------------------
// game server入口文件
//------------------------------------------------------------------

package main

import (
	"fmt"
	//"github.com/go-ini/ini"
	"strconv"
	//"./Games"
	//"./Logic/UserModel"
	"./CSV"
	"time"
	//"math"
	//"./Core/NetWork"
	"./Core/Utils/zLog"
	//"./Core/Utils/zIP"
	"./Core/Utils/zRedis"
	"./Core/ZServer"
	"./Core/Utils/ztimer"
	//"github.com/yuin/gopher-lua"
	//"./GlobalVar"
	"runtime"
	"flag"
	"net/http"
	_ "net/http/pprof"
	oldLog "log"
	"./Games"
	"./Games/Model/PlayerModel"
	"./Core/GameCore"
	"./Games/CommonLogic"
	"./Games/DataBase/SQLSeverDB/SqlServerBy"
	"./Games/DataBase/SQLSeverDB/SqlServerLog"
	"./Games/DataBase/SQLSeverDB/SqlServerFriend"
	"./Games/DataBase/MySqlDB"
	//"github.com/go-ini/ini"
)


//func checkPanic(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//var ServerOpen bool		//服务器开启完毕

// ---------------------------程序入口-----------------------------------




var WebSocketPort int
var SocketPort int
var GameRoomServerID int	// 游戏原有的ServerID(来源于GameRoomInfo中对应的)
var ServerAddress string    // ServerAddress 服务器地址
//var WebSocketAddress string // WebSocketAddress 服务器地址


//

//var GameManagerLua *ZServer.MyLua // 公共部分lua脚本
//var GameManagerLuaReloadTime int    // 公共逻辑处理的lua更新时间



//var GoroutineMax int 			// 给lua的游戏桌子使用的协程数量		暂时没用
//var GoroutineTableLua *ZServer.MyLua		// 桌子lua脚本
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
	zLog.ServerPort = SocketPort // 传递给log日志，让日志记录的时候区分服务器端口
	fmt.Println("-------------------读取本地配置文件---------------------------")
	if !initSetting() {
		return
	}


	fmt.Println("-------------------读取CVS数据文件---------------------------")
	CSV.LoadFishServerExcel()

	fmt.Println("-------------------服务器初始化---------------------------")
	initVar()
	//ZServer.QueueInit()
	//UserModel.GetALLUserUUID()			// 获取玩家的总体分配UUID

	fmt.Println("-------------------Redis 数据库连接---------------------------")
	if zRedis.InitRedis(RedisAddress, RedisPass) == false {
		return
	}
	fmt.Println("-------------------MySql 数据库连接---------------------------")
	if !MySqlDB.ConnectMySqlDB(MySqlServerIP, MySqlServerPort, MySqlDatabase, MySqlUid, MySqlPwd){
		return
	}
	fmt.Println("-------------------Sql server 数据库连接---------------------------")
	if !SqlServerBy.BYConnectSqlDB(SqlServerIP,  SqlServerDatabase, SqlServerUid, SqlServerPwd) {
		return
	}
	fmt.Println("-------------------Sql server Log 数据库连接---------------------------")
	if !SqlServerLog.LogConnectSqlDB(SqlServerIPLog,  SqlServerDatabaseLog, SqlServerUidLog, SqlServerPwdLog)  {
		return
	}
	fmt.Println("-------------------Sql server Friend 数据库连接---------------------------")
	if !SqlServerFriend.FriendConnectSqlDB(SqlServerIPFriend,  SqlServerDatabaseFriend, SqlServerUidFriend, SqlServerPwdFriend)  {
		return
	}



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



	fmt.Println("-------------------游戏服务器初始化---------------------------")
	GameManagerInit()
	//fmt.Println("ZServer 代码初始化完成")

	//fmt.Println("-------------------启动 ZServer 访问 MySql ---------------------------")
	//if GameManagerLua.GoCallLuaConnectMysql(MySqlServerIP, MySqlServerPort , MySqlDatabase,MySqlUid,MySqlPwd) == false{
	//	fmt.Println("lua zMysqlForLua 数据库没有连接成功")
	//	return
	//}
	//fmt.Println("-------------------启动 ZServer 访问 Sql Server---------------------------")
	//if GameManagerLua.GoCallLuaConnectSqlServer(SqlServerIP, "" , SqlServerDatabase, SqlServerUid,SqlServerPwd) == false{
	//	fmt.Println("lua sql  server 数据库没有连接成功")
	//	return
	//}

	//fmt.Println("-------------------多核桌子逻辑处理器---------------------------")
	//CreateGoroutineForLuaGameTable()

	//fmt.Println("-------------------启动gameManager---------------------------")
	//if GameManagerLua.GoCallLuaConnectMysql(MySqlServerIP,Database,MySqlUid,MySqlPwd) == false{
	//	fmt.Println("lua zMysqlForLua 数据库没有连接成功")
	//	return
	//}
	//GameManagerLua.GoCallLuaLogic("GoCallLuaStartGamesServers")
	//StartMultiThreadChannelPlayerToGameManager()



	//fmt.Println("-------------------读取数据库设置---------------------------")
	//UpdateLuaReload()

	fmt.Println("-------------------游戏服务器开始建立连接---------------------------")
	ZServer.NetWorkServerStart(ServerAddress , SocketPort , WebSocketPort)


	//-------------------------------------创建各个游戏，以后新增游戏，要在这里增加进去即可-----------------------------------
	fmt.Println("-------------------启动游戏服务器---------------------------")
	Games.StartGamesServers()
	fmt.Println("-------------------启动Common定时处理器---------------------------")
	TimerCommonLogicStart()

	//service := "127.0.0.1:"+strconv.Itoa(SocketPort)
	//listener, err := net.Listen("tcp", service)
	//zLog.CheckError(err)
	//for {
	//	conn, err := listener.Accept()
	//	if zLog.CheckError(err){
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

		//ztimer.CheckRunTimeCost(func() {
		//		GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 桌子的run
		//	}, "桌子循环GoCallLuaGoRoutineForLuaGameTable"		)

		//startTime := ztimer.GetOsTimeMillisecond()
		//GameManagerLua.GoCallLuaLogic("GoCallLuaGoRoutineForLuaGameTable") // 桌子的run
		//if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		//	zLog.PrintfLogger("--------!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!![ 警告 ] 桌子循环 消耗时间过长: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
		//}
		runtime.GC()
		time.Sleep(time.Millisecond * 1000 * 60)                                //给其他协程让出60秒的时间， 这个可以后期调整
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




//----------------------------变量的初始化---------------------------------------------------------------
func initVar(){
	//Client.AllClientsList = make(map[*Client.Client]struct{})
	//Client.AllUserClientList = make(map[uint32]*Client.Client)
	//Games.AllGamesList = make(map[int]*Games.Games)
	ZServer.InitGlobalVar()
}




//-----------------------------------游戏公共逻辑处理-------------------------------------------------------
func GameManagerInit() {

	Games.NetWorkFuncRegister()		// 回调注册
	GameCore.AllGamesList = make(map[int]*GameCore.Games)
	GameCore.AllPlayerList = make(map[int] *PlayerModel.Player)
	//GameManagerLua = ZServer.NewMyLua()

	//GameManagerLua.Init() // 绑定lua脚本
	//ZServer.GoCallLuaTest(GameManagerLua.L,1)

	//ZServer.GameManagerLuaHandle = GameManagerLua // 把句柄传递给lua保存一份
	//GameManagerLuaReloadTime = GlobalVar.LuaReloadTime
	//GameManagerLua.GoCallLuaSetStringVar("ServerIP_Port", ServerAddress+ ":" + strconv.Itoa(SocketPort)) 	//把服务器地址传递给lua
	//GameManagerLua.GoCallLuaSetIntVar("GameRoomServerID", GameRoomServerID) 								//把服务器地址传递给lua
}

//// 检查通用逻辑部分的lua是否需要更新
//func GameManagerLuaReloadCheck() {
//	if GameManagerLuaReloadTime == GlobalVar.LuaReloadTime {
//		//return
//	}
//	// 如果跟本地的lua时间戳不一致，就更新
//	err = GameManagerLua.GoCallLuaReload()
//	if err == nil{
//		// 热更新成功
//		GameManagerLuaReloadTime = GlobalVar.LuaReloadTime
//	}
//}

func TimerCommonLogicStart() {
	// -------------------创建计时器，定期去run公共逻辑---------------------
	ztimer.TimerCheckUpdate(func() {
		connectShow, _, _, _ := GetAllConnectMsg()
		memoryShow, _ := zLog.GetSysMemInfo()

		zLog.PrintfLogger("[%s]   %s   %s      处理消息平均时间：%d  ", ServerAddress+":"+strconv.Itoa(SocketPort),
			connectShow, memoryShow, ZServer.StaticNetWorkReceiveToSendCostTime)

		// 把服务器的状态信息，传递给lua
		//GameManagerLua.GoCallLuaSetIntVar("ServerStateSendNum", successSendMsg)
		//GameManagerLua.GoCallLuaSetIntVar("ServerStateReceiveNum", successRecMsg)
		//GameManagerLua.GoCallLuaSetIntVar("ServerSendWriteChannelNum", WriteChan)
		//GameManagerLua.GoCallLuaSetIntVar("ServerDataHeadErrorNum", ZServer.StaticDataPackageHeadFlagError)
		//GameManagerLua.GoCallLuaSetIntVar("ServerHeapInUse", heapInUse)
		//GameManagerLua.GoCallLuaSetIntVar("ServerNetWorkDelay", ZServer.StaticNetWorkReceiveToSendCostTime)

		for k,v := range GameCore.AllGamesList {
			zLog.PrintfLogger("游戏id   %d   桌子数量：%d   玩家数量：  %d", k, len(v.AllTableList),  len(GameCore.AllPlayerList))
		}


		ztimer.CheckRunTimeCost(func() {
			//GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogicRun") //公共逻辑处理循环
			CommonLogic.CommonLogicRun()
		}, "GoCallLuaCommonLogicRun")

	}, 10)

	//// ---------------------创建计时器，定期去保存服务器状态---------------------
	//ztimer.TimerCheckUpdate(func() {
	//	ztimer.CheckRunTimeCost(func() {
	//		GameManagerLua.GoCallLuaLogic("GoCallLuaSaveServerState") // 保存服务器的状态
	//	}, "GoCallLuaSaveServerState")
	//
	//
	//	runtime.GC()
	//}, 60)	// 1 分钟

	//// ---------------------创建计时器，定期去更新lua脚本reload---------------------
	//ztimer.TimerCheckUpdate(func() {
	//	// 定期更新后台的配置数据
	//	ztimer.CheckRunTimeCost(func() {
	//		UpdateLuaReload()
	//	}, "UpdateLuaReload")
	//
	//
	//},  20 * 1)  // 1 分钟

	//---------------------创建计时器，夜里12点触发---------------------
	ztimer.TimerClock0(func() {
		ztimer.CheckRunTimeCost(func() {
			//GameManagerLua.GoCallLuaLogic("GoCallLuaCommonLogic12clock") //公共逻辑处理循环
			CommonLogic.CommonLogic12clock()
		}, "GoCallLuaCommonLogic12clock")
	})


	////---------------------创建计时器，定期执行所有的接受消息队列---------------------
	//ztimer.TimerMillisecondCheckUpdate(func() {
	//	ztimer.CheckRunTimeCost(func() {
	//		ZServer.QueueRun()
	//	}, "ZServer.QueueRun")
	//},  20)  //

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

	ZServer.RWMutexServerIdConnect.RLock()
	for _,v := range ZServer.ServerIdConnectMyServer {
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
	ZServer.RWMutexServerIdConnect.RUnlock()
	//GameManagerLua.GoCallLuaSetIntVar("ServerSendWriteChannelNum", WriteChan)                           // 发送缓冲区大小
	//GameManagerLua.GoCallLuaSetIntVar("ServerDataHeadErrorNum", ZServer.StaticDataPackageHeadFlagError) // 把数据头尾错误发送给lua
	str:=fmt.Sprintf(" 发送连接数量 %d  接收连接数量  %d 每秒发送 %d  每秒接收 %d    发送缓存WriteChan %d  ",   successSendClients, successRecClients, successSendMsg , successRecMsg, WriteChan)
	return "所有连接数量："+ strconv.Itoa(AllConnect) + str , successSendMsg , successRecMsg, WriteChan
}
