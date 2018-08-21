package main

import (
	"log"
	"github.com/gorilla/websocket"
	"time"
	"strconv"
)

const w_max_num = 20

type CWclients struct {
	conn  *websocket.Conn
	Index int
}

func main() {
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	//log.Printf("connecting to %s", u.String())
	for i := 0; i < w_max_num; i++ {
		c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/echo", nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		c.WriteMessage(websocket.TextMessage, []byte(" 1我是客户端   "+ strconv.Itoa(i)))


		//done := make(chan struct{})

		go func(index int) {
			clients := CWclients{c, index}
			//defer close(done)
			for {
				// 接收数据的部分
				_, message, err := clients.conn.ReadMessage()
				if err != nil {
					log.Println("接收数据错误:", err)
					return
				}
				log.Printf("收到数据: %s", message)

				time.Sleep(time.Second)
				err = clients.conn.WriteMessage(websocket.TextMessage, []byte(" 我是客户端   "+ strconv.Itoa(clients.Index)))
				if err != nil {
					log.Println("定时发送错误:", err)
					return
				}

			}
		}(i)

		//ticker := time.NewTicker(time.Second * 10)
		//defer ticker.Stop()
		//
		//for {
		//	select {
		//	//case <-done:
		//	//	return
		//	case t := <-ticker.C:
		//		err := clients.conn.WriteMessage(websocket.TextMessage, []byte(t.String()+"       "+ strconv.Itoa(clients.Index)))
		//		if err != nil {
		//			log.Println("定时器发送错误:", err)
		//			return
		//		}
		//
		//	}
		//}

	}
	for{
		select {

		}
	}
}
