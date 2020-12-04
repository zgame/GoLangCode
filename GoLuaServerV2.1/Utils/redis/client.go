//------------------------------------------------------------------------------------------
// redis 的数据库
//------------------------------------------------------------------------------------------

package redis

import (
	"github.com/gomodule/redigo/redis"
	lua "github.com/yuin/gopher-lua"
	"time"
)

const (
	CLIENT_TYPENAME = "redis{client}"
)

type Client struct {
	redis   redis.Conn
	Timeout time.Duration
}

var clientMethods = map[string]lua.LGFunction{
	"connect": clientConnectMethod,
	"close":   clientCloseMethod,
	"cmd":     CmdForRedis,
	"stringList": GetStringListFromRedis,
	"script":  RunLuaScript,
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

	client.redis.Close()

	// always clean
	client.redis = nil
	//if err != nil {
	//	L.Push(lua.LBool(false))
	//	L.Push(lua.LString(err.Error()))
	//	return 2
	//}

	L.Push(lua.LBool(true))
	return 1
}
