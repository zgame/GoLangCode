package BY

import (
	"fmt"
	"time"
	"../../Const"
	"../../Core/GameCore"
	."../Model/PlayerModel"
	. "./BYModel"
	"sync"
)

//--------------------------------------------------------------------------------------
// 该类为桌子的通用父类， 这里可以定义一些通用的函数，每个游戏具体的桌子不定义的话，就用父类， 也可以重新定义实现重载
//--------------------------------------------------------------------------------------



// -------------------------桌子的统一结构-------------------------------------
type BYTable struct {
	GameCore.TableBase				// 继承桌子基类
	//GameID    int // 游戏类型
	//TableID   int // 桌子id
	//TableMax  int // 桌子容纳玩家数量
	//RoomScore int // 房间分数
	////State int			//桌子状态		0 开放中，有玩家正在玩 ； 1 关闭中，没有玩家在游戏
	//
	//UserSeatArray       map[int]*Player // 座椅对应玩家uid的哈希表 ， key ： seat ，value： 玩家uid

	//UserCnt int			// 玩家数量
	//StateChgTime time.Time	 // 状态变更时间

	//
	GenerateFishUid       int                     //生成鱼的uid
	GenerateBulletUid     int                     // 生成子弹的uid

	FishArray             map[int]*CommonFish   // 鱼的哈希表
	BulletArray           map[int]*CommonBullet // 子弹的哈希表
	RWMutexFishArray       sync.RWMutex                     // 主要用于针对map进行读写时候的锁
	RWMutexBulletArray       sync.RWMutex                     // 主要用于针对map进行读写时候的锁

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
func (table *BYTable) NewTable(uid int, gameId int) GameCore.TableInterface {
	fmt.Println("创建 捕鱼 BY  桌子")
	byTable := &BYTable{}
	byTable.TableID = uid
	byTable.GameID = gameId
	byTable.SetTableMax(Const.BY_TABLE_MAX_PLAYER)
	return byTable
	//return GameCore.TableBase.NewTable(uid , gameId)
}

// 当获取到桌子接口的时候， 我们可以强转成我们当前table的指针
func GetBYTableHandle(tableInterface GameCore.TableInterface) *BYTable{
	return tableInterface.(*BYTable)
}

//---------------------------------桌子的状态------------------------------------------



// 初始化桌子
func (table *BYTable) InitTable()  {
	if table.CheckTableEmpty(){

		table.BulletArray = make(map[int] *CommonBullet )
		table.FishArray = make(map[int]*CommonFish )
		table.UserSeatArray = make(map[int]*Player)

		table.InitDistributeInfo(table.GetRoomScore())	//玩家进来之后，如果是第一个玩家，那么房间开始初始化一下
	}
}

// --------------------------- 玩家离开椅子 -----------------------
func (table *BYTable)PlayerStandUp(seatId int, user *Player) {
	table.RWMutexSeatArray.Lock()
	table.UserSeatArray[seatId] = nil
	table.RWMutexSeatArray.Unlock()

	// 清理掉玩家的所有子弹
	table.DelBullets(user.GetUID())
	// 如果是空桌子的话，清理一下桌子
	if table.CheckTableEmpty() {
		table.ClearTable()
	}
}
// 清理桌子
func (table *BYTable)ClearTable()  {
	// 清理掉生成鱼群的结构
	table.DistributeArray = nil
	table.BossDistributeArray = nil		// 因为玩家进来之后会初始化会make，所以就不用make了

	// 清理掉所有的子弹和鱼群
	table.DelBullets(-1)
	table.DelFishes()

	table.RWMutexBulletArray.Lock()
	table.BulletArray = nil
	table.RWMutexBulletArray.Unlock()

	table.RWMutexFishArray.Lock()
	table.FishArray = nil
	table.RWMutexFishArray.Unlock()

	table.RWMutexSeatArray.Lock()
	table.UserSeatArray = nil
	table.RWMutexSeatArray.Unlock()
}


//----------------------------------------------------------------------------
// -----------------------------  桌子主循环逻辑 -------------------------------
//----------------------------------------------------------------------------
func (table *BYTable)RunTable() {
	for {
		if table.CheckTableEmpty()  {
			//fmt.Println("这是一个空桌子")
		}else{
			//	fmt.Println("运行桌子", table.TableID, "游戏类型", table.GameID)
			table.RunDistributeInfo(table.GetRoomScore())
			table.RunBossDistributeInfo(table.GetRoomScore())


			table.RWMutexBulletArray.RLock()
			for _,bullet := range table.BulletArray{
				table.BulletRun(bullet)				// 遍历所有子弹，并且run
			}
			table.RWMutexBulletArray.RUnlock()

			table.RWMutexFishArray.RLock()
			for _,fish := range table.FishArray{
				table.FishRun(fish)					// 遍历所有鱼，并且run
			}
			table.RWMutexFishArray.RUnlock()

			fmt.Println(table.GetTableUID(),"共有子弹数量",len(table.BulletArray))
			fmt.Println(table.GetTableUID(),"共有鱼数量",len(table.FishArray))
		}
		time.Sleep(time.Millisecond * 1000)
	}
}


// -------------------------子弹的主循环------------------------
func (table *BYTable) BulletRun(bullet *CommonBullet)   {
	//fmt.Println("子弹的run", bullet.GetBulletUID())
	// 子弹的生存时间到了
	if time.Now().After(bullet.DeadTime){
		table.DelBullet(bullet.BulletUID)		// 生存时间已经到了，销毁
		return
	}
	// 子弹的移动

	// 子弹是否击中的判断

}

// -------------------------鱼的主循环-------------------------
func (table *BYTable) FishRun(fish *CommonFish) {
	//fmt.Println("fish的run")

	// 鱼的生存时间到了
	if time.Now().After(fish.DeadTime){
		fmt.Println("生存时间到了")
		table.DelFish(fish.FishUID)		// 生存时间已经到了，销毁
		return
	}
	// 鱼的移动

	// 生成鱼


}



//----------------------------------------------------------------------------
//------------------------------bullet 子弹-----------------------------------------
//----------------------------------------------------------------------------
// 玩家发射一个新的子弹
func (table *BYTable)FireBullet(player *Player, lockFishId int) int{
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
	//bullet := table.GetBulletInterface()				// 获取子弹接口类型
	bulletT := NewBullet(table.GenerateBulletUid) // 生成子弹
	table.GenerateBulletUid ++

	table.RWMutexBulletArray.Lock()
	table.BulletArray[bulletT.GetBulletUID()] = bulletT // 把bullet加入列表
	table.RWMutexBulletArray.Unlock()

	bulletT.SetBulletPlayerUID(player.GetUID())         //子弹的主人
	bulletT.SetLockFishID(lockFishId)                   // 锁定鱼
	player.SetActivityBulletNum(num+1)                  //玩家已激活子弹增加
	return bulletT.GetBulletUID()
}

//  击中一条鱼
func (table *BYTable) HitFish(userid *Player, bulletid int, fishid int) {
	// 增加CD判断，不可以太频繁

	// 删除子弹
	table.DelBullet(bulletid)

	// 获得鱼的金币
	fish := table.GetFish(fishid)
	if fish != nil {
		fmt.Println("捕获鱼成功，金币加")
	}
	// 删除鱼
	table.DelFish(fishid)

}


//  删除特定uid的子弹
func (table *BYTable) DelBullet(bulletId int) {
	//for key, bullet := range table.BulletArray {
	//	if bullet.GetBulletUID() == bulletId {
	//		delete(table.BulletArray, key)
	//	}
	//}

	table.RWMutexBulletArray.Lock()
	bulletHandle := table.BulletArray[bulletId]
	if bulletHandle != nil{
		delete(table.BulletArray, bulletId)
	}
	table.RWMutexBulletArray.Unlock()


	if len(table.BulletArray) == 0 {
		table.GenerateBulletUid = 0 //重置一下生成uuid
	}
}

//  删除所有子弹， 1 如果传入玩家uid，删除玩家的  ； 2  如果传入 -1 ，那么删除所有的
func (table *BYTable)DelBullets(userId int) {
	table.RWMutexBulletArray.Lock()
	defer  table.RWMutexBulletArray.Unlock()

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
func (table *BYTable)CreateFish() *CommonFish{
	//var fish * CommonFish
	//fish := table.GetFishInterface()				// 获取鱼接口类型
	fishT := NewFish(table.GenerateFishUid) // 生成鱼
	table.GenerateFishUid ++

	table.RWMutexFishArray.Lock()
	table.FishArray[fishT.GetFishUID()] = fishT // 把fish加入列表
	table.RWMutexFishArray.Unlock()

	return fishT
}


// 获取鱼的句柄
func (table *BYTable)GetFish(fishId int)  *CommonFish {
	table.RWMutexFishArray.RLock()
	fish := table.FishArray[fishId]
	table.RWMutexFishArray.RUnlock()
	return fish
}

//  删除特定uid的鱼
func (table *BYTable) DelFish(fishId int) {
	//for key, fish := range table.FishArray {
	//	if fish.GetFishUID() == fishId {
	//		delete(table.FishArray, key)
	//	}
	//}
	fish := table.GetFish(fishId)
	if fish != nil {
		table.RWMutexFishArray.Lock()
		delete(table.FishArray, fishId)
		table.RWMutexFishArray.Unlock()
	}

	if len(table.FishArray) == 0 {
		table.GenerateFishUid = 0 //重置一下生成uuid
	}
}

// 清空所有的鱼群
func (table *BYTable)DelFishes() {
	table.RWMutexFishArray.Lock()
	for key := range table.FishArray{
		delete(table.FishArray, key)
	}
	defer table.RWMutexFishArray.Unlock()

	if len(table.FishArray) == 0 {
		table.GenerateFishUid = 0 //重置一下生成uuid
	}
}




//----------------------------------------------------------------------------
//---------------------------------------------------------------------
//----------------------------------------------------------------------------