package NetWork

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	//"github.com/name5566/leaf/zLog"
	"net"
	"net/http"
	"sync"
	"time"
)
//---------------------------------------------------------------------------------------------------
// WebSocket 服务器部分
//---------------------------------------------------------------------------------------------------



type WSServer struct {
	Addr            string		// 服务器地址
	MaxConnNum      int			// 最大连接数量
	PendingWriteNum int
	MaxMsgLen       uint32			// 消息的最大长度
	HTTPTimeout     time.Duration		// timeout
	CertFile        string
	KeyFile         string
	NewAgent        func(*WSConn) Agent // 代理回调，提供run函数，保存conn
	ln              net.Listener
	handler         *WSHandler
}

type WSHandler struct {
	maxConnNum      int			// 最大连接数量
	pendingWriteNum int
	maxMsgLen       uint32              // 消息的最大长度
	newAgent        func(*WSConn) Agent // 代理回调，提供run函数，保存conn
	upgrader        websocket.Upgrader
	conns           WebsocketConnSet
	mutexConns      sync.Mutex				// 互斥锁， 用在保持多线程对map的操作安全上
	wg              sync.WaitGroup
}

// webSocket server Conn 每当有新用户连接， 就会调用一次该函数
func (handler *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("new  ServeHTTP")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("upgrade error: %v", err)
		return
	}
	conn.SetReadLimit(int64(handler.maxMsgLen))

	handler.wg.Add(1)
	defer handler.wg.Done()

	handler.mutexConns.Lock()
	if handler.conns == nil {
		handler.mutexConns.Unlock()
		conn.Close()
		return
	}
	if len(handler.conns) >= handler.maxConnNum {		// 超过最大的连接数量了
		handler.mutexConns.Unlock()
		conn.Close()
		fmt.Println("已经达到最大的连接数量限制",handler.maxConnNum)
		return
	}
	handler.conns[conn] = struct{}{}
	handler.mutexConns.Unlock()

	wsConn := newWSConn(conn, handler.pendingWriteNum, handler.maxMsgLen) // 创建连接
	agent := handler.newAgent(wsConn)
	agent.Run()								// 代理run函数


	// cleanup
	wsConn.Close()						// 关闭连接
	handler.mutexConns.Lock()
	delete(handler.conns, conn)			// 去掉当前连接列表中这个连接
	handler.mutexConns.Unlock()
	agent.OnClose()						// 关闭代理
	//fmt.Println("当前连接数量为：",len(handler.conns))

}

// 服务器的启动
func (server *WSServer) Start() {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Printf("开始服务器出错，%v", err)
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		fmt.Printf("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}
	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = 100
		fmt.Printf("invalid PendingWriteNum, reset to %v", server.PendingWriteNum)
	}
	if server.MaxMsgLen <= 0 {
		server.MaxMsgLen = 4096
		fmt.Printf("invalid MaxMsgLen, reset to %v", server.MaxMsgLen)
	}
	if server.HTTPTimeout <= 0 {
		server.HTTPTimeout = 10 * time.Second
		fmt.Printf("invalid HTTPTimeout, reset to %v", server.HTTPTimeout)
	}
	if server.NewAgent == nil {
		fmt.Printf("指定用那种服务器的代理不能为空啊")
	}

	if server.CertFile != "" || server.KeyFile != "" {
		config := &tls.Config{}
		config.NextProtos = []string{"http/1.1"}

		var err error
		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = tls.LoadX509KeyPair(server.CertFile, server.KeyFile)
		if err != nil {
			fmt.Printf("%v", err)
		}

		ln = tls.NewListener(ln, config)
	}

	server.ln = ln
	server.handler = &WSHandler{
		maxConnNum:      server.MaxConnNum,
		pendingWriteNum: server.PendingWriteNum,
		maxMsgLen:       server.MaxMsgLen,
		newAgent:        server.NewAgent,
		conns:           make(WebsocketConnSet),
		upgrader: websocket.Upgrader{
			HandshakeTimeout: server.HTTPTimeout,
			CheckOrigin:      func(_ *http.Request) bool { return true },
		},
	}

	httpServer := &http.Server{
		Addr:           server.Addr,
		Handler:        server.handler,
		ReadTimeout:    server.HTTPTimeout,
		WriteTimeout:   server.HTTPTimeout,
		MaxHeaderBytes: 1024,
	}

	fmt.Println("开始websocket服务器")
	go httpServer.Serve(ln)				// 开始服务器的协程
}

func (server *WSServer) Close() {
	server.ln.Close()

	server.handler.mutexConns.Lock()
	for conn := range server.handler.conns {
		conn.Close()
	}
	server.handler.conns = nil
	server.handler.mutexConns.Unlock()

	server.handler.wg.Wait()
}
