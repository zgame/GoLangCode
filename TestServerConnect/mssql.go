package main

import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"github.com/go-xorm/xorm"
)

var sqlEngine *xorm.Engine

func GetEngine()  {
	//************************************************************************************
	// 读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return
	}
	ServerIP := f.Section("MsSql").Key("ServerIP").Value()
	Database := f.Section("MsSql").Key("Database").Value()
	uid := f.Section("MsSql").Key("uid").Value()
	pwd := f.Section("MsSql").Key("pwd").Value()

	//**************************************************************************************

	fmt.Println("开始连接数据库！")

	// 从sql中获取数据
	sqlEngine, err = xorm.NewEngine("mssql", "server="+ServerIP+";user id="+uid+";password="+pwd+";Database="+Database)
	//sqlEngine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")
	//sqlEngine.ShowSQL(true)
	err=sqlEngine.Ping()

	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		log.Fatal(err)
		return
	}
}

// 往数据库里面插入技能数量

func GMAddSkill(uid int ,skillid int)  {
	// 删除全部
	sql:=  fmt.Sprintf("delete  from userskillinfo where userid = '%d' and itemid = '%d'", uid, skillid)
	//fmt.Println(" sql,",sql)
	_, err := sqlEngine.Query(sql)
	//fmt.Printf("result : %v", result)
	//fmt.Println("")

	if err != nil {
		fmt.Println("addSkill数据库查询出错！", err)
		log.Fatal(err)
		return
	}
	//  插入
	sql =  fmt.Sprintf("insert into userskillinfo (userid,skillid,used,total,usedtime,itemid,todaycatch) values ('%d','0','0','99999','0','%d','0')", uid, skillid)
	//fmt.Println(" sql,",sql)
	_,err = sqlEngine.Query(sql)
	//fmt.Printf("result : %v", result)
	//fmt.Println("")

	if err != nil {
		fmt.Println("addSkill数据库插入出错！", err)
		log.Fatal(err)
		return
	}

}