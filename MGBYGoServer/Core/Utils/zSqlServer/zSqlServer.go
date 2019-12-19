package zSqlServer

import (
"fmt"
"github.com/go-xorm/xorm"
_"github.com/denisenkom/go-mssqldb"
"log"
//_ "github.com/go-sql-driver/"
"../ztimer"
)
//var SqlEngine *xorm.Engine
//------------------------------------------------------------------------------------------
//  sql server database
//------------------------------------------------------------------------------------------


// 这个是go调用mysql
func ConnectDB(ServerIP string,Database string, uid string, pwd string) (bool, *xorm.Engine){
	fmt.Println("开始连接 sql server 数据库！")

	// 从sql中获取数据
	engine, err := xorm.NewEngine("mssql", "server="+ServerIP+";Database="+Database+";user id="+uid+";password="+pwd+";")
	//engine, err := xorm.NewEngine("zMysqlForLua", uid+":"+pwd+"@tcp("+ServerIP+":"+ ServerPort +")/"+Database+"?charset=utf8")


	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		log.Fatal(err)
		return false,nil
	}
	engine.ShowSQL(false)
	err = engine.Ping()
	if err != nil {
		fmt.Println("数据库ping不通！", err)
		return false,nil
	}
	//SqlEngine = engine

	//SqlExec("use "+ Database)
	//SqlExec("select * from game_state")

	return true,engine
}


// 执行sql语句
func SqlExec(engine *xorm.Engine, sql string){
	ztimer.CheckRunTimeCost(func() {
		_,err := engine.Exec(sql)
		if err!=nil{
			fmt.Println("SqlExec error ", err.Error())
		}
	},"SqlExec")

}

// 查询sql语句
func SqlQuery(engine *xorm.Engine, sql string) []map[string]interface{} {
	var resultsSlice []map[string]interface{}
	ztimer.CheckRunTimeCost(func() {
		data, err := engine.QueryInterface(sql)
		if err != nil {
			fmt.Println("SqlQuery error ", err.Error())
		}
		resultsSlice = data
	}, "SqlQuery")

	return resultsSlice
}


