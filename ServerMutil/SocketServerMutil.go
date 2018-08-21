package main

import (
	"net"
	"fmt"
	"time"
	//"os"

	"strconv"
)

type Clients struct{
	conn  map[net.Conn]struct{}
	Index int
}

var clientMsg Clients

func main() {

	clientMsg = Clients{make(map[net.Conn]struct{},0),0}
	go loop_timer()
	service := "127.0.0.1:8301"
	listener, err := net.Listen("tcp", service)
	if err != nil{
		fmt.Println("net.Listen", "监听tcp端口出现问题")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// 
		clientMsg.conn[conn] = struct{}{}		// 哈希格式新增一个连接作为key
		go handlerClientM(conn, clientMsg.Index)
		clientMsg.Index ++

	}
}

// 定期显示当前链接数量
func loop_timer()  {
	for {
		fmt.Println("当前活跃连接数量为:"+ strconv.Itoa(len(clientMsg.conn)))

		time.Sleep(2*time.Second)
	}
}



func handlerClientM(conn net.Conn, id int) {
	//defer conn.Close()
	for {

		//fmt.Println("ssssssssssss")
		// 设置读取超时
		//conn.SetReadDeadline(time.Now().Add(1e10))

		buf := make([]byte, 1024) //定义一个切片的长度是1024。
		n, err := conn.Read(buf)
		//result, err := ioutil.ReadAll(conn)
		if err != nil {
			fmt.Println("接收数据错误", err.Error())
			clearClients(conn)
			return
			//os.Exit(1)
		}
		//fmt.Printf("收到客户端消息： %x ", buf[:n])

		sInfo := "你的消息是: " + string(buf[:n]) + ",编号是：" + strconv.Itoa(id)
		clientInfo := []byte(sInfo)
		_, err = conn.Write(clientInfo)


		// 读取到数据
		if nil != err {
			fmt.Println("网络发送错误 ", strconv.Itoa(id) , "***" , err.Error())

			clearClients(conn)
			return
		}

		// 因超时, 未能读取到数据
		if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
			fmt.Println(" 超时，未能读取数据 " ,strconv.Itoa(id) , "------" , err.Error())

			clearClients(conn)
			return
		}

	}

}

func clearClients(conn net.Conn)  {
	delete(clientMsg.conn, conn)		//删掉哈希表中的连接
	conn.Close()
}


//// 发送消息
//func Sendmsg(conn net.Conn, msg string){
//	_, err := conn.Write([]byte(msg))
//	checkError(err)
//}


//func checkError(err error) {
//	if err != nil {
//		fmt.Println("数据错误", err.Error())
//		//os.Exit(1)
//	}
//}
