package Games

import (
	"./BY"
	"./BY2"
	"./BY3"
	. "../Core/GameCore"

)

//----------------------------------------------------------------
// 该文件是所有游戏的管理类， 掌握所有游戏的列表
// 单个游戏的结构也掌握所有桌子的列表
//----------------------------------------------------------------


// 所有游戏的编号
const(
	GameTypeBY  = 3
	GameTypeBY2 = 2999
	GameTypeBY3 = 3999
	// 更多游戏编号添加到此处...
	// ...
	// ...
)

// 增加一个游戏， 指定这个游戏的类型， 并且创建一个桌子，并启动桌子逻辑
func AddGame(name string, gameType int) {
	var games *Games
	games = NewGame(name, gameType, true) // 创建游戏房间
	AllGamesList[gameType] = games

	// 根据游戏类型进行判断
	if gameType == GameTypeBY{
		var table  * BY.BYTable
		games.CreateTable(table,1)		//创建1底分桌子
		//games.CreateTable(table,100)	//创建100底分桌子
		//games.CreateTable(table,10000)	//创建10000底分桌子
	}else if gameType == GameTypeBY2{
		var table   * BY2.BY2Table
		games.CreateTable(table,1)
	}else if gameType == GameTypeBY3{
		var table * BY3.BY3Table
		games.CreateTable(table,1)
	}
	// 增加游戏类型...
	//...
	//...
	//...

}

