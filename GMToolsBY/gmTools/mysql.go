package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"github.com/go-xorm/xorm"
)


type Test_ip struct {
	Ip string `xorm:"varchar(40)"`
}

func main1() {
	get_data()
}



func get_data()  {
	//************************************************************************************
	// 读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return
	}
	ServerIP := f.Section("Mysql").Key("ServerIP").Value()
	Database := f.Section("Mysql").Key("Database").Value()
	uid := f.Section("Mysql").Key("uid").Value()
	pwd := f.Section("Mysql").Key("pwd").Value()

	//**************************************************************************************

	fmt.Println("开始连接数据库！")

	// 从sql中获取数据
	//Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server="+ServerIP+";Database="+Database+";uid="+uid+";pwd="+pwd+";")
	engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")
	engine.ShowSQL(true)
	err=engine.Ping()

	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		log.Fatal(err)
		return
	}




	fmt.Println("判断表是否存在")

	has,err := engine.IsTableExist(new(Test_ip))
	if has{
		fmt.Println("表存在，那么读一下")

		fmt.Println("开始同步表结构")
		err = engine.Sync2(new(Test_ip)) //同步表跟结构

		fmt.Println("开始获取单条数据")
		var test_ii Test_ip
		engine.Get(&test_ii)		//获取单条数据
		println(test_ii.Ip)


		var test_iii []Test_ip
		engine.Find(&test_iii)		//获取多条
		println(len(test_iii))
		for i ,v := range test_iii{
			fmt.Println(i)
			fmt.Printf("%d:%s",i,v.Ip)
			fmt.Println("")
		}



		println("----------")

		// update更新数据
		_,err = engine.Exec("update Test_ip set Ip = ? where Ip = ?", "192.2.0.0", "192.1.0.0")


		// insert 单条数据
		test_ii.Ip = "1111"
		_, err = engine.Insert(&test_ii)

		// 插入多条数据
		test_iii = append(test_iii, Test_ip{"2222"})
		fmt.Printf("%v",test_iii)

		_, err = engine.Insert(&test_iii)

		// 删除全部
		engine.Query("delete  from Test_ip")


		if err != nil {
			fmt.Println("数据库查询出错！", err)
			log.Fatal(err)
			return
		}
	}else{
		fmt.Println("不存在，创建表")
		engine.CreateTables(new(Test_ip))

	}




}

