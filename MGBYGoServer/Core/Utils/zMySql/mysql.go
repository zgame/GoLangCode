package zMySql

import (
	"fmt"
	"github.com/go-xorm/xorm"
	//"log"
	_ "github.com/go-sql-driver/mysql"
	"../ztimer"
)

//------------------------------------------------------------------------------------------
// my sql database
//------------------------------------------------------------------------------------------

// 这个是go调用mysql
func ConnectDB(ServerIP string,ServerPort string ,Database string, uid string, pwd string) (*xorm.Engine,bool){
	fmt.Println("开始连接mysql数据库！")

	engine, err := xorm.NewEngine("mysql", uid+":"+pwd+"@tcp("+ServerIP+":"+ ServerPort +")/"+Database+"?charset=utf8")


	if err != nil {
		fmt.Println("数据库引擎出错！", err)
		return engine,false
	}
	engine.ShowSQL(false)
	err = engine.Ping()
	if err != nil {
		fmt.Println("数据库ping不通！", err)
		return engine,false
	}


	return engine,true
}

// 执行sql语句
func SqlExec(engine *xorm.Engine,sql string){
	ztimer.CheckRunTimeCost(func() {
		_,err := engine.Exec(sql)
		if err!=nil{
			fmt.Println("my SqlExec error ", err.Error())
		}
	},"my SqlExec  "+ sql)

}

// 查询sql语句
func SqlQuery(engine *xorm.Engine, sql string) []map[string]interface{} {
	var resultsSlice []map[string]interface{}
	ztimer.CheckRunTimeCost(func() {
		data, err := engine.QueryInterface(sql)
		if err != nil {
			fmt.Println("my SqlQuery error ", err.Error())
		}
		resultsSlice = data
	}, "my SqlQuery   "+ sql)

	return resultsSlice
}


