package SqlServerBy


import (
	"github.com/go-xorm/xorm"
	"../../../../Core/Utils/zSqlServer"
)

//------------------------------------------------------------------------------------
//// 捕鱼 数据库
//------------------------------------------------------------------------------------

var BySqlEngine *xorm.Engine		// 捕鱼 数据库

// 连接捕鱼数据库
func BYConnectSqlDB(ServerIP string,Database string, uid string, pwd string) bool{
	var result bool
	result , BySqlEngine = zSqlServer.ConnectDB(ServerIP ,Database  ,uid  ,pwd )

	// debug info
	//re:= BYSqlDBQuery("select * from MinorLimit")
	//fmt.Println("",re)
	//fmt.Println("",re[0]["startplaytime"])

	return result
}

// 执行sql
func BYSqlDBExec(sql string)  {
	zSqlServer.SqlExec(BySqlEngine,sql)

}

// 查询sql
func BYSqlDBQuery(sql string) []map[string]interface{} {
	return zSqlServer.SqlQuery(BySqlEngine,sql)

}
