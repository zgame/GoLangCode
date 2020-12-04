//------------------------------------------------------------------------------------------
// redis 的数据库查询
//------------------------------------------------------------------------------------------
package redis

import (
	"GoLuaServerV2.1/Utils/zLog"
	"fmt"
	"github.com/gomodule/redigo/redis"
	gluasql_util "github.com/tengattack/gluasql/util"
	"github.com/yuin/gopher-lua"
	"reflect"
	"time"
)


//-------------------------------------------通用函数-----------------------------------------------

// 通用函数， 第一个返回值是字符串， 第二个是数字
func CmdForRedis(L *lua.LState ) int {
	client := checkClient(L)
	cmd := L.ToString(2)
	args,ok := gluasql_util.GetValue(L, 3).([]interface{})		//强转为数组
	if !ok {
		zLog.PrintfLogger("redis cmd :%s 参数转换成数组出错 ",cmd )
		return 0
	}

	//args := make([]interface{},0)
	//luaScript := L.ToString(1)
	//name := L.ToString(2)

	ret, err := client.redis.Do(cmd, args...)
	if err != nil {
		zLog.PrintfLogger("=======redis  CmdForRedis ========= %s  %s   出错了: %s", cmd, args, err.Error())
		//panic("redis    出错 " + cmd  + "  " + err.Error())
		return 0
	}
	re1,_ := redis.String(ret,err)
	re2,_ := redis.Int64(ret,err)

	//if re1 != "" {
	//	fmt.Printf("redes do : %s  %v   result: %s \n", cmd, args, re1)
	//}else{
	//	fmt.Printf("redes do : %s  %v   result: %d \n", cmd, args, re2)
	//}
	L.Push(lua.LString(re1))
	L.Push(lua.LNumber(re2))

	return 2
}

// 获取多个string返回值， json格式的数组形式
func GetStringListFromRedis(L *lua.LState ) int {
	client := checkClient(L)
	cmd := L.ToString(2)
	args,ok := gluasql_util.GetValue(L, 3).([]interface{})		//强转为数组
	if !ok {
		zLog.PrintfLogger("redis string list :%s 参数转换成数组出错 ",cmd )
		return 0
	}

	ret, err := redis.Values(client.redis.Do(cmd, args...))
	if err != nil {
		zLog.PrintfLogger("=======redis  GetStringListFromRedis =========  出错了 %s %v : %s", cmd,args, err.Error())
		return 0
		//panic("redis list 读取出错 " + cmd + "  " + err.Error())
	}
	//result := ""
	tb := L.NewTable()
	if ret != nil {
		//var str []string
		for _ , v := range ret {
			tb.Append(lua.LString(string(v.([]byte))))
			//fmt.Println("",string(v.([]byte)))
			//str = append(str, string(v.([]byte)))
		}
		//data, _ := json.MarshalIndent(str, "", " ")
		//fmt.Printf("redes do : %s  %v   result: %s \n", cmd, args, string(data))
		//result =  string(data)
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


//------------------------------------------------------------------------------------
// utils
//------------------------------------------------------------------------------------

// map to  lua table
func ToTableFromMap(l *lua.LState, v reflect.Value) lua.LValue {
	tb := &lua.LTable{}
	for _, k := range v.MapKeys(){
		key := ToArbitraryValue(l, k.Interface())
		if key!= lua.LString("_id") {
			tb.RawSet(key, ToArbitraryValue(l, v.MapIndex(k).Interface()))
		}
	}
	return tb
}
// slice to lua table
func ToTableFromSlice(l *lua.LState, v reflect.Value) lua.LValue {
	tb := &lua.LTable{}
	for j := 0; j < v.Len(); j++ {
		tb.RawSet(ToArbitraryValue(l, j+1), // because lua is 1-indexed
			ToArbitraryValue(l, v.Index(j).Interface()))
	}
	return tb
}

func ToArbitraryValue(l *lua.LState, i interface{}) lua.LValue {
	if i == nil {
		return lua.LNil
	}

	switch ii := i.(type) {
	case bool:
		return lua.LBool(ii)
	case int:
		return lua.LNumber(ii)
	case int8:
		return lua.LNumber(ii)
	case int16:
		return lua.LNumber(ii)
	case int32:
		return lua.LNumber(ii)
	case int64:
		return lua.LNumber(ii)
	case uint:
		return lua.LNumber(ii)
	case uint8:
		return lua.LNumber(ii)
	case uint16:
		return lua.LNumber(ii)
	case uint32:
		return lua.LNumber(ii)
	case uint64:
		return lua.LNumber(ii)
	case float64:
		return lua.LNumber(ii)
	case float32:
		return lua.LNumber(ii)
	case string:
		return lua.LString(ii)
	case []byte:
		return lua.LString(ii)
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Ptr:
			return ToArbitraryValue(l, v.Elem().Interface())

		//case reflect.Struct:
		//	return ToTableFromStruct(l, v)

		case reflect.Map:
			return ToTableFromMap(l, v)

		case reflect.Slice:
			return ToTableFromSlice(l, v)

		//case mgo.Index{}

		default:
			return lua.LString("")
		}
	}
}

// 用来调试打印的
func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
	fmt.Print("\t", reflect.TypeOf(*pval))

}
