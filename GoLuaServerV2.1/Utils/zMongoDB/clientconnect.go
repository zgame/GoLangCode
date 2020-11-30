//------------------------------------------------------------------------------------------
//MongoDB 的数据库连接
//------------------------------------------------------------------------------------------

package zMongoDB

import (
	"gopkg.in/mgo.v2"
	"time"

	//"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/tengattack/gluasql/util"
	"github.com/yuin/gopher-lua"

	//_ "github.com/denisenkom/go-mssqldb"
)

func clientConnectMethod(L *lua.LState) int {

	client := checkClient(L)
	tb := gluasql_util.GetValue(L, 2)
	options, ok := tb.(map[string]interface{})

	if tb == nil || !ok {
		L.ArgError(2, "options excepted")
		return 0
	}

	host, _ := options["host"].(string)
	if host == "" {
		host = "127.0.0.1:27017"
	}

	database, _ := options["database"].(string)
	user, _ := options["user"].(string)
	password, _ := options["password"].(string)


	var err error
	dialInfo := &mgo.DialInfo{
		Addrs: []string{host}, //远程(或本地)服务器地址及端口号
		Direct: false,
		Timeout: time.Second * 1,
		Database: database, 		//数据库
		Source: "admin",
		Username: user,
		Password: password,
		PoolLimit: 100, // Session.SetPoolLimit
	}
	client.MongoSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))

		fmt.Println("Mongo  数据库连接错误", err.Error())
		return 2
	}

	client.MongoSession.SetMode(mgo.Monotonic, true)

	//GlobalDB = client.DB
	fmt.Println("Mongo 数据库连接成功")
	L.Push(lua.LBool(true))
	return 1
}
