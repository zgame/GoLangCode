// test server入口文件
package main

import (
	"net"
	"sync"
	"net/http"
	//"util/logs"
	//"github.com/astaxie/beego/logs"
	"fmt"
	_ "net/http/pprof"			// 注意这里
	"time"
	"github.com/go-ini/ini"
	"log"
	"runtime"
	"strconv"
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
var LoginServer string
var ShowLog int
var SkillID int
var FireCD int
//var SkillCD int
var GoldFishMap map[int]int		//key : fishid   value: fishtype
//var OnlyUseSkill bool
var UseSkill bool
var UseFire bool
var GameKindID int
var ReLoginTime int

var AllUserNum int			// 同时在线玩家数量

// 程序入口
func main() {
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("配置文件出错")
		return
	}
	LoginServer = f.Section("Server").Key("LoginServer").Value()
	GameServer  = f.Section("Server").Key("GameServer").Value()
	ClientStart,err   := f.Section("Server").Key("ClientStart").Int()
	ClientEnd ,err   := f.Section("Server").Key("ClientEnd").Int()

	AllUserNum = ClientEnd - ClientStart

	ShowLog ,err   = f.Section("Server").Key("ShowLog").Int()
	SkillID ,err   = f.Section("Server").Key("SkillID").Int()
	FireCD ,err   = f.Section("Server").Key("FireCD").Int()
	GameKindID ,err   = f.Section("Server").Key("GameKindID").Int()
	ReLoginTime ,err   = f.Section("Server").Key("ReLoginTime").Int()
	//SkillCD ,err   = f.Section("Server").Key("SkillCD").Int()
	//OnlyUseSkill ,err   = f.Section("Server").Key("OnlyUseSkill").Bool()
	UseFire ,err   = f.Section("Server").Key("UseFire").Bool()
	UseSkill ,err   = f.Section("Server").Key("UseSkill").Bool()


	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

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



			clients := &Client{conn, i, nil, nil, nil, 0, false,
			time.Now(), time.Now(), time.Now(), time.Now(), 0, 0, 0,
				time.Now().Add(time.Second * time.Duration(ReLoginTime)),false,nil}

			clients.Gameinfo = clients.Gameinfo.New()

			//fmt.Println("发送登录请求",i)
			clients.LoginSend() //开始登录请求
			//fmt.Println("发送登录完成")
			startClient(clients)

		}(i)
	}

	go func() {
		for{
			PrintfLogger("内存情况：%s",GetSysMemInfo())
			time.Sleep(time.Second)
		}
	}()

	w.Wait()
	//for{
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("timer ")
	//}
	fmt.Println("-----------------------------------------------------")
	fmt.Println("---------全部连接已经关闭 -------  ")
	fmt.Println("---------压测已经结束! ----- ")
	fmt.Println("-----------------------------------------------------")




}
func startClient(c *Client) {
	//var e error
	for {

		//c.Conn.SetDeadline(time.Now().Add(1e10))
		//fmt.Println(" receive ")
		if c.Quit == true{
			fmt.Println("连接结束",c.Index)
			return
		}

		if c.Receive() == false{
			// 连接关闭， 那么退出吧
			//fmt.Println("-------关闭--------")
			return
		}

		c.GameAI(ReLoginTime)


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


func (c *Client) ConnectLoginServer() {
	c.Conn.Close()

	addr := LoginServer

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
	c.ReloginTime = time.Now().Add(time.Second * time.Duration(ReLoginTime))
	c.Gameinfo = nil
	c.Gameinfo = c.Gameinfo.New()

	//clients := &Client{conn, i, nil,nil , nil}
	//clients.Gameinfo = clients.Gameinfo.New()

	//fmt.Println("发送登录游戏服务器请求",c.Index)
	c.LoginSend()
	//fmt.Println("发送登录游戏服务器完成")
	startClient(c)
}



func GetSysMemInfo()  string{
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)

	str:= ""
	//str += "   Lookups:" + strconv.Itoa( int(memStat.Lookups))
	//str += "M   TotalAlloc:" + strconv.Itoa( int(memStat.TotalAlloc/1000000))//从服务开始运行至今分配器为分配的堆空间总和
	str += "  Sys:" + strconv.Itoa( int(memStat.Sys/1000000) )+ "M"
	//str += "M   Mallocs:" + strconv.Itoa( int(memStat.Mallocs))//服务malloc的次数
	//str += "次   Frees:" + strconv.Itoa( int(memStat.Frees))//服务回收的heap objects
	str += "   HeapAlloc:" + strconv.Itoa( int(memStat.HeapAlloc/1000000)) + "M"//服务分配的堆内存
	str += "   HeapSys:" + strconv.Itoa( int(memStat.HeapSys/1000000))+ "M"//系统分配的堆内存
	str += "   HeapIdle:" + strconv.Itoa( int(memStat.HeapIdle/1000000))+ "M"//申请但是为分配的堆内存，（或者回收了的堆内存）
	str += "   HeapInuse:" + strconv.Itoa( int(memStat.HeapInuse/1000000))+ "M"//正在使用的堆内存
	str += "   HeapReleased:" + strconv.Itoa( int(memStat.HeapReleased))+ "M"//返回给OS的堆内存，类似C/C++中的free。
	//str += "   HeapObjects:" + strconv.Itoa( int(memStat.HeapObjects))+ "个"//堆内存块申请的量
	str += "   StackInuse:" + strconv.Itoa( int(memStat.StackInuse/1000000)) + "M"//正在使用的栈
	str += "   StackSys:" + strconv.Itoa( int(memStat.StackSys/1000000)) + "M"//系统分配的作为运行栈的内存
	//str += "   NumGC:" + strconv.Itoa( int(memStat.NumGC))+ "次"////垃圾回收的内存大小
	//str += "   NumForcedGC:" + strconv.Itoa( int(memStat.NumForcedGC))
	//str += "   LastGC:" + strconv.Itoa( int(memStat.LastGC))
	return str

}
