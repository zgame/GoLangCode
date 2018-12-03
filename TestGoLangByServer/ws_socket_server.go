package main

import (
	"time"
	"math"
	"fmt"
	"./NetWork"
	//"io"
	. "./const"
	"sync"
	"strconv"
)
var clients []*NetWork.TCPClient
var wsclients []*NetWork.WSClient
var receiveMsgNum int		// 接收包数量
var sendMsgNum int			// 发送包的数量
var GlobalMutex sync.Mutex // 全局互斥锁


var StaticDataPackageHeadLess = 0  // 统计信息，数据包 头部数据不全
var StaticDataPackageProtoDataLess = 0  // 统计信息，数据包 pb数据不全
var StaticDataPackagePasteNum = 0   // 统计信息，拼接次数
var StaticDataPackagePasteSuccess = 0   // 统计信息，成功拼接后，解析成功
var StaticDataPackageHeadFlagError = 0   // 统计信息，数据包头部标识不正确



func StartClient(ConnNum int , IsWebSocket bool) {
	//IsWebSocket := false
	if !IsWebSocket {
		// socket client----------------------------------------------------------


		client := new(NetWork.TCPClient)
		client.Addr = GameServerAddress+":"+ strconv.Itoa(SocketPort)
		client.ConnNum = ConnNum
		client.ConnectInterval = 3 * time.Second
		client.PendingWriteNum = 100
		client.LenMsgLen = 4
		client.MaxMsgLen = math.MaxUint32
		client.NewAgent = func(conn *NetWork.TCPConn,index int) NetWork.Agent {
			a := &Client{Conn: conn,Index:index}
			return a
		}

		fmt.Println("开始连接", client.Addr)
		client.Start()
		clients = append(clients, client)
	}
	if IsWebSocket{
		// websocket client------------------------------------------------------------------


		wsclient := new(NetWork.WSClient)
		wsclient.Addr = "ws://"+GameServerAddress+":"+ strconv.Itoa(WebSocketPort)+"/"
		wsclient.ConnNum = ConnNum
		wsclient.ConnectInterval = 3 * time.Second
		wsclient.PendingWriteNum = 100
		wsclient.HandshakeTimeout = 10 * time.Second
		wsclient.MaxMsgLen = math.MaxUint32
		wsclient.NewAgent = func(conn *NetWork.WSConn,index int) NetWork.Agent {
			a := &Client{Conn: conn, Index:index}
			return a
		}

		fmt.Println("开始连接",wsclient.Addr)
		wsclient.Start()
		wsclients = append(wsclients, wsclient)
	}



}



// wsServer.NewAgent 服务器连接的代理
//type Client struct {
//	conn NetWork.Conn
//	//gate     *Gate
//	//userData interface{}
//}

func (a *Client)init()  {
	if a.Index == 0{
		a.ShowMsgSendTime = true	// 第一个才显示
	}
	a.SendTokenID = 1
	a.Gameinfo = a.Gameinfo.New()
	a.loginGS()

	go func() {
		for {
			a.GameAI()
			time.Sleep(time.Millisecond * 200)

		}
	}()
	go func() {
		for {
			if !a.StartAI{
				time.Sleep(time.Millisecond * 100)
				continue
			}
			a.SendMsgTime = a.GetOsTime()
			a.Send("",MDM_GF_GAME, SUB_S_BOSS_COME)
			time.Sleep(time.Millisecond * 2000)

		}
	}()
}


func (a *Client) Run() {
	//a.WriteMsg([]byte("我是客户端哟"))
	a.init()

	for {
		//a.ClientMutex.Lock()
		buf,bufLen, err := a.Conn.ReadMsg()
		//a.ClientMutex.Unlock()

		if err != nil {
			fmt.Println("跟对方的连接中断了", a.Index)
			break
		}
		//if err != nil && err != io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
		//	fmt.Println("接收时候对方服务器链接关闭了！")
		//	//log.Println(err)
		//	this.Conn.Close()
		//
		//	return false
		//}
		if bufLen <= 0{
			fmt.Println("收到的数据为空！", bufLen)
			break
		}
		bufHead := 0
		//num:=0
		for {
			if a.ReceiveBuf !=nil {
				//str:= fmt.Sprintf("%d上次buf: %x ", this.Index,this.ReceiveBuf)
				//this.Zlog(str)
				//str= fmt.Sprintf("%d本次buf: %x ", this.Index,buf)
				//this.Zlog(str)

				GlobalMutex.Lock()
				StaticDataPackagePasteNum++
				GlobalMutex.Unlock()

				buf2 := make([]byte,len(a.ReceiveBuf)+bufLen)		//缓存从新组合包
				copy(buf2, a.ReceiveBuf)
				copy(buf2[len(a.ReceiveBuf):],buf[:bufLen])
				//str= fmt.Sprintf("%d合并后buf2: %x ", this.Index,buf2)
				//this.Zlog(str)
				buf = buf2
				bufLen= len(buf2)
			}
			//fmt.Println(" buf ",buf)
			bufTemp := buf[bufHead:bufLen]   //要处理的buffer
			bufHeadTemp := a.handlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
			bufHead += bufHeadTemp
			time.Sleep(time.Millisecond * 2)

			//fmt.Println("bufHead:",bufHead, " bufLen", bufLen)
			if bufHeadTemp == 0 {
				break		// 包不全，等待下一次继续接受包
			}else if bufHeadTemp > 0 {				// 解析完成
				if a.ReceiveBuf != nil {			// 如果是拼接包，清理一下
					//str := fmt.Sprintf("%d 拼接后成功解析%x", this.Index, buf)
					//this.Zlog(str)
					GlobalMutex.Lock()
					StaticDataPackagePasteSuccess++
					GlobalMutex.Unlock()

					a.ReceiveBuf = nil
				}
			}else if bufHeadTemp == -1 {
				break 		//数据包不正确，放弃
			}

			if bufHead >= bufLen {
				break //解析结束，等待下一次继续接受
			}
		}


	}

}

func (a *Client) OnClose() {

}

func (a *Client) WriteMsg(msg... []byte) {
	err := a.Conn.WriteMsg(msg...)
	if err != nil {
		//fmt.Printf("write message %v error: %v", reflect.TypeOf(msg), err)
	}
	//}
}
//
//func (a *agentClient) LocalAddr() net.Addr {
//	return a.conn.LocalAddr()
//}
//
//func (a *agentClient) RemoteAddr() net.Addr {
//	return a.conn.RemoteAddr()
//}
//
//func (a *agentClient) Close() {
//	a.conn.Close()
//}
//
//func (a *agentClient) Destroy() {
//	a.conn.Destroy()
//}
//
//func (a *agentClient) UserData() interface{} {
//	return a.userData
//}
//
//func (a *agentClient) SetUserData(data interface{}) {
//	a.userData = data
//}
//
