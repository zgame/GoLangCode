package Common

import (
	"fmt"
	"time"
	"../../Const"
	. "../../Const"
	"github.com/golang/protobuf/proto"
	//"../../Core/NetWork"
	"../../Core/ZServer"
	"../../ProtocolBuffer/CMD"
	//. "../../Const"

)

//--------------------------------------------------------------------------------------
// 该类为桌子的通用父类， 这里可以定义一些通用的函数，每个游戏具体的桌子不定义的话，就用父类， 也可以重新定义实现重载
//--------------------------------------------------------------------------------------

// -------------------------桌子的统一定义接口-------------------------------------
type TableInterface interface{
	NewTable(uid int, gameId int)  TableInterface 	//桌子构造函数
	GetTableUID()	int			// 获取桌子uid
	RunTable()						// 桌子的逻辑启动

	// 椅子逻辑
	GetEmptySeatInTable()	int           // 获取空椅子的编号
	PlayerSeat(seatId int, user *Player)    // 玩家坐到椅子上
	PlayerStandUp(seatId int, user *Player) // 玩家离开椅子

	FireBullet(user *Player, lockFishId int)	int // 发射一个新的子弹
	CreateFish() FishInterface                     // 新建一个新的鱼

	GetUsersSeatInTable() []int		// 获取桌子上的所有玩家uid
	CheckTableEmpty() bool			// 判断桌子是有人，还是空桌子
	InitTable()						// 初始化桌子
	ClearTable()					// 清空桌子


	GetBulletInterface() BulletInterface // 获取桌子管理子弹的句柄类型
	SetBulletInterface()                 // 设置桌子管理子弹的句柄类型
	GetFishInterface() FishInterface     // 获取桌子管理鱼的句柄类型
	SetFishInterface()                   // 设置桌子管理鱼的句柄类型


	GetRoomScore() int						// 获取房间的分数
	SetRoomScore(roomScore int)				// 设定房间的分数


	SendSceneFishes(user *Player)      // 玩家登陆的时候，同步给鱼群数据
	SendNewFishes(fish FishInterface) // 同步新建一条鱼的数据

	InitDistributeInfo(roomScore int)			// 初始化鱼池的刷新
	//GetTableState() int		// 获取桌子状态
	//SetTableState(state int )	// 设置桌子状态
	//GetTableBulletArray() map[int]GetBulletInterface		// 获取桌子的子弹列表
	HitFish(userId *Player, bulletId int, fishId int) // 击中一条鱼

	DelBullet(bulletId int) //  删除特定uid的子弹
	DelBullets(userId int)  //  删除所有子弹， 1 如果传入玩家uid，删除玩家的  ； 2  如果传入 -1 ，那么删除所有的

	DelFish(fishId int)				// 删除特定uid的鱼
	DelFishes()						// 清空所有的鱼群
	GetFish(fishId int) FishInterface		// 获取鱼的句柄

	SendMsgToOtherUsers(uid int, sendCmd proto.Message, mainCmd int, subCmd int)		// 给桌上的其他玩家同步消息
	SendMsgToAllUsers(sendCmd proto.Message, mainCmd int, subCmd int)		// 给桌上的所有玩家同步消息
}


// -------------------------桌子的统一结构-------------------------------------
type CommonTable struct {
	GameID    int // 游戏类型
	TableID   int // 桌子id
	TableMax  int // 桌子容纳玩家数量
	RoomScore int // 房间分数
	//State int			//桌子状态		0 开放中，有玩家正在玩 ； 1 关闭中，没有玩家在游戏

	UserSeatArray       map[int]*Player // 座椅对应玩家uid的哈希表 ， key ： seat ，value： 玩家uid

	//UserCnt int			// 玩家数量
	//StateChgTime time.Time	 // 状态变更时间

	//
	GenerateFishUid       int                     //生成鱼的uid
	GenerateBulletUid     int                     // 生成子弹的uid
	FishInterfaceHandle   FishInterface           // 鱼的句柄类型
	BulletInterfaceHandle BulletInterface         // 子弹的句柄类型
	FishArray             map[int]FishInterface   // 鱼的哈希表
	BulletArray           map[int]BulletInterface // 子弹的哈希表

	DistributeArray []FishDistributeInfo 	// 鱼的生成信息数据
	BossDistributeArray []FishDistributeInfo //Boss鱼的生成信息数组
	//PathArray []int 		//路径数组

	//SceneStatus int 		// 场景状态
	//LastUpdateTime time.Time	//上次更新时间
	//SceneElapsed time.Time		//场景累计时间

	//ByFishFormNum	int		// 当前鱼阵轮数
	//CurrentFishForm int 	// 当前鱼阵
	//CurrentSceneID int 	// 当前场景id
	//FishFormTime int		// 鱼阵的持续时间


}


// -------------------------构造函数-------------------------
func (table *CommonTable) NewTable(uid int, gameId int) TableInterface{
	return &CommonTable{TableID:uid, GameID: gameId, TableMax:Const.BY_TABLE_MAX_PLAYER}
}


//---------------------------------桌子的状态------------------------------------------

// 返回桌子的UID
func  (table *CommonTable)GetTableUID() int{
	return table.TableID
}

// 获取桌子管理子弹的句柄类型
func (table *CommonTable) GetBulletInterface() BulletInterface {
	return table.BulletInterfaceHandle
}
// 设置桌子管理子弹的句柄类型
func (table *CommonTable) SetBulletInterface() {
	var bullet * CommonBullet
	table.BulletInterfaceHandle = bullet
}
// 获取桌子管理鱼的句柄类型
func (table *CommonTable) GetFishInterface() FishInterface {
	return table.FishInterfaceHandle
}
// 设置桌子管理鱼的句柄类型
func (table *CommonTable) SetFishInterface() {
	var fish * CommonFish
	table.FishInterfaceHandle = fish
}



// 获取房间的分数
func (table *CommonTable) GetRoomScore() int {
	return table.RoomScore
}
// 设定房间的分数
func (table *CommonTable) SetRoomScore(roomScore int) {
	table.RoomScore = roomScore
}
//----------------------------------------------------------------------------
// -------------------------玩家逻辑--------------------------------------------------
//----------------------------------------------------------------------------

// -------------------------判断桌子是有人，还是空桌子---------------------------------------------------
func (table *CommonTable)CheckTableEmpty() bool{
	// 遍历所有的座位
	for i:= 0 ;i< table.TableMax;i++{
		if table.UserSeatArray[i] != nil{
			return false		// 有人在玩
		}
	}
	return true		// 空桌子
}


// -----------------获取桌子的所有玩家----------------
func (table *CommonTable)GetUsersSeatInTable() []int{
	userList := make([]int,0)
	// 遍历所有的座位
	for i:= 0 ;i< table.TableMax;i++{
		if table.UserSeatArray[i] != nil {
			//如果座位上有玩家uid，说明有人
			userList = append(userList,table.UserSeatArray[i].GetUID())
		}
	}
	return userList
}
// -----------------获取桌子的空座位, 返回座椅的编号，从0开始到tableMax， 如果返回-1说明满了----------------
func (table *CommonTable)GetEmptySeatInTable() int{
	// 遍历所有的座位
	for i:= 0 ;i< table.TableMax;i++{
		if table.UserSeatArray[i] == nil {
			//如果座位上没有玩家uid，说明没有人，返回座位id
			fmt.Println(i,"这个位置有空着椅子")
			return i
		}
	}
	// 如果座椅都满了， 那么返回-1
	fmt.Println(table.GetTableUID(), "都满了")
	return -1
}

// ---------------------------玩家坐到椅子上-------------------------------
func (table * CommonTable)PlayerSeat(seatId int, user *Player)  {
	table.UserSeatArray[seatId] = user
}

// 初始化桌子
func (table *CommonTable) InitTable()  {
	if table.CheckTableEmpty(){
		table.BulletArray = make(map[int]BulletInterface )
		table.FishArray = make(map[int]FishInterface )
		table.UserSeatArray = make(map[int]*Player)

		table.InitDistributeInfo(table.GetRoomScore())	//玩家进来之后，如果是第一个玩家，那么房间开始初始化一下
	}
}

// --------------------------- 玩家离开椅子 -----------------------
func (table * CommonTable)PlayerStandUp(seatId int, user *Player) {
	table.UserSeatArray[seatId] = nil
	// 清理掉玩家的所有子弹
	table.DelBullets(user.GetUID())
	// 如果是空桌子的话，清理一下桌子
	if table.CheckTableEmpty() {
		table.ClearTable()
	}
}
// 清理桌子
func (table * CommonTable)ClearTable()  {
	// 清理掉生成鱼群的结构
	table.DistributeArray = nil
	table.BossDistributeArray = nil		// 因为玩家进来之后会初始化会make，所以就不用make了

	// 清理掉所有的子弹和鱼群
	table.DelBullets(-1)
	table.DelFishes()

	table.BulletArray = nil
	table.FishArray = nil
	table.UserSeatArray = nil
}

//----------------------------------------------------------------------------
// -----------------------------  桌子主循环逻辑 -------------------------------
//----------------------------------------------------------------------------
func (table *CommonTable)RunTable() {
	for {
		if table.CheckTableEmpty()  {
			fmt.Println("这是一个空桌子")
		}else{
		//	fmt.Println("运行桌子", table.TableID, "游戏类型", table.GameID)
			table.RunDistributeInfo(table.GetRoomScore())
			table.RunBossDistributeInfo(table.GetRoomScore())


			for _,bullet :=range table.BulletArray{
				bullet.BulletRun(table)				// 遍历所有子弹，并且run
			}
			for _,fish :=range table.FishArray{
				fish.FishRun(table)					// 遍历所有鱼，并且run
			}
			//fmt.Println("共有子弹数量",len(table.BulletArray))
			fmt.Println("共有鱼数量",len(table.FishArray))
		}
		time.Sleep(time.Millisecond * 1000)
	}
}


//----------------------------------------------------------------------------
//------------------------------bullet 子弹-----------------------------------------
//----------------------------------------------------------------------------
// 玩家发射一个新的子弹
func (table *CommonTable)FireBullet(player *Player, lockFishId int) int{
	num := player.GetActivityBulletNum()
	if num > Const.MAX_BULLET_NUMBER{
		fmt.Println("子弹数量超上限了")
		return -1
	}

	cost := table.RoomScore  // 底分
	if player.GetUser().Score < int64(cost) {
		fmt.Println("玩家没钱了")
		return -1
	}


	//var bullet * CommonBullet
	bullet := table.GetBulletInterface()				// 获取子弹接口类型
	bulletT := bullet.NewBullet(table.GenerateBulletUid) // 生成子弹
	table.GenerateBulletUid ++
	table.BulletArray[bulletT.GetBulletUID()] = bulletT // 把bullet加入列表
	bulletT.SetBulletPlayerUID(player.GetUID())         //子弹的主人
	bulletT.SetLockFishID(lockFishId)                   // 锁定鱼
	player.SetActivityBulletNum(num+1)                  //玩家已激活子弹增加
	return bulletT.GetBulletUID()
}

//  击中一条鱼
func (table *CommonTable) HitFish(userid *Player, bulletid int, fishid int) {
	// 增加CD判断，不可以太频繁

	// 删除子弹
	table.DelBullet(bulletid)

	// 获得鱼的金币
	fish := table.FishArray[fishid]
	if fish != nil {
		fmt.Println("捕获鱼成功，金币加")
	}
	// 删除鱼
	table.DelFish(fishid)

}


//  删除特定uid的子弹
func (table *CommonTable) DelBullet(bulletId int) {
	//for key, bullet := range table.BulletArray {
	//	if bullet.GetBulletUID() == bulletId {
	//		delete(table.BulletArray, key)
	//	}
	//}

	bulletHandle := table.BulletArray[bulletId]
	if bulletHandle != nil{
		delete(table.BulletArray, bulletId)
	}


	if len(table.BulletArray) == 0 {
		table.GenerateBulletUid = 0 //重置一下生成uuid
	}
}

//  删除所有子弹， 1 如果传入玩家uid，删除玩家的  ； 2  如果传入 -1 ，那么删除所有的
func (table *CommonTable)DelBullets(userId int) {
	for key, bullet := range table.BulletArray{
		if userId == -1 || bullet.GetBulletPlayerUID() == userId {
			delete(table.BulletArray, key)
		}
	}
	if len(table.BulletArray) == 0 {
		table.GenerateBulletUid = 0 //重置一下生成uuid
	}
}

//----------------------------------------------------------------------------
//------------------------------鱼-----------------------------------------
//----------------------------------------------------------------------------

// 新建一个新的鱼
func (table *CommonTable)CreateFish() FishInterface{
	//var fish * CommonFish
	fish := table.GetFishInterface()				// 获取鱼接口类型
	fishT := fish.NewFish(table.GenerateFishUid) // 生成鱼
	table.GenerateFishUid ++
	table.FishArray[fishT.GetFishUID()] = fishT // 把fish加入列表
	return fishT
}

// 获取鱼的句柄
func (table * CommonTable)GetFish(fishId int) FishInterface {
	fish := table.FishArray[fishId]
	return fish
}

//  删除特定uid的鱼
func (table *CommonTable) DelFish(fishId int) {
	//for key, fish := range table.FishArray {
	//	if fish.GetFishUID() == fishId {
	//		delete(table.FishArray, key)
	//	}
	//}
	fish := table.FishArray[fishId]
	if fish != nil {
		delete(table.FishArray, fishId)
	}

	if len(table.FishArray) == 0 {
		table.GenerateFishUid = 0 //重置一下生成uuid
	}
}

// 清空所有的鱼群
func (table *CommonTable)DelFishes() {
	for key := range table.FishArray{
		delete(table.FishArray, key)
	}
	if len(table.FishArray) == 0 {
		table.GenerateFishUid = 0 //重置一下生成uuid
	}
}



// 玩家登陆的时候， 同步给玩家场景中目前鱼群的信息
func (table *CommonTable) SendSceneFishes(user *Player){
	var SceneFishs []*CMD.TagSceneFish

	fmt.Println("玩家登陆，鱼数量",len(table.FishArray))

	for _,fish := range table.FishArray{
		Uid:= uint32(fish.GetFishUID())
		KindId:= uint32(fish.GetFishKindID())

		fishT := &CMD.TagSceneFish{Uid:&Uid,KindId:&KindId}
		SceneFishs = append(SceneFishs,fishT)
	}
	sendCmd1 := &CMD.CMD_S_SCENE_FISH{
		SceneFishs:SceneFishs,
	}
	ZServer.NetWorkSendByUid(user.GetUID(), sendCmd1, MDM_GF_GAME, SUB_S_SCENE_FISH)
	//NetWork.Send(UserSave.GetConn(), sendCmd1, MDM_GF_GAME, SUB_S_SCENE_FISH,"")  // 场景鱼刷新
}

// 给所有玩家同步新建的鱼的信息
func (table *CommonTable) SendNewFishes(fish FishInterface) {
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



//----------------------------------------------------------------------------
//-----------------------------消息同步-----------------------------------------
//----------------------------------------------------------------------------

// 给桌上的所有玩家同步消息
func (table *CommonTable) SendMsgToAllUsers(sendCmd proto.Message, mainCmd int, subCmd int) {
	for _, player := range table.UserSeatArray {
		if player != nil && player.CheckIsRobot() == false { // 有玩家，不是我，并且还不是机器人，那么发送
			//fmt.Println("给玩家",UserSave.GetUID(),"发送消息", subCmd)
			ZServer.NetWorkSendByUid(player.GetUID(), sendCmd,mainCmd,subCmd)
		}
	}
}

// 给桌上的其他玩家同步消息
func (table *CommonTable) SendMsgToOtherUsers(uid int, sendCmd proto.Message, mainCmd int, subCmd int) {
	for _, player := range table.UserSeatArray {
		if player != nil && player.GetUID() != uid && player.CheckIsRobot() == false { // 有玩家，不是我，并且还不是机器人，那么发送
			//fmt.Println("给玩家",UserSave.GetUID(),"发送消息", subCmd)
			ZServer.NetWorkSendByUid(player.GetUID(), sendCmd,mainCmd,subCmd)
		}
	}
}



//----------------------------------------------------------------------------
//---------------------------------------------------------------------
//----------------------------------------------------------------------------