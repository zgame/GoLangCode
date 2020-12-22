package MySql

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)
var DataBaseEngine *xorm.Engine		// 全局变量

// 初始化数据库
func InitDataBase() * xorm.Engine{
	//************************************************************************************
	// 读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return nil
	}
	ServerIP := f.Section("author").Key("ServerIP").Value()
	Database := f.Section("author").Key("Database").Value()
	uid := f.Section("author").Key("uid").Value()
	pwd := f.Section("author").Key("pwd").Value()

	//**************************************************************************************

	fmt.Println("开始连接mysql数据库！")

	// 从sql中获取数据
	//Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server="+ServerIP+";Database="+Database+";uid="+uid+";pwd="+pwd+";")
	engine, err2 := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")
	//engine.ShowSQL(true)
	err2 = engine.Ping()

	if err2 != nil {
		fmt.Println("数据库引擎出错！", err2)
		log.Fatal(err2)
		return nil
	}

	return engine
}


// 同步表结构
func InitSycTables() {
	SyncUserInfoTable()		// userinfo 表
	SyncRechargeTable()		// recharge
	SyncUserItemTable()		// userItem
	SyncMallInfoTable()		// shop mall
}