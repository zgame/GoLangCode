package NetWork

import (
	"bytes"
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
	sync.Mutex  				// 互斥锁 ， 关闭的时候，写入的时候用
	Conn       *net.UDPConn
	Buffer     bytes.Buffer
	//closeFlag bool
	UDPAddr *net.UDPAddr
	//msgParser *MsgParser
}

func newUDPConn(conn *net.UDPConn, pUDPAddr *net.UDPAddr, data []byte) *UdpConn {
	udpConn := new(UdpConn)
	udpConn.Conn = conn
	udpConn.UDPAddr = pUDPAddr
	udpConn.Buffer.Write(data)

	return udpConn
}

func (udpConn *UdpConn) doDestroy() {
	//udpConn.Conn.(*net.TCPConn).SetLinger(0)
	udpConn.Conn.Close()

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

}

func (udpConn *UdpConn) Read(b []byte) (int, error) {
	return udpConn.Conn.Read(b)
}

func (udpConn *UdpConn) LocalAddr() net.Addr {
	return udpConn.Conn.LocalAddr()
}

func (udpConn *UdpConn) RemoteAddr() net.Addr {
	return udpConn.Conn.RemoteAddr()
}

func (udpConn *UdpConn) ReadMsg() ([]byte, int, error) {

	msgData := make([]byte, 1024*1)
	//if _, err := io.ReadFull(udpConn.Conn, msgData); err != nil {
	//	return nil,0, err
	//}
	//Len:= len(msgData)
	Len,_,err := udpConn.Conn.ReadFromUDP(msgData)
	if err != nil {
		return nil,0, err
	}

	return msgData,Len, nil
	//return udpConn.msgParser.Read(udpConn)
}

func (udpConn *UdpConn) WriteMsg(args ...[]byte) error {

	_,_,err :=  udpConn.Conn.WriteMsgUDP(args[0],nil, nil)
	if err!=nil{
		println(err.Error())
		return nil
	}
	return nil
}

func (udpConn *UdpConn) GetWriteChanCap() int {
	return  999
}