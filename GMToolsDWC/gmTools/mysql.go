package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"github.com/go-xorm/xorm"
	"time"
	"strconv"
)


type Serverlist struct {
	Serverid	int
	Gameid	int
	Room_name string `xorm:"varchar(40)"`
	Domain string `xorm:"varchar(40)"`
	Ipaddr2	string `xorm:"varchar(40)"`
	Ipaddr string `xorm:"varchar(40)"`
	Listen_port int
	Ssh_port int
	Ssh_user string `xorm:"varchar(40)"`
	Ssh_passwd string `xorm:"varchar(40)"`
	Max_player int
	Nginx_addr string `xorm:"varchar(40)"`
	Daemon int

	//Ip string `xorm:"varchar(40)"`
}

type Serverlogerror struct {
	Idx      int
	Gameid   int
	Loglevel string
	Logtxt   string
	Logtime  string
}

func getServerList() []Serverlist{
	engine := dbConnect()
	if engine != nil {
		return selectServerList(engine, new(Serverlist))
	}
	return nil
}

func insertLogToSQL(gameId int, logLevel string , logTxt string)  {
	engine := dbConnect()
	if engine != nil {
		insertLog(engine, new(Serverlogerror),gameId,logLevel,logTxt)
	}
}


//*************************数据库连接*************************
func dbConnect() *xorm.Engine {

	//*************************读取配置文件
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		log.Fatal(err)
		return nil
	}
	ServerIP := f.Section("Mysql").Key("ServerIP").Value()
	Database := f.Section("Mysql").Key("Database").Value()
	uid := f.Section("Mysql").Key("uid").Value()
	pwd := f.Section("Mysql").Key("pwd").Value()

	//*****************************

	//fmt.Println("开始连接数据库！")

	// 从sql中获取数据
	//Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server="+ServerIP+";Database="+Database+";uid="+uid+";pwd="+pwd+";")
	engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+")/"+Database+"?charset=utf8")
	//engine.ShowSQL(true)
	err = engine.Ping()

	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		log.Fatal(err)
		return nil
	}
	return engine
}

//***************************数据库表检查*********************************************************
func tableCheck(engine *xorm.Engine, table interface{}) bool {

	//fmt.Println("判断表是否存在")
	has,err := engine.IsTableExist(table)
	if has{
		//fmt.Println("表存在，那么读一下")
		//fmt.Println("开始同步表结构")
		err = engine.Sync(table) //同步表跟结构
		if err != nil {
			fmt.Println("数据库查询出错！", err)
			log.Fatal(err)
			return false
		}
	}else{
		fmt.Println("不存在，创建表?")
		return false
		//engine.CreateTables(new(Test_ip))
	}
	return true
}

//*************************** select *********************************************************
func selectServerList(engine *xorm.Engine, table interface{}) []Serverlist{
	if !tableCheck(engine,table){
		fmt.Println("表格同步未通过， 不执行sql")
		return nil
	}

	//fmt.Println("开始获取单条数据")
	//var test_ii Serverlist
	//_,err := engine.Get(&test_ii)		//获取单条数据
	//fmt.Printf("test_ii,  %v",test_ii)
	//fmt.Println("--------------------------------")
	////println(test_ii.Ip)


	var sList []Serverlist
	err := engine.Find(&sList) //获取多条
	//fmt.Printf("sList : %v", sList)
	//fmt.Println("--------------------------------")
	//println(len(sList))
	//for i ,v := range sList{
	//	fmt.Println(i)
	//	fmt.Printf("%d:%s",i,v.Ip)
	//	fmt.Println("")
	//}

	//println("---------select-------")
	if err != nil {
		fmt.Println("数据库查询出错！", err)
		log.Fatal(err)
		return nil
	}
	return sList
}

//*************************** insert *********************************************************
func insertLog(engine *xorm.Engine, table interface{}, gameId int, logLevel string , logTxt string)  {
	if !tableCheck(engine,table){
		fmt.Println("表格同步未通过， 不执行sql")
		return
	}
	//println("------insert----")

	//insert 单条数据
	var logOne Serverlogerror
	logOne.Gameid = gameId
	logOne.Loglevel = logLevel
	logOne.Logtxt = logTxt
	logOne.Logtime = strconv.FormatInt(time.Now().UnixNano(),10)
	_, err := engine.Insert(&logOne)

	//// 插入多条数据
	//test_iii = append(test_iii, Test_ip{"2222"})
	//fmt.Printf("%v",test_iii)
	_, err = engine.Insert(&logOne)

	if err != nil {
		fmt.Println("数据库插入出错！", err)
		log.Fatal(err)
	}
}

//*************************** update *********************************************************
func update(engine *xorm.Engine, table interface{})  {
	if !tableCheck(engine,table){
		fmt.Println("表格同步未通过， 不执行sql")
		return
	}

	println("--- update -------")
	// update更新数据
	_,err := engine.Exec("update Test_ip set Ip = ? where Ip = ?", "192.2.0.0", "192.1.0.0")

	if err != nil {
		fmt.Println("数据库更新出错！", err)
		log.Fatal(err)

	}
}
