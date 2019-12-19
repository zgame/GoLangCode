package CommonLogin

import (
	"github.com/golang/protobuf/proto"
	//. "./constd"
	"fmt"
	"../../ProtocolBuffer/CMD"
	. "../../Const"
	"../../Core/Utils/zLog"
	//"../../Games"
	//"../../Core/Utils/zRedis"
	"../../Core/ZServer"
	"../../Core/GameCore"
	"../Model/UserModel"
	"../Model/PlayerModel"
	"strconv"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"crypto/md5"

)


//***********************************************游客登录服申请*******************************************
//func (client *Client)SevLoginGuest(buf []byte){
//	fmt.Println("------------客户端申请login登录服务器，创建user信息-------------")
//	// 游戏登录申请，获取玩家信息
//	protocolBuffer := buf
//	msg := &CMD.CMD_MB_GuestLogin{}
//	err := proto.Unmarshal(protocolBuffer, msg)
//	log.CheckError(err)
//
//	fmt.Println("客户端申请login登录服务器, GetMachineId: ",string(msg.GetMachineId()),"DitchNumber:", strconv.Itoa(int(msg.GetDitchNumber())))
//
//	// --------------------游戏类型-------------------
//	fmt.Println("gamekind id :", msg.GetGameKindId())
//	client.Games = GetGameByID(int(msg.GetGameKindId())) //根据客户端要登录的游戏类型，赋值游戏类型指针
//
//	// ----------------------处理用户------------------------
//	// 如果不存在， 就创建新的
//
//	client.PlayerModel = client.PlayerModel.NewUser()
//	client.PlayerModel.FaceId = 0
//	client.PlayerModel.Gender = 0
//	client.PlayerModel.UserId = 2027445
//	client.PlayerModel.GameId = 320395999
//	client.PlayerModel.Exp = 254
//	client.PlayerModel.Loveliness = 0
//	client.PlayerModel.Score = 100000009
//	client.PlayerModel.NickName = "玩家320395999"
//	client.PlayerModel.Level = 1
//	client.PlayerModel.VipLevel = 0
//	client.PlayerModel.AccountLevel = 3
//	client.PlayerModel.SiteLevel = 0
//	client.PlayerModel.CurLevelExp = 0
//	client.PlayerModel.NextLevelExp = 457
//	client.PlayerModel.PayTotal = 0
//	client.PlayerModel.Diamond = 29
//
//	// 发送用户信息给客户端
//	// .....
//	sendCmd := &CMD.CMD_MB_LogonSuccess{
//		FaceId: &(client.PlayerModel.FaceId),
//		Gender: &(client.PlayerModel.Gender),
//		UserId: &(client.PlayerModel.UserId),
//		GameId: &(client.PlayerModel.GameId),
//		Exp: &(client.PlayerModel.Exp),
//		LoveLiness: &(client.PlayerModel.Loveliness),
//		UserScore: &(client.PlayerModel.Score),
//		NickName: []byte(client.PlayerModel.NickName),
//		Level: &(client.PlayerModel.Level),
//		VipLev: &(client.PlayerModel.VipLevel),
//		AccountLevel: &(client.PlayerModel.AccountLevel),
//		Sitelevel: &(client.PlayerModel.SiteLevel),
//		CurLevelExp: &(client.PlayerModel.CurLevelExp),
//		NextLevelExp: &(client.PlayerModel.NextLevelExp),
//		PayTotal: &(client.PlayerModel.PayTotal),
//		UserDiamond: &(client.PlayerModel.Diamond),
//	}
//	client.Send(sendCmd, MDM_MB_LOGON, SUB_MB_LOGON_SUCCESS)
//	client.SendCmd(MDM_MB_LOGON, SUB_MB_LOGON_FINISH) // 发送登录结束
//}

//****************************************************游客登录游戏服申请****************************************************
func HandleLoginGameServerGuest(serverId int,  buf []byte){
	protocolBuffer := buf
	msg := &CMD.CMD_GR_LogonUserID{}
	err := proto.Unmarshal(protocolBuffer, msg)
	zLog.CheckError(err)

	fmt.Println("客户端申请登录游戏服务器, UserId: ",msg.GetUserId())

	// --------------------游戏类型-------------------
	fmt.Println("gamekind id :", msg.GetKindId())
	gameID := int(msg.GetKindId())						 //根据客户端要登录的游戏类型，赋值游戏类型指针


	//uid := 2027445
	uid := int(msg.GetUserId())

	// 从数据库读取玩家信息


	user := UserModel.NewUser()
	user.FaceId = 0
	user.Gender = 0
	user.UserId = uint32(uid)
	user.GameId = 320395999
	user.Exp = 254
	user.Loveliness = 0
	user.Score = 100000009
	user.NickName = "玩家320395999"
	user.Level = 1
	user.VipLevel = 0
	user.AccountLevel = 3
	user.SiteLevel = 0
	user.CurLevelExp = 0
	user.NextLevelExp = 457
	user.PayTotal = 0
	user.Diamond = 29

	player := PlayerModel.NewPlayer(user)
	player.SetUser(user)
	player.SetGameID(gameID)


	GameCore.AddPlayerToAllPlayerList(player)		//把新客户端地址增加到哈希表中保存，方便以后遍历

	//zRedis.SavePlayerToRedis(client.PlayerModel)
	//re := zRedis.GetPlayerFromRedis(int(client.PlayerModel.UserId))
	//fmt.Println("",re.UserId)

	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)

	// ----------------------处理用户------------------------

	// 发送用户信息给客户端
	// .....
	// 判断白名单， 房间状态是否开启
	// 发送中心服请求,中心服判断重复登录，断线重连

	serverid := int32(99099)
	sendCmd := &CMD.CMD_GR_LogonSuccess{
		ServerId: &serverid,
	}

	ZServer.ResisterUID(serverId , uid )
	ZServer.NetWorkSendByUid(uid,sendCmd, MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS)
	ZServer.NetWorkSendByUid(uid,nil, MDM_GR_LOGON, SUB_GR_LOGON_FINISH)
	//client.Send(sendCmd, MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS,"1111111111111111111111玩家账号出现异常，需要重新登录，本次提示为测试提示，客户端需要打印出来，看看是否正常显示！！！")
	//client.SendCmd(MDM_GR_LOGON, SUB_GR_LOGON_FINISH) // 发送登录结束


	//ZServer.NetWorkSendByUid(player.GetUID(), sendCmd,mainCmd,subCmd)
}







//-----------------------------游客密码的生成规则---------------------------------------------
func luaCallGoGetPWD(user_id int, mac string) string {
	pwd := mac + "               " + strconv.Itoa(user_id)
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
	//L.Push(lua.LString(pwdmd5))
	return pwdmd5
}

//-----------------------------------------------游客密码的生成规则---------------------------------------------
//func (client *Client)getPWD(user_id int, mac string) string {
//	pwd := ""
//	//mac := this.build_mac_addr(0)
//	pwd = mac + "               " + strconv.Itoa(user_id)
//
//	buffertt := new(bytes.Buffer)
//	for _,v := range pwd{
//		binary.Write(buffertt, binary.LittleEndian, uint16(v))
//	}
//
//	tokenBuf := buffertt.Bytes()
//	//fmt.Println("", tokenBuf)
//
//	h:= md5.NewUser()
//	h.Write(tokenBuf)
//	cips := h.Sum(nil)			// h.Sum(nil) 将h的hash转成[]byte格式
//	pwdmd5 := hex.EncodeToString(cips)
//	pwdmd5 = strings.ToUpper(pwdmd5)
//	//fmt.Println("pwdmd5: ", pwdmd5)
//
//	return pwdmd5
//}
