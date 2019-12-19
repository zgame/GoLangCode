package BY

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"../../ProtocolBuffer/CMD"
	"../../Core/Utils/zLog"
	. "../../Const"
	"encoding/json"
	"../Model/PlayerModel"

)



// -------------------------------------游戏房间-------------------------------------------

//
//// 进入场景
//func (this *Client)handleEnterScence(buf []byte, bufferSize int){
//	protocolBuffer := buf
//	msg := &CMD.CMD_S_ENTER_SCENE{}
//	err := proto.Unmarshal(protocolBuffer, msg)
//	log.CheckError(err)
//	//dataJ, _ := json.MarshalIndent(msg, "", " ")
//	//fmt.Printf("%s", dataJ)
//	//fmt.Println("----------------进入场景--------------", this.Index)
//	for _,v := range msg.GetTableUsers(){
//		if v.GetUserId() == uint32(this.PlayerModel.user_id){
//			this.MyGameInfo.chair_id = int(v.GetChairId())
//		}

//	}
//	fmt.Printf("进入场景,场景id:%d,桌子id:%d,椅子id:%d",msg.GetSceneId(), msg.GetTableId(),this.MyGameInfo.chair_id )
//	fmt.Println("")
//}
//
////# 其他人进入房间
//func (this *Client)handleOtherEnterScence(buf []byte, bufferSize int){
//	protocolBuffer := buf
//	msg := &CMD.CMD_S_OTHER_ENTER_SCENE{}
//	err := proto.Unmarshal(protocolBuffer, msg)
//	log.CheckError(err)
//	//dataJ, _ := json.MarshalIndent(msg, "", " ")
//	//fmt.Printf("%s", dataJ)
//	//fmt.Printf("----------------%s进入了房间--------------%d",  msg.GetUserInfo().GetNickName(),this.Index,)
//	//fmt.Println("")
//
//}
//
//// # 场景鱼刷新
//func (this *Client)handleSceneFish(buf []byte, bufferSize int){
//	protocolBuffer := buf
//	msg := &CMD.CMD_S_SCENE_FISH{}
//	err := proto.Unmarshal(protocolBuffer, msg)
//	log.CheckError(err)
//	//dataJ, _ := json.MarshalIndent(msg, "", " ")
//	//fmt.Printf("%s", dataJ)
//	//fmt.Printf("----------------场景鱼刷新-------------%d", this.Index)
//	//fmt.Println("")
//
//	fish_cnt :=0
//	if this.MyGameInfo.fish_pool == nil {
//		this.MyGameInfo.fish_pool = make([]*FishObj, 0)
//	}
//	fish_cnt = len(this.MyGameInfo.fish_pool)
//	if fish_cnt > MAX_FISH_CNT {
//		return
//	}
//	for _,v :=range msg.GetSceneFishs(){
//		var ffish FishObj
//		ffish.kind_id = int(v.GetKindId())
//		ffish.uid = int(v.GetUid())
//		ffish.tick = time.Now()
//		this.MyGameInfo.fish_pool = append(this.MyGameInfo.fish_pool, &ffish)
//		fish_cnt += 1
//		if fish_cnt > MAX_FISH_CNT{
//			return
//		}
//	}
//
//}
//
//// #新增鱼
//func (this *Client)handleNewFish(buf []byte, bufferSize int){
//	protocolBuffer := buf
//	msg := &CMD.CMD_S_DISTRIBUTE_FISH{}
//	err := proto.Unmarshal(protocolBuffer, msg)
//	log.CheckError(err)
//	//dataJ, _ := json.MarshalIndent(msg, "", " ")
//	//fmt.Printf("%s", dataJ)
//	//fmt.Printf("----------------新增鱼-------------%d", this.Index)
//	//fmt.Println("")
//
//	fish_cnt :=0
//	if this.MyGameInfo.fish_pool == nil {
//		this.MyGameInfo.fish_pool = make([]*FishObj, 0)
//	}
//	fish_cnt = len(this.MyGameInfo.fish_pool)
//	if fish_cnt > MAX_FISH_CNT {
//		return
//	}
//	for _,v :=range msg.GetFishs(){
//		var ffish FishObj
//		ffish.kind_id = int(v.GetKindId())
//		ffish.uid = int(v.GetUid())
//		ffish.tick = time.Now()
//		this.MyGameInfo.fish_pool = append(this.MyGameInfo.fish_pool, &ffish)
//		fish_cnt += 1
//		if fish_cnt > MAX_FISH_CNT{
//			return
//		}
//	}
//
//}
//

// -------------------------------------捕鱼-------------------------------------------

//func (this *Client)do_fire() {
//	this.select_fish()
//
//	tick_count := uint64(time.Now().Unix())
//	angle := float32(45)
//	lock_fish_id := uint32(this.Fish_id)			//选择一个鱼
//	//lock_fish_id := uint32(0)			//选择一个鱼
//	bullet_mulriple := uint32(1)
//	bullet_temp_id := uint32(this.MyGameInfo.bullet_id)
//	this.MyGameInfo.bullet_id +=1
//	bullet_num := uint32(1)
//	is_broadcast := true
//
//	this.Failed_cnt ++
//
//	sendCmd := &CMD.CMD_C_USER_FIRE{
//		TickCount:&tick_count,
//		Angle:&angle,
//		LockFishId:&lock_fish_id,
//		BulletMulriple:&bullet_mulriple,
//		BulletTempId:&bullet_temp_id,
//		BulletNum:&bullet_num,
//		IsBroadcast:&is_broadcast,
//	}
//
//	this.Send(sendCmd, MDM_GF_GAME, SUB_C_USER_FIRE)
//	//fmt.Println("发子弹")
//}
//
//func (this *Client)do_catch(bullet *BulletObj) {
//	fish_uid :=  uint32(bullet.fish_id)
//	bullet_id := uint32(bullet.bullet_id)
//	bullet_temp_id := uint32(bullet.bullet_local_id)
//	chair_id := uint32(this.MyGameInfo.chair_id)
//
//	sendCmd := &CMD.CMD_C_CATCH_FISH{
//		FishUid:&fish_uid,
//		BulletId:&bullet_id,
//		BulletTempId:&bullet_temp_id,
//		ChairId:&chair_id,
//	}
//
//	this.Send(sendCmd, MDM_GF_GAME, SUB_C_CATCH_FISH)
//}


// --------------------------------------------客户端发子弹-----------------------------------------
func HandleUserFire(player *PlayerModel.Player,   table  *BYTable,  buf []byte){
	protocolBuffer := buf
	msg := &CMD.CMD_C_USER_FIRE{}
	err := proto.Unmarshal(protocolBuffer, msg)
	zLog.CheckError(err)
	dataJ, _ := json.MarshalIndent(msg, "", " ")
	fmt.Printf("%s", dataJ)
	fmt.Printf("----------------开火完成，处理子弹消息,------------%d",player.GetUID())
	fmt.Println("")

	// ----------------------------------------逻辑处理-----------------------------------------------
	newBullet := uint32(table.FireBullet( player ,int(msg.GetLockFishId())))		// 获取子弹的uid
	// -------------------------------------------发送消息--------------------------------------------
	BulletId := newBullet
	ChairId := int32(player.GetChairID())
	LockFishId := uint32(msg.GetLockFishId())
	CurrScore := player.GetUser().Score
	sendCmd := &CMD.CMD_S_USER_FIRE{
		BulletId:&BulletId,
		ChairId:&ChairId,
		LockFishId:&LockFishId,
		CurrScore:&CurrScore,
	}
	//client.Send(sendCmd, MDM_GF_GAME, SUB_S_USER_FIRE)  		// 确认客户端发子弹了
	// ----------------------------------------同步给所有玩家该消息-----------------------------------------------

	sendCmd = &CMD.CMD_S_USER_FIRE{
		BulletId:&BulletId,
		ChairId:&ChairId,
		LockFishId:&LockFishId,
		CurrScore:&CurrScore,
	}



	table.SendMsgToAllUsers(sendCmd, MDM_GF_GAME, SUB_S_USER_FIRE)		//发给桌子上面的所有玩家，广播发子弹
	//client.SendToTableOtherUsers(sendCmd, MDM_GF_GAME, SUB_S_USER_FIRE)

}

// ---------------------------------------------客户端申请捕获鱼--------------------------------------------------------
func HandleCatchFish(player *PlayerModel.Player,  table *BYTable, buf []byte){
	protocolBuffer := buf
	msg := &CMD.CMD_C_CATCH_FISH{}
	err := proto.Unmarshal(protocolBuffer, msg)
	zLog.CheckError(err)
	dataJ, _ := json.MarshalIndent(msg, "", " ")
	fmt.Printf("%s", dataJ)
	fmt.Printf("----------------申请捕获鱼,------------%d",player.GetUID())
	fmt.Println("")



	//CatchFishs = &CMD.TagCatchFish{
	//	BulletId:&BulletId,
	//	ChairId:&ChairId,
	//	CatchFishs:&CatchFishs,
	//}
	//BulletId := uint32(client.TableBase.FireBullet(client.PlayerModel,int(msg.GetLockFishId())))		// 获取子弹的uid
	//ChairId := int32(client.PlayerModel.GetChairID())
	//LockFishId := uint32(msg.GetLockFishId())

	// 捕鱼成功
	chairId := int32(player.GetChairID())
	bulletId:= msg.GetBulletId()
	fishUid := msg.GetFishUid()
	addExp := int32(20)

	CurrScore := uint32(220)
	fishI:= table.GetFish(int(fishUid))
	if fishI != nil {
		fmt.Println("鱼的分数",fishI.GetFishScore())
		CurrScore = uint32(fishI.GetFishScore())
	}

	table.DelBullet(int(bulletId))		//删掉子弹
	table.DelFish(int(fishUid))			// 清理掉鱼
	player.GetUser().Score += int64(CurrScore)			// 给玩家把分数加上
	// 把分数保存到数据库


	var CatchFishs []*CMD.TagCatchFish
	fish := &CMD.TagCatchFish{FishUid:&fishUid,FishScore:&CurrScore}
	CatchFishs = append(CatchFishs,fish)

	sendCmd := &CMD.CMD_S_CATCH_FISH{
		ChairId:&chairId,
		AddExp:&addExp,
		CurrScore:&player.GetUser().Score,
		CatchFishs:CatchFishs,

	}
	//client.Send(sendCmd, MDM_GF_GAME, SUB_S_CATCH_FISH)  		// 确认客户端申请捕鱼
	table.SendMsgToAllUsers(sendCmd,MDM_GF_GAME, SUB_S_CATCH_FISH)		//发给桌子上面的所有玩家，广播玩家捕到了鱼
	//for _,v := range msg.GetCatchFishs(){
	//	for i,v2 := range this.MyGameInfo.fish_pool{
	//		if v.GetFishUid() == uint32(v2.uid) {
	//			catch_cnt +=1
	//			this.MyGameInfo.fish_pool = append(this.MyGameInfo.fish_pool[:i], this.MyGameInfo.fish_pool[i+1:]...)				//删除鱼
	//			total_score += int(v.GetFishScore())
	//			this.Fish_id = 0
	//			this.Failed_cnt = 0
	//			break
	//		}
	//	}
	//}
	//if msg.GetChairId() == int32(this.MyGameInfo.chair_id){
	//	this.ShowLog++
	//	//if this.ShowLog % uint64(ShowLog) ==0 {
	//	//	fmt.Printf("----------------捕获到%d条鱼,获得金币:%d,经验:%d------------%d", catch_cnt, total_score, msg.GetAddExp(), this.Index)
	//	//	fmt.Println("")
	//	//}
	//}

}


