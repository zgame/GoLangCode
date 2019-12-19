package BY2

import (
	"../BY"
	"time"
	"../../Core/GameCore"
	"fmt"
)

//---------------------------------------------------------------------
// 这是一个模板例子
//---------------------------------------------------------------------



type BY2Table struct {
	BY.BYTable

}
// -------------------------构造函数-------------------------
func (table *BY2Table) NewTable(uid int, gameId int) GameCore.TableInterface {
	fmt.Println("创建 捕鱼 BY2  桌子")
	//var tableBase *GameCore.TableBase
	//tableInterface := tableBase.NewTable(uid, gameId)
	//by2Table := GetBY2TableHandle(tableInterface)
	//by2Table.SetTableMax(Const.BY_TABLE_MAX_PLAYER)
	return table.BYTable.NewTable(uid,gameId)
	//return GameCore.TableBase.NewTable(uid , gameId)
}

// 当获取到桌子接口的时候， 我们可以强转成我们当前table的指针
func GetBY2TableHandle(tableInterface GameCore.TableInterface) *BY2Table{
	return tableInterface.(*BY2Table)
}


//// ----------------------------- Run 桌子 -------------------------------
func (table *BY2Table)RunTable() {
	for {
		//select {
		//
		//}

		//fmt.Println("yuyun运行桌子", "by2")
		time.Sleep(time.Second*3)

	}
}
