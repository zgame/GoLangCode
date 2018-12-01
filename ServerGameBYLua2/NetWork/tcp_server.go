package NetWork

import (
	//"github.com/name5566/leaf/log"
	"net"
	"sync"
	"time"
	"fmt"
)
//---------------------------------------------------------------------------------------------------
// Socket 服务器
//---------------------------------------------------------------------------------------------------

type TCPServer struct {
	Addr            string
	MaxConnNum      int
	PendingWriteNum int
	NewAgent        func(*TCPConn) Agent
	ln              net.Listener
	conns           ConnSet
	mutexConns      sync.Mutex
	wgLn            sync.WaitGroup
	wgConns         sync.WaitGroup

	// msg parser
	LenMsgLen    int
	MinMsgLen    uint32
	MaxMsgLen    uint32
	LittleEndian bool
	//msgParser    *MsgParser
}

func (server *TCPServer) Start() {
	fmt.Println("开始socket服务器")
	server.init()
	go server.run()
}

func (server *TCPServer) init() {
	ln, err := net.Listen("tcp", server.Addr)
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
	server.conns = make(ConnSet)

	// msg parser
	//msgParser := NewMsgParser()
	//msgParser.SetMsgLen(server.LenMsgLen, server.MinMsgLen, server.MaxMsgLen)
	//msgParser.SetByteOrder(server.LittleEndian)
	//server.msgParser = msgParser
}

func (server *TCPServer) run() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()

	var tempDelay time.Duration
	for {
		conn, err := server.ln.Accept()				// 一直在监听，如果有新的连接进来了， 那么就一直走到后面建立新的连接
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

		server.mutexConns.Lock()
		if len(server.conns) >= server.MaxConnNum {
			server.mutexConns.Unlock()
			conn.Close()
			fmt.Println("连接数量超过了最大上限：",server.MaxConnNum)
			continue
		}
		server.conns[conn] = struct{}{}
		server.mutexConns.Unlock()

		server.wgConns.Add(1)

		tcpConn := newTCPConn(conn, server.PendingWriteNum)		// 建立新连接
		agent := server.NewAgent(tcpConn)
		go func() {
			agent.Run()

			// cleanup
			tcpConn.Close()
			server.mutexConns.Lock()
			delete(server.conns, conn)
			server.mutexConns.Unlock()
			agent.OnClose()

			server.wgConns.Done()
			//fmt.Println("当前连接数量为：",len(server.conns))
		}()
	}
}

func (server *TCPServer) Close() {
	server.ln.Close()
	server.wgLn.Wait()

	server.mutexConns.Lock()
	for conn := range server.conns {
		conn.Close()
	}
	server.conns = nil
	server.mutexConns.Unlock()
	server.wgConns.Wait()
}
