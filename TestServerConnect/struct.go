package main

import (
	"time"
	"net"
)

type Client struct {
	Conn       net.Conn
	Index      int
	User       *User
	Serverlist []*GameServerInfo
	Gameinfo   * GameInfo
	SendTokenID int32
	StartAI 	bool
	Last_fire_tick time.Time		// 开火CD
	Last_skill_tick time.Time		// 技能CD
	Select_tick time.Time			// 选鱼CD
	Last_check_due_tick time.Time    // 定期清理过期鱼
	Fish_id	int				//锁定的鱼id
	Failed_cnt int			// 锁定鱼打了几炮

	ShowLog uint64 			//打鱼的记录

	ReloginTime time.Time	// 志勇要模拟玩家玩1分钟再退出重新进入的测试
	Quit  bool

	ReceiveBuf  []byte
}



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


