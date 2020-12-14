package udp

import (
	"fmt"
	"net"
)

func Server()  {

	fmt.Println("udp server")
	pUDPAddr, err := net.ResolveUDPAddr("udp", "10.96.8.121:8124")
	listener, err := net.ListenUDP("udp", pUDPAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//logger.Printf("listening on addr=%s with block size=%d", listener.LocalAddr(), *blockSize)

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("接收到<%s> :%s\n", remoteAddr, data[:n])
		str:=fmt.Sprintf(   "我已经接受到你的消息 %s :%s " ,remoteAddr , string(data[:n]))
		listener.WriteToUDP([]byte( str),remoteAddr)
		fmt.Println("发送")
	}
	
}