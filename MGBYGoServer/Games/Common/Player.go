package Common


import (
	"../../Const"
	"../Model/UserSave"
)

//----------------------------------------------------------------------------------
// Common 定义的结构（类） 是基类， 具体游戏有共同的逻辑部分可以归纳到Common里面， 具体游戏可以继承，也可以重载
//----------------------------------------------------------------------------------




//// -------------------------玩家的统一定义接口-------------------------------------
//type Player interface{
//	NewPlayer(uid int) Player
//	GetUID() int			// 获取玩家的uid
//	GetTableID() int		// 获取玩家的桌子
//	SetTableID(tid int)	// 设置玩家的桌子
//	GetChairID() int		// 获取玩家的椅子
//	SetChairID(ChairID int)	// 设置玩家的椅子
//
//	GetGameID() int
//	SetGameID(gameID int)
//
//	GetActivityBulletNum() int		// 获取玩家的发射子弹数量
//	SetActivityBulletNum(num int)	// 设置玩家的发射子弹数量
//
//	GetUser() *UserSave.UserSave       // 获取玩家的总数据
//	SetUser(player *UserSave.UserSave) // 设定玩家的总数据
//
//	//GetConn() net.Conn			// 获取玩家的socket
//	//SetConn(conn net.Conn)			// 设置玩家的socket句柄
//
//	CheckIsRobot() bool			// 是不是机器人判断
//	SetIsRobot(isRobot bool)	// 设定是机器人
//}



// ------------------------玩家的结构----------------------------
type Player struct {
	//conn    net.Conn
	UserSave * UserSave.UserSave
	//UID      int  // 用户uid
	GameID   int  // 游戏id
	TableID  int  // 桌子id
	ChairID  int  // 椅子id
	IsRobot  bool // 是不是机器人

	ActivityBulletNum int  //当前已经发射的子弹数量
	//BulletMulriple int 	// 炮的倍率
	//MissFishFixLib int 	// miss 库
	//NewSuportSwitch	int  //新手扶持开关
	//NewSuportBulletNum int //新手扶持子弹数量
	//

}


// -------------------------构造函数-------------------------
func  NewPlayer(user *UserSave.UserSave) *Player {
	return &Player{UserSave:user,TableID:Const.TABLE_CHAIR_NOBODY,ChairID:Const.TABLE_CHAIR_NOBODY, IsRobot:false}
}

// 获取玩家的uid
func (player *Player) GetUID() int {
	return int(player.UserSave.UserId)
}

// 获取玩家游戏类型
func (player *Player) GetGameID() int{
	return player.GameID
}

//  设置玩家游戏类型
func (player *Player) SetGameID(gameID int) {
	player.GameID = gameID
}



// 获取玩家的总数据
func (player *Player) GetUser() *UserSave.UserSave {
	return player.UserSave
}
// 设置玩家的总数据
func (player *Player) SetUser(save *UserSave.UserSave) {
	player.UserSave = save
}


// 是不是机器人判断
func (player *Player) CheckIsRobot() bool {
	return player.IsRobot
}
// 设定是机器人
func (player *Player) SetIsRobot(isRobot bool) {
	player.IsRobot = isRobot
}


// 获取玩家的桌子
func (player *Player) GetTableID() int {
	return player.TableID
}
// 设置玩家的桌子
func (player *Player) SetTableID(tid int)  {
	player.TableID = tid
}
// 获取玩家的椅子
func (player *Player) GetChairID() int {
	return player.ChairID
}
// 设置玩家的椅子
func (player *Player) SetChairID(ChairID int)  {
	player.ChairID = ChairID
}

// 获取玩家的发射子弹数量
func (player *Player) GetActivityBulletNum() int {
	return player.ActivityBulletNum
}
// 设置玩家的发射子弹数量
func (player *Player) SetActivityBulletNum(num int)  {
	player.ActivityBulletNum = num
}




