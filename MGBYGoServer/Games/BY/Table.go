package BY

import (
	"../Common"
	"../../Const"
)

type Table struct {
	Common.CommonTable
}

func (table *Table) NewTable(uid int, gameId int)  Common.TableInterface{
	return &Table{Common.CommonTable{TableID:uid, GameID: gameId, TableMax:Const.BY_TABLE_MAX_PLAYER}}
}

// 设定子弹的句柄类型
func (table *Table) SetBulletInterface() {
	var bullet * Bullet
	table.BulletInterfaceHandle = bullet
}

// 设定鱼的句柄类型
func (table *Table) SetFishInterface() {
	var fish * Fish
	table.FishInterfaceHandle = fish
}


// 设定玩家的句柄类型
func (table *Table) SetUserInterfaceHandle() {
	var user * User
	table.UserInterfaceHandle = user
}

//
////---------------------------------桌子------------------------------------------
//
//// 返回桌子的UID
////func  (tt *Table)GetTableUID() uint64{
////	return tt.TableID
////
////}
//
//
//// 新建桌子
//func CreateTable() {
//
//}
//
//// 销毁桌子
//func ClearTable() {
//
//}
//
// ----------------------------- Run 桌子 -------------------------------
//func (table *Table)RunTable() {
//	for {
//		//select {
//		//
//		//}
//
//		fmt.Println("yuyun运行桌子", "by1")
//		time.Sleep(time.Second*3)
//
//	}
//}
