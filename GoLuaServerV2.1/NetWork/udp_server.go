package NetWork

import (
	//"github.com/name5566/leaf/zLog"
	"net"
	"sync"
	"time"
	"fmt"
)
//---------------------------------------------------------------------------------------------------
// Socket udp 服务器
//---------------------------------------------------------------------------------------------------

type UDPServer struct {
	Addr            string
	UdpAddr         *net.UDPAddr
	MaxConnNum      int
	PendingWriteNum int
	NewAgent        func(*UdpConn) Agent
	Listen          *net.UDPConn
	//conns           UdpConnSet
	//mutexConns      sync.Mutex		// 互斥锁， 用在保持多线程对map的操作安全上
	wgLn            sync.WaitGroup
	wgConns         sync.WaitGroup
	//AddrMap			sync.Map

	// msg parser
	LenMsgLen    int
	MinMsgLen    uint32
	MaxMsgLen    uint32
	LittleEndian bool
	//msgParser    *MsgParser
}

func (server *UDPServer) Start() {
	fmt.Println("开始socket Udp 服务器")
	server.init()
	go server.run()
}

func (server *UDPServer) init() {
	server.UdpAddr, _ = net.ResolveUDPAddr("udp", server.Addr)
	ln, err := net.ListenUDP("udp", server.UdpAddr)
	if err != nil {
		fmt.Printf("%v", err)
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		fmt.Printf("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}
	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = 100
		fmt.Printf("invalid PendingWriteNum, reset to %v", server.PendingWriteNum)
	}
	if server.NewAgent == nil {
		fmt.Println("NewAgent must not be nil")
	}

	server.Listen = ln
}

func (server *UDPServer) run() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()

	var tempDelay time.Duration
	data := make([]byte, 1024*1)
	for {
		// udp 是无连接状态的
		_, remoteAddr, err := server.Listen.ReadFromUDP(data)
		//println("消息：",string(data[:n]), remoteAddr.String())
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				fmt.Printf("accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0


		server.wgConns.Add(1)

		udpConn := newUDPConn(server.Listen,remoteAddr, data) // 传递数据给lua
		agent := server.NewAgent(udpConn)
		go func() {
			agent.Run()

			// cleanup
			udpConn.Close() // 接收的线程关闭的时候， 也会关闭发送的线程
			agent.OnClose()

			server.wgConns.Done()
			//fmt.Println("当前连接数量为：",len(server.conns))
		}()
	}
}

func (server *UDPServer) Close() {
	server.Listen.Close()
	server.wgLn.Wait()

	//server.mutexConns.Lock()
	//for Conn := range server.conns {
	//	Conn.Close()
	//}
	//server.conns = nil
	//server.mutexConns.Unlock()
	server.wgConns.Wait()
}
