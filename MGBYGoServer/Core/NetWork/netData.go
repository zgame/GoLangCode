//---------------------------------------------------------------------------------

// 这是服务器的部分，  服务器发给客户端的不要加密， 但是服务器接受要解密 ， 客户端发来带token， 服务器发回去不用token

//---------------------------------------------------------------------------------

package NetWork

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)


// lua 使用的是下面的方法
// 把要发送出去的proto buf -- data， 错误消息 msg， 命令，组合成buffer，准备发送
func DealSendData(sendCmd proto.Message, msg string, mainCmd int, subCmd int,token int) []byte {
	// 服务器不发token
	//buffertt := new(bytes.Buffer)
	//binary.Write(buffertt, binary.LittleEndian, int16(token))
	//tokenBuf := buffertt.Bytes()
	tokenSize :=  0

	//  proto buffer 处理
	protoData, _ := proto.Marshal(sendCmd)
	protoDataSize := len(protoData)
	//// 增加服务器的错误提示msg
	//msgData := []byte(msg)
	//msgSize := len(msgData)

	// 生成数据包头部信息
	bufferHead := getSendTcpHeaderData(uint16(mainCmd), uint16(subCmd), uint16(protoDataSize))
	headSize := len(bufferHead)


	// 开始加密
	bufferData := make([]byte, protoDataSize + tokenSize)
	//copy(bufferData, tokenBuf)
	copy(bufferData[0:], protoData)
	//bufferEncryp := Encryp(bufferData)
	bufferEncryp := bufferData
	//fmt.Printf("发送 buffer : %x \n",bufferData)


	// 发送最后数据包
	bufferEnd := make([]byte, protoDataSize + headSize + tokenSize)
	copy(bufferEnd, bufferHead)// copy 数据包头部
	copy(bufferEnd[headSize:], bufferEncryp)// copy proto buffer 数据
	//copy(bufferEnd[protoDataSize+headSize:protoDataSize+headSize+msgSize], msgData)  // copy msg 数据
	//copy(bufferEnd[protoDataSize+headSize+msgSize:], endData)  // copy  数据包尾部信息



	//// 发送最后数据包
	//bufferEnd := make([]byte, dataSize+  headSize + tokenSize)
	//copy(bufferEnd, bufferH)
	//copy(bufferEnd[headSize:], bufferEncryp)
	//_, err := this.Conn.Write(bufferEnd)
	//checkError(err)

	return bufferEnd
}

// 加密数据包 token + protocol buffer
func Encryp(buffer []byte) []byte{
	buffercryp := make([]byte,len(buffer))
	size := len(buffercryp)
	for i:=0; i<size; i++{
		tmp:= size - i
		//c:= buffer[i]

		//cc,_ := fmt.Printf("%x",buffer[i])
		cc := int(buffer[i])
		cc ^= 0xE9 * tmp + tmp % 14
		cc = cc % 256

		buffercryp[i] = byte(cc)

		//
		//tmp = length-i
		//c = ord(s[i])
		//c ^= 0xE9 * tmp + tmp % 14
		//lst.append(chr(c % 256))
	}
	return buffercryp
}
func Decryp(buffer []byte) []byte{
	buffercryp := make([]byte,len(buffer))
	size := len(buffercryp)
	for i:=0; i<size; i++{
		tmp:= size - i
		//c:= buffer[i]

		//cc,_ := fmt.Printf("%x",buffer[i])
		cc := int(buffer[i])
		cc ^= 0xE9 * tmp + tmp % 14
		//cc = cc % 256

		buffercryp[i] = byte(cc)

		//
		//tmp = length-i
		//c = ord(s[i])
		//c ^= 0xE9 * tmp + tmp % 14
		//lst.append(chr(c % 256))
	}
	return buffercryp
}


//--------------------------------------------------------------------------------------------------
//处理头部数据
//--------------------------------------------------------------------------------------------------
var TCPHeaderSize = 10


// 发送数据包组合成头部信息
func getSendTcpHeaderData(maincmd uint16, childcmd uint16, size uint16) []byte {
	ver := 0           // version 版本0 不带token
	verSize := 2 * ver	 // token占位

	bufferT := new(bytes.Buffer)
	binary.Write(bufferT,binary.LittleEndian,uint8(0))
	binary.Write(bufferT,binary.LittleEndian,uint8(1))
	binary.Write(bufferT,binary.LittleEndian,size + uint16(verSize)) // 注意这里+2 是因为有一个token
	binary.Write(bufferT,binary.LittleEndian,maincmd)
	binary.Write(bufferT,binary.LittleEndian,childcmd)
	binary.Write(bufferT,binary.LittleEndian,uint16(ver))			// 编号1 是要增加一个token

	return bufferT.Bytes()
}

//# 接收数据获取TCPHead头部信息
func DealReceiveTcpDeaderData(msg []byte) (uint16, uint16,uint16,uint16){
	var hh TCPHeader
	buf1 := bytes.NewBuffer(msg[:TCPHeaderSize])
	binary.Read(buf1,binary.LittleEndian,&hh)
	bufferSize := hh.PackSize
	subCmd := hh.SubCMDID
	mainCmd := hh.MainCMDID
	ver := hh.PackerVer
	return mainCmd, subCmd, bufferSize, ver
}

type TCPHeader struct {
	DateKind  uint8  //数据类型
	CheckCode uint8  //效验字段
	PackSize  uint16 //数据大小
	MainCMDID uint16 // 主命令码
	SubCMDID  uint16 // 子命令码
	PackerVer uint16 // 封包版本号
}