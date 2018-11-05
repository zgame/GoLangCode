package main

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"./CMD"
	"time"
	"math/rand"
	. "./const"
)
// -------------------------------------游戏大厅-------------------------------------------

func (this *Client)EnterScence() {
	allow_lookon := int32(0)
	frame_version := uint32(0)
	client_version := uint32(101056515)

	sendCmd := &CMD.CMD_GF_GameOption{
		AllowLookon:&allow_lookon,
		FrameVersion:&frame_version,
		ClientVersion:&client_version,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_FRAME, SUB_GF_GAME_OPTION, uint16(size))
	this.Send(bufferT, data)
}

// 更新游戏状态
func (this *Client)handleGameStatus(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_GF_GameStatus{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Println("----------------更新游戏状态--------------", this.Index)

}

func (this *Client)handleGameMessage(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_CR_SystemMessage{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Println("----------------游戏消息--------------", this.Index, "+", msg.Text)

}

// 接收技能
func (this *Client)handleUserSkill(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_CF_S_UserSkill{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------接收技能--------------%d, %v", this.Index, msg.UserSkills)

}

// -------------------------------------游戏房间-------------------------------------------


// 进入场景
func (this *Client)handleEnterScence(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_ENTER_SCENE{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Println("----------------进入场景--------------", this.Index)
	for _,v := range msg.GetTableUsers(){
		if v.GetUserId() == uint32(this.User.user_id){
			this.Gameinfo.chair_id = int(v.GetChairId())
		}
	}
	fmt.Printf("进入场景,场景id:%d,桌子id:%d,椅子id:%d",msg.GetSceneId(), msg.GetTableId(),this.Gameinfo.chair_id )
	fmt.Println("")
}

//# 其他人进入房间
func (this *Client)handleOtherEnterScence(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_OTHER_ENTER_SCENE{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------%s进入了房间--------------%d",  msg.GetUserInfo().GetNickName(),this.Index,)
	//fmt.Println("")

}

// # 场景鱼刷新
func (this *Client)handleSceneFish(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_SCENE_FISH{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------场景鱼刷新-------------%d", this.Index)
	//fmt.Println("")

	fish_cnt :=0
	if this.Gameinfo.fish_pool == nil {
		this.Gameinfo.fish_pool = make([]*FishObj, 0)
	}
	fish_cnt = len(this.Gameinfo.fish_pool)
	if fish_cnt > MAX_FISH_CNT {
		return
	}
	for _,v :=range msg.GetSceneFishs(){
		var ffish FishObj
		ffish.kind_id = int(v.GetKindId())
		ffish.uid = int(v.GetUid())
		ffish.tick = time.Now()
		this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool, &ffish)
		fish_cnt += 1
		if fish_cnt > MAX_FISH_CNT{
			return
		}
	}

}

// #新增鱼
func (this *Client)handleNewFish(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_DISTRIBUTE_FISH{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------新增鱼-------------%d", this.Index)
	//fmt.Println("")

	fish_cnt :=0
	if this.Gameinfo.fish_pool == nil {
		this.Gameinfo.fish_pool = make([]*FishObj, 0)
	}
	fish_cnt = len(this.Gameinfo.fish_pool)
	if fish_cnt > MAX_FISH_CNT {
		return
	}
	for _,v :=range msg.GetFishs(){
		var ffish FishObj
		ffish.kind_id = int(v.GetKindId())
		ffish.uid = int(v.GetUid())
		ffish.tick = time.Now()
		this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool, &ffish)
		fish_cnt += 1
		if fish_cnt > MAX_FISH_CNT{
			return
		}
	}

}


// -------------------------------------开炮普通捕鱼-------------------------------------------
// 开炮
func (this *Client)do_fire() {
	this.select_fish()

	tick_count := uint64(time.Now().Unix())
	angle := float32(45)
	lock_fish_id := uint32(this.Fish_id)			//选择一个鱼
	//lock_fish_id := uint32(0)			//选择一个鱼
	bullet_mulriple := uint32(1)
	bullet_temp_id := uint32(this.Gameinfo.bullet_id)
	this.Gameinfo.bullet_id +=1
	bullet_num := uint32(1)
	is_broadcast := true

	this.Failed_cnt ++

	sendCmd := &CMD.CMD_C_USER_FIRE{
		TickCount:&tick_count,
		Angle:&angle,
		LockFishId:&lock_fish_id,
		BulletMulriple:&bullet_mulriple,
		BulletTempId:&bullet_temp_id,
		BulletNum:&bullet_num,
		IsBroadcast:&is_broadcast,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_USER_FIRE, uint16(size))
	this.Send(bufferT, data)
	//fmt.Println("发子弹")
}



// 捕获鱼
func (this *Client)do_catch(bullet *BulletObj) {
	fish_uid :=  uint32(bullet.fish_id)
	bullet_id := uint32(bullet.bullet_id)
	bullet_temp_id := uint32(bullet.bullet_local_id)
	chair_id := uint32(this.Gameinfo.chair_id)

	sendCmd := &CMD.CMD_C_CATCH_FISH{
		FishUid:&fish_uid,
		BulletId:&bullet_id,
		BulletTempId:&bullet_temp_id,
		ChairId:&chair_id,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_CATCH_FISH, uint16(size))
	this.Send(bufferT, data)
}


// #处理子弹消息,
func (this *Client)handleUserFire(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_USER_FIRE{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------开火完成，处理子弹消息,------------%d", this.Index)
	//fmt.Println("")

	if msg.GetChairId() == int32(this.Gameinfo.chair_id){
		var bullt BulletObj
		bullt.bullet_id = int(msg.GetBulletId())
		bullt.fish_id = int(msg.GetLockFishId())
		bullt.bullet_local_id = int(msg.GetBulletTempId())
		bullt.tick = time.Now()

		this.User.score = msg.GetCurrScore()
		this.do_catch(&bullt)
	}
}

//------------------------------------------------------------发技能弹头捕鱼------------------------------------------

// 申请开始使用技能
func (this *Client)do_start_skill() {

	SkillID_ := int32(SkillID)
	IsTrigger_ := false
	sendCmd := &CMD.CMD_C_START_SKILL{
		SkillId:&SkillID_,
		IsTrigger:&IsTrigger_,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_START_SKILL, uint16(size))
	this.Send(bufferT, data)
	//fmt.Println("发弹头")
}
// #处理技能消息,
func (this *Client)handleUserStartUseSkill(buf []byte){
	protocolBuffer := buf
	msg := &CMD.CMD_S_START_SKILL{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------handleUserStartUseSkill技能完成，处理技能消息,------------%d", this.User.user_id)
	//fmt.Println("")

	if msg.GetChairId() == int32(this.Gameinfo.chair_id){
		this.do_skill()
	}

}

// 使用技能
func (this *Client)do_skill() {

	TargetId := make([]int,0)
	for _,v:=range this.Gameinfo.fish_pool {
		if GoldFishMap[v.kind_id] != 0	{
			//fmt.Println("是黄金鱼", v.uid ,"  kindid:",v.kind_id)
			TargetId= append(TargetId, v.uid)
			//break
		}
	}
	//fmt.Printf("黄金鱼%v", TargetId)
	var targetFish uint32
	if len(TargetId)>0 {
		rand.Seed(time.Now().UnixNano())
		targetFish = uint32(TargetId[rand.Intn(len(TargetId))])
		//fmt.Println("选中黄金鱼",targetFish, "  kindid:", )
	}

	SkillID_ := uint32(SkillID)
	IsTrigger_ := false
	sendCmd := &CMD.CMD_C_USE_SKILL{
		SkillId:&SkillID_,
		TargetId:&targetFish,
		IsTrigger:&IsTrigger_,
	}

	//dataJ, _ := json.MarshalIndent(sendCmd, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("---------------do_skill处理技能消息,------------%d", this.Index)
	//fmt.Println("")

	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_USE_SKILL, uint16(size))
	this.Send(bufferT, data)
	//fmt.Println("发弹头")
}


// #处理使用技能消息,
func (this *Client)handleUserUseSkill(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_USE_SKILL{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//if msg.GetSkillStatus() >0 {
	//	dataJ, _ := json.MarshalIndent(msg, "", " ")
	//	fmt.Printf("%s", dataJ)
	//	fmt.Printf("----------------handleUserUseSkill处理技能消息,------------%d", this.Index)
	//	fmt.Println("")
	//}

	if msg.GetChairId() == int32(this.Gameinfo.chair_id){
		this.do_skill_catch()
	}
}

// 技能捕获鱼
func (this *Client)do_skill_catch() {
	fish_uid :=  make([]uint32,0)
	for _,v :=range this.Gameinfo.fish_pool{
		fish_uid = append(fish_uid, uint32(v.uid))
	}

	chair_id := uint32(this.Gameinfo.chair_id)
	SkillID_ := uint32(SkillID)
	sendCmd := &CMD.CMD_C_SKILL_CATCH_FISH{
		FishsUid:fish_uid,
		SkillId:&SkillID_,
		ChairId:&chair_id,
	}
	data, _ := proto.Marshal(sendCmd)
	size := len(data)
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_SKILL_CATCH_FISH, uint16(size))
	this.Send(bufferT, data)

	//fmt.Println(this.User.user_id,"使用技能捕获鱼")
}

// #成功捕获鱼
func (this *Client)handleCatchFish(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_CATCH_FISH{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	catch_cnt := 0
	total_score := 0

	for _,v := range msg.GetCatchFishs(){
		for i,v2 := range this.Gameinfo.fish_pool{
			if v.GetFishUid() == uint32(v2.uid) {
				catch_cnt +=1
				this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool[:i], this.Gameinfo.fish_pool[i+1:]...)				//删除鱼
				total_score += int(v.GetFishScore())
				this.Fish_id = 0
				this.Failed_cnt = 0
				break
			}
		}
	}
	if msg.GetChairId() == int32(this.Gameinfo.chair_id){
		this.ShowLog++
		if this.ShowLog % uint64(ShowLog) ==0 {
			fmt.Printf("----------------捕获到%d条鱼,获得金币:%d,经验:%d------------%d", catch_cnt, total_score, msg.GetAddExp(), this.Index)
			fmt.Println("")
		}

		// 普通子弹捕获鱼之后，说明已经有鱼出现了， 发技能捕鱼
		//this.do_skill()

	}

}

// 获取救济金
func (this *Client)handleDrawAlm(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_C_GET_ALMS{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	fmt.Printf("----------------获取救济金,------------%d", this.Index)

	_, err = this.Conn.Write(getSendTcpHeaderData(MDM_GF_GAME, SUB_C_GET_ALMS, 0)) //发送注册成为客户端请求
	checkError(err)
}



// game AI
func (this *Client)GameAI(ReLoginTime int)  {
	if !this.StartAI{
		return
	}

	if ReLoginTime >0 && time.Now().After(this.ReloginTime){
		// 在client初始化的时候创建了时间， 1分钟左右， 这里到期之后， 就断线重连
		fmt.Println(".......到时间了..........")

		this.StartAI = false




		// index ++ 这样这个连接就会变更下一个mac地址，继续连接
		this.Index += AllUserNum

		rad := rand.Intn(2)
		if rad == 1 {
			// 正常退出
			this.logoutGS()
			fmt.Println("正常退出")
		}else {
			// 不正常退出，不发消息
			this.ConnectLoginServer() // 重新登录登录服务器
			fmt.Println("不正常退出")
		}

		return
	}


	//fmt.Println("AI-----")

	if UseFire == true {
		if time.Now().After(this.Last_fire_tick) {
			this.Last_fire_tick = time.Now().Add(time.Microsecond * time.Duration(FireCD))
			this.do_fire()
		}
	}


	if UseSkill == true {
		if time.Now().After(this.Last_skill_tick) {
			this.Last_skill_tick = time.Now().Add(time.Second * 2)
			for _, v := range this.Gameinfo.fish_pool {
				if GoldFishMap[v.kind_id] != 0 {
					//fmt.Println("开始使用技能", this.User.user_id)
					this.do_start_skill()
					return
				}
			}
		}
	}

	//# 处理过期的鱼
	if time.Now().After(this.Last_check_due_tick) {
		this.Last_check_due_tick = time.Now().Add(time.Second)
		this.check_overdue_fish()
	}



}


// 选择鱼
func (this *Client)select_fish() {
	//# 1s 选择一次

	if time.Now().After(this.Select_tick) {
		this.Select_tick = time.Now().Add(time.Second)

		if this.Failed_cnt < 10 && this.Fish_id != 0 {
			return
		}


		if this.Failed_cnt > 10 && this.Fish_id != 0 {		// 删除鱼
			for i,v:=range this.Gameinfo.fish_pool {
				if v.uid == this.Fish_id {
					//fmt.Println("开了10炮打不死",v.uid)
					this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool[:i], this.Gameinfo.fish_pool[i+1:]...)
				}
			}
		}

		if len(this.Gameinfo.fish_pool) > 0 {
			rand.Seed(time.Now().UnixNano())
			x := rand.Intn(len(this.Gameinfo.fish_pool))
			this.Fish_id = this.Gameinfo.fish_pool[x].uid

			//fmt.Println("重新选择鱼id", this.Fish_id)
		}
	}

}

// 定期删除过期的鱼

func (this *Client)check_overdue_fish(){
	for i,v:= range this.Gameinfo.fish_pool{
		if time.Now().After(v.tick.Add(time.Second*10)){
			if this.Fish_id == v.uid{
				this.Fish_id = 0
			}
			//fmt.Println("鱼过期了", v.uid)
			//fmt.Println("删除定期的鱼", i, "--len" , len(this.Gameinfo.fish_pool) )
			//if len(this.Gameinfo.fish_pool) <= 1{
			//	this.Gameinfo.fish_pool = make([]*FishObj, 0)
			//}else{
			this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool[:i], this.Gameinfo.fish_pool[i+1:]...)
			//}
			return
		}
	}

}
