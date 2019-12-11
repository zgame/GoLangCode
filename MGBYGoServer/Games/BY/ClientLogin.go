package BY

import (
	"github.com/golang/protobuf/proto"
	//. "./constd"
	"fmt"
	"../CMD"
	. "../Const"
	"../Utils/log"
	"../Games"
	"../Utils/zRedis"
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
//	client.Player = client.Player.NewPlayer()
//	client.Player.FaceId = 0
//	client.Player.Gender = 0
//	client.Player.UserId = 2027445
//	client.Player.GameId = 320395999
//	client.Player.Exp = 254
//	client.Player.Loveliness = 0
//	client.Player.Score = 100000009
//	client.Player.NickName = "玩家320395999"
//	client.Player.Level = 1
//	client.Player.VipLevel = 0
//	client.Player.AccountLevel = 3
//	client.Player.SiteLevel = 0
//	client.Player.CurLevelExp = 0
//	client.Player.NextLevelExp = 457
//	client.Player.PayTotal = 0
//	client.Player.Diamond = 29
//
//	// 发送用户信息给客户端
//	// .....
//	sendCmd := &CMD.CMD_MB_LogonSuccess{
//		FaceId: &(client.Player.FaceId),
//		Gender: &(client.Player.Gender),
//		UserId: &(client.Player.UserId),
//		GameId: &(client.Player.GameId),
//		Exp: &(client.Player.Exp),
//		LoveLiness: &(client.Player.Loveliness),
//		UserScore: &(client.Player.Score),
//		NickName: []byte(client.Player.NickName),
//		Level: &(client.Player.Level),
//		VipLev: &(client.Player.VipLevel),
//		AccountLevel: &(client.Player.AccountLevel),
//		Sitelevel: &(client.Player.SiteLevel),
//		CurLevelExp: &(client.Player.CurLevelExp),
//		NextLevelExp: &(client.Player.NextLevelExp),
//		PayTotal: &(client.Player.PayTotal),
//		UserDiamond: &(client.Player.Diamond),
//	}
//	client.Send(sendCmd, MDM_MB_LOGON, SUB_MB_LOGON_SUCCESS)
//	client.SendCmd(MDM_MB_LOGON, SUB_MB_LOGON_FINISH) // 发送登录结束
//}

//****************************************************游客登录游戏服申请****************************************************
func (client *Client)SevLoginGSGuest(buf []byte){
	protocolBuffer := buf
	msg := &CMD.CMD_GR_LogonUserID{}
	err := proto.Unmarshal(protocolBuffer, msg)
	log.CheckError(err)

	fmt.Println("客户端申请登录游戏服务器, UserId: ",msg.GetUserId())

	// --------------------游戏类型-------------------
	fmt.Println("gamekind id :", msg.GetKindId())
	client.Games = Games.GetGameByID(int(msg.GetKindId())) //根据客户端要登录的游戏类型，赋值游戏类型指针


	client.Player = client.Player.NewPlayer()
	client.Player.FaceId = 0
	client.Player.Gender = 0
	client.Player.UserId = 2027445
	client.Player.GameId = 320395999
	client.Player.Exp = 254
	client.Player.Loveliness = 0
	client.Player.Score = 100000009
	client.Player.NickName = "玩家320395999"
	client.Player.Level = 1
	client.Player.VipLevel = 0
	client.Player.AccountLevel = 3
	client.Player.SiteLevel = 0
	client.Player.CurLevelExp = 0
	client.Player.NextLevelExp = 457
	client.Player.PayTotal = 0
	client.Player.Diamond = 29


	AllUserClientList[client.Player.UserId] = client //把新客户端地址增加到哈希表中保存，方便以后遍历

	zRedis.SavePlayerToRedis(client.Player)
	re := zRedis.GetPlayerFromRedis(int(client.Player.UserId))
	fmt.Println("",re.UserId)

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
	client.Send(sendCmd, MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS,"1111111111111111111111玩家账号出现异常，需要重新登录，本次提示为测试提示，客户端需要打印出来，看看是否正常显示！！！")
	client.SendCmd(MDM_GR_LOGON, SUB_GR_LOGON_FINISH) // 发送登录结束
}



//*******************************************************进入大厅申请******************************************************
func (client *Client)SevEnterScence(buf []byte){
	fmt.Println("------------客户端申请进入大厅-------------")
	protocolBuffer := buf
	msg := &CMD.CMD_GF_GameOption{}
	err := proto.Unmarshal(protocolBuffer, msg)
	log.CheckError(err)

	fmt.Println("客户端申请进入大厅, GetClientVersion: ", msg.GetClientVersion())

	//-------------------------------------------逻辑------------------------------------------------------------
	tableUid := client.Games.PlayerLoginGame(client.Player, 1) // 玩家登陆游戏，分配桌子
	client.Table = client.Games.GetTableByUID(tableUid)     // 设置桌子句柄
	client.User = client.Games.GetUserByUID(int(client.Player.UserId))	// 设置玩家的句柄
	client.User.SetConn(client.Conn)								// 设置玩家的连接句柄
	client.User.SetIsRobot(false)							// 连接进来的都不是机器人
	//-------------------------------------------------------------------------------------------------------
	// 发送用户信息给客户端
	sendCmd := &CMD.CMD_S_ENTER_SCENE{

	}
	//this.Send(sendCmd, MDM_GF_FRAME, SUB_GF_GAME_STATUS)	//更新游戏状态
	//this.Send(sendCmd, MDM_GF_FRAME, SUB_GF_SYSTEM_MESSAGE)  // 系统消息
	//this.Send(sendCmd, MDM_GF_FRAME, SUB_GF_USER_SKILL)  // 玩家技能

	client.Send(sendCmd, MDM_GF_GAME, SUB_S_ENTER_SCENE,"") // 进入房间
	//this.Send(sendCmd, MDM_GF_GAME, SUB_S_OTHER_ENTER_SCENE)  // 其他玩家进入房间


	//-------------------------------------------------------------------------------------------------------
	client.Table.SendSceneFishes(client.User)			// 同步一下渔场的所有鱼
	//sendCmd1 := &CMD.CMD_S_SCENE_FISH{
	//
	//}
	//client.Send(sendCmd1, MDM_GF_GAME, SUB_S_SCENE_FISH)  // 场景鱼刷新
	//sendCmd = &CMD.CMD_S_ENTER_SCENE{
	//
	//}
	//client.Send(sendCmd, MDM_GF_GAME, SUB_S_DISTRIBUTE_FISH)  // 新增鱼

	//this.Send(sendCmd, MDM_GF_GAME, SUB_S_USER_FIRE)  // 开火
	//this.Send(sendCmd, MDM_GF_GAME, SUB_S_CATCH_FISH)  // 抓鱼

	//this.SendCmd(MDM_MB_LOGON, SUB_MB_LOGON_FINISH)  // 发送登录结束
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
//	h:= md5.NewPlayer()
//	h.Write(tokenBuf)
//	cips := h.Sum(nil)			// h.Sum(nil) 将h的hash转成[]byte格式
//	pwdmd5 := hex.EncodeToString(cips)
//	pwdmd5 = strings.ToUpper(pwdmd5)
//	//fmt.Println("pwdmd5: ", pwdmd5)
//
//	return pwdmd5
//}
