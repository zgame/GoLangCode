package client

import (
	"net"
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

type Client struct {
	Conn net.Conn
	Index int
}

func (this *Client) quit() {
	this.Conn.Close()
}
func (this *Client) Receive(isHeadLoaded bool) (bool , []byte){
	bodyLen := 0
	//var rebodySl []byte
	//reader := bufio.NewReader(this.Conn)

	//if !isHeadLoaded {
		headLenSl := make([]byte, 4)
		fmt.Println("读取包头....")

		//已经读取的包头字节数
		//readedHeadLen := 0

		//for readedHeadLen < 4 {
		_, err := io.ReadFull(this.Conn, headLenSl)

		if err != nil {
			fmt.Println("读取包头出错, ", err.Error())
			this.quit()
			return false, nil
		}
		//	readedHeadLen += len
		//}

		bodyLen = int(binary.BigEndian.Uint32(headLenSl))
		fmt.Println("读取包头成功, 包体字节长度: ", bodyLen)
	//	isHeadLoaded = true
	//}

	//if isHeadLoaded {
		fmt.Println("解析包体")
		bodySl := make([]byte, bodyLen)

		//已经读取的包体字节数
		readedBodyLen := 0
		//for readedBodyLen < bodyLen {
		//len, err := reader.Read(bodySl)
		_, err = io.ReadFull(this.Conn, bodySl)

		if err != nil {
			fmt.Println("读取包体出错,: ", err.Error())
			return false, nil
		}
		//fmt.Println("buff: %v",bodySl)
		//readedBodyLen += len
		//}
		//rebodySl = bodySl
		fmt.Println("读取包体完成,包体字节长度: ", readedBodyLen)
	//	isHeadLoaded = false
	//}
	return isHeadLoaded,bodySl
}

func (this *Client) Send(msg string) {
	writer := bufio.NewWriter(this.Conn)
	msgLen := len(msg)
	//写入2个字节字符串长度.以供flash读取便利
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(msgLen))
	writer.Write(buf)
	writer.WriteString(msg)
	writer.Flush()

	//fmt.Println("send buf %v",buf)

}
