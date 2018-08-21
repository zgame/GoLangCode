package main

import (
	"fmt"
	"net"
	"io"
	log "github.com/jeanphorn/log4go"
	"github.com/go-ini/ini"
	//"os"
	"bytes"
	"encoding/binary"
	"./BY_proto"
	"github.com/golang/protobuf/proto"
	"github.com/mahonia"
	"time"
	"strings"
	"strconv"
)


//--------------------------------------------------------------------------------------------------
//处理头部数据
//--------------------------------------------------------------------------------------------------
func getSendTcpHeaderData(maincmd uint16, childcmd uint16, size uint16) []byte {
	bufferT := new(bytes.Buffer)
	binary.Write(bufferT,binary.LittleEndian,uint8(0))
	binary.Write(bufferT,binary.LittleEndian,uint8(2))
	binary.Write(bufferT,binary.LittleEndian,size)
	binary.Write(bufferT,binary.LittleEndian,maincmd)
	binary.Write(bufferT,binary.LittleEndian,childcmd)
	binary.Write(bufferT,binary.LittleEndian,uint16(0))

	//buffer_t = struct.pack("BBHHHH", 0, 1, size, maincmd, childcmd, 0)
	//fmt.Printf("Send bytes: %b", bufferT.Bytes())
	//fmt.Println("")
	return bufferT.Bytes()
}

//# 获取TCPHead头部信息
func dealRecvTcpDeaderData(msg []byte) (uint16, uint16){
	var hh TCPHeader
	buf1 := bytes.NewBuffer(msg[:10])
	binary.Read(buf1,binary.LittleEndian,&hh)
	bufferSize := hh.PackSize
	subCmd := hh.SubCMDID
	return subCmd, bufferSize
}

//--------------------------------------------------------------------------------------------------
// TCP主循环
//--------------------------------------------------------------------------------------------------

var connectServer net.Conn
var MailOpen bool


func startTcp() {
	fmt.Println("--------------------startTcp--------------------------")
	f, err := ini.Load("Setting.ini")
	if err !=nil{
		fmt.Println("配置文件读取错误！！")
		return
	}
	ip := f.Section("LogServer").Key("ServerIP").Value()
	fmt.Println("ServerIP:",ip)
	port := f.Section("LogServer").Key("port").Value()
	fmt.Println("port:",port)


	MailOpen ,_ = f.Section("Mail").Key("mail_open").Bool()

	//service := "172.16.140.63:8081"
	address := ip+":"+port
	fmt.Println("address:"+address)
	//var err error
	var timer int
	timer = 0

	log.LoadConfiguration("log4go.json")

	for {
		if connectServer == nil {
			fmt.Println("开始连接服务器！！！")
			connectServer, err = net.Dial("tcp", address)
			checkError(err)
			if err != nil {
				fmt.Println("连接服务器失败！！！")
			}else{
				fmt.Println("连接服务器成功！！！")
				_, err = connectServer.Write(getSendTcpHeaderData(MAIN_CMD_ID, SUB_C_MONITOR_REG, 0)) //发送注册成为客户端请求
				checkError(err)
			}
		}else{
			handlerReceiveBuf(connectServer)
			timer++
			if timer > 10 {
				// 10秒一个心跳包
				_,err = connectServer.Write(getSendTcpHeaderData(MAIN_CMD_ID, SUB_C_MONITOR_KEEPLIVE, 0)) //发送心跳包
				if err != nil && err != io.EOF {  	//io.EOF在网络编程中表示对端把链接关闭了。
					fmt.Println("发送时候对方服务器链接关闭了！")
					//log.Println(err)
					connectServer = nil
				}
				timer = 0
			}
			if timer%3 == 1 {
				mwGlobal.model.PublishRowsReset() //5秒刷新一次列表
				mwGlobal.model.PublishRowChanged(0)
				ShowAllServerNum()
			}


		}
		time.Sleep(1000 * time.Millisecond)
	}
	defer connectServer.Close() //断开TCP链接。
}


//--------------------------------------------------------------------------------------------------
// 接收逻辑
//--------------------------------------------------------------------------------------------------
func handlerReceiveBuf(conn net.Conn){
	buf := make([]byte,1024 * 180) //定义一个切片的长度是1024 * 8
	bufLen,err :=conn.Read(buf)
	if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
		fmt.Println("接收时候对方服务器链接关闭了！")
		//log.Println(err)
		connectServer = nil
		return
	}
	if bufLen <= 0{
		fmt.Println("收到的数据为空！")
		return
	}
	bufHead := 0
	num:=0
	for {
		bufTemp := buf[bufHead:bufLen]   //要处理的buffer
		bufHead += handlerRead(bufTemp)   //处理结束之后返回，接下来要开始的范围
		//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
		num++
		//fmt.Println("num",num)
		if bufHead >= bufLen{
			return
		}
	}


}

// 处理单个包内容
func handlerRead(buf []byte) int{
	var err error
	//fmt.Printf("Receive buf: %x",buf)
	//fmt.Println(" ")
	subCmd, bufferSize:=dealRecvTcpDeaderData(buf)


	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")

	if subCmd == SUB_S_MONITOR_ITEMS{
		// 注册成功返回服务器列表
		protocolBuffer := buf[10:10 + bufferSize]
		serverL := &CMD.CMD_MONITOR_ITEM_LST{}
		err = proto.Unmarshal(protocolBuffer, serverL)
		checkError(err)
		//dataJ, _ := json.MarshalIndent(serverL, "", " ")
		//fmt.Printf("%s", dataJ)
		fmt.Println("----------------返回服务器列表--------------")

		// 保存到ServerListAll中
		ServerListAll = make([]ServerState, 0)
		for _,v := range serverL.Items{
			output := convertStringCode(string(v.ServerName))

			ServerListAll = append(ServerListAll, ServerState{
				ServerId:    int(v.ServerId),
				ServerName:  output,
				ServerState: int(v.RoomState),
				Online:0,
			})
		}
		// 刷新listbox
		mwGlobal.model.ResetRows()



	}else if subCmd == SUB_S_MONITOR_KEEPLIVE{
		// 心跳
		//fmt.Println("----------心跳-----------")
		//fmt.Printf("serverlistall: %v", ServerListAll)

	}else if subCmd == SUB_S_MONITOR_STATE{
		// 更新服务器状态
		protocolBuffer := buf[10:10 + bufferSize]
		serverState := &CMD.CMD_MONITOR_ITEM_STATE{}
		err = proto.Unmarshal(protocolBuffer, serverState)
		checkError(err)
		//dataJ, _ := json.MarshalIndent(serverState, "", " ")
		//fmt.Printf("%s", dataJ)
		//fmt.Println("----------刷新服务器状态-----------", serverState.ServerId,"---RoomState--" ,serverState.RoomState, "---Online--" ,serverState.Online)

		if err!= nil {
			// 保存更新到ServerListAll中
			for i, v := range mwGlobal.model.items {
				if v.ServerId == int(serverState.ServerId) {
					//如果id相同，那么更新一下数据
					v.ServerState = int(serverState.RoomState)
					v.Online = int(serverState.Online)
					v.Cpu = int(serverState.Cpu)
					v.Memory = int(serverState.Memory)
					v.IoRead = int(serverState.IoRead)
					v.IoWrite = int(serverState.IoWrite)
					// 刷新listbox
					mwGlobal.model.UpdateRows(v.ServerState, v.Online, v.Cpu, v.Memory, v.IoRead, v.IoWrite, i)

				}
			}
		}

	}else if subCmd == SUB_S_NEW_MONITOR_ITEM{
		// 新增服务器
		protocolBuffer := buf[10:10 + bufferSize]
		newServer := &CMD.CMD_MONITOR_NEW_ITEM{}
		err = proto.Unmarshal(protocolBuffer, newServer)
		checkError(err)
		//dataJ, _ := json.MarshalIndent(newServer, "", " ")
		//fmt.Printf("%s", dataJ)
		fmt.Println("----------新增服务器-----------")
		output := convertStringCode(string(newServer.Item.ServerName))

		ServerListAll = append(ServerListAll, ServerState{
			ServerId:    int(newServer.Item.ServerId),
			ServerName:  output,
			ServerState: int(newServer.Item.RoomState),
		})
		mwGlobal.model.InsertRows(&ServerState{
			ServerId:    int(newServer.Item.ServerId),
			ServerName:  output,
			ServerState: 0,
		})

	} else if subCmd == SUB_S_DEL_MONITOR_ITEM {
		// 减少服务器
		protocolBuffer := buf[10 : 10+bufferSize]
		redServer := &CMD.CMD_MONITOR_DEL_ITEM{}
		err = proto.Unmarshal(protocolBuffer, redServer)
		checkError(err)
		//dataJ, _ := json.MarshalIndent(redServer, "", " ")
		//fmt.Printf("%s", dataJ)
		fmt.Println("----------删除服务器-----------")

		for i, v := range mwGlobal.model.items {
			if v.ServerId == int(redServer.ServerId) {
				// 删掉ServerListAll中的服务器数据
				//ServerListAll = append(ServerListAll[:i], ServerListAll[i+1:]...)
				// 删掉界面显示的listbox数据
				mwGlobal.model.DeleteRows(i)
			} else {
				//fmt.Println("Error----------界面显示服务器名字与后台数据不符合!!!!-------------")
			}
		}
	} else if subCmd == SUB_S_MONITOR_CMD {
		// GM命令
		protocolBuffer := buf[10 : 10+bufferSize]
		gmCMD := &CMD.CMD_MONITOR_CMD_RESP{}
		err = proto.Unmarshal(protocolBuffer, gmCMD)
		checkError(err)
		//dataJ, _ := json.MarshalIndent(gmCMD, "", " ")
		//fmt.Printf("%s", dataJ)
		//fmt.Println("")
	} else if subCmd == SUB_S_MONITOR_LOG {
		// log日志接收
		protocolBuffer := buf[10 : 10+bufferSize]
		logR := &CMD.CMD_MONITOR_LOG{}
		err = proto.Unmarshal(protocolBuffer, logR)
		checkError(err)
		if logR.LogLevel == LogLevelCritical {//|| logR.LogLevel == LogLevelException {
			//dataJ, _ := json.MarshalIndent(logR, "", " ")
			//fmt.Printf("%s", dataJ)
			//fmt.Println("")

			LogText := convertStringCode(string(logR.LogText))
			//fmt.Println("----------------------------------------------------------")
			//fmt.Println("LogText:", LogText)
			//fmt.Println("----------------------------------------------------------")

			// 保存成日志
			//fmt.Printf("logMap:%v",logMap)
			//fmt.Println("")

			if true{ //!checkDuplicateLog(LogText,int(logR.ServerId)){	//判断一下是否重复
				slog := LogText + "--from--"+ strconv.Itoa(int(logR.ServerId))
				log.LOGGER("Exception").Info(slog)
				// 发送警报邮件
				if MailOpen  {
					sendMail(slog)
				}

			}else{
				fmt.Println("重复了，不记录了！")
			}

			//file, _ := os.OpenFile("GameServerError.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
			//logger := log.New(file, "", log.LstdFlags)
			//logger.Println("gameError:"+LogText)
		}
	}

	return int(bufferSize + 10)
}

var logMap 	map[string]string		 //去掉重复日志

// 检查重复日志
func checkDuplicateLog(LogText string, serverId int)  bool{
	var Log string
	// 判断是否有去重标志
	logSplit := strings.Split(LogText,"((")
	if len(logSplit) > 1{
		//说明有去重标志
		split2 := strings.Split(logSplit[1],"))")
		Log = split2[0] + strconv.Itoa(serverId)
	}else{
		//没有去重标志，则全文去重
		Log = LogText
	}
	//fmt.Println("",Log)

	if logMap == nil{
		logMap  = make(map[string]string)
	}
	if logMap[Log] != "has"{
		// 集合没有，那么添加
		logMap[Log] = "has"
		return false
	}
	return true		//重复了
}



// 编码转换
func convertStringCode(old string ) string {
	enc := mahonia.NewDecoder("gb2312")
	// 转换一下编码格式
	output := enc.ConvertString(old)
	return output
}
//--------------------------------------------------------------------------------------------------
// 发送逻辑
//--------------------------------------------------------------------------------------------------
func send_gm_cmd(Serverid int, cmd string){
	sendCmd := &CMD.CMD_MONITOR_CMD{
		ServerId:int32(Serverid),
		ClientId:0,
		Cmd:[]byte(cmd),
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MAIN_CMD_ID, SUB_C_MONITOR_CMD, uint16(size))


	bufferEnd := make([]byte,size+10)
	copy(bufferEnd, bufferT)
	copy(bufferEnd[len(bufferT):], data)
	_, err := connectServer.Write(bufferEnd)
	checkError(err)

}
//--------------------------------------------------------------------------------------------------
// 错误日志处理
//--------------------------------------------------------------------------------------------------

func checkError(e error) {
	if e!=nil{
		fmt.Println("...error:...",e.Error())
		log.LOGGER("Tools Error").Info("-----------------------"+e.Error())

		//file, _ := os.OpenFile("error.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		//logger := log.New(file, "", log.LstdFlags|log.Llongfile)
		//logger.Println("...error:...",e.Error())
	}
}

