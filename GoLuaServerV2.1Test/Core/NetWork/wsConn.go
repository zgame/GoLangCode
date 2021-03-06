package NetWork

import (
	"github.com/gorilla/websocket"
	"net"
	"sync"
)

//---------------------------------------------------------------------------------------------------
// WebSocket 连接部分
//---------------------------------------------------------------------------------------------------
type WebsocketConnSet map[*websocket.Conn]struct{}

type WSConn struct {
	sync.Mutex			// 互斥锁 ， 关闭的时候，写入的时候用
	conn      *websocket.Conn
	writeChan chan []byte
	maxMsgLen uint32
	closeFlag bool
}

// 创建websocket连接
func newWSConn(conn *websocket.Conn, pendingWriteNum int, maxMsgLen uint32) *WSConn {
	wsConn := new(WSConn)
	wsConn.conn = conn
	wsConn.writeChan = make(chan []byte, pendingWriteNum)
	wsConn.maxMsgLen = maxMsgLen

	go func() {
		for b := range wsConn.writeChan {
			if b == nil {
				//fmt.Println("wsConn.writeChan is null              Quit!")
				break
			}
			err := conn.WriteMessage(websocket.BinaryMessage, b)
			if err != nil {
				println("Conn.WriteMessage  Error  发送数据出错 %s", err.Error())
				break
			}
		}
		conn.Close()
		wsConn.Lock()
		wsConn.closeFlag = true
		wsConn.Unlock()
	}()
	return wsConn
}

func (wsConn *WSConn) doDestroy() {
	wsConn.conn.UnderlyingConn().(*net.TCPConn).SetLinger(0)
	wsConn.conn.Close()

	if !wsConn.closeFlag {
		close(wsConn.writeChan)
		wsConn.closeFlag = true
	}
}

func (wsConn *WSConn) Destroy() {
	wsConn.Lock()
	defer wsConn.Unlock()

	wsConn.doDestroy()
}

func (wsConn *WSConn) Close() {
	wsConn.Lock()
	defer wsConn.Unlock()
	if wsConn.closeFlag {
		return
	}

	wsConn.doWrite(nil)
	wsConn.closeFlag = true
	//wsConn.doDestroy()
}


//将byte数组写入到websocket发送
func (wsConn *WSConn) doWrite(b []byte) {
	if len(wsConn.writeChan) == cap(wsConn.writeChan) {
		println("发送数据包的缓冲区已经满了，关闭该连接!!!")
		wsConn.doDestroy()
		return
	}

	wsConn.writeChan <- b
}

func (wsConn *WSConn) LocalAddr() net.Addr {
	return wsConn.conn.LocalAddr()
}

func (wsConn *WSConn) RemoteAddr() net.Addr {
	return wsConn.conn.RemoteAddr()
}

// goroutine not safe
func (wsConn *WSConn) ReadMsg() ([]byte, int,error) {
	_, b, err := wsConn.conn.ReadMessage()
	Len:= len(b)
	return b, Len,err
}


// args must not be modified by the others goroutines
func (wsConn *WSConn) WriteMsg(args ...[]byte) error {
	wsConn.Lock()
	defer wsConn.Unlock()
	if wsConn.closeFlag {
		return nil
	}

	wsConn.doWrite(args[0])

	return nil
}
func (wsConn *WSConn) GetWriteChanCap() int {
	return  len(wsConn.writeChan)
}