package Client

import (
	//"encoding/binary"
	"fmt"
	"io"
	//"bytes"
	"github.com/golang/protobuf/proto"
	"../CMD"
	. "../Const"
	"../Utils/log"
	"net"
	"../Logic/Player"
	"../Games/Common"
	"../NetWork"
	"../Games"
)

// -------------------------------------------------------------------
// 接收和发送的业务逻辑
// 注意： 客户端发过来带token  ，  服务器发过去不要带token
// -------------------------------------------------------------------

// 全局变量，保存connect的哈希表，方便以后遍历或者查询
var AllClientsList map[*Client]struct{}
var AllUserClientList map[uint32]*Client
//--------------------------------------------------------------------------------------------------
// 结构定义
//--------------------------------------------------------------------------------------------------
type Client struct {
	Conn             net.Conn			// socket连接
	Games            *Games.Games				// 对应游戏的句柄
	Table			 Common.TableInterface		// 桌子的句柄
	User 			 Common.UserInterface	// 用户数据的句柄，包含Player数据
	Player           *Player.Player 	// 玩家的数据库永久数据

}

func (client *Client) NewClient(conn net.Conn)  *Client{
	return &Client{Conn:conn}
}

//--------------------------------------------------------------------------------------------------
// 接受数据包逻辑
//--------------------------------------------------------------------------------------------------

func (client *Client) Receive()  bool{
	buf := make([]byte,1024 * 80) //定义一个切片的长度是1024 * 8
	bufLen,err := client.Conn.Read(buf)
	if err != nil && err == io.EOF {  //io.EOF在网络编程中表示对端把链接关闭了。
		fmt.Println("接收时候对方服务器链接关闭了！")
		//log.Println(err)
		client.Conn.Close()            // 关闭连接
		if client.Player != nil {
			delete(AllUserClientList, client.Player.UserId) // 把自己从所有客户端哈希表中删除
		}
		delete(AllClientsList, client) // 把自己从所有客户端哈希表中删除
		if client.Games != nil {
			client.Games.PlayerLogOutGame(client.User)
		}
		return false
	}
	if bufLen <= 0 || err != nil{
		fmt.Println("收到的数据为！", bufLen)
		fmt.Println("错误为", err.Error())
		client.Conn.Close()            // 关闭连接
		if client.Player != nil {
			delete(AllUserClientList, client.Player.UserId) // 把自己从所有客户端哈希表中删除
		}
		delete(AllClientsList, client) // 把自己从所有客户端哈希表中删除
		if client.Games != nil {
			client.Games.PlayerLogOutGame(client.User)
		}
		return false
	}
	bufHead := 0
	num:=0
	for {
		//fmt.Println(" buf ",buf)
		//fmt.Println(" bufsize ",bufLen)
		bufTemp := buf[bufHead:bufLen]         //要处理的buffer
		bufHead += client.handlerRead(bufTemp) //处理结束之后返回，接下来要开始的范围
		//time.Sleep(time.Millisecond * 100)
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
func (client *Client)Send(sendCmd proto.Message, mainCmd uint16, subCmd uint16, msg string) {
	NetWork.Send(client.Conn,sendCmd,mainCmd,subCmd,msg)
}

//--------------------------------------------------------------------------------------------------
// 发送特定消息逻辑
//--------------------------------------------------------------------------------------------------
func (client *Client)SendGmCmd(cmd string){
	sendCmd := &CMD.CMD_GM_CMD{
		Cmd:[]byte(cmd),
	}
	client.Send(sendCmd, MDM_GF_GMCMD, GMCMD_CMD,"")
}

// 发送不带数据的消息包
func (client *Client) SendCmd(main uint16, sub uint16) {
	_, err := client.Conn.Write(NetWork.GetSendTcpHeaderData(main,sub, 0,0))
	log.CheckError(err)

}

// 同步发给桌子上面的其他玩家
func (client *Client) SendToTableOtherUsers(sendCmd proto.Message, mainCmd uint16, subCmd uint16){
	for _,uid := range client.Table.GetUsersSeatInTable(){
		if uid != int(client.Player.UserId){		// 不是我自己
			client.Send(sendCmd, mainCmd, subCmd,"")
		}
	}
}




//--------------------------------------------------------------------------------------------------
// 加密数据包 token + protocol buffer
//--------------------------------------------------------------------------------------------------
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
