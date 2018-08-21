package main

import (
	"fmt"
	"net"
	"io"
	//"log"

	"os"
	//"time"
	"sync"
	"strconv"
	"time"
)

const max_num = 5000


type Cclients struct{
	conn  net.Conn
	Index int
}

func main() {
	service := "127.0.0.1:8301"
	w := sync.WaitGroup{}

	for i:=0;i<max_num;i++{
		conn, err := net.Dial("tcp", service)
		checkError(err)
		w.Add(1)

		clients := Cclients{conn, i}
		go func(i int) {
			defer w.Done()
			defer conn.Close()				//断开TCP链接。


			//clients.Sendmsg("我是客户端!"+strconv.Itoa(i))
			if clients.handlerRead() == false{
				return
			}

		}(i)
		go Send_loop_timer(clients)
	}
	w.Wait()



}


// 接收消息
func (this *Cclients)handlerRead()bool{
	for{
		//fmt.Println("eeeee")
		buf := make([]byte,1024) //定义一个切片的长度是1024。
		n,err :=this.conn.Read(buf)

		if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
			fmt.Println("对方关闭了连接")
			return false
		}
		fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
		//fmt.Println("")

		//time.Sleep(time.Millisecond * 500)

		//sInfo := "我是id："+ strconv.Itoa(this.Index)
		//this.Sendmsg(sInfo)
	}
	return true

}

// 定期给客户端发送消息
func Send_loop_timer(clients Cclients)  {
	for {
		//fmt.Println("ddddddddd")
		sInfo := "我是id："+ strconv.Itoa(clients.Index)
		clients.Sendmsg(sInfo)
		time.Sleep(time.Millisecond * 5000)
	}
}


// 发送消息
func (this *Cclients)Sendmsg(msg string){
	_, err := this.conn.Write([]byte(msg))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
