package GameCore

import (
	"../../Const"
	"../../Games/Model/PlayerModel"
	"github.com/golang/protobuf/proto"
	"../ZServer"
	"fmt"
	"time"
	"sync"
)

// -------------------------桌子的统一定义接口-------------------------------------
type TableInterface interface {
	NewTable(uid int, gameId int) TableInterface //桌子构造函数
	GetTableUID() int                            // 获取桌子uid
	GetRoomScore() int                           // 获取房间的分数
	SetRoomScore(roomScore int)                  // 设定房间的分数
	GetTableMax() int                           // 获取桌子容纳玩家数量
	SetTableMax(tableMax int)                  // 设定桌子容纳玩家数量

	// 椅子逻辑
	GetEmptySeatInTable() int   // 获取空椅子的编号
	GetUsersSeatInTable() []int // 获取桌子上的所有玩家uid
	CheckTableEmpty() bool      // 判断桌子是有人，还是空桌子

	PlayerSeat(seatId int, user *PlayerModel.Player)    // 玩家坐到椅子上
	PlayerStandUp(seatId int, user *PlayerModel.Player) // 玩家离开椅子

	InitTable()  // 初始化桌子
	ClearTable() // 清空桌子
	RunTable()   // 桌子的逻辑启动

	EnterSceneSyncMsg(player *PlayerModel.Player)
	SendMsgToOtherUsers(uid int, sendCmd proto.Message, mainCmd int, subCmd int) // 给桌上的其他玩家同步消息
	SendMsgToAllUsers(sendCmd proto.Message, mainCmd int, subCmd int)            // 给桌上的所有玩家同步消息
}

// -------------------------桌子的统一结构-------------------------------------
type TableBase struct {
	GameID    int // 游戏类型
	TableID   int // 桌子id
	TableMax  int // 桌子容纳玩家数量
	RoomScore int // 房间分数
	//State int			//桌子状态		0 开放中，有玩家正在玩 ； 1 关闭中，没有玩家在游戏

	UserSeatArray map[int]*PlayerModel.Player // 座椅对应玩家uid的哈希表 ， key ： seat ，value： 玩家uid

	//UserCnt int			// 玩家数量
	//StateChgTime time.Time	 // 状态变更时间
	RWMutexSeatArray sync.RWMutex // 主要用于针对map进行读写时候的锁
}

// -------------------------构造函数-------------------------
func (table *TableBase) NewTable(uid int, gameId int) TableInterface {
	fmt.Println("创建桌子类型： TableBase")
	return &TableBase{TableID: uid, GameID: gameId, TableMax: Const.BY_TABLE_MAX_PLAYER}
}

// 返回桌子的UID
func (table *TableBase) GetTableUID() int {
	return table.TableID
}

// 获取房间的分数
func (table *TableBase) GetRoomScore() int {
	return table.RoomScore
}

// 设定房间的分数
func (table *TableBase) SetRoomScore(roomScore int) {
	table.RoomScore = roomScore
}
// 获取桌子的座位数
func (table *TableBase) GetTableMax() int {
	return table.RoomScore
}
// 设定桌子的座位数
func (table *TableBase) SetTableMax(tableMax int) {
	table.TableMax = tableMax
}

// 获取椅子上面的玩家
func (table *TableBase)GetSeatArray(index int) *PlayerModel.Player {
	table.RWMutexSeatArray.RLock()
	defer 	table.RWMutexSeatArray.RUnlock()
	return table.UserSeatArray[index]
}


//----------------------------------------------------------------------------
// -------------------------玩家逻辑--------------------------------------------------
//----------------------------------------------------------------------------

// -------------------------判断桌子是有人，还是空桌子---------------------------------------------------
func (table *TableBase) CheckTableEmpty() bool {
	// 遍历所有的座位
	for i := 0; i < table.TableMax; i++ {
		if table.GetSeatArray(i) != nil {
			return false // 有人在玩
		}
	}
	return true // 空桌子
}

// -----------------获取桌子的所有玩家----------------
func (table *TableBase) GetUsersSeatInTable() []int {
	userList := make([]int, 0)
	// 遍历所有的座位
	for i := 0; i < table.TableMax; i++ {
		player := table.GetSeatArray(i)
		if player != nil {
			//如果座位上有玩家uid，说明有人
			userList = append(userList, player.GetUID())
		}
	}
	return userList
}

// -----------------获取桌子的空座位, 返回座椅的编号，从0开始到tableMax， 如果返回-1说明满了----------------
func (table *TableBase) GetEmptySeatInTable() int {
	// 遍历所有的座位
	for i := 0; i < table.TableMax; i++ {
		if table.GetSeatArray(i) == nil {
			//如果座位上没有玩家uid，说明没有人，返回座位id
			fmt.Println(i, "这个位置有空着椅子")
			return i
		}
	}
	// 如果座椅都满了， 那么返回-1
	fmt.Println(table.GetTableUID(), "都满了")
	return -1
}

// ---------------------------玩家坐到椅子上-------------------------------
func (table *TableBase) PlayerSeat(seatId int, user *PlayerModel.Player) {
	table.RWMutexSeatArray.Lock()
	table.UserSeatArray[seatId] = user
	table.RWMutexSeatArray.Unlock()
}

// --------------------------- 玩家离开椅子 -----------------------
func (table *TableBase) PlayerStandUp(seatId int, user *PlayerModel.Player) {
	table.RWMutexSeatArray.Lock()
	table.UserSeatArray[seatId] = nil
	table.RWMutexSeatArray.Unlock()
	// 如果是空桌子的话，清理一下桌子
	if table.CheckTableEmpty() {
		table.ClearTable()
	}
}

// 初始化桌子
func (table *TableBase) InitTable() {
	if table.CheckTableEmpty() {
		table.UserSeatArray = make(map[int]*PlayerModel.Player)
	}
}

// 清理桌子
func (table *TableBase) ClearTable() {
	table.RWMutexSeatArray.Lock()
	defer table.RWMutexSeatArray.Unlock()
	table.UserSeatArray = nil
}

//----------------------------------------------------------------------------
// -----------------------------  桌子主循环逻辑 -------------------------------
//----------------------------------------------------------------------------
func (table *TableBase) RunTable() {
	for {
		if table.CheckTableEmpty() {
			//fmt.Println("这是一个空桌子")
		} else {
			fmt.Println("运行桌子", table.TableID, "游戏类型", table.GameID)
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

//----------------------------------------------------------------------------
//-----------------------------消息同步-----------------------------------------
//----------------------------------------------------------------------------

// 进入场景之后的首次消息同步， 每个游戏不一样， 都会重载该方法
func (table *TableBase) EnterSceneSyncMsg(player *PlayerModel.Player)  {
	// 其他子类会重写该方法
	fmt.Println(" 这是 父类 EnterSceneSyncMsg 方法")
}


// 给桌上的所有玩家同步消息
func (table *TableBase) SendMsgToAllUsers(sendCmd proto.Message, mainCmd int, subCmd int) {
	table.RWMutexSeatArray.RLock()
	defer 	table.RWMutexSeatArray.RUnlock()
	for _, player := range table.UserSeatArray {
		if player != nil && player.CheckIsRobot() == false { // 有玩家，不是我，并且还不是机器人，那么发送
			//fmt.Println("给玩家",UserModel.GetUID(),"发送消息", subCmd)
			ZServer.NetWorkSendByUid(player.GetUID(), sendCmd, mainCmd, subCmd)
		}
	}
}

// 给桌上的其他玩家同步消息
func (table *TableBase) SendMsgToOtherUsers(uid int, sendCmd proto.Message, mainCmd int, subCmd int) {
	table.RWMutexSeatArray.RLock()
	defer 	table.RWMutexSeatArray.RUnlock()
	for _, player := range table.UserSeatArray {
		if player != nil && player.GetUID() != uid && player.CheckIsRobot() == false { // 有玩家，不是我，并且还不是机器人，那么发送
			//fmt.Println("给玩家",UserModel.GetUID(),"发送消息", subCmd)
			ZServer.NetWorkSendByUid(player.GetUID(), sendCmd, mainCmd, subCmd)
		}
	}
}


