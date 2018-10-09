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
)


//func checkPanic(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//var ServerOpen bool		//服务器开启完毕

// ---------------------------程序入口-----------------------------------
func main() {
	//ServerOpen = false
	fmt.Println("-------------------读取配置文件---------------------------")
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("配置文件出错")
		return
	}
	Port ,err   := f.Section("Server").Key("Port").Int()
	log.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()

	service := "127.0.0.1:"+strconv.Itoa(Port)
	listener, err := net.Listen("tcp", service)
	log.CheckError(err)
	fmt.Println("-------------------数据库连接---------------------------")
	zRedis.InitRedis()

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


	fmt.Println("-------------------游戏服务器开始运行---------------------------")
	//---------------------------------监听网络接口---------------------------------------
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
	// --------------------------------------------------------------
}

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

// --------------------------------------------------------------