package GameCore

import ("sync"
"../../Games/Model/PlayerModel"
	"fmt"
	"../Utils/zLog"
)

//----------------------------------------------------------------
//----------------------------定义--------------------------------
//----------------------------------------------------------------

// 全部游戏
var AllGamesList map[int]*Games // 保存所有游戏的列表
// 玩家
var AllPlayerList map[int] *PlayerModel.Player // 玩家的总列表 ， key： user id  ，  value：  user的指针
var AllPlayerListRWMutex       sync.RWMutex // 主要用于针对map进行读写时候的锁


// 游戏的类型定义
type Games struct {
	GameName string  // 游戏名字，用来显示的
	GameID  int // 游戏类型id
	GameSwitch bool	// 游戏开启的开关

	// 桌子
	AllTableList     map[int]TableInterface // 游戏的桌子列表,  key: table的uuid， value： table的指针
	TableUUID        int                    // 桌子的UUID生成, 自增的， 不用存数据库， 每次重新启动服务器之后，重新生成即可
	RWMutexTableList sync.RWMutex           // 主要用于针对map进行读写时候的锁

}

func NewGame(name string, gameid int, open bool) *Games {
	return &Games{GameName:name, GameID:gameid, GameSwitch:open, AllTableList:make(map[int]TableInterface), TableUUID:1}
}


// 通过gameID获取是哪个游戏
func GetGameByID(gameId int) *Games {
	game := AllGamesList[gameId] // 客户端会保留对应登录游戏的句柄， 方便对游戏的调用
	if game == nil {
		zLog.PrintfLogger("找不到  game ID ： %d   ， 请检查游戏 ID 是否正确  \n" , gameId)
	}
	return game
}


// 根据桌子uid 返回桌子的句柄
func  (games *Games)GetTableByUID(tableId int) TableInterface {
	//fmt.Println("table Id :" ,tableId , " game id:", games.GameID)

	games.RWMutexTableList.RLock()
	table := games.AllTableList[tableId]
	games.RWMutexTableList.RUnlock()

	//println("GetTableByUID table id ", table.GetTableUID(),  "  " ,table.GetRoomScore())
	if table == nil {
		zLog.PrintfLogger("找不到  table ID ： %d   ， 请检查 table ID 是否正确  \n" , tableId)
	}
	return table
}


// 根据user uid 返回user的句柄
func  GetPlayerByUID(uid int) *PlayerModel.Player {
	AllPlayerListRWMutex.RLock()
	player := AllPlayerList[uid]
	AllPlayerListRWMutex.RUnlock()
	if player == nil{
		zLog.PrintfLogger("找不到  player ID ： %d   ， 请检查 player ID 是否正确  \n" , uid)
	}
	return player
}

// 添加到所有玩家列表中
func AddPlayerToAllPlayerList(player *PlayerModel.Player)  {
	AllPlayerListRWMutex.Lock()
	AllPlayerList[int(player.GetUID())] = player
	AllPlayerListRWMutex.Unlock()
}




//----------------------------------------------------------------
//-----------------------------管理桌子---------------------------------------------
//----------------------------------------------------------------




// 创建桌子，并启动它
func (games * Games) CreateTable(table TableInterface, gameScore int) int {


	//  创建游戏针对性的桌子内存，并获取到了具体的游戏的桌子的句柄
	thisTable := table.NewTable(games.TableUUID, games.GameID)
	thisTable.SetRoomScore(gameScore)
	fmt.Println("创建", games.GameName, "的一个新桌子 ID:", games.TableUUID)

	// 增加该桌子到总列表中
	games.RWMutexTableList.Lock()
	games.AllTableList[games.TableUUID] = thisTable // 加入桌子列表
	games.TableUUID ++
	games.RWMutexTableList.Unlock()

	// 桌子开始自行启动计算
	thisTable.InitTable()
	go thisTable.RunTable()

	return thisTable.GetTableUID()
}






//----------------------------------------------------------------
//-----------------------------管理玩家---------------------------------------------
//----------------------------------------------------------------




// 有玩家登陆游戏，想进入对应分数的房间
func (games *Games)PlayerLoginGame(newPlayer *PlayerModel.Player,  gameScore int) int{
	// 根据游戏类型创建游戏中玩家的句柄
	uid := newPlayer.GetUID()
	//userHandle := games.AllTableList[0].GetUserInterfaceHandle() // 获取玩家类型句柄
	//newPlayer := CommonLogin.NewPlayer(uid)                       // 创建游戏中玩家数据
	//newPlayer.SetUser(User)                                      // 设置player的句柄给user
	//newPlayer.SetGameID(gameID)


	//// 创建好之后加入玩家总列表
	//AllPlayerListRWMutex.Lock()
	//AllPlayerList[uid] = newPlayer
	//AllPlayerListRWMutex.Unlock()

	//newPlayer := games.GetPlayerByUID(uid)
	//games := GetGameByID(gameid)

	// 然后找一个有空位的桌子让玩家加入游戏
	games.RWMutexTableList.RLock()
	for _, table := range games.AllTableList {
		if table.GetRoomScore() == gameScore {		// 找分数一致的房间
			seatId := table.GetEmptySeatInTable() //先查找一下所有开放着的桌子是否有空位
			if seatId >= 0 {
				fmt.Println(uid, "---想坐下----有空的座位----")
				table.InitTable()						// 看看是不是空桌子，如果是空桌子需要初始化

				table.PlayerSeat(seatId, newPlayer)       //桌子上让玩家坐下
				newPlayer.SetTableID(table.GetTableUID()) // 玩家坐下
				newPlayer.SetChairID(seatId)              // 玩家坐到椅子上

				return table.GetTableUID()
			}
		}
	}
	games.RWMutexTableList.RUnlock()

	fmt.Println("------没有空的座位了，新建一个桌子吧----------底分",gameScore)
	//tableUid := games.CreateTableByType(games.GameID)
	tableUid := games.CreateTable(games.AllTableList[0],gameScore)
	table := games.GetTableByUID(tableUid)
	seatId := table.GetEmptySeatInTable()		//获取空椅位
	//table.InitTable()						// 桌子初始化

	table.PlayerSeat(seatId, newPlayer)       //桌子上让玩家坐下
	newPlayer.SetTableID(table.GetTableUID()) // 玩家坐下
	newPlayer.SetChairID(seatId)              // 玩家坐到椅子上
	return tableUid
}

// 玩家登出
func (games *Games)PlayerLogOutGame( player *PlayerModel.Player) {
	//player := games.GetPlayerByUID(uid)
	table := games.GetTableByUID(player.GetTableID())
	table.PlayerStandUp(player.GetChairID(), player) // 玩家离开桌子
	//player := games.GetPlayerByUID(uid)
	player.SetTableID(-1) //玩家桌子数据清空
	player.SetChairID(-1) //玩家椅子数据清空

	// 从总列表中删除玩家
	AllPlayerListRWMutex.Lock()
	delete(AllPlayerList, player.GetUID())
	AllPlayerListRWMutex.Unlock()
	fmt.Println("玩家", player.GetUID(), "离开了")
}


