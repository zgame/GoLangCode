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
	UdpAddr            *net.UDPAddr
	MaxConnNum      int
	PendingWriteNum int
	NewAgent        func(*UdpConn) Agent
	ln              *net.UDPConn
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

	server.ln = ln
	//server.conns = make(UdpConnSet)

	// msg parser
	//msgParser := NewMsgParser()
	//msgParser.SetMsgLen(server.LenMsgLen, server.MinMsgLen, server.MaxMsgLen)
	//msgParser.SetByteOrder(server.LittleEndian)
	//server.msgParser = msgParser
}

func (server *UDPServer) run() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()

	var tempDelay time.Duration
	data := make([]byte, 1024*1)
	for {
		// udp 是无连接状态的
		n, remoteAddr, err := server.ln.ReadFromUDP(data)
		println("消息：",string(data[:n]))
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

		//server.mutexConns.Lock()
		////server.conns[remoteAddr] = struct{}{}
		//server.mutexConns.Unlock()

		server.wgConns.Add(1)

		//if v,ok := server.AddrMap.Load(remoteAddr.String());ok{
		//	// 有记录
		//}else{
		//	// 没有记录
		//}


		udpConn := newUDPConn(server.ln ,remoteAddr, data)  			// 传递数据给lua
		agent := server.NewAgent(udpConn)
		go func() {
			agent.Run()

			// cleanup
			udpConn.Close() // 接收的线程关闭的时候， 也会关闭发送的线程
			//server.mutexConns.Lock()
			//delete(server.conns, conn)
			//server.mutexConns.Unlock()
			agent.OnClose()

			server.wgConns.Done()
			//fmt.Println("当前连接数量为：",len(server.conns))
		}()
	}
}

func (server *UDPServer) Close() {
	server.ln.Close()
	server.wgLn.Wait()

	//server.mutexConns.Lock()
	//for conn := range server.conns {
	//	conn.Close()
	//}
	//server.conns = nil
	//server.mutexConns.Unlock()
	server.wgConns.Wait()
}
