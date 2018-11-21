package main

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"./CMD"
	"time"
	"math/rand"
)
// -------------------------------------子弹 和 鱼-------------------------------------------
//# 子弹对象
type BulletObj struct {
	bullet_local_id int     //# 本地id
	bullet_id int             // # 服务器id
	tick  time.Time                    //# 发射时间
	fish_id  int              // # 锁定鱼id
}


//# 鱼对象
type FishObj struct {
	uid  int
	kind_id int
	tick  time.Time
}
// -------------------------------------玩家info-------------------------------------------

type User struct {
	face_id int // # 头像id
	gender int //  # 性别
	user_id int //  # 用户id
	game_id int //  # 游戏id
	exp int //  # 经验
	loveliness int //  # 魅力
	score int64 //  # 分数
	nick_name string //  # 昵称
	level int //  # 等级
	vip_level int //  # vip等级
	account_level int //  # 账号等级
	site_level int //  # 炮等级
	cur_level_exp int //  # 当前等级经验
	next_level_exp int //  # 下一等级经验
	pay_total int //  # 充值总金额
	diamond int //  # 钻石数量
}

func (self *User)New() *User {
	return &User{face_id:0,gender:0,user_id:0,game_id:0,exp:0,loveliness:0,score:0,nick_name:"",level:0,vip_level:0,account_level:0,site_level:0,cur_level_exp:0,next_level_exp:0,pay_total:0,diamond:0}
}
// -------------------------------------服务器info-------------------------------------------
type GameServerInfo struct {
	kind_id int  //  # 名称索引
	node_id int  //  # 节点索引
	sort_id int  //  # 排序索引
	server_id int  //  # 房间索引
	server_port int  //  # 房间端口
	online_count int  //  # 在线人数
	online_pc int  //
	online_andriod int  //
	online_ios int  //  # 在线人数
	full_count int  //  # 满员人数
	server_addr string  //# 房间地址(抛弃)
	server_name string  //# 房间名称
	cell_score int  //
	min_enter_score int  //  # 最低进入分数
	max_enter_score int  //  # 最高进入分数
	server_type int  //  # 服务器类型1为财富4为比赛
	min_enter_vip int  //  # 最低进入VIP
	min_enter_cannon_lev int  //  # 最低进入炮等级
	server_rule int  //  # 服务器规则

}

func NewGameServerInfo() *GameServerInfo {
	return &GameServerInfo{kind_id:0,node_id:0,sort_id:0,server_id:0,server_port:0,online_count:0,online_pc:0,server_addr:"",server_name:"",online_ios:0,cell_score:0,full_count:0,min_enter_score:0,max_enter_score:0,server_type:0,min_enter_vip:0,min_enter_cannon_lev:0,server_rule:0}
}

// -------------------------------------游戏info-------------------------------------------

type GameInfo struct {
	fish_pool []*FishObj // # 场景中的鱼
	last_check_due_tick int //  # 检查鱼过期
	last_ai_update_tick int // ai更新时间
	chair_id int //  #椅子id
	fire_cnt int	//# 发射炮弹数量
	catch_cnt int	//# 捕获鱼的数量
	chat_server_ip string	//# 聊天服务器ip
	chat_server_port int 	//# 聊天服务器端口
	chat_token string	//# 聊天服务器token
	chat_rand int	//# 聊天服务器rand
	//friend_list []		//# 好友列表

	bullet_id int  //子弹id，自增
}

func (self *GameInfo)New() *GameInfo {
	return &GameInfo{fish_pool:nil,last_check_due_tick:0,last_ai_update_tick:0,chair_id:0,fire_cnt:0,catch_cnt:0,chat_server_port:0,chat_server_ip:"",chat_token:"",chat_rand:0,bullet_id:1}
}


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
	bufferT := getSendTcpHeaderData(MDM_GF_FRAME, SUB_GF_GAME_OPTION, uint16(size),uint16(this.SendTokenID))
	this.SendTokenID++
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
	//fmt.Println("----------------进入场景--------------", this.User.user_id )
	for _,v := range msg.GetTableUsers(){
		if v.GetUserId() == uint32(this.User.user_id){
			this.Gameinfo.chair_id = int(v.GetChairId())
		}
	}
	//fmt.Printf("进入场景,场景id:%d,桌子id:%d,椅子id:%d  \n",msg.GetSceneId(), msg.GetTableId(),this.Gameinfo.chair_id )
	//fmt.Println("")
}

//# 其他人进入房间
func (this *Client)handleOtherEnterScence(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_OTHER_ENTER_SCENE{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------%d进入了房间%d--椅子 ：%d \n",  msg.GetUserInfo().GetUserId(), msg.GetUserInfo().GetTableId(), msg.GetUserInfo().GetChairId())
	//fmt.Println("")

}

// # 场景鱼刷新
func (this *Client)handleSceneFish(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_SCENE_FISH{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s\n", dataJ)

	//fmt.Println("")

	fish_cnt :=0
	if this.Gameinfo.fish_pool == nil {
		this.Gameinfo.fish_pool = make([]*FishObj, 0)
	}
	fish_cnt = len(this.Gameinfo.fish_pool)
	if fish_cnt > MAX_FISH_CNT {
		return
	}
	//for _,v :=range msg.GetSceneFishs(){
	//	var ffish FishObj
	//	ffish.kind_id = int(v.GetKindId())
	//	ffish.uid = int(v.GetUid())
	//	ffish.tick = time.Now()
	//	this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool, &ffish)
	//	fish_cnt += 1
	//	if fish_cnt > MAX_FISH_CNT{
	//		return
	//	}
	//}
	//fmt.Printf("----------------场景鱼刷新-------------%d", fish_cnt)

}

// #新增鱼
func (this *Client)handleNewFish(buf []byte, bufferSize int){
	protocolBuffer := buf
	msg := &CMD.CMD_S_DISTRIBUTE_FISH{}
	err := proto.Unmarshal(protocolBuffer, msg)
	checkError(err)
	//dataJ, _ := json.MarshalIndent(msg, "", " ")
	//fmt.Printf("%s", dataJ)
	//fmt.Printf("----------------新增鱼-------------%d\n", this.Index)


	fish_cnt :=0
	if this.Gameinfo.fish_pool == nil {
		this.Gameinfo.fish_pool = make([]*FishObj, 0)
	}
	fish_cnt = len(this.Gameinfo.fish_pool)
	if fish_cnt > MAX_FISH_CNT {
		return
	}
	for _,v :=range msg.GetFishs(){
		//var ffish FishObj
		//ffish.kind_id = int(v.GetKindId())
		//ffish.uid = int(v.GetUid())
		//ffish.tick = time.Now()
		//this.Gameinfo.fish_pool = append(this.Gameinfo.fish_pool, &ffish)
		//fish_cnt += 1
		//if fish_cnt > MAX_FISH_CNT{
		//	return
		//}

		this.Fish_id = int(v.GetUid())
	}

}


// -------------------------------------捕鱼-------------------------------------------

func (this *Client)do_fire() {
	//this.select_fish()

	if this.Fish_id <=0 {
		return
	}

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
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_USER_FIRE, uint16(size),uint16(this.SendTokenID))
	this.SendTokenID++
	this.Send(bufferT, data)
	//fmt.Println("发子弹")





}

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
	bufferT := getSendTcpHeaderData(MDM_GF_GAME, SUB_C_CATCH_FISH, uint16(size),uint16(this.SendTokenID))
	this.SendTokenID++
	this.Send(bufferT, data)
	//fmt.Println("申请捕鱼",fish_uid,bullet_id,chair_id)
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
		//bullt.bullet_local_id = int(msg.GetBulletTempId())
		bullt.tick = time.Now()

		//this.User.score = msg.GetCurrScore()


		this.do_catch(&bullt)
	}

}

// #捕获鱼
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
			fmt.Printf("----------------捕获到%d条鱼,获得金币:%d,经验:%d------------%d \n", catch_cnt, total_score, msg.GetAddExp(), this.Index)
			//fmt.Println("")
		}
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

	_, err = this.Conn.Write(getSendTcpHeaderData(MDM_GF_GAME, SUB_C_GET_ALMS, 0,uint16(this.SendTokenID))) //发送注册成为客户端请求
	this.SendTokenID++
	checkError(err)
}



// game AI
func (this *Client)GameAI()  {
	if !this.StartAI{
		return
	}
	//fmt.Println("AI-----")
	if time.Now().After(this.Last_fire_tick)   {
		this.Last_fire_tick = time.Now().Add( time.Microsecond * 200)
		this.do_fire()


		//var bullt BulletObj
		//bullt.bullet_id = 1
		//bullt.fish_id = 1
		//bullt.tick = time.Now()
		//this.do_catch(&bullt)
	}

	//# 处理过期的鱼
	//if time.Now().After(this.Last_check_due_tick) {
	//	this.Last_check_due_tick = time.Now().Add(time.Second)
	//	this.check_overdue_fish()
	//}

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
