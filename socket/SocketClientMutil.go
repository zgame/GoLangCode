package main

import (
	"fmt"
	"net"
	"io"
	"log"

	"os"
	"time"
)

func main() {
	service := "127.0.0.1:8081"
	conn, err := net.Dial("tcp", service)
	checkError(err)

	Sendmsg(conn,"我是客户端!")

	for{
		go handlerRead(conn)
		Sendmsg(conn,"我是客户端!")
		time.Sleep(1*time.Second)
	}
	defer conn.Close()  //断开TCP链接。
}


// 接收消息
func handlerRead(conn net.Conn) {
	buf := make([]byte,1024) //定义一个切片的长度是1024。
	n,err :=conn.Read(buf)

	if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	fmt.Println("")
}

// 发送消息
func Sendmsg(conn net.Conn, msg string){
	_, err := conn.Write([]byte(msg))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
