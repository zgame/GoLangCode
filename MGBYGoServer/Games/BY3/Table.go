package BY3

//---------------------------------------------------------------------
// 这是一个继承的例子， 由于继承自BY， 所以BY的Table的函数都是直接调用的， 如果再写一遍同名函数， 那么就重载了
//---------------------------------------------------------------------


import (
	"../BY"
	"../Common"
	"time"
	"../../Const"
)

type Table struct {
	BY.Table // 这里的例子是继承自BY这个， 这说明BY3跟BY的逻辑还是比较像的，  也就是说继承的时候，选择逻辑近似的去继承即可


}
//
func (table *Table) NewTable(uid int, gameId int)  Common.TableInterface{
	return &Table{BY.Table{Common.CommonTable{TableID:uid,GameID: gameId, TableMax:Const.BY_TABLE_MAX_PLAYER}}}
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

////---------------------------------桌子------------------------------------------
//
//
//// 返回桌子的UID
//func  (tt *CommonTable)GetTableUID() uint64{
//	return tt.TableID
//
//}
//
//// 新建桌子
//func CreateTable() {
//
//}
//
//
//// 销毁桌子
//func ClearTable() {
//
//}
//
// ----------------------------- Run 桌子 -------------------------------
func (table *Table)RunTable() {
	//fmt.Println("我重载了， 如果把我函数改名字，或者删除， 那么调用父类的函数")
	for {
		//select {
		//
		//}

		//fmt.Println("yuyun运行桌子", "by3")
		time.Sleep(time.Second*3)

	}
}
