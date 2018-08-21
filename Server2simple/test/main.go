// test server入口文件
package main

import (
	"net"
	"os"
	"strconv"
	"sync"

	//"util/logs"
	"github.com/astaxie/beego/logs"
	"fmt"
	"../client"
)

//
func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// 程序入口
func main() {
	addr := "127.0.0.1:7981"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)

	num := 2					// 压测客户端数量

	if len(os.Args) > 1 {
		num, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Println("max conn:", num)

	w := sync.WaitGroup{}
	for i := 0; i < num; i++ {
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

			client := &client.Client{conn, i}
			TestClient(client)
		}(i)
	}

	w.Wait()

	logs.Info("test finished!")
}
func TestClient(c *client.Client) {
	//var e error

	//
	//logon := &NcLogin{
	//	AccId:  proto.Uint32(uint32(c.index)),
	//	Digest: proto.String("aaa"),
	//	Time:   proto.Uint32(9898),
	//}
	str := "hello I am client: " + strconv.Itoa(c.Index)
	fmt.Println(str)
	c.Send(str)
	//checkPanic(e)

	//var m NsLogin
	_, bytes := c.Receive(false)
	fmt.Println("client receive:", string(bytes))
	//checkPanic(e)

	_, bytes = c.Receive(false)
	fmt.Println("client receive:", string(bytes))
	//c.Infoln("nslogin:", m.String())

	//
	//create := &NcCreatePlayer{Name: proto.String("nnd"), PlantIndex: proto.Int(1)}
	//e = c.Send(EMsgId_ID_NcCreatePlayer, create)
	//checkPanic(e)
	//
	//var screate NsCreatePlayer
	//e = c.Recv(EMsgId_ID_NsCreatePlayer, &screate)
	//checkPanic(e)
	//
	//c.Infoln("nscreatepalyer:", screate.String())

	//time.Sleep(time.Second * 500)
}
