package NetWork

import (
	"bytes"
	"encoding/binary"
)

// 这个函数在lua中没有使用
//func Send(Conn net.Conn, sendCmd proto.Message, mainCmd uint16, subCmd uint16, msg string) {
//	// 把外面组合好的protocol数据包传进来
//	protoData, _ := proto.Marshal(sendCmd)
//	protoDataSize := len(protoData)
//
//	// 增加服务器的错误提示msg
//	msgData := []byte(msg)
//	msgSize :=  len(msgData)
//
//	// 生成数据包头部信息
//	bufferHead := GetSendTcpHeaderData(mainCmd, subCmd, uint16(protoDataSize), uint8(len(msgData)))
//	headSize :=  len(bufferHead)
//
//	//// 增加token id 作为发送的编号， 因为是服务器端， 所以不加token给客户端
//	//token1 := this.SendTokenID
//	//this.SendTokenID ++
//
//	// 增加一个token的byte
//	//buffertt := new(bytes.Buffer)
//	//binary.Write(buffertt, binary.LittleEndian, int16(token1))
//	//tokenBuf := buffertt.Bytes()
//
//
//	// 开始加密
//	//bufferData := make([]byte, protoDataSize+msgSize)
//	////copy(bufferData, tokenBuf)
//	//copy(bufferData[msgSize:], protoData)
//	//
//	////bufferEncryp := Encryp(bufferData)
//	//bufferEncryp := bufferData		// 给客户端发消息不加密
//
//
//	// 发送最后数据包
//	bufferEnd := make([]byte, protoDataSize + headSize + msgSize)
//	copy(bufferEnd, bufferHead)		// copy 数据包头部
//	copy(bufferEnd[headSize:protoDataSize + headSize], protoData)	// copy protobuffer 数据
//	copy(bufferEnd[protoDataSize + headSize:], msgData)		// copy msg 数据
//	_, err := Conn.Write(bufferEnd)
//	log.CheckError(err)
//
//
//	//fmt.Printf("send msg: %x", bufferEnd)
//	//fmt.Println("")
//}


// lua 使用的是下面的方法
// 把要发送出去的proto buf -- data， 错误消息 msg， 命令，组合成buffer，准备发送
func DealSendData(data string, msg string, mainCmd int, subCmd int,token int) []byte {
	// 组合一下str
	protoData := []byte(data)
	protoDataSize := len(protoData)
	// 增加服务器的错误提示msg
	msgData := []byte(msg)
	msgSize := len(msgData)

	// 生成数据包头部信息
	bufferHead := GetSendTcpHeaderData(uint16(mainCmd), uint16(subCmd), uint16(protoDataSize), uint16(msgSize),uint16(token))
	headSize := len(bufferHead)
	// 发送最后数据包
	bufferEnd := make([]byte, protoDataSize+headSize+msgSize)
	copy(bufferEnd, bufferHead)
	// copy 数据包头部
	copy(bufferEnd[headSize:protoDataSize+headSize], protoData)
	// copy protobuffer 数据
	copy(bufferEnd[protoDataSize+headSize:], msgData)
	// copy msg 数据
	return bufferEnd
}


//--------------------------------------------------------------------------------------------------
//处理头部数据
//--------------------------------------------------------------------------------------------------
var TCPHeaderSize = 11

type TCPHeader struct {
	DateKind  uint8  //数据类型			始终是0
	CheckCode uint16  //效验字段			msg错误消息的长度
	PackSize  uint16 //数据大小			数据包的大小
	MainCMDID uint16 // 主命令码
	SubCMDID  uint16 // 子命令码
	PackerVer uint16 // 封包版本号		数据包的一个客户端编号，用来防止重复的，客户端可以自己记录，然后到9999清零
}


//# 接收消息的时候获取TCPHead头部信息
func DealRecvTcpDeaderData(msg []byte) (uint8,uint16, uint16,uint16,uint16,uint16){
	var hh TCPHeader
	buf1 := bytes.NewBuffer(msg[:TCPHeaderSize])
	binary.Read(buf1,binary.LittleEndian,&hh)
	bufferSize := hh.PackSize
	subCmd := hh.SubCMDID
	mainCmd := hh.MainCMDID
	ver := hh.PackerVer
	msgSize := hh.CheckCode
	HeadFlag := hh.DateKind
	return 	HeadFlag,mainCmd, subCmd, bufferSize, ver,msgSize
}


//---------------------------------------------------------------
// 发送消息的时候，组合头部数据
// 00 0000 0700 0700 6900 0000 0a0508331099010000070007006900
// 00 0000 0700 0700 6900 0000 0a05082f10970100000007000700690000000a050830109801000000070007006
// 0 - msgSize - size -  maincmd - childcmd - 0
//---------------------------------------------------------------

func GetSendTcpHeaderData(maincmd uint16, childcmd uint16, size uint16, msgSize uint16, token uint16) []byte {

	bufferT := new(bytes.Buffer)
	binary.Write(bufferT,binary.LittleEndian,uint8(254))
	binary.Write(bufferT,binary.LittleEndian,msgSize)		// msg错误消息的长度
	binary.Write(bufferT,binary.LittleEndian,size)			// 服务器发送不带token
	binary.Write(bufferT,binary.LittleEndian,maincmd)
	binary.Write(bufferT,binary.LittleEndian,childcmd)
	binary.Write(bufferT,binary.LittleEndian,uint16(token))			// 服务器给客户端不用带token
	//binary.Write(bufferT,binary.LittleEndian,SendTokenID)			//注意头有2个结构

	//fmt.Printf("Send head bytes: %x", bufferT.Bytes())
	//fmt.Println("")
	return bufferT.Bytes()
}
