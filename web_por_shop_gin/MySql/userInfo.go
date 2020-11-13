package MySql

import (
	"web_gin/MiddleWare/zLog"
)

// 根据用户表结构创建， 大写代表_
type Userinfo struct {
	Uid      int    `xorm:"int"`
	Openid   string `xorm:"varchar(100)"`
	Psw      string `xorm:"varchar(15)"`
	Mac      string `xorm:"varchar(50)"`
	Nickname string `xorm:"varchar(30)"`
}

// 同步表结构
func SyncUserInfoTable() bool {
	err := DataBaseEngine.Sync2(new(Userinfo)) //同步表跟结构
	if err != nil {
		zLog.PrintfLogger("同步表结构出错！", err)
		return false
	}
	return true
}

// 查询数据
func GetUserInfoData(openId string ) *Userinfo{

	selectData := &Userinfo{Openid: openId}
	result,err := DataBaseEngine.Get(selectData) //获取单条数据

	if err != nil {
		zLog.PrintfLogger("数据库查询出错！  %s", err)
		return nil
	}
	if result == false{
		return nil
	}
	return selectData
}

// 更新数据   &Userinfo{Uid:1111}
func UpdateUserInfoData(updateData *Userinfo, condition *Userinfo) {

	_, err := DataBaseEngine.Update(updateData, condition)
	if err != nil {
		zLog.PrintfLogger("数据库更新出错！  %s", err)
		return
	}

}
// 插入单行数据
func InsertUserInfoData(insertData *Userinfo) {
	// insert 单条数据
	_, err := DataBaseEngine.Insert(insertData)

	if err != nil {
		zLog.PrintfLogger("数据库插入出错！ %s ", err)
		return
	}
}
