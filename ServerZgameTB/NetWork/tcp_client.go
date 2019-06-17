package NetWork

import (
	//"github.com/name5566/leaf/log"
	"net"
	"sync"
	"time"
	"fmt"
)
//---------------------------------------------------------------------------------------------------
// Socket 的客户端代码， 用来做测试用的，服务器用不上
//---------------------------------------------------------------------------------------------------

type TCPClient struct {
	sync.Mutex				// 互斥锁 ，作用就是用来防止多线程的map冲突,  conns 读写操作的时候用
	Addr            string
	ConnNum         int
	ConnectInterval time.Duration
	PendingWriteNum int
	AutoReconnect   bool
	NewAgent        func(*TCPConn,int) Agent
	conns           ConnSet
	wg              sync.WaitGroup
	closeFlag       bool

	// msg parser
	LenMsgLen    int
	MinMsgLen    uint32
	MaxMsgLen    uint32
	LittleEndian bool
	//msgParser    *MsgParser
}

func (client *TCPClient) Number() int{
	return len(client.conns)
}

func (client *TCPClient) Start(start int ,end int) {
	//fmt.Println("start")
	client.init()

	for i := start; i < end; i++ {
		client.wg.Add(1)
		//fmt.Println("for")
		go client.connect(i)
		//time.Sleep(time.Millisecond * 10)
	}
}

func (client *TCPClient) init() {
	//fmt.Println("init")
	client.Lock()
	defer client.Unlock()

	if client.ConnNum <= 0 {
		client.ConnNum = 1
		fmt.Printf("invalid ConnNum, reset to %v", client.ConnNum)
	}
	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 3 * time.Second
		fmt.Printf("invalid ConnectInterval, reset to %v", client.ConnectInterval)
	}
	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 100
		fmt.Printf("invalid PendingWriteNum, reset to %v", client.PendingWriteNum)
	}
	if client.NewAgent == nil {
		fmt.Println("NewAgent must not be nil")
	}
	if client.conns != nil {
		fmt.Println("client is running")
	}

	client.conns = make(ConnSet)
	client.closeFlag = false

	// msg parser
	//msgParser := NewMsgParser()
	//msgParser.SetMsgLen(client.LenMsgLen, client.MinMsgLen, client.MaxMsgLen)
	//msgParser.SetByteOrder(client.LittleEndian)
	//client.msgParser = msgParser
}

func (client *TCPClient) dial() net.Conn {
	for {
		conn, err := net.Dial("tcp", client.Addr)
		if err == nil || client.closeFlag {
			return conn
		}

		fmt.Printf("connect to %v error: %v \n", client.Addr, err)
		time.Sleep(client.ConnectInterval)
		continue
	}
}

func (client *TCPClient) connect(index int) {

	//fmt.Println("开始连接...")
	defer client.wg.Done()

reconnect:
	conn := client.dial()
	if conn == nil {
		return
	}

	client.Lock()
	if client.closeFlag {
		client.Unlock()
		conn.Close()
		return
	}
	client.conns[conn] = struct{}{}
	client.Unlock()

	tcpConn := newTCPConn(conn, client.PendingWriteNum)
	agent := client.NewAgent(tcpConn,index)
	agent.Run()

	// cleanup
	tcpConn.Close()
	client.Lock()
	delete(client.conns, conn)
	client.Unlock()
	agent.OnClose()

	if client.AutoReconnect {
		time.Sleep(client.ConnectInterval)
		goto reconnect
	}
}

func (client *TCPClient) Close() {
	client.Lock()
	client.closeFlag = true
	for conn := range client.conns {
		conn.Close()
	}
	client.conns = nil
	client.Unlock()

	client.wg.Wait()
}
