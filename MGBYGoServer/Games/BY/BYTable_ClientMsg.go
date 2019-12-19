package BY


import (
	"fmt"
	."../../Const"
	"../../Core/ZServer"
	."../Model/PlayerModel"
	. "./BYModel"
	"../../ProtocolBuffer/CMD"
)

// 进入场景之后的首次消息同步， 每个游戏不一样， 都会重载该方法
func (table *BYTable) EnterSceneSyncMsg(player *Player)  {
	//-------------------------------------------------------------------------------------------------------
	// 发送用户信息给客户端
	fmt.Println("---------------------------进入场景同步信息发送-------------------------------")

	var TableUsers []*CMD.TagUserInfo

	// 获取所有桌子上玩家的信息
	table.RWMutexSeatArray.RLock()
	for chairId , player := range table.UserSeatArray{
		if player != nil {
			uid := uint32(player.GetUID())
			chairId32 := int32(chairId)
			userInfo := &CMD.TagUserInfo{UserId:&uid, ChairId:&chairId32}
			TableUsers = append(TableUsers, userInfo)
		}
	}
	table.RWMutexSeatArray.RUnlock()


	tableUid := int32(table.GetTableUID())
	sendCmd := &CMD.CMD_S_ENTER_SCENE{
		TableId:&tableUid,
		TableUsers:TableUsers,
	}
	ZServer.NetWorkSendByUid(player.GetUID(),sendCmd, MDM_GF_GAME, SUB_S_ENTER_SCENE) // 进入房间

	//this.Send(sendCmd, MDM_GF_FRAME, SUB_GF_GAME_STATUS)	//更新游戏状态
	//this.Send(sendCmd, MDM_GF_FRAME, SUB_GF_SYSTEM_MESSAGE)  // 系统消息
	//this.Send(sendCmd, MDM_GF_FRAME, SUB_GF_USER_SKILL)  // 玩家技能


	//this.Send(sendCmd, MDM_GF_GAME, SUB_S_OTHER_ENTER_SCENE)  // 其他玩家进入房间


	//-------------------------------------------------------------------------------------------------------
	table.SendSceneFishes(player) // 同步一下渔场的所有鱼
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



// 玩家登陆的时候， 同步给玩家场景中目前鱼群的信息
func (table *BYTable) SendSceneFishes(player *Player){
	var SceneFishs []*CMD.TagSceneFish

	fmt.Println("玩家登陆，鱼数量",len(table.FishArray))

	table.RWMutexFishArray.RLock()
	for _,fish := range table.FishArray{
		Uid:= uint32(fish.GetFishUID())
		KindId:= uint32(fish.GetFishKindID())
		fishT := &CMD.TagSceneFish{Uid:&Uid,KindId:&KindId}
		SceneFishs = append(SceneFishs,fishT)
	}
	table.RWMutexFishArray.RUnlock()

	sendCmd1 := &CMD.CMD_S_SCENE_FISH{
		SceneFishs:SceneFishs,
	}
	ZServer.NetWorkSendByUid(player.GetUID(), sendCmd1, MDM_GF_GAME, SUB_S_SCENE_FISH)
	//NetWork.Send(UserModel.GetConn(), sendCmd1, MDM_GF_GAME, SUB_S_SCENE_FISH,"")  // 场景鱼刷新
}

// 给所有玩家同步新建的鱼的信息
func (table *BYTable) SendNewFishes(fish  *CommonFish) {
	var SceneFishs []*CMD.TagSceneFish

	Uid := uint32(fish.GetFishUID())
	KindId := uint32(fish.GetFishKindID())

	fishT := &CMD.TagSceneFish{Uid: &Uid, KindId: &KindId}
	SceneFishs = append(SceneFishs, fishT)

	sendCmd := &CMD.CMD_S_DISTRIBUTE_FISH{
		Fishs: SceneFishs,
	}

	table.SendMsgToAllUsers(sendCmd, MDM_GF_GAME, SUB_S_DISTRIBUTE_FISH) // 告诉所有玩家，新增鱼了
}

