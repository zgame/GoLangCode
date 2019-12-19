package SqlServerLog


import (
	"github.com/go-xorm/xorm"
	"../../../../Core/Utils/zSqlServer"
)

//------------------------------------------------------------------------------------
//// 捕鱼 数据库
//------------------------------------------------------------------------------------

var LogSqlEngine *xorm.Engine		// 捕鱼 数据库

// 连接捕鱼数据库
func LogConnectSqlDB(ServerIP string,Database string, uid string, pwd string) bool{
	var result bool
	result , LogSqlEngine = zSqlServer.ConnectDB(ServerIP ,Database  ,uid  ,pwd )
	return result
}

// 执行sql
func LogSqlDBExec(sql string)  {
	zSqlServer.SqlExec(LogSqlEngine,sql)

}

// 查询sql
func LogSqlDBQuery(sql string)  []map[string]interface{}{
	return zSqlServer.SqlQuery(LogSqlEngine,sql)

}

