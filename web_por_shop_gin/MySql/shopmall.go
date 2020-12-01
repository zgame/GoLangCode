package MySql

import (
	"web_gin/MiddleWare/zLog"
)

// 根据用户表结构创建， 大写代表_
type Shopmall struct {
	Id                int     `xorm:"int"`
	Sellingway        int     `xorm:"int"`
	Recommend         int     `xorm:"int"`
	Recommendactivity int     `xorm:"int"`
	Price             int     `xorm:"int"`
	Discountprice     float64 `xorm:"double"`
	Starttime         string  `xorm:"varchar(20)"`
	Endtime           string  `xorm:"varchar(20)"`
}

// 同步表结构
func SyncMallInfoTable() bool {
	err := DataBaseEngine.Sync2(new(Shopmall)) //同步表跟结构
	if err != nil {
		zLog.PrintfLogger("同步表结构出错！", err)
		return false
	}
	return true
}

// 查询数据
func GetMallInfoData() []Shopmall {

	selectData := make([]Shopmall, 0)
	err := DataBaseEngine.Find(&selectData) //获取单条数据

	if err != nil {
		zLog.PrintfLogger("数据库查询出错！  %s", err)
		return nil
	}

	return selectData
}

//
//// 更新数据   &Userinfo{Uid:1111}
//func UpdateUserInfoData(updateData *Userinfo, condition *Userinfo) {
//
//	_, err := DataBaseEngine.Update(updateData, condition)
//	if err != nil {
//		zLog.PrintfLogger("数据库更新出错！  %s", err)
//		return
//	}
//
//}

// 插入单行数据
func InsertMallInfoData(insertData *Shopmall) {
	// insert 单条数据
	_, err := DataBaseEngine.Insert(insertData)
	if err != nil {
		zLog.PrintfLogger("数据库插入出错！ %s ", err)
		return
	}
}