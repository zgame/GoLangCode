package Games

import (
	"fmt"
	"./BY"
	"./BY2"
	"./BY3"
	"./Common"
	//"../Const"
	"../Logic/Player"
)

//----------------------------------------------------------------
// 该文件是所有游戏的管理类， 掌握所有游戏的列表
// 单个游戏的结构也掌握所有桌子的列表
//----------------------------------------------------------------

// 所有游戏的编号
const(
	GameTypeBY  = 1
	GameTypeBY2 = 2
	GameTypeBY3 = 3
	// 更多游戏编号添加到此处...
	// ...
	// ...
)
//----------------------------------------------------------------
//----------------------------定义------------------------------------
//----------------------------------------------------------------

// 全部游戏
var AllGamesList map[int]*Games // 保存所有游戏的列表


// 游戏的类型定义
type Games struct {
	GameName string  // 游戏名字，用来显示的
	GameID  int // 游戏类型id
	GameSwitch bool	// 游戏开启的开关

	// 桌子
	AllTableList map[int]Common.TableInterface // 游戏的桌子列表,  key: table的uuid， value： table的指针
	TableUUID    int                           // 桌子的UUID生成, 自增的， 不用存数据库， 每次重新启动服务器之后，重新生成即可

	// 玩家
	AllUserList map[int] Common.UserInterface // 玩家的总列表 ， key： user id  ，  value：  user的指针

}

func (games *Games)NewGame(name string, gameid int, open bool) *Games {
	return &Games{GameName:name, GameID:gameid, GameSwitch:open, AllTableList:make(map[int]Common.TableInterface), TableUUID:0, AllUserList:make(map[int]Common.UserInterface)}
}

//----------------------------------------------------------------
//---------------------------管理游戏----------------------------------------
//----------------------------------------------------------------

// 增加一个游戏， 指定这个游戏的类型， 并且创建一个桌子，并启动桌子逻辑
func AddGame(name string, gameType int) {
	var games *Games
	games = games.NewGame(name, gameType, true) // 创建游戏房间
	AllGamesList[gameType] = games

	// 根据游戏类型进行判断
	if gameType == GameTypeBY{
		var table  * BY.Table
		games.CreateTable(table,1)		//创建1底分桌子
		//games.CreateTable(table,100)	//创建100底分桌子
		//games.CreateTable(table,10000)	//创建10000底分桌子
	}else if gameType == GameTypeBY2{
		var table   * BY2.Table
		games.CreateTable(table,1)
	}else if gameType == GameTypeBY3{
		var table * BY3.Table
		games.CreateTable(table,1)
	}
	// 增加游戏类型...
	//...
	//...
	//...

}

// 通过gameID获取是哪个游戏
func GetGameByID(gameid int) *Games {
	return AllGamesList[gameid] // 客户端会保留对应登录游戏的句柄， 方便对游戏的调用
}

//----------------------------------------------------------------
//-----------------------------管理桌子---------------------------------------------
//----------------------------------------------------------------

//// 根据游戏类型，增加桌子
//func (games *Games) CreateTableByType(gameType int) int{
//	var tableUid int
//	if gameType == GameTypeBY{
//		var table  * BY.Table
//		tableUid = games.CreateTable(table)
//	}else if gameType == GameTypeBY2{
//		var table   * BY2.Table
//		tableUid = games.CreateTable(table)
//	}else if gameType == GameTypeBY3{
//		var table * BY3.Table
//		tableUid = games.CreateTable(table)
//	}
//	// 增加游戏类型...
//	//...
//	//...
//	//...
//
//	return tableUid
//}

// 创建桌子，并启动它
func (games *Games) CreateTable(table Common.TableInterface, gameScore int) int {

	//  创建游戏针对性的桌子内存，并获取到了具体的游戏的桌子的句柄
	thisTablePoint := table.NewTable(games.TableUUID, games.GameID)
	thisTablePoint.SetBulletInterface()
	thisTablePoint.SetFishInterface()
	thisTablePoint.SetUserInterfaceHandle()
	thisTablePoint.SetRoomScore(gameScore)
	fmt.Println("创建", games.GameName, "的一个新桌子", games.TableUUID)

	// 增加该桌子到总列表中
	games.AllTableList[games.TableUUID] = thisTablePoint
	games.TableUUID ++
	// 桌子开始自行启动计算
	thisTablePoint.InitTable()
	go thisTablePoint.RunTable()

	return thisTablePoint.GetTableUID()
}




// 根据桌子uid 返回桌子的句柄
func  (games *Games)GetTableByUID(uid int) Common.TableInterface{
	return games.AllTableList[uid]
}



//----------------------------------------------------------------
//-----------------------------管理玩家---------------------------------------------
//----------------------------------------------------------------

// 根据user uid 返回user的句柄
func  (games *Games)GetUserByUID(uid int) Common.UserInterface{
	return games.AllUserList[uid]
}


// 有玩家登陆游戏，想进入对应分数的房间
func (games *Games)PlayerLoginGame(player *Player.Player, gameScore int) int{
	// 根据游戏类型创建游戏中玩家的句柄
	uid := int(player.UserId)
	userHandle := games.AllTableList[0].GetUserInterfaceHandle() // 获取玩家类型句柄
	user := userHandle.NewUser(uid)                              // 创建游戏中玩家数据
	user.SetPlayer(player)									// 设置player的句柄给user

	// 创建好之后加入玩家总列表
	games.AllUserList[uid] = user

	//user := games.GetUserByUID(uid)
	//games := GetGameByID(gameid)

	// 然后找一个有空位的桌子让玩家加入游戏
	for _, table := range games.AllTableList {
		if table.GetRoomScore() == gameScore {		// 找分数一致的房间
			seatId := table.GetEmptySeatInTable() //先查找一下所有开放着的桌子是否有空位
			if seatId >= 0 {
				fmt.Println(uid, "---想坐下----有空的座位----")
				table.InitTable()						// 看看是不是空桌子，如果是空桌子需要初始化

				table.PlayerSeat(seatId, user)       //桌子上让玩家坐下
				user.SetTableID(table.GetTableUID()) // 玩家坐下
				user.SetChairID(seatId)              // 玩家坐到椅子上

				return table.GetTableUID()
			}
		}
	}
	fmt.Println("------没有空的座位了，新建一个桌子吧----------底分",gameScore)
	//tableUid := games.CreateTableByType(games.GameID)
	tableUid := games.CreateTable(games.AllTableList[0],gameScore)
	table := games.GetTableByUID(tableUid)
	seatId := table.GetEmptySeatInTable()		//获取空椅位
	table.InitTable()						// 桌子初始化

	table.PlayerSeat(seatId, user)		//桌子上让玩家坐下
	user.SetTableID(table.GetTableUID())	// 玩家坐下
	user.SetChairID(seatId)					// 玩家坐到椅子上
	return tableUid
}

// 玩家登出
func (games *Games)PlayerLogOutGame( user Common.UserInterface)  {
	//user := games.GetUserByUID(uid)
	table := games.GetTableByUID(user.GetTableID())
	table.PlayerStandUp(user.GetChairID(), user)		// 玩家离开桌子
	//user := games.GetUserByUID(uid)
	user.SetTableID(-1)			//玩家桌子数据清空
	user.SetChairID(-1)		//玩家椅子数据清空
	fmt.Println("玩家",user.GetUID(), "离开了")
}


//---------------------------------------消息分发-------------------------
//
//// 发送消息给桌子的其他用户
//func (games *Games)SendMsgToTablePlayers()  {
//
//
//}