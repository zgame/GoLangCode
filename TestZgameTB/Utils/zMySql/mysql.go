package zMySql

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)
var MySqlEngine *xorm.Engine



// 这个是go调用mysql， 暂时没用， 而使用lua掉mysql
func ConnectDB() bool{
	// 读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return false
	}
	ServerIP := f.Section("Server").Key("MySqlServerIP").Value()
	Database := f.Section("Server").Key("Database").Value()
	uid := f.Section("Server").Key("uid").Value()
	pwd := f.Section("Server").Key("pwd").Value()

	fmt.Println("开始连接mysql数据库！")

	// 从sql中获取数据
	//Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server="+ServerIP+";Database="+Database+";uid="+uid+";pwd="+pwd+";")
	engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")


	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		log.Fatal(err)
		return false
	}
	engine.ShowSQL(true)
	err=engine.Ping()
	MySqlEngine = engine
	
	test()
	
	return false
}

func test()  {

	//fmt.Println("判断表是否存在")
	engine:= MySqlEngine
	//has,err := engine.IsTableExist(new(Test_ip))
	//if has{
	//	fmt.Println("表存在，那么读一下")
	//
	//	fmt.Println("开始同步表结构")
	//	err = engine.Sync2(new(Test_ip)) //同步表跟结构
	//
	//	fmt.Println("开始获取单条数据")
	//	var test_ii Test_ip
	//	engine.Get(&test_ii)		//获取单条数据
	//	println(test_ii.Ip)
	//
	//
	//	var test_iii []Test_ip
	//	engine.Find(&test_iii)		//获取多条
	//	println(len(test_iii))
	//	for i ,v := range test_iii{
	//		fmt.Println(i)
	//		fmt.Printf("%d:%s",i,v.Ip)
	//		fmt.Println("")
	//	}
	//
	//
	//
	//	println("----------")
	//
	//	// update更新数据
	//	_,err = engine.Exec("update Test_ip set Ip = ? where Ip = ?", "192.2.0.0", "192.1.0.0")
	//
	//
	//	// insert 单条数据
	//	test_ii.Ip = "1111"
	//	_, err = engine.Insert(&test_ii)
	//
	//	// 插入多条数据
	//	test_iii = append(test_iii, Test_ip{"2222"})
	//	fmt.Printf("%v",test_iii)
	//
	//	_, err = engine.Insert(&test_iii)

		// 删除全部
	result,err:=	engine.Query("SELECT * FROM user")
	if err!=nil{
		fmt.Println("Query error ", err.Error())
	}
	fmt.Printf("result %v \n",result)
	for i:=range result{
		fmt.Printf("result %s \n", string(result[i]["login_time"]))
	}




	//	if err != nil {
	//		fmt.Println("数据库查询出错！", err)
	//		log.Fatal(err)
	//		return
	//	}
	//}else{
	//	fmt.Println("不存在，创建表")
	//	engine.CreateTables(new(Test_ip))
	//
	//}

}

