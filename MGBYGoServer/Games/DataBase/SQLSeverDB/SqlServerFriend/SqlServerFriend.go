package SqlServerFriend


import (
	"github.com/go-xorm/xorm"
	"../../../../Core/Utils/zSqlServer"
)

//------------------------------------------------------------------------------------
//// 捕鱼 数据库
//------------------------------------------------------------------------------------

var FriendSqlEngine *xorm.Engine		// 捕鱼 数据库

// 连接捕鱼数据库
func FriendConnectSqlDB(ServerIP string,Database string, uid string, pwd string) bool{
	var result bool
	result , FriendSqlEngine = zSqlServer.ConnectDB(ServerIP ,Database  ,uid  ,pwd )
	return result
}

// 执行sql
func FriendSqlDBExec(sql string)  {
	zSqlServer.SqlExec(FriendSqlEngine,sql)

}

// 查询sql
func FriendSqlDBQuery(sql string) []map[string]interface{} {
	return zSqlServer.SqlQuery(FriendSqlEngine,sql)
}


