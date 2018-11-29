package NetWork

import (
	"net"
)

type Conn interface {
	ReadMsg() ([]byte, int, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
