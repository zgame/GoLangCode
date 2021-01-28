//------------------------------------------------------------------------------------------
// redis 的数据库查询
//------------------------------------------------------------------------------------------
package redis

import (
	"GoLuaServerV2.1/Core/Utils/zLog"
	"GoLuaServerV2.1/Core/Utils/zLua"
	"github.com/gomodule/redigo/redis"
	"github.com/yuin/gopher-lua"
)


//-------------------------------------------通用函数-----------------------------------------------

// 通用函数， 第一个返回值是字符串， 第二个是数字
func CmdForRedis(L *lua.LState ) int {
	client := checkClient(L)
	cmd := L.ToString(2)

	array := zLua.LuaGetValue(L, 3)
	args,ok := array.([]interface{}) //强转为数组
	if !ok {
		zLog.PrintfLogger("redis cmd :%s 参数转换成数组出错 ",cmd )
		return 0
	}
	//fmt.Printf("%v \n",args)
	if client.redis == nil {
		return 0
	}
	ret, err := client.redis.Do(cmd, args...)
	if err != nil {
		zLog.PrintfLogger("=======redis  CmdForRedis ========= %s  %s   出错了: %s", cmd, args, err.Error())
		//panic("redis    出错 " + cmd  + "  " + err.Error())
		return 0
	}
	re1,_ := redis.String(ret,err)
	re2,_ := redis.Int64(ret,err)

	L.Push(lua.LString(re1))
	L.Push(lua.LNumber(re2))

	return 2
}

// 获取多个string返回值， json格式的数组形式
func GetStringListFromRedis(L *lua.LState ) int {
	client := checkClient(L)
	cmd := L.ToString(2)
	args,ok := zLua.LuaGetValue(L, 3).([]interface{}) //强转为数组
	if !ok {
		zLog.PrintfLogger("redis string list :%s 参数转换成数组出错 ",cmd )
		return 0
	}


	ret, err := redis.Values(client.redis.Do(cmd, args...))
	if err != nil {
		zLog.PrintfLogger("=======redis  GetStringListFromRedis =========  出错了 %s %v : %s", cmd,args, err.Error())
		return 0
	}
	tb := L.NewTable()
	if ret != nil {
		for _ , v := range ret {
			tb.Append(lua.LString(string(v.([]byte))))
		}
	}

	L.Push(tb)
	return 1
}

//------------------------------------------- 脚本方式 -----------------------------------------------

// redis直接运行lua的脚本， 这个主要是用来进行分布式的统一性， 可以避免加分布式锁， 广泛用在处理跨服的活动上面，比如分布式抢红包，世界boss受伤，boss击杀， 主要是保证数值的增加，或者减少是分布式统一协调的
func RunLuaScript(L *lua.LState) int {
	client := checkClient(L)
	luaScript := L.ToString(2)
	name := L.ToString(3)

	var AddScript = redis.NewScript(0, luaScript)
	v, err := AddScript.Do(client.redis)
	if err != nil {
		zLog.PrintLogger(name + "  RedisRunLuaScript Error: " + err.Error())
		L.Push(lua.LNil)
		L.Push(lua.LString(name + "  RedisRunLuaScript Error: " + err.Error()))
		return 2
		//panic(name + "  RedisRunLuaScript Error: " + err.Error())
		//os.Exit(0)
	}
	re := int(v.(int64))
	L.Push(lua.LNumber(re))
	return 1

}

