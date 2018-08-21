package main

import (
	"net"
	"fmt"
	"os"
	"time"
)



func main() {

	service := "127.0.0.1:8081"
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		handlerClient(conn)
		go handler_timer(conn)
	}
}

// 定期给客户端发送消息
func handler_timer(conn net.Conn)  {
	for {
		sInfo := "I am server, tick！"

		fmt.Println("send:"+sInfo)
		Sendmsg(conn, sInfo)

		time.Sleep(1*time.Second)
	}
}


func handlerClient(conn net.Conn) {
	//defer conn.Close()

	buf := make([]byte,1024) //定义一个切片的长度是1024。
	n,err :=conn.Read(buf)
	//result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("收到客户端消息："+string(buf[:n]))

	sInfo := "你已经连接服务器，你的消息是！"+string(buf[:n])
	clientInfo := []byte(sInfo)
	_,err = conn.Write(clientInfo)
	checkError(err)
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
