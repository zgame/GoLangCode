package Player

import (
	"math/rand"
)
//------------------------------------------------------------------------------------------------------
// 玩家数据结构， 这里会整体保存玩家的完整数据结构到数据库中， 为了通用，可以采用转换成json格式之后保存， 可以是redis，也可以是mongodb
// ---------------------------------------------------------------------------------------------------

type Player struct {
	FaceId       int32  // # 头像id
	Gender       int32  //  # 性别
	UserId       uint32 //  # 用户id
	GameId       uint32 //  # 游戏id
	Exp          int64  //  # 经验
	Loveliness   uint32 //  # 魅力
	Score        int64  //  # 分数
	NickName     string //  # 昵称
	Level        int32  //  # 等级
	VipLevel     int32  //  # vip等级
	AccountLevel int32  //  # 账号等级
	SiteLevel    int32  //  # 炮等级
	CurLevelExp  int64  //  # 当前等级经验
	NextLevelExp int64  //  # 下一等级经验
	PayTotal     int32  //  # 充值总金额
	Diamond      int64  //  # 钻石数量
}

func (player *Player) NewPlayer() *Player {
	return &Player{FaceId:0,Gender:0, UserId:0, GameId:0,Exp:0,Loveliness:0,Score:0, NickName:"",Level:0, VipLevel:0, AccountLevel:0, SiteLevel:0, CurLevelExp:0, NextLevelExp:0, PayTotal:0,Diamond:0}
}

// ---------------------------------玩家的逻辑-------------------------------------------

// 玩家升级


// 玩家充值


// 玩家炮等级提升






// ----------------------------玩家UID的设定------------------------------------
var ALLUserUUID int	//玩家uid的自增

// 服务器启动的时候， 从数据库中读取玩家最后的uid
func GetALLUserUUID() {
	// 这个要从数据库读取


	// 如果读取数据是0，那么就重置
	if ALLUserUUID == 0 {
		ALLUserUUID = 1000000000
	}
}

// 有一个新的玩家注册了，那么给他分配一个UID
func GetLastUserID() int{
	r := rand.Intn(5)		//返回[0,5)的随机整数
	ALLUserUUID += r			// 玩家的UID 中间会隔一些数字， 防止玩家挨个去猜UID
	return ALLUserUUID
}


//------------------------------------------------------------------------------
