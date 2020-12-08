//------------------------------------------------------------------------------------------
// sql server 的数据库连接
//------------------------------------------------------------------------------------------

package sqlServer

import (
	"GoLuaServerV2.1/Utils"
	"database/sql"
	"fmt"
	"github.com/yuin/gopher-lua"
	_ "github.com/denisenkom/go-mssqldb"
)

func clientConnectMethod(L *lua.LState) int {

	client := checkClient(L)
	tb := Utils.LuaGetValue(L, 2)
	options, ok := tb.(map[string]interface{})

	if tb == nil || !ok {
		L.ArgError(2, "options excepted")
		return 0
	}

	host, _ := options["host"].(string)
	if host == "" {
		host = "127.0.0.1"
	}
	port, _ := options["port"].(int)
	if port == 0 {
		port = 3306
	}
	database, _ := options["database"].(string)
	user, _ := options["user"].(string)
	password, _ := options["password"].(string)
	//charset, _ := options["charset"].(string)

	//fmt.Println("dsn :")
	//
	//// current support tcp connection only
	//dsn := fmt.Sprintf("tcp(%s:%d)/%s", host, port, database)
	//if user != "" {
	//	if password != "" {
	//		dsn = fmt.Sprintf("%s:%s@", user, password) + dsn
	//	} else {
	//		dsn = fmt.Sprintf("%s@", user) + dsn
	//	}
	//}

	//query := url.Values{}
	//if charset != "" {
	//	query.Set("charset", charset)
	//}
	//if client.Timeout > 0 {
	//	stimeout := client.Timeout.String()
	//	query.Set("readTimeout", stimeout)
	//	query.Set("writeTimeout", stimeout)
	//}
	//
	//s := query.Encode()
	//if s != "" {
	//	dsn += "?" + s
	//}

	var err error
	dsn := "server=" + host + ";user id=" + user + ";password=" + password + ";database=" + database

	//fmt.Println("dsn ",dsn)
	client.DB, err = sql.Open("mssql", dsn)
	client.DB.SetMaxOpenConns(200)		// 设置最大连接池数量
	client.DB.SetMaxIdleConns(100)

	if err != nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))

		fmt.Println("sql server  数据库连接错误", err.Error())
		return 2
	}

	//GlobalDB = client.DB
	fmt.Println("sql server 数据库连接成功")
	L.Push(lua.LBool(true))
	return 1
}
