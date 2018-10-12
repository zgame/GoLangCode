//------------------------------------------------------------------
// game server入口文件
//------------------------------------------------------------------

package main

import (
	"net"
	"fmt"
	"./Client"
	"github.com/go-ini/ini"
	"strconv"
	"./Utils/log"
	"./Utils/zRedis"
	"./Games"
	"./Logic/Player"
	"./CSV"
	"time"
	"math"
	"github.com/yuin/gopher-lua"
	"./NetWork"
	"./Lua"
	"./Utils/ztimer"
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
var luaCodeToShare *lua.FunctionProto

var WebSocketServer bool	// websocket 开启
var SocketServer bool		// socket 开启
var WebSockewtPort int
var SockewtPort int

var RedisAddress string		// redis 服务器地址
var err error
var LuaReloadTime int	//lua脚本当前最新版本的时间戳，后台设置的，保存在服务器中，定期去更新一次

var CommonLua *lua.LState		// 公共部分lua脚本
var CommonLogicLuaReloadTime int  // 公共逻辑处理的lua更新时间


func main() {

	fmt.Println("-------------------读取本地配置文件---------------------------")
	initSetting()

	fmt.Println("-------------------lua编译---------------------------")
	luaCodeToShare,err = Lua.CompileLua("Script/main.lua")
	if err!=nil{
		fmt.Println("加载main.lua文件出错了！")
	}

	fmt.Println("-------------------数据库连接---------------------------")
	zRedis.InitRedis(RedisAddress)

	fmt.Println("-------------------读取数据库设置---------------------------")
	UpdateDBSetting()

	fmt.Println("-------------------读取CVS数据文件---------------------------")
	CSV.LoadFishServerExcel()

	fmt.Println("-------------------服务器初始化---------------------------")
	Client.AllClientsList = make(map[*Client.Client]struct{})
	Client.AllUserClientList = make(map[uint32]*Client.Client)
	Games.AllGamesList = make(map[int]*Games.Games)
	Player.GetALLUserUUID()			// 获取玩家的总体分配UUID

	//-------------------------------------创建各个游戏，以后新增游戏，要在这里增加进去即可-----------------------------------
	Games.AddGame("满贯捕鱼", Games.GameTypeBY)
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
	CommonLogicInit()
	CommonLogicStart()


	fmt.Println("-------------------游戏服务器开始建立连接---------------------------")
	NetWorkServerStart()


	service := "127.0.0.1:"+strconv.Itoa(SockewtPort)
	listener, err := net.Listen("tcp", service)
	log.CheckError(err)
	for {
		conn, err := listener.Accept()
		if log.CheckError(err){
			continue
		}

		//clientNew := &Client.Client{conn, 0, nil,nil , nil, 0, false, time.Now(), time.Now(),  time.Now(),0 ,0,0}
		var clientNew *Client.Client
		clientNew = clientNew.NewClient(conn)
		Client.AllClientsList[clientNew] = struct{}{} //把新客户端地址增加到哈希表中保存，方便以后遍历

		go startClient(clientNew)
	}

	// ----------------------主循环计时器----------------------------------------
	tickerCheckUpdateData := time.NewTicker(time.Second * 5)		// 每5秒触发一次计时器，用于定期更新数据库的通用配置信息
	defer tickerCheckUpdateData.Stop()

	for {
		select {
		case <-tickerCheckUpdateData.C:
			// 定期更新后台的配置数据
			UpdateDBSetting()
		}
	}
	


}


//-------------------------------------------------------------------------------------
//  这里是每一个连接新建一个协程，有接收和发送的时候，内部会通知，其他时候自动阻塞
func startClient(client *Client.Client) {
	for {
		if client.Receive() == false{
			// 连接关闭， 那么退出吧
			client = nil
			//fmt.Println("-------连接关闭--------")
			return
		}
	}
}







//-----------------------------本地配置文件---------------------------------------------------
func initSetting()  {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("读取配置文件出错")
		return
	}
	WebSockewtPort ,err  = f.Section("Server").Key("WebSockewtPort").Int()
	SockewtPort ,err   = f.Section("Server").Key("SockewtPort").Int()
	log.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()
	WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()
	RedisAddress = f.Section("Server").Key("SocketServer").String()

	log.CheckError(err)
}

//-------------------------定期更新后台的配置数据---------------------------------------------------------
func UpdateDBSetting() {
	// 以后增加读取数据库的代码
	//...

	// 从服务器更新lua热更新的时间戳
	LuaReloadTime = 1111
	CommonLogicLuaReloadCheck()		//共有逻辑检查一下是否需要更新, 玩家部分每个连接自己检查
	


}

//-----------------------------------建立服务器的网络功能---------------------------------------------------------------
func NetWorkServerStart()  {
	if WebSocketServer {
		// websocket 服务器开启---------------------------------
		wsServer = new(NetWork.WSServer)
		wsServer.Addr = "localhost:"+strconv.Itoa(WebSockewtPort)
		wsServer.MaxConnNum = 2000
		wsServer.PendingWriteNum = 100
		wsServer.MaxMsgLen = 4096
		wsServer.HTTPTimeout = 10 * time.Second
		wsServer.CertFile = ""
		wsServer.KeyFile = ""
		wsServer.NewAgent = func(conn *NetWork.WSConn) NetWork.Agent {
			a := &myServer{conn: conn}
			return a
		}

		wsServer.Start()
	}
	if SocketServer{
		// socket 服务器开启----------------------------------
		server = new(NetWork.TCPServer)
		server.Addr = "127.0.0.1:"+strconv.Itoa(SockewtPort)
		server.MaxConnNum = int(math.MaxInt32)
		server.PendingWriteNum = 100
		server.LenMsgLen = 4
		server.MaxMsgLen = math.MaxUint32
		server.NewAgent = func(conn *NetWork.TCPConn) NetWork.Agent {
			a := &myServer{conn: conn}
			return a
		}
		server.Start()

	}
}

//-----------------------------------游戏公共逻辑处理-------------------------------------------------------
func CommonLogicInit() {
	CommonLua = Lua.Init(luaCodeToShare) // 绑定lua脚本
	Lua.GoCallLua(CommonLua,1)
	CommonLogicLuaReloadTime = LuaReloadTime
}

func CommonLogicLuaReloadCheck()  {
	if CommonLogicLuaReloadTime == LuaReloadTime {
		return
	}
	// 如果跟本地的lua时间戳不一致，就更新
	err = Lua.GoCallLuaReload(CommonLua)
	if err == nil{
		// 热更新成功
		CommonLogicLuaReloadTime = LuaReloadTime
	}
}

func CommonLogicStart() {

	// 创建计时器，定期去run公共逻辑
	ztimer.TimerCheckUpdate(func() {
		Lua.GoCallLuaCommonLogic(CommonLua,"GoCallLuaCommonLogicRun") //公共逻辑处理循环
	}, 5)

	ztimer.TimerClock12(func() {		// 创建计时器，夜里12点触发
		Lua.GoCallLuaCommonLogic(CommonLua,"GoCallLuaCommonLogic12clock") //公共逻辑处理循环
	})


}
