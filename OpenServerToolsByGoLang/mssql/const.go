package main

import "github.com/go-xorm/xorm"
import "github.com/lxn/walk"

func GameRoomListSelect() []*Species {
	return []*Species{
		{0,"满贯捕鱼公会房.sql"},
		{1,"满贯捕鱼传奇起航新手场.sql"},
		{2,"满贯捕鱼海盗宝藏30倍场.sql"},
		{3,"满贯捕鱼大师赛.sql"},
		{4,"满贯捕鱼海妖领域300倍房.sql"},
		{5,"满贯捕鱼导弹场飞禽走兽.sql"},
		{6,"满贯捕鱼妖怪哪里跑.sql"},
		{7,"满贯捕鱼大奖赛.sql"},
		{8,"满贯捕鱼活动赛.sql"},
		{9,"满贯捕鱼飞禽走兽.sql"},
		{10,"满贯捕鱼百人牛牛.sql"},
		{11,"满贯捕鱼水果机.sql"},
		{12,"满贯捕鱼水浒传.sql"},
		{13,"满贯捕鱼龙穴探宝.sql"},
		{14,"满贯捕鱼斗地主.sql"},
		{15,"满贯捕鱼打海兽.sql"},
		{16,"满贯捕鱼食人鱼谷1区.sql"},
		{17,"满贯捕鱼食人鱼谷2区.sql"},
		{18,"满贯捕鱼食人鱼谷3区.sql"},
	}
}

func SqlFileRobotListSelect() []*Species {
	return []*Species{
		{0,"不生成机器人"},
		{1,"生成8人中级场机器人.sql"},
		{2,"生成8人免费场机器人.sql"},
		{3,"生成8人初级场机器人.sql"},
		{4,"生成8人高级场机器人.sql"},
		{5,"生成40人专家场机器人.sql"},
		{6,"生成40人中级场机器人.sql"},
		{7,"生成40人初级场机器人.sql"},
		{8,"生成40人高级场机器人.sql"},
		{9,"生成十倍房机器人.sql"},
		{10,"生成千倍房机器人.sql"},
		{11,"生成大师赛排名机器人.sql"},
		{12,"生成大师赛机器人.sql"},
		{13,"生成新手房机器人.sql"},
		{14,"生成杀分游戏机器人.sql"},
		{15,"生成百倍房机器人.sql"},

	}
}


type MyMainWindow struct {
	*walk.MainWindow
	model *TableViewModel
	tv    *walk.TableView
	teid  *walk.LineEdit
	temachine  *walk.LineEdit
	teandorid  *walk.LineEdit
	teversion  *walk.LineEdit
	tegametype  *walk.ComboBox
	tesqltype  *walk.ComboBox
}

type TableViewModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*ServerState
}

type ServerState struct {
	ServerId    int
	ServerName  string
	ServerMachine string
	AndroidCount      int
	GameRoomListIndex      int
	SqlFileRobotListIndex         int

	//Quux    time.Time
	checked bool
}



// 这个是数据库中的表结构，多行返回这个是一个本结构体的数组
type AaWhiteIPList struct {
	Ip       string `xorm:"varchar(15)"`
	Comments string `xorm:"varchar(100)"`
}
//var ServerListAll []ServerState

var ServerMachine string
var ServerID int
var AndroidNum int
var Version string
var Engine *xorm.Engine

type Species struct {
	Id int
	Name string
}