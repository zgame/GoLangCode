package udp

import (
	"fmt"
	"net"
)

func Client()  {
	pUDPAddr, err := net.ResolveUDPAddr("udp", "10.96.8.121:8124")
	conn, err := net.DialUDP("udp", nil, pUDPAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("我是客户端!"))

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == nil {
			fmt.Printf("接收：%s",buf[:n])
			conn.Close()
			return
		} else {
			return
		}
	}
	
}
