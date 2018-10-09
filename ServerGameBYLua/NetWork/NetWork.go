package NetWork

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"../Utils/log"
	"net"
)

func Send(Conn net.Conn, sendCmd proto.Message, mainCmd uint16, subCmd uint16, msg string) {
	// 把外面组合好的protocol数据包传进来
	protoData, _ := proto.Marshal(sendCmd)
	protoDataSize := len(protoData)

	// 增加服务器的错误提示msg
	msgData := []byte(msg)
	msgSize :=  len(msgData)

	// 生成数据包头部信息
	bufferHead := GetSendTcpHeaderData(mainCmd, subCmd, uint16(protoDataSize), uint8(len(msgData)))
	headSize :=  len(bufferHead)

	//// 增加token id 作为发送的编号， 因为是服务器端， 所以不加token给客户端
	//token1 := this.SendTokenID
	//this.SendTokenID ++

	// 增加一个token的byte
	//buffertt := new(bytes.Buffer)
	//binary.Write(buffertt, binary.LittleEndian, int16(token1))
	//tokenBuf := buffertt.Bytes()


	// 开始加密
	//bufferData := make([]byte, protoDataSize+msgSize)
	////copy(bufferData, tokenBuf)
	//copy(bufferData[msgSize:], protoData)
	//
	////bufferEncryp := Encryp(bufferData)
	//bufferEncryp := bufferData		// 给客户端发消息不加密


	// 发送最后数据包
	bufferEnd := make([]byte, protoDataSize + headSize + msgSize)
	copy(bufferEnd, bufferHead)		// copy 数据包头部
	copy(bufferEnd[headSize:protoDataSize + headSize], protoData)	// copy protobuffer 数据
	copy(bufferEnd[protoDataSize + headSize:], msgData)		// copy msg 数据
	_, err := Conn.Write(bufferEnd)
	log.CheckError(err)


	//fmt.Printf("send msg: %x", bufferEnd)
	//fmt.Println("")
}


//--------------------------------------------------------------------------------------------------
//处理头部数据
//--------------------------------------------------------------------------------------------------

type TCPHeader struct {
	DateKind  uint8  //数据类型
	CheckCode uint8  //效验字段
	PackSize  uint16 //数据大小
	MainCMDID uint16 // 主命令码
	SubCMDID  uint16 // 子命令码
	PackerVer uint16 // 封包版本号
}


//# 接收消息的时候获取TCPHead头部信息
func DealRecvTcpDeaderData(msg []byte) (uint16, uint16,uint16,uint16){
	var hh TCPHeader
	buf1 := bytes.NewBuffer(msg[:10])
	binary.Read(buf1,binary.LittleEndian,&hh)
	bufferSize := hh.PackSize
	subCmd := hh.SubCMDID
	mainCmd := hh.MainCMDID
	ver := hh.PackerVer
	//msgSize := hh.CheckCode
	return mainCmd, subCmd, bufferSize, ver//,msgSize
}

// 发送消息的时候，组合头部数据
func GetSendTcpHeaderData(maincmd uint16, childcmd uint16, size uint16, msgSize uint8) []byte {

	bufferT := new(bytes.Buffer)
	binary.Write(bufferT,binary.LittleEndian,uint8(0))
	binary.Write(bufferT,binary.LittleEndian,uint8(msgSize))	// msg错误消息的长度
	binary.Write(bufferT,binary.LittleEndian,size)			// 服务器发送不带token
	binary.Write(bufferT,binary.LittleEndian,maincmd)
	binary.Write(bufferT,binary.LittleEndian,childcmd)
	binary.Write(bufferT,binary.LittleEndian,uint16(0))			// 编号1 是要增加一个token ; 0是不带token
	//binary.Write(bufferT,binary.LittleEndian,SendTokenID)			//注意头有2个结构

	//fmt.Printf("Send head bytes: %x", bufferT.Bytes())
	//fmt.Println("")
	return bufferT.Bytes()
}
