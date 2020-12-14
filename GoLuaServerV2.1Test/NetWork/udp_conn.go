package NetWork

import (
	//"GoLuaServerV2.1Test/Utils/log"
	//"github.com/name5566/leaf/zLog"
	"net"
	"sync"
)

//---------------------------------------------------------------------------------------------------
// Socket udp连接部分
// udp 分有连接和无连接两种udp 读写函数不同
// 读统一使用 ReadFromUDP，写则统一使用 WriteMsgUDP,但是地址必须为 nil
//---------------------------------------------------------------------------------------------------

type UdpConnSet map[*net.UDPConn]struct{}

type UdpConn struct {
	sync.Mutex				// 互斥锁 ， 关闭的时候，写入的时候用
	conn      *net.UDPConn
	//writeChan chan []byte
	//closeFlag bool
	UDPAddr *net.UDPAddr
	//msgParser *MsgParser
}

func newUDPConn(conn *net.UDPConn, pendingWriteNum int, pUDPAddr *net.UDPAddr) *UdpConn {
	udpConn := new(UdpConn)
	udpConn.conn = conn
	udpConn.UDPAddr = pUDPAddr
	//udpConn.writeChan = make(chan []byte, pendingWriteNum)

	//udpConn.msgParser = nil

	//go func() {
	//	for b := range udpConn.writeChan {
	//		if b == nil {
	//			//fmt.Println("udpConn.writeChan is null              Quit!")
	//			break
	//		}
	//
	//		_, err := conn.WriteToUDP(b,pUDPAddr)
	//		if err != nil {
	//			log.PrintfLogger("udpConn.writeChan Error 发送数据出错 %s", err.Error())
	//			break
	//		}
	//	}
	//
	//	conn.Close()
	//	udpConn.Lock()
	//	udpConn.closeFlag = true
	//	udpConn.Unlock()
	//}()

	return udpConn
}

func (udpConn *UdpConn) doDestroy() {
	//udpConn.conn.(*net.TCPConn).SetLinger(0)
	udpConn.conn.Close()

	//if !udpConn.closeFlag {
	//	close(udpConn.writeChan)
	//	udpConn.closeFlag = true
	//}
}

func (udpConn *UdpConn) Destroy() {
	udpConn.Lock()
	defer udpConn.Unlock()

	udpConn.doDestroy()
}

func (udpConn *UdpConn) Close() {
	udpConn.Lock()
	defer udpConn.Unlock()
	//if udpConn.closeFlag {
	//	return
	//}
	//
	////udpConn.doWrite(nil)
	//udpConn.closeFlag = true
	//udpConn.doDestroy()
}

//func (udpConn *UdpConn) doWrite(b []byte) {
//	//if len(udpConn.writeChan) > cap(udpConn.writeChan)/2 {
//	//	zLog.PrintfLogger("发送数据包的缓冲区大于1/2!!!")
//	//	time.Sleep(time.Millisecond * 500)
//	//}
//	//if len(udpConn.writeChan) > cap(udpConn.writeChan)*2/3 {
//	//	zLog.PrintfLogger("发送数据包的缓冲区大于2/3!!!")
//	//	time.Sleep(time.Millisecond * 2000)
//	//}
//	if len(udpConn.writeChan) == cap(udpConn.writeChan) {
//		log.PrintfLogger("发送数据包的缓冲区已经满了，关闭该连接!!!")
//		udpConn.doDestroy()
//		return
//	}
//
//	udpConn.writeChan <- b
//}

// b must not be modified by the others goroutines
//func (udpConn *UdpConn) Write(b []byte) {
//	udpConn.Lock()
//	defer udpConn.Unlock()
//	if udpConn.closeFlag || b == nil {
//		return
//	}
//
//	udpConn.doWrite(b)
//}

func (udpConn *UdpConn) Read(b []byte) (int, error) {
	return udpConn.conn.Read(b)
}

func (udpConn *UdpConn) LocalAddr() net.Addr {
	return udpConn.conn.LocalAddr()
}

func (udpConn *UdpConn) RemoteAddr() net.Addr {
	return udpConn.conn.RemoteAddr()
}

func (udpConn *UdpConn) ReadMsg() ([]byte, int, error) {

	msgData := make([]byte, 1024*1)
	//if _, err := io.ReadFull(udpConn.conn, msgData); err != nil {
	//	return nil,0, err
	//}
	//Len:= len(msgData)
	Len,_,err := udpConn.conn.ReadFromUDP(msgData)
	if err != nil {
		return nil,0, err
	}

	return msgData,Len, nil
	//return udpConn.msgParser.Read(udpConn)
}

func (udpConn *UdpConn) WriteMsg(args ...[]byte) error {

	_,_,err :=  udpConn.conn.WriteMsgUDP(args[0],nil, nil)
	if err!=nil{
		println(err.Error())
		return nil
	}
	return nil
}

func (udpConn *UdpConn) GetWriteChanCap() int {
	return  999
}