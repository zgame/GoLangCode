package main
//**********************************************
// 底层服务器主循环
//**********************************************

import (
	"net"
	"time"
	//"../msg"
	//"../logs"
	"fmt"
)

// 常量
const (
	xReadWriteDeadline = 1e10 // 连接读写等待时间
)

// socket server
type Server struct {
	chClose    chan bool // 关服标识
	*clientMgr           // 客户端管理器
	//msg.Parser           // 消息处理接口
}

//
func NewServer() *Server {
	return &Server{make(chan bool), nil}
}

//
func (s *Server) Serve(lsnAddr string, maxClients int) error {
	// addr
	tcpAddr, e := net.ResolveTCPAddr("tcp", lsnAddr)
	if e != nil {
		return e
	}

	// listen
	listener, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		return e
	}

	// 初始化客户端管理
	s.clientMgr = NewClientMgr(maxClients)

	// 处理连接
	go s.handleClient(listener)

	return nil
}

//
func (s *Server) Stop() {
	close(s.chClose)
	s.clientMgr.Destroy()
}

// 处理连接
func (s *Server) handleClient(listener *net.TCPListener) {
	// log
	fmt.Println("server listen start")
	defer fmt.Println("server listen end")

	// 关闭监听
	defer listener.Close()

	// 协程等待标识
	s.clientMgr.wgClose.Add(1)
	defer s.clientMgr.wgClose.Done()

	for {
		// 一直监听是否有新连接进来
		select {
		case <-s.chClose:
			return

		default:				//非阻塞的轮询
		}

		// debug log
		fmt.Println("wait accept!")

		// 设置超时, 并监听
		listener.SetDeadline(time.Now().Add(xReadWriteDeadline))
		conn, err := listener.Accept()
		if err != nil {
			if e, ok := err.(net.Error); ok && e.Temporary() {
				continue
			}
			fmt.Println(err.Error())
			s.Stop()
			return
		}

		// 创建新的客户端
		client, err := s.clientMgr.createClient(conn)
		if err != nil || nil == client {
			fmt.Println("create new client failed! error:", err, ",client:", client)
			continue
		}

		// log
		fmt.Println("client accepted!id:%v,ip:%v", client.id, conn.RemoteAddr())

		// 处理
		go client.RecvMsg(s.chClose)		// 让该连接的收消息自己去跑
		go client.SendMsg(s.chClose)		// 让该连接的发送消息自己去跑
	}
}