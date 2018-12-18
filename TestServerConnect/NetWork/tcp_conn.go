package NetWork

import (
	//"github.com/name5566/leaf/log"
	"net"
	"sync"
	"../log"
)

//---------------------------------------------------------------------------------------------------
// Socket 连接部分
//---------------------------------------------------------------------------------------------------

type ConnSet map[net.Conn]struct{}

type TCPConn struct {
	sync.Mutex				// 互斥锁 ， 关闭的时候，写入的时候用
	conn      net.Conn
	writeChan chan []byte
	closeFlag bool
	//msgParser *MsgParser
}

func newTCPConn(conn net.Conn, pendingWriteNum int) *TCPConn {
	tcpConn := new(TCPConn)
	tcpConn.conn = conn
	tcpConn.writeChan = make(chan []byte, pendingWriteNum)
	//tcpConn.msgParser = nil

	go func() {
		for b := range tcpConn.writeChan {
			if b == nil {
				//fmt.Println("tcpConn.writeChan is null              Quit!")
				break
			}

			_, err := conn.Write(b)
			if err != nil {
				log.PrintfLogger("tcpConn.writeChan   Error %s", err.Error())
				break
			}
		}

		conn.Close()
		tcpConn.Lock()
		tcpConn.closeFlag = true
		tcpConn.Unlock()
	}()

	return tcpConn
}

func (tcpConn *TCPConn) doDestroy() {
	tcpConn.conn.(*net.TCPConn).SetLinger(0)
	tcpConn.conn.Close()

	if !tcpConn.closeFlag {
		close(tcpConn.writeChan)
		tcpConn.closeFlag = true
	}
}

func (tcpConn *TCPConn) Destroy() {
	tcpConn.Lock()
	defer tcpConn.Unlock()

	tcpConn.doDestroy()
}

func (tcpConn *TCPConn) Close() {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	if tcpConn.closeFlag {
		return
	}

	tcpConn.doWrite(nil)
	tcpConn.closeFlag = true
	//tcpConn.doDestroy()
}

func (tcpConn *TCPConn) doWrite(b []byte) {
	if len(tcpConn.writeChan) == cap(tcpConn.writeChan) {
		log.PrintfLogger("发送数据包的缓冲区已经满了，关闭该连接!!!")
		tcpConn.doDestroy()
		return
	}

	tcpConn.writeChan <- b
}

// b must not be modified by the others goroutines
func (tcpConn *TCPConn) Write(b []byte) {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	if tcpConn.closeFlag || b == nil {
		return
	}

	tcpConn.doWrite(b)
}

func (tcpConn *TCPConn) Read(b []byte) (int, error) {
	return tcpConn.conn.Read(b)
}

func (tcpConn *TCPConn) LocalAddr() net.Addr {
	return tcpConn.conn.LocalAddr()
}

func (tcpConn *TCPConn) RemoteAddr() net.Addr {
	return tcpConn.conn.RemoteAddr()
}

func (tcpConn *TCPConn) ReadMsg() ([]byte, int, error) {

	msgData := make([]byte, 1024*1)
	//if _, err := io.ReadFull(tcpConn.conn, msgData); err != nil {
	//	return nil,0, err
	//}
	//Len:= len(msgData)
	Len,err := tcpConn.conn.Read(msgData)
	if err != nil {
		return nil,0, err
	}

	return msgData,Len, nil
	//return tcpConn.msgParser.Read(tcpConn)
}

func (tcpConn *TCPConn) WriteMsg(args ...[]byte) error {

	//return tcpConn.msgParser.Write(tcpConn, args...)
	//_, err :=tcpConn.conn.Write(args[0])
	//if err != nil {
	//	//log.PrintfLogger("TCPConn .写入错误 %s", err.Error())
	//	return err
	//}

	tcpConn.Write(args[0])
	return nil
}
func (tcpConn *TCPConn) GetWriteChanCap() int {
	return  len(tcpConn.writeChan)
}