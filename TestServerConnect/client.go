package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"bytes"
	"os"
	"log"
	"github.com/golang/protobuf/proto"
	"./CMD"
	. "./const"
	"time"
)



//func (this *Client) quit() {
//	this.Conn.Close()
//}

//--------------------------------------------------------------------------------------------------
// 接受数据包逻辑
//--------------------------------------------------------------------------------------------------

func (this *Client) Receive()  bool{
	buf := make([]byte,1024 * 80) //定义一个切片的长度是1024 * 8
	bufLen,err := this.Conn.Read(buf)
	if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
		fmt.Println("接收时候对方服务器链接关闭了！")
		this.Quit = true
		//log.Println(err)
		this.Conn.Close()
		return false
	}
	if bufLen <= 0{
		this.Conn.Close()        // 关闭连接
		fmt.Println("收到的数据为空！", bufLen)
		this.Quit = true
		return false
	}
	bufHead := 0
	num:=0
	for {
		//fmt.Println(" buf ",buf)
		if this.Quit {
			return false
		}

		bufTemp := buf[bufHead:bufLen]   //要处理的buffer
		bufHead += this.handlerRead(bufTemp)   //处理结束之后返回，接下来要开始的范围
		time.Sleep(time.Millisecond * 100)
		//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
		num++
		//fmt.Println("num",num)
		if bufHead >= bufLen{
			return true
		}
	}

}
//--------------------------------------------------------------------------------------------------
// 发送消息逻辑， 逻辑需要组合好数据的格式， 这里只管发送
//--------------------------------------------------------------------------------------------------
func (this *Client) Send(bufferH []byte, data []byte) {
	// 增加token id 作为发送的编号
	token1 := this.SendTokenID
	this.SendTokenID ++

	buffertt := new(bytes.Buffer)
	binary.Write(buffertt, binary.LittleEndian, int16(token1))
	tokenBuf := buffertt.Bytes()


	dataSize :=  len(data)
	headSize :=  len(bufferH)
	tokenSize :=  len(tokenBuf)


	// 开始加密
	bufferData := make([]byte, dataSize + tokenSize)
	copy(bufferData, tokenBuf)
	copy(bufferData[tokenSize:], data)
	bufferEncryp := Encryp(bufferData)


	// 发送最后数据包
	bufferEnd := make([]byte, dataSize+  headSize + tokenSize)
	copy(bufferEnd, bufferH)
	copy(bufferEnd[headSize:], bufferEncryp)
	_, err := this.Conn.Write(bufferEnd)
	checkError(err)


	//fmt.Printf("send msg: %x", bufferEnd)
	//fmt.Println("")
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



//--------------------------------------------------------------------------------------------------
// 发送特定消息逻辑
//--------------------------------------------------------------------------------------------------
func (this *Client)SendGmCmd(cmd string){
	sendCmd := &CMD.CMD_GM_CMD{
		Cmd:[]byte(cmd),
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_GMCMD, GMCMD_CMD, uint16(size))

	this.Send(bufferT, data)


}

func (this *Client) SendReg(msg string) {

	_, err := this.Conn.Write(getSendTcpHeaderData(MAIN_CMD_ID, SUB_C_MONITOR_REG, 0)) //发送注册成为客户端请求
	checkError(err)

}


//--------------------------------------------------------------------------------------------------
// 错误日志处理
//--------------------------------------------------------------------------------------------------

func checkError(e error) {
	if e!=nil{
		file, _ := os.OpenFile("error.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		logger := log.New(file, "", log.LstdFlags|log.Llongfile)
		logger.Println("...error:...",e.Error())
	}
}



//--------------------------------------------------------------------------------------------------
//处理头部数据
//--------------------------------------------------------------------------------------------------

// 发送数据包组合成头部信息
func getSendTcpHeaderData(maincmd uint16, childcmd uint16, size uint16) []byte {

	bufferT := new(bytes.Buffer)
	binary.Write(bufferT,binary.LittleEndian,uint8(0))
	binary.Write(bufferT,binary.LittleEndian,uint8(1))
	binary.Write(bufferT,binary.LittleEndian,size+2)		// 注意这里+2 是因为有一个token
	binary.Write(bufferT,binary.LittleEndian,maincmd)
	binary.Write(bufferT,binary.LittleEndian,childcmd)
	binary.Write(bufferT,binary.LittleEndian,uint16(1))			// 编号1 是要增加一个token
	//binary.Write(bufferT,binary.LittleEndian,SendTokenID)			//注意头有2个结构


	//buffer_t = struct.pack("BBHHHH", 0, 1, size, maincmd, childcmd, 0)
	//fmt.Printf("Send head bytes: %x", bufferT.Bytes())
	//fmt.Println("")
	return bufferT.Bytes()
}

//# 接收数据获取TCPHead头部信息
func dealRecvTcpDeaderData(msg []byte) (uint16, uint16,uint16,uint16){
	var hh TCPHeader
	buf1 := bytes.NewBuffer(msg[:10])
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