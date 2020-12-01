//------------------------------------------------------------------------------------------
// MongoDB 的数据库
//------------------------------------------------------------------------------------------


package mongoDB

import (
	"gopkg.in/mgo.v2"
	"time"
	//_ "github.com/go-sql-driver/mySql"
	"github.com/yuin/gopher-lua"
)

const (
	CLIENT_TYPENAME = "mongodb{client}"
)


type Client struct {
	MongoSession * mgo.Session
	Timeout time.Duration
}

var clientMethods = map[string]lua.LGFunction{
	"connect":       clientConnectMethod,
	"close":         clientCloseMethod,
	//"query":         clientQueryMethod,
	//"exec":         clientExecMethod,
	"insert":         clientInsertMethod,
	"del":         clientDelMethod,
	"update":         clientUpdateMethod,
	"find":         clientFindMethod,
	"finds":         clientFindsMethod,
}

func checkClient(L *lua.LState) *Client {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Client); ok {
		return v
	}
	L.ArgError(1, "client expected")
	return nil
}

func clientCloseMethod(L *lua.LState) int {
	client := checkClient(L)

	client.MongoSession.Close()

	// always clean
	client.MongoSession = nil
	//if err != nil {
	//	L.Push(lua.LBool(false))
	//	L.Push(lua.LString(err.Error()))
	//	return 2
	//}

	L.Push(lua.LBool(true))
	return 1
}
