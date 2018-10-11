// test server入口文件
package main

import (
	"net"
	"sync"

	//"util/logs"
	//"github.com/astaxie/beego/logs"
	"fmt"

	"time"
	"github.com/go-ini/ini"
)


//const (
//	LoginServer = "192.168.101.109:8300"		// 登录服务器地址
//	ClientNum = 200								// 压测客户端数量
//	GameServer =  "192.168.101.109:9902"		// 游戏服务器
//)

//
func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

var GameServer string
var ShowLog int
var SkillID int
var FireCD int
//var SkillCD int
var GoldFishMap map[int]int		//key : fishid   value: fishtype
//var OnlyUseSkill bool
var UseSkill bool
var UseFire bool
var GameKindID int
// 程序入口
func main() {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("配置文件出错")
		return
	}
	LoginServer := f.Section("Server").Key("LoginServer").Value()
	GameServer  = f.Section("Server").Key("GameServer").Value()
	ClientStart,err   := f.Section("Server").Key("ClientStart").Int()
	ClientEnd ,err   := f.Section("Server").Key("ClientEnd").Int()
	ShowLog ,err   = f.Section("Server").Key("ShowLog").Int()
	SkillID ,err   = f.Section("Server").Key("SkillID").Int()
	FireCD ,err   = f.Section("Server").Key("FireCD").Int()
	GameKindID ,err   = f.Section("Server").Key("GameKindID").Int()
	//SkillCD ,err   = f.Section("Server").Key("SkillCD").Int()
	//OnlyUseSkill ,err   = f.Section("Server").Key("OnlyUseSkill").Bool()
	UseFire ,err   = f.Section("Server").Key("UseFire").Bool()
	UseSkill ,err   = f.Section("Server").Key("UseSkill").Bool()

	GetEngine()
	getGoldFishMap()

	addr := LoginServer
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)

	//num := ClientNum					// 压测客户端数量

	//if len(os.Args) > 1 {
	//	num, _ = strconv.Atoi(os.Args[1])		// 或者是命令行输入数量
	//}

	fmt.Println("max conn start :", ClientStart, "--------", ClientEnd)

	w := sync.WaitGroup{}


	for i := ClientStart; i < ClientEnd; i++ {
		time.Sleep(time.Millisecond * 50)			//申请登录的时候给一个延迟
		w.Add(1)
		go func(i int) {
			defer w.Done()

			fmt.Println("connection:", i)
			conn, e := net.DialTCP("tcp", nil, tcpAddr)
			if e != nil {
				fmt.Println(i, e)
				return
			}
			defer conn.Close()


			clients := &Client{conn, i, nil,nil , nil, 0, false, time.Now(), time.Now(),time.Now(),  time.Now(),0 ,0,0}
			clients.Gameinfo = clients.Gameinfo.New()

			//fmt.Println("发送登录请求",i)
			clients.LoginSend()		//开始登录请求
			//fmt.Println("发送登录完成")
			startClient(clients)

		}(i)
	}

	w.Wait()
	//for{
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("timer ")
	//}
	fmt.Println("-----------------------------------------------------")
	fmt.Println("---------全部连接已经关闭 -------  ")
	fmt.Println("---------压测已经结束! ----- ")
	fmt.Println("-----------------------------------------------------")

	for{
		select {

		}
		time.Sleep(time.Second)
	}


}
func startClient(c *Client) {
	//var e error
	for {

		//c.Conn.SetDeadline(time.Now().Add(1e10))
		//fmt.Println(" receive ")

		if c.Receive() == false{
			// 连接关闭， 那么退出吧
			//fmt.Println("-------关闭--------")
			return
		}

		c.GameAI()


	}

}

func (c *Client) ConnectGameServer(addr string)  {
	c.Conn.Close()

	addr = GameServer
	//addr =  "192.168.101.109:9902"

	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
	//fmt.Println("connection:", c.Index,  "------",  addr)
	conn, e := net.DialTCP("tcp", nil, tcpAddr)
	if e != nil {
		fmt.Println(c.Index, e)
		return
	}
	defer conn.Close()
	c.Conn = conn
	c.SendTokenID = 0
	//clients := &Client{conn, i, nil,nil , nil}
	//clients.Gameinfo = clients.Gameinfo.New()

	//fmt.Println("发送登录游戏服务器请求",c.Index)
	c.loginGS()
	//fmt.Println("发送登录游戏服务器完成")
	startClient(c)
}

