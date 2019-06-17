package zMySql

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"../ztimer"
)
var MySqlEngine *xorm.Engine
//------------------------------------------------------------------------------------------
//--- 这里有2种方式调用mysql
//--- 第一种是通过go调用，适用于对性能有要求，并且没有返回的情况    zMySql
//--- 第二种是通过lua调用， 性能会差很多，会有延迟， 但是可以获取到select的返回值，并且是table类型的     gluasql_mysql
//------------------------------------------------------------------------------------------

// 这个是go调用mysql
func ConnectDB(ServerIP string,ServerPort string ,Database string, uid string, pwd string) bool{
	// 读取配置文件
	//f, err := ini.Load("Setting.ini")
	//if err != nil {
	//	fmt.Println("ini配置文件出错！", err)
	//	log.Fatal(err)
	//	return false
	//}
	//ServerIP := f.Section("Server").Key("MySqlServerIP").Value()
	//ServerPort := f.Section("Server").Key("MySqlServerPort").Value()
	//Database := f.Section("Server").Key("Database").Value()
	//uid := f.Section("Server").Key("uid").Value()
	//pwd := f.Section("Server").Key("pwd").Value()

	fmt.Println("开始连接mysql数据库！")

	// 从sql中获取数据
	//Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server="+ServerIP+";Database="+Database+";uid="+uid+";pwd="+pwd+";")
	engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+":"+ ServerPort +")/"+Database+"?charset=utf8")


	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		log.Fatal(err)
		return false
	}
	engine.ShowSQL(false)
	err = engine.Ping()
	if err != nil {
		fmt.Println("数据库ping不通！", err)
		return false
	}
	MySqlEngine = engine

	//SqlExec("use "+ Database)
	//SqlExec("select * from game_state")

	return true
}

//// 保存服务器的游戏房间的当前状态， 这个作废了， 效率低， 后面采用组合的唯一主键， 然后采用replace into 或者 duplicate来进行更新
//func SqlSaveGameState(ServerIP_Port string,gameType int,tableId  int,FishNum int,BulletNum int,SeatArray int)  {
//
//	// 这个太麻烦，效率也不高， 作废了
//	return
//
//	ztimer.CheckRunTimeCost(func() {
//		// 这里没有用duplicate是因为没有主键
//		select_sql := fmt.Sprintf("select table_id from game_state where server_ip = '%s' and game_id = %d and table_id = %d limit 1",ServerIP_Port,gameType,tableId )
//		insert_sql := fmt.Sprintf("insert into game_state (server_ip,game_id,table_id,fish_num,bullet_num,seat_array) values ('%s',%d, %d,%d,%d,%d)",ServerIP_Port,gameType,tableId ,FishNum,BulletNum,SeatArray)
//		update_sql := fmt.Sprintf("update  game_state   set  fish_num =%d ,bullet_num =%d,seat_array=%d where server_ip = '%s' and game_id = %d and table_id = %d",FishNum,BulletNum,SeatArray,ServerIP_Port,gameType,tableId )
//
//		// 后面改成下面的语句，然后用lua去调用了， go这个废了
//		replace_into := "insert into game_state (zkey,server_ip,game_id,table_id,fish_num,bullet_num,seat_array) values ('1_1_1','1',1, 1,1,1,1) on DUPLICATE key update fish_num =99 ,bullet_num =99,seat_array=99"
//
//		//fmt.Println("select_sql",select_sql)
//
//		result,err:=	MySqlEngine.Query(select_sql)
//		if err!=nil{
//			fmt.Println("Query error ", err.Error())
//		}
//
//		if len(result) == 0 {
//			_,err = MySqlEngine.Exec(insert_sql)
//		}else {
//			_,err = MySqlEngine.Exec(update_sql)
//		}
//		if err!=nil{
//			fmt.Println("Exec error ", err.Error())
//		}
//	},"SqlSaveGameState")
//
//}

// 执行sql语句
func SqlExec(sql string){

	ztimer.CheckRunTimeCost(func() {
		_,err := MySqlEngine.Exec(sql)
		if err!=nil{
			fmt.Println("SqlExec error ", err.Error())
		}
	},"SqlExec")

}




//func test()  {
//	//fmt.Println("判断表是否存在")
//	engine:= MySqlEngine
//	//has,err := engine.IsTableExist(new(Test_ip))
//	//if has{
//	//	fmt.Println("表存在，那么读一下")
//	//
//	//	fmt.Println("开始同步表结构")
//	//	err = engine.Sync2(new(Test_ip)) //同步表跟结构
//	//
//	//	fmt.Println("开始获取单条数据")
//	//	var test_ii Test_ip
//	//	engine.Get(&test_ii)		//获取单条数据
//	//	println(test_ii.Ip)
//	//
//	//
//	//	var test_iii []Test_ip
//	//	engine.Find(&test_iii)		//获取多条
//	//	println(len(test_iii))
//	//	for i ,v := range test_iii{
//	//		fmt.Println(i)
//	//		fmt.Printf("%d:%s",i,v.Ip)
//	//		fmt.Println("")
//	//	}
//	//
//	//
//	//
//	//	println("----------")
//	//
//	//	// update更新数据
//	//	_,err = engine.Exec("update Test_ip set Ip = ? where Ip = ?", "192.2.0.0", "192.1.0.0")
//	//
//	//
//	//	// insert 单条数据
//	//	test_ii.Ip = "1111"
//	//	_, err = engine.Insert(&test_ii)
//	//
//	//	// 插入多条数据
//	//	test_iii = append(test_iii, Test_ip{"2222"})
//	//	fmt.Printf("%v",test_iii)
//	//
//	//	_, err = engine.Insert(&test_iii)
//
//		// 删除全部
//	result,err:=	engine.Query("SELECT * FROM user")
//	if err!=nil{
//		fmt.Println("Query error ", err.Error())
//	}
//	fmt.Printf("result %v \n",result)
//	for i:=range result{
//		fmt.Printf("result %s \n", string(result[i]["login_time"]))
//	}
//
//
//
//
//	//	if err != nil {
//	//		fmt.Println("数据库查询出错！", err)
//	//		log.Fatal(err)
//	//		return
//	//	}
//	//}else{
//	//	fmt.Println("不存在，创建表")
//	//	engine.CreateTables(new(Test_ip))
//	//
//	//}
//}
//
