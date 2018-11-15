package main
import "github.com/golang/protobuf/proto"
import (
	//. "./constd"
	"fmt"
	"./CMD"
	"./proto_global"
	"strconv"
	"encoding/hex"
	"crypto/md5"
	"bytes"
	"encoding/binary"
	"strings"
	"log"
	"encoding/json"
)


func (this *Client)build_mac_addr(index int)string{
	str := "74-D4-36-AD-"
	i1 := index/256
	i2 := index%256

	ss1 := fmt.Sprintf("%x",i1)
	if len(ss1) ==1{
		str += "0"
	}
	str += ss1
	str += "-"
	ss2 := fmt.Sprintf("%x",i2)
	if len(ss2) ==1{
		str += "0"
	}
	str += ss2

	return str
	//return "74-D4-36-AD-1E-EE"
}


//---------------------------------------游客登录--------------------------------
func (this *Client)LoginSend() {
	dit := int32(9)
	clik := int32(3)
	gamek := uint32(8)
	plaza := uint32(0)
	sendCmd := &CMD.CMD_MB_GuestLogin{
		MachineId:[]byte(this.build_mac_addr(this.Index)),
		DitchNumber: &dit,
		ClientKind:&clik,
		GameKindId:&gamek,
		PlazaVersion:&plaza,
		IpAddr:[]byte("192.168.101.109"),
		DeviceType:[]byte("virtual"),
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_MB_LOGON, SUB_MB_GUESTLOGIN, uint16(size),uint16(this.SendTokenID))  // size要增加一个token的位置
	this.SendTokenID++

	this.Send(bufferT, data)
}

func (this *Client)handleLoginSucess(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_MB_LogonSuccess{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Println("----------------登录成功--------------", this.Index,"+", msg.UserId)
	this.User = this.User.New()
	this.User.face_id = int(msg.GetFaceId())
	this.User.gender = int(msg.GetGender())
	this.User.user_id = int(msg.GetUserId())
	this.User.game_id =int(msg.GetGameId())
	this.User.exp = int(msg.GetExp())
	this.User.loveliness = int(msg.GetLoveLiness())
	this.User.score = int64(msg.GetUserScore())
	this.User.nick_name = string(msg.GetNickName())
	this.User.level = int(msg.GetLevel())
	this.User.vip_level = int(msg.GetVipLev())
	this.User.account_level = int(msg.GetAccountLevel())
	this.User.site_level = int(msg.GetSitelevel())
	this.User.cur_level_exp = int(msg.GetCurLevelExp())
	this.User.next_level_exp = int(msg.GetNextLevelExp())
	this.User.pay_total = int(msg.GetPayTotal())
	this.User.diamond = int(msg.GetUserDiamond())


	//fmt.Println("玩家金币数量", this.User.score)
}

func (this *Client)handleLoginFailed(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_MB_LogonFailure{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	fmt.Println("----------------登录失败--------------", this.Index,"+", msg.ResultCode)
}


// ----------------------------------登录聊天服务器--------------------------------------------------------------

func (this *Client)loginCS() {
	dit := uint32(this.User.user_id)
	git := uint32(this.User.game_id)
	token := ""
	sendCmd := &CMD_GLOBAL.CMD_C_LOGIN{
		UserId:&dit,
		GameId:&git,
		Token:&token,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MAIN_CHAT_CMD, SUB_C_LOGIN, uint16(size),uint16(this.SendTokenID))
	this.SendTokenID++
	this.Send(bufferT, data)
}

// 登录聊天服务器结果
func (this *Client)handleLoginCS(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD_GLOBAL.CMD_S_LOGIN{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	fmt.Println("----------------登录聊天服务器结果  0 是成功--------------", this.Index,"+", msg.GetResult())

}

// 聊天服务器信息
func (this *Client)handleCsInfo(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_CHAT_SERVER_INFO{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Println("----------------聊天服务器信息--------------", this.Index,"+", msg.GetAddr())

}



// -------------------------------------登录成功后下发服务器列表-------------------------------------------

func (this *Client)handleServerList(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_MB_LIST_SERVER{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------获取服务器列表--------------%d   %v",this.Index, msg.GetGameServerList())
	//fmt.Println("")

	//this.Serverlist = make([]*GameServerInfo,len(msg.GetGameServerList()))
	for _,v :=range msg.GetGameServerList(){
		var server *GameServerInfo
		server = NewGameServerInfo()

		server.kind_id = int(v.GetKindId())
		server.node_id = int(v.GetNodeId())
		server.sort_id = int(v.GetSortId())
		server.server_id = int(v.GetServerId())
		server.server_port = int(v.GetServerPort())
		server.online_count = int(v.GetOnlineCount())
		server.online_pc = int(v.GetOnlinePc())
		server.online_andriod = int(v.GetOnlineAndriod())
		server.online_ios = int(v.GetOnlineIos())
		server.full_count = int(v.GetFullCount())
		server.server_name = string(v.GetServerName())
		server.cell_score = int(v.GetCellScore())
		server.min_enter_score = int(v.GetMinEnterScore())
		server.max_enter_score = int(v.GetMaxEnterScore())
		server.server_type = int(v.GetServerType())
		server.min_enter_vip = int(v.GetMinEnterVip())
		server.min_enter_cannon_lev = int(v.GetMinEnterCannonLev())
		server.server_rule = int(v.GetServerRule())
		server.server_addr = string(v.GetServerAddr()[0])

		this.Serverlist = append(this.Serverlist, server)
	}
}

// -------------------------------------游戏服务器登录-------------------------------------------



func (this *Client)loginGS() {

	this.User = this.User.New()
	this.User.user_id = 2027445


	user_id := uint32(this.User.user_id)
	plaza_version := uint32(0)
	frame_version := uint32(101056515)
	process_version := uint32(101056515)
	client_type := int32(3)
	kind_id:= int32(1)				// 游戏类型by1
	ditch_number:= int32(9)
	packet_index:= uint32(1)
	is_android:= false
	cannon_mulriple:= uint32(0)

	sendCmd := &CMD.CMD_GR_LogonUserID{
		UserId:&user_id,
		PlazaVersion:&plaza_version,
		FrameVersion:&frame_version,
		ProcessVersion:&process_version,
		ClientType:&client_type,
		Password:[]byte(""),
		MachineId:[]byte(this.build_mac_addr(this.Index)),
		KindId:&kind_id,
		IpAddr:[]byte("192.168.101.109"),
		DitchNumber:&ditch_number,
		DeviceType:[]byte("virtual"),
		PacketIndex:&packet_index,
		IsAndroid:&is_android,
		CannonMulriple:&cannon_mulriple,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GR_LOGON, SUB_GR_LOGON_USERID, uint16(size),uint16(this.SendTokenID))
	this.SendTokenID++

	this.Send(bufferT, data)
}

// gs 登录成功
func (this *Client)handleLoginSucessGs(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_GR_LogonSuccess{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	dataJ, _ := json.MarshalIndent(msg, "", " ")
	fmt.Printf("%s", dataJ)
	fmt.Println("----------------登录游戏服务器成功--------------", this.Index, msg.GetUserRight())
	this.User.user_id = int(msg.GetUserRight())

}
// gs 登录失败
func (this *Client)handleLoginFailedGs(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_GR_LogonFailure{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Println("GetDescribe", string(msg.GetDescribe()))
	//fmt.Println("GetErrorCode", int(msg.GetErrorCode()))

	log.Println("----------------登录游戏服务器失败-------------------", this.Index, "--------------错误信息------------------", string(msg.GetDescribe()))

}

//-----------------------------游客密码的生成规则---------------------------------------------
func (this *Client)getPWD(user_id int) string {
	pwd := ""
	mac := this.build_mac_addr(this.Index)
	pwd = mac + "               " + strconv.Itoa(user_id)

	buffertt := new(bytes.Buffer)
	for _,v := range pwd{
		binary.Write(buffertt, binary.LittleEndian, uint16(v))
	}

	tokenBuf := buffertt.Bytes()
	//fmt.Println("", tokenBuf)

	h:= md5.New()
	h.Write(tokenBuf)
	cips := h.Sum(nil)			// h.Sum(nil) 将h的hash转成[]byte格式
	pwdmd5 := hex.EncodeToString(cips)
	pwdmd5 = strings.ToUpper(pwdmd5)
	//fmt.Println("pwdmd5: ", pwdmd5)

	return pwdmd5
}
