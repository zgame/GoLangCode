package MySql

import (
	"encoding/json"
	"web_gin/MiddleWare/zLog"
)

// 根据充值表结构创建， 大写代表_
type Recharge struct {
	Uid          int    `xorm:"int"`
	Openid       string `xorm:"varchar(100)"`
	Payno        string `xorm:"varchar(100)"`
	RechargeTime string `xorm:"varchar(20)"`
	Rmb          string `xorm:"varchar(20)"`
	ItemId       int    `xorm:"int"`
	Channel      string `xorm:"varchar(10)"`
}

// 同步表结构
func SyncRechargeTable() bool {
	err := DataBaseEngine.Sync2(new(Recharge)) //同步表跟结构
	if err != nil {
		zLog.PrintfLogger("充值 同步表结构出错！", err)
		return false
	}
	return true
}

// 查询订单数据
func GetRechargeData(payno string) *Recharge {

	selectData := &Recharge{Payno: payno}
	//total, err := DataBaseEngine.Where("Payno =?", Payno).Sum(selectData, "rmb")		// 获取总数
	result, err := DataBaseEngine.Get(selectData) //获取单条数据

	if err != nil {
		zLog.PrintfLogger("充值 数据库查询出错！  %s", err)
		return nil
	}
	if result == false {
		return nil
	}
	return selectData
}

// 根据充值更新该玩家的道具列表
func UpdateAllItems(openId string) {
	itemList := make([]int, 0)
	list := make([]Recharge, 0)
	err := DataBaseEngine.Where("openid = ? ", openId).Find(&list)
	if err != nil {
		zLog.PrintfLogger("充值 数据库查询出错！  %s", err)
		return
	}
	for _, item := range list {
		itemList = append(itemList, item.ItemId)
	}
	data, _ := json.Marshal(itemList)
	updata := new(Useritem)
	updata.Openid = openId
	updata.ShopList = string(data)
	UpdateUserItemData(updata, openId)
}



// 插入单行数据
func InsertRechargeData(insertData *Recharge) bool{
	_, err := DataBaseEngine.Insert(insertData)
	if err != nil {
		zLog.PrintfLogger("充值数据库插入充值出错！ %s ", err)
		return false
	}
	zLog.PrintLogger("充值数据插入数据库充值ok！  ")
	return true
}
