// test server入口文件
package main

import (
	//"util/logs"
	//"github.com/astaxie/beego/logs"
	"fmt"

	"time"
	"github.com/go-ini/ini"
	"runtime"
	"flag"
	"./log"
)


//const (
//	LoginServer = "192.168.101.109:8300"		// 登录服务器地址
//	ClientNum = 200								// 压测客户端数量
//	GameServerAddress =  "192.168.101.109:9902"		// 游戏服务器
//)

//
func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

var GameServerAddress string
//var GameServerWebSocketAddress string
var WebSocketPort int
var SocketPort int

var ShowLog int
var IsWebSocket bool

// 程序入口
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("配置文件出错")
		return
	}



	//LoginServer := f.Section("Server").Key("LoginServer").Value()
	GameServerAddress = f.Section("Server").Key("GameServerAddress").Value()
	//GameServerWebSocketAddress = f.Section("Server").Key("GameServerWebSocketAddress").Value()
	//ClientStart,err   := f.Section("Server").Key("ClientStart").Int()
	//ClientEnd ,err   := f.Section("Server").Key("ClientEnd").Int()
	ShowLog ,err   = f.Section("Server").Key("ShowLog").Int()
	IsWebSocket ,err   = f.Section("Server").Key("IsWebSocket").Bool()

	// -------------------------读取命令行参数--------------------------
	wsPort := flag.Int("WebSocketPort", 0, "")
	sPort := flag.Int("SocketPort", 0, "")
	start := flag.Int("ClientStart", 0, "")
	end := flag.Int("ClientEnd", 0, "")
	flag.Parse()
	WebSocketPort = *wsPort
	SocketPort = *sPort
	ClientStart := * start
	ClientEnd := *end

	if WebSocketPort == 0 || SocketPort == 0 || ClientStart == 0 || ClientEnd == 0{

		for{
			fmt.Println("缺少命令行参数！ 参数要设置类似 -WebSocketPort=8089 -SocketPort=8123 -ClientStart=1 -ClientEnd=100")
			time.Sleep(time.Second)
		}
	}

	fmt.Println("max conn start :", ClientStart, "--------", ClientEnd)
	StartClient(ClientStart,ClientEnd, IsWebSocket)

	//
	//w := sync.WaitGroup{}
	//for i := ClientStart; i < ClientEnd; i++ {
	//	w.Add(1)
	//	//GlobalMutex.Lock()
	//	go func(i int) {
	//		defer w.Done()
	//
	//		//fmt.Println("connection:", i)
	//		//conn, e := net.DialTCP("tcp", nil, tcpAddr)
	//		//if e != nil {
	//		//	fmt.Println(i, e)
	//		//	return
	//		//}
	//		//defer conn.Close()
	//
	//
	//		clients := &Client{nil, i, nil,nil , nil, 0, false, time.Now(), time.Now(),  time.Now(),0 ,0,0,0,false,nil}
	//		clients.Gameinfo = clients.Gameinfo.New()
	//		if i==ClientStart{
	//			clients.ShowMsgSendTime = true	// 第一个才显示
	//		}
	//
	//		//fmt.Println("发送登录请求",i)
	//		//clients.LoginSend()		//开始登录请求
	//		clients.ConnectGameServer("")  // 直接登录游戏服务器
	//		//fmt.Println("发送登录完成")
	//		startClient(clients)
	//
	//	}(i)
	//	//GlobalMutex.Unlock()
	//	time.Sleep(time.Millisecond * 50)
	//}
	//
	//w.Wait()
	//for{
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("timer ")
	//}
	//fmt.Println("-----------------------------------------------------")
	//fmt.Println("---------全部连接已经关闭 -------  ")
	//fmt.Println("---------压测已经结束! ----- ")
	//fmt.Println("-----------------------------------------------------")

	for{
		GetStaticPrint()
		time.Sleep(time.Second)
	}
}

func GetStaticPrint()  {
	successSendClients := 0
	successRecClients := 0
	successSendMsg := 0
	successRecMsg := 0
	WriteChan := 0
	AllConnect :=0

	GlobalMutex.Lock()
	for k,_:=range GlobalClients{
		AllConnect++
		if k.SendMsgNum>0{
			successSendClients++
			successSendMsg += k.SendMsgNum
			k.SendMsgNum = 0
		}
		if k.ReceiveMsgNum>0{
			successRecClients ++
			successRecMsg += k.ReceiveMsgNum
			k.ReceiveMsgNum = 0
		}
		WriteChan += k.Conn.GetWriteChanCap()
	}
	//if AllConnect>0{
	//	WriteChan = WriteChan/AllConnect		// 求一个平均值
	//}
	GlobalMutex.Unlock()
	log.PrintfLogger("连接数量 %d 用户正常发送消息数量 %d  正常接收  %d 每秒发送 %d  每秒接收 %d  goroutine数量 %d  WriteChan数量 %d ",  AllConnect, successSendClients, successRecClients, successSendMsg , successRecMsg,  runtime.NumGoroutine(),WriteChan)
}


//func startClient(c *Client) {
//	//var e error
//	for {
//
//		//c.Conn.SetDeadline(time.Now().Add(1e10))
//		//fmt.Println(" receive ")
//
//		if c.Receive() == false{
//			// 连接关闭， 那么退出吧
//			//fmt.Println("-------关闭--------")
//			return
//		}
//
//
//
//		time.Sleep(time.Millisecond * 200)
//		c.GameAI()
//	}
//
//}
//
//func (c *Client) ConnectGameServer(addr string)  {
//	//c.Conn.Close()
//
//	addr = GameServerAddress
//	//addr =  "192.168.101.109:9902"
//
//	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
//	//fmt.Println("connection:", c.Index,  "------",  addr)
//	conn, e := net.DialTCP("tcp", nil, tcpAddr)
//	if e != nil {
//		fmt.Println(c.Index, "服务器连接不上",e)
//		return
//	}
//	defer conn.Close()
//	//c.Conn = conn
//	c.SendTokenID = 1
//	//clients := &Client{conn, i, nil,nil , nil}
//	//clients.Gameinfo = clients.Gameinfo.New()
//
//	//fmt.Println("发送登录游戏服务器请求",c.Index)
//	c.loginGS()
//	//fmt.Println("发送登录游戏服务器完成")
//	startClient(c)
//}


func TimerCheckUpdate(f func(), timer time.Duration)  {
	go func() {
		tickerCheckUpdateData := time.NewTicker(time.Second * timer)
		defer tickerCheckUpdateData.Stop()

		for {
			select {
			case <-tickerCheckUpdateData.C:
				f()
			}
		}
	}()
}
