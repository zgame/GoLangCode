//**************************************
// test server入口文件
// 模拟一定数量的客户端连接服务器
//**************************************
package main

import (
	"net"
	"os"
	"strconv"
	"sync"

	"../core/logs"
)

//
func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// 程序入口
func main() {
	addr := "127.0.0.1:7080"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)

	num := 10     // 压力测试
	if len(os.Args) > 1 {
		num, _ = strconv.Atoi(os.Args[1])
	}

	logs.Infoln("max conn:", num)

	w := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()

			logs.Infoln("connection:", i)
			conn, e := net.DialTCP("tcp", nil, tcpAddr)
			if e != nil {
				logs.Warnln(strconv.Itoa(i), e)
				return
			}
			defer conn.Close()

			client := &Client{conn, i}
			TestClient(client)
		}(i)
	}

	w.Wait()

	logs.Info("test finished!")
}
