package MySqlDB


import (
	"github.com/go-xorm/xorm"
	"../../../Core/Utils/zMySql"
)


//------------------------------------------------------------------------------------
//
//------------------------------------------------------------------------------------

var MySqlEngine *xorm.Engine


// 连接my sql数据库
func ConnectMySqlDB(ServerIP string,ServerPort string ,Database string, uid string, pwd string) bool{
	var result bool
	MySqlEngine, result = zMySql.ConnectDB(ServerIP ,ServerPort,Database  ,uid  ,pwd )
	return result
}

// 执行sql
func MySqlDBExec(sql string)  {
	zMySql.SqlExec(MySqlEngine,sql)

}

// 查询sql
func MYSqlDBQuery(sql string)  []map[string]interface{}{
	return zMySql.SqlQuery(MySqlEngine,sql)

}
