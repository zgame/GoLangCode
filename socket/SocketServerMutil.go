package main

import (
	"net"
	"fmt"
	"time"
	"os"

	"strconv"
)

type Clients struct{
	conn []net.Conn
	currentNum int
}


func main() {

	clientMsg := Clients{make([]net.Conn,0),0}

	service := "127.0.0.1:8081"
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// 
		clientMsg.conn =  append(clientMsg.conn, conn)
		clientMsg.currentNum ++

		go handlerClientM(conn, clientMsg.currentNum)
		//go loop_timer(conn)
	}
}

// 定期给客户端发送消息
func loop_timer(conn net.Conn)  {
	for {
		sInfo := "I am server, tick！"

		fmt.Println("send:"+sInfo)
		Sendmsg(conn, sInfo)

		time.Sleep(1*time.Second)
	}
}

func handlerClientM(conn net.Conn, id int) {
	//defer conn.Close()
	for {
	//	select {
	//	default:
	//
	//}

		// 设置读取超时
		conn.SetReadDeadline(time.Now().Add(1e10))

		buf := make([]byte, 1024) //定义一个切片的长度是1024。
		n, err := conn.Read(buf)
		//result, err := ioutil.ReadAll(conn)
		checkError(err)
		fmt.Printf("收到客户端消息： %x ", buf[:n])

		sInfo := "你已经连接服务器，你的消息是！ " + string(buf[:n]) + "你的编号是：" + strconv.Itoa(id)
		clientInfo := []byte(sInfo)
		_, err = conn.Write(clientInfo)
		checkError(err)


		// 读取到数据
		if nil == err {
			continue
		}

		// 因超时, 未能读取到数据
		if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
			continue
		}

		fmt.Println("receive msg failed! id:%v,ip:%v,error:%v", id, conn.RemoteAddr(), err)

		return
	}

}




// 发送消息
func Sendmsg(conn net.Conn, msg string){
	_, err := conn.Write([]byte(msg))
	checkError(err)
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//os.Exit(1)
	}
}
