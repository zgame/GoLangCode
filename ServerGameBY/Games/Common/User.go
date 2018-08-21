package Common

import "../../Logic/Player"
import (
	"../../Const"
	"net"
)

//----------------------------------------------------------------------------------
// Common 定义的结构（类） 是基类， 具体游戏有共同的逻辑部分可以归纳到Common里面， 具体游戏可以继承，也可以重载
//----------------------------------------------------------------------------------




// -------------------------玩家的统一定义接口-------------------------------------
type UserInterface interface{
	NewUser(uid int)  UserInterface
	GetUID() int			// 获取玩家的uid
	GetTableID() int		// 获取玩家的桌子
	SetTableID(tid int)	// 设置玩家的桌子
	GetChairID() int		// 获取玩家的椅子
	SetChairID(ChairID int)	// 设置玩家的椅子

	GetActivityBulletNum() int		// 获取玩家的发射子弹数量
	SetActivityBulletNum(num int)	// 设置玩家的发射子弹数量

	GetPlayer() *Player.Player			// 获取玩家的总数据
	SetPlayer(player *Player.Player)	// 设定玩家的总数据

	GetConn() net.Conn			// 获取玩家的socket
	SetConn(conn net.Conn)			// 设置玩家的socket句柄

	CheckIsRobot() bool			// 是不是机器人判断
	SetIsRobot(isRobot bool)	// 设定是机器人
}



// ------------------------玩家的结构----------------------------
type CommonUser struct {
	conn    net.Conn
	player  * Player.Player
	UID     int  // 用户uid
	TableID int  // 桌子id
	ChairID int  // 椅子id
	IsRobot bool // 是不是机器人

	ActivityBulletNum int  //当前已经发射的子弹数量
	//BulletMulriple int 	// 炮的倍率
	//MissFishFixLib int 	// miss 库
	//NewSuportSwitch	int  //新手扶持开关
	//NewSuportBulletNum int //新手扶持子弹数量
	//

}


// -------------------------构造函数-------------------------
func (user *CommonUser) NewUser(uid int) UserInterface{
	return &CommonUser{UID:uid,TableID:Const.TABLE_CHAIR_NOBODY,ChairID:Const.TABLE_CHAIR_NOBODY, IsRobot:false}
}

// 获取玩家的uid
func (user *CommonUser) GetUID() int {
	return user.UID
}

// 获取玩家的socket
func (user *CommonUser) GetConn() net.Conn {
	return user.conn
}
// 设置玩家的socket句柄
func (user *CommonUser) SetConn(conn net.Conn) {
	user.conn = conn
}


// 获取玩家的总数据
func (user *CommonUser) GetPlayer() *Player.Player {
	return user.player
}
// 设置玩家的总数据
func (user *CommonUser) SetPlayer(player *Player.Player) {
	user.player = player
}


// 是不是机器人判断
func (user *CommonUser) CheckIsRobot() bool {
	return user.IsRobot
}
// 设定是机器人
func (user *CommonUser) SetIsRobot(isRobot bool) {
	user.IsRobot = isRobot
}


// 获取玩家的桌子
func (user *CommonUser) GetTableID() int {
	return user.TableID
}
// 设置玩家的桌子
func (user *CommonUser) SetTableID(tid int)  {
	user.TableID = tid
}
// 获取玩家的椅子
func (user *CommonUser) GetChairID() int {
	return user.ChairID
}
// 设置玩家的椅子
func (user *CommonUser) SetChairID(ChairID int)  {
	user.ChairID = ChairID
}

// 获取玩家的发射子弹数量
func (user *CommonUser) GetActivityBulletNum() int {
	return user.ActivityBulletNum
}
// 设置玩家的发射子弹数量
func (user *CommonUser) SetActivityBulletNum(num int)  {
	user.ActivityBulletNum = num
}




