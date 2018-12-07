package main

import (
	"./NetWork"

	//"io"
	//"bytes"
	"os"
	"log"
	//"github.com/golang/protobuf/proto"
	//"./CMD"
	//. "./const"
	"time"

	"sync"
)

type Client struct {

	Conn       NetWork.Conn
	Index      int
	User       *User
	Serverlist []*GameServerInfo
	Gameinfo   * GameInfo
	SendTokenID int32
	StartAI 	bool
	Last_fire_tick time.Time		// 开火CD
	Select_tick time.Time			// 选鱼CD
	Last_check_due_tick time.Time    // 定期清理过期鱼
	Fish_id	int				//锁定的鱼id
	Failed_cnt int			// 锁定鱼打了几炮

	ShowLog uint64 			//打鱼的记录
	SendMsgTime int64	// 发送消息时间
	ShowMsgSendTime bool  // 是否显示
	ShowMsgFire int64  // 子弹平均发多少颗
	ShowMsgReFire int64  // 子弹平均回复多少颗
	ShowMsgCatchFish int64  // 打多少条鱼了
	ShowMsgReCatchFish int64  // 打多少条鱼了


	ReceiveBuf  []byte
	ClientMutex sync.Mutex // 用来保证发送和接收的线程安全的

	ReceiveMsgNum int		// 接收包数量
	SendMsgNum int			// 发送包的数量
	SuccessBuf []byte
	LastBuf []byte
}

func (this *Client) quit() {
	this.Conn.Close()
}



//--------------------------------------------------------------------------------------------------
// 接受数据包逻辑
//--------------------------------------------------------------------------------------------------
//
//func (this *Client) Receive()  bool{
//	if this.Conn == nil {
//		return false
//	}
//
//	//GlobalMutex.Lock()
//	buf := make([]byte,1024 * 1) //定义一个切片的长度是1024 * 8
//	bufLen,err := this.Conn.Read(buf)
//
//	//_,err := io.ReadFull(this.Conn, buf)
//
//	//buf,err:= ioutil.ReadAll(this.Conn)
//	//bufLen:= len(buf)
//
//	//GlobalMutex.Unlock()
//
//	if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
//		fmt.Println("接收时候对方服务器链接关闭了！")
//		//log.Println(err)
//		this.Conn.Close()
//
//		return false
//	}
//	if bufLen <= 0{
//		fmt.Println("收到的数据为空！", bufLen)
//		return false
//	}
//	bufHead := 0
//	//num:=0
//	for {
//		if this.ReceiveBuf !=nil {
//			//str:= fmt.Sprintf("%d上次buf: %x ", this.Index,this.ReceiveBuf)
//			//this.PrintLogger(str)
//			//str= fmt.Sprintf("%d本次buf: %x ", this.Index,buf)
//			//this.PrintLogger(str)
//
//			buf2 := make([]byte,len(this.ReceiveBuf)+bufLen)		//缓存从新组合包
//			copy(buf2, this.ReceiveBuf)
//			copy(buf2[len(this.ReceiveBuf):],buf[:bufLen])
//			//str= fmt.Sprintf("%d合并后buf2: %x ", this.Index,buf2)
//			//this.PrintLogger(str)
//			buf = buf2
//			bufLen= len(buf2)
//		}
//		//fmt.Println(" buf ",buf)
//		bufTemp := buf[bufHead:bufLen]   //要处理的buffer
//		bufHeadTemp := this.handlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
//		bufHead += bufHeadTemp
//		time.Sleep(time.Millisecond * 2)
//
//		//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
//		if bufHeadTemp == 0 {
//			return true 		// 包不全，等待下一次继续接受包
//		}else if bufHeadTemp > 0 {				// 解析完成
//			if this.ReceiveBuf != nil {			// 如果是拼接包，清理一下
//				//str := fmt.Sprintf("%d 拼接后成功解析%x", this.Index, buf)
//				//this.PrintLogger(str)
//				this.ReceiveBuf = nil
//			}
//		}else if bufHeadTemp == -1 {
//			return true 		//数据包不正确，放弃
//		}
//
//		if bufHead >= bufLen {
//			return true //解析结束，等待下一次继续接受
//		}
//	}
//
//
//}
//--------------------------------------------------------------------------------------------------
// 发送消息逻辑， 逻辑需要组合好数据的格式， 这里只管发送
//--------------------------------------------------------------------------------------------------
func (this *Client) Send(data string, mainCmd int, subCmd int) {
	// 增加token id 作为发送的编号
	//token1 := this.SendTokenID
	//this.SendTokenID ++

	//buffertt := new(bytes.Buffer)
	//binary.Write(buffertt, binary.LittleEndian, int16(token1))
	//tokenBuf := buffertt.Bytes()


	//dataSize :=  len(data)
	//headSize :=  len(bufferH)
	////tokenSize :=  len(tokenBuf)
	//
	//
	////// 开始加密
	////bufferData := make([]byte, dataSize + tokenSize)
	////copy(bufferData, tokenBuf)
	////copy(bufferData[tokenSize:], data)
	//
	////fmt.Printf("send buf: %x",bufferData)
	////fmt.Println("")
	////bufferEncryp := Encryp(bufferData)
	//bufferEncryp := (data)
	//
	//
	//// 发送最后数据包
	////fmt.Println("数据包大小：", strconv.Itoa(dataSize))
	////fmt.Println("token大小：", strconv.Itoa( tokenSize))
	//bufferEnd := make([]byte, dataSize+  headSize )
	//copy(bufferEnd, bufferH)
	//copy(bufferEnd[headSize:], bufferEncryp)
	//
	//_, err := this.Conn.Write(bufferEnd)
	GlobalMutex.Lock()
	this.SendMsgNum++
	GlobalMutex.Unlock()


	bufferEnd := NetWork.DealSendData(data , "" , mainCmd , subCmd ,int(this.SendTokenID))

	this.ClientMutex.Lock()
	this.WriteMsg(bufferEnd)		// 这里要加锁，因为多个线程可能同时进行发送消息
	this.SendTokenID++
	this.ClientMutex.Unlock()
	//checkError(err)

	time.Sleep(time.Millisecond * 20)
	//fmt.Printf("send msg: %x", bufferEnd)
	//fmt.Println("")
}

//// 加密数据包 token + protocol buffer
//func Encryp(buffer []byte) []byte{
//	buffercryp := make([]byte,len(buffer))
//	size := len(buffercryp)
//	for i:=0; i<size; i++{
//		tmp:= size - i
//		//c:= buffer[i]
//
//		//cc,_ := fmt.Printf("%x",buffer[i])
//		cc := int(buffer[i])
//		cc ^= 0xE9 * tmp + tmp % 14
//		cc = cc % 256
//
//		buffercryp[i] = byte(cc)
//
//		//
//		//tmp = length-i
//		//c = ord(s[i])
//		//c ^= 0xE9 * tmp + tmp % 14
//		//lst.append(chr(c % 256))
//	}
//	return buffercryp
//}



//--------------------------------------------------------------------------------------------------
// 发送特定消息逻辑
//--------------------------------------------------------------------------------------------------
//func (this *Client)SendGmCmd(cmd string){
//	sendCmd := &CMD.CMD_GM_CMD{
//		Cmd:[]byte(cmd),
//	}
//	data, _ := proto.Marshal(sendCmd)
//	size := len(data)
//	bufferT := getSendTcpHeaderData(MDM_GF_GMCMD, GMCMD_CMD, uint16(size), uint16(this.SendTokenID))
//	this.SendTokenID ++
//
//	this.Send(bufferT, data)
//
//
//}

//func (this *Client) SendReg(msg string) {
//
//	_, err := this.Conn.Write(getSendTcpHeaderData(MAIN_CMD_ID, SUB_C_MONITOR_REG, 0,uint16(this.SendTokenID))) //发送注册成为客户端请求
//	this.SendTokenID++
//	checkError(err)
//
//}


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


func (this *Client)GetOsTime()  int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//////--------------------------------------------------------------------------------------------------
//////处理头部数据
//////--------------------------------------------------------------------------------------------------
//func getSendTcpHeaderData(maincmd uint16, childcmd uint16, size uint16 , token uint16) []byte {
//
//	bufferT := new(bytes.Buffer)
//	binary.Write(bufferT,binary.LittleEndian,uint8(254))
//	binary.Write(bufferT,binary.LittleEndian,uint16(0))
//	binary.Write(bufferT,binary.LittleEndian,size)
//	binary.Write(bufferT,binary.LittleEndian,maincmd)
//	binary.Write(bufferT,binary.LittleEndian,childcmd)
//	binary.Write(bufferT,binary.LittleEndian,token)			// 是要增加一个token
//
//
//	//buffer_t = struct.pack("BBHHHH", 0, 1, size, maincmd, childcmd, 0)
//	//fmt.Printf("Send head bytes: %x", bufferT.Bytes())
//	//fmt.Println("")
//	return bufferT.Bytes()
//}
//var TCPHeaderSize = 11
//
////# 获取TCPHead头部信息
//func dealRecvTcpDeaderData(msg []byte) (uint8,uint16, uint16,uint16,uint16,uint16){
//	var hh TCPHeader
//	buf1 := bytes.NewBuffer(msg[:TCPHeaderSize])
//	binary.Read(buf1,binary.LittleEndian,&hh)
//	HeadFlag := hh.DateStart
//	bufferSize := hh.PackSize
//	subCmd := hh.SubCMDID
//	mainCmd := hh.MainCMDID
//	ver := hh.PackerVer
//	msgSize := hh.CheckCode
//	return HeadFlag,mainCmd, subCmd, bufferSize, ver,msgSize
//}
//
