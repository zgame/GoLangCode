package BY3

//---------------------------------------------------------------------
// 这是一个继承的例子， 由于继承自BY， 所以BY的Table的函数都是直接调用的， 如果再写一遍同名函数， 那么就重载了
//---------------------------------------------------------------------


import (
	"../BY"
	"time"
	"../../Core/GameCore"
	"fmt"
)

//---------------------------------------------------------------------
// 这是一个模板例子
//---------------------------------------------------------------------


type BY3Table struct {
	BY.BYTable

}
// -------------------------构造函数-------------------------
func (table *BY3Table) NewTable(uid int, gameId int) GameCore.TableInterface {
	fmt.Println("创建 捕鱼 BY3  桌子")
	//var tableBase *GameCore.TableBase
	//tableInterface := tableBase.NewTable(uid, gameId)
	//by3Table := GetBY3TableHandle(tableInterface)
	//by3Table.SetTableMax(Const.BY_TABLE_MAX_PLAYER)
	return table.BYTable.NewTable(uid,gameId)
	//return GameCore.TableBase.NewTable(uid , gameId)
}

// 当获取到桌子接口的时候， 我们可以强转成我们当前table的指针
func GetBY3TableHandle(tableInterface GameCore.TableInterface) *BY3Table{
	return tableInterface.(*BY3Table)
}


// ----------------------------- Run 桌子 -------------------------------
func (table *BY3Table)RunTable() {
	//fmt.Println("我重载了， 如果把我函数改名字，或者删除， 那么调用父类的函数")
	for {
		//select {
		//
		//}

		//fmt.Println("yuyun运行桌子", "by3")
		time.Sleep(time.Second*3)

	}
}
