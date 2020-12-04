//------------------------------------------------------------------------------------------
//redis 的数据库连接
//------------------------------------------------------------------------------------------

package redis

import (
	"GoLuaServerV2.1/Utils/zLog"
	"github.com/gomodule/redigo/redis"

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

	address, _ := options["host"].(string)
	password, _ := options["password"].(string)


	re, err := redis.Dial("tcp", address)
	if err != nil {
		zLog.PrintLogger("Redis 服务器连接不上"+ address + err.Error())
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	// 密码验证
	if _, err := re.Do("AUTH", password); err != nil {
		re.Close()
		zLog.PrintLogger("Redis 服务器密码不正确")
		L.Push(lua.LBool(false))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	client.redis = re

		//GlobalDB = client.DB
	fmt.Println("redis 数据库连接成功")
	L.Push(lua.LBool(true))
	return 1
}

/*
	//----------------------------------------------------------------------------------------
	// test  这是测试用的例子 ， 可以作为参考
	//----------------------------------------------------------------------------------------

	fmt.Println("//-------------------------------------------- string --------------------------------------------------")
	CmdForRedis("set" ,"zsw", "zsw_value1")			// 成功返回  "OK"
	CmdForRedis("get" ,"zsw")							// 存在返回 (value,0)   如果不存在返回("",0)
	CmdForRedis("setnx" ,"zsw", "zsw_value11") 		// 如果不存在，那么set   成功返回1 ，失败返回0
	fmt.Println("//-------------------------------------------- key --------------------------------------------------")
	CmdForRedis("exists" ,"zsw")						// 存在返回1 ，不存在返回0
	CmdForRedis("setex" ,"zsw", "50","zsw_value11") 	// set  value ，并且设定过期时间 *秒   成功返回 "OK"
	CmdForRedis("expire" ,"zsw", "100")              	// 设定过期时间   成功返回1 ，失败返回0
	CmdForRedis("expireat" ,"zsw", "2293840000")   	// 设定过期时间 在某个系统时间戳定义的时间过期    成功返回1 ，失败返回0
	CmdForRedis("ttl" ,"zsw")                          // 返回过期时间
	CmdForRedis("del" ,"zsw")							// 返回删除key的数量
	fmt.Println("//-------------------------------------------- hash --------------------------------------------------")
	CmdForRedis("hset" ,"zsw_hash", "zsw_key1","zsw_value11")	// 成功返回1 ，失败返回0
	CmdForRedis("hget" ,"zsw_hash", "zsw_key1")		// 成功返回value ，失败返回 ""
	CmdForRedis("hexists" ,"zsw_hash", "zsw_key1")		// 存在返回1 ，不存在返回0
	CmdForRedis("hdel" ,"zsw_hash", "zsw_key1")		// 成功返回1 ，失败返回0
	fmt.Println("//-------------------------------------------- list --------------------------------------------------")
	CmdForRedis("lpush" ,"zsw_list", "zsw_key1")                          // 从头部插入 返回list 长度
	CmdForRedis("rpush" ,"zsw_list", "zsw_key1")                          // 从尾部插入 返回list 长度
	CmdForRedis("llen" ,"zsw_list")                                       // 返回list 长度
	CmdForRedis("lindex" ,"zsw_list",0)                                   // 获取 index 元素
	CmdForRedis("lset" ,"zsw_list",0, "zsw_key2")                         // 成功返回  "OK"
	GetStringListFromRedis("lrange", "zsw_list", 0, -1)				   // 获取所有元素
	CmdForRedis("lrem" ,"zsw_list",0 ,"zsw_key1")                         // 移除 一定数量的 value , > 0 从头计数 ， < 0 从尾部计数 , 0 是所有 , 返回值是移除元素数量
	CmdForRedis("rpop" ,"zsw_list")                                       // 删尾部一个，返回删除的值， 不存在的话 ( "",0)
	CmdForRedis("lpop" ,"zsw_list")                                       // 删头部一个，返回删除的值， 不存在的话 ( "",0)
	fmt.Println("//-------------------------------------------- set --------------------------------------------------")
	CmdForRedis("sadd" ,"zsw_set","zsw_key1")			// 添加成员 , 返回成员数量
	CmdForRedis("scard" ,"zsw_set")					// 成员数量
	CmdForRedis("sismember" ,"zsw_set", "zsw_key1")	// 是否是成员 存在返回1 ，不存在返回0
	GetStringListFromRedis("smembers", "zsw_set")		// 获取所有元素
	CmdForRedis("srem" ,"zsw_set", "zsw_key1")			// 移除
	fmt.Println("//-------------------------------------------- sorted set --------------------------------------------------")
	CmdForRedis("zadd" ,"zsw_zset", 100, "zsw_key1")			// 添加成员 , 返回成功添加成员数量
	CmdForRedis("zadd" ,"zsw_zset", 1100, "zsw_key2")			// 添加成员 , 返回成功添加成员数量
	CmdForRedis("zadd" ,"zsw_zset", 3100, "zsw_key3")			// 添加成员 , 返回成功添加成员数量
	CmdForRedis("zadd" ,"zsw_zset", 4100, "zsw_key4")			// 添加成员 , 返回成功添加成员数量
	CmdForRedis("zadd" ,"zsw_zset", 110, "zsw_key1")			// 更新成员 , 返回成功添加成员数量 , 更新返回 0
	CmdForRedis("zscore" ,"zsw_zset", "zsw_key1")				// 返回成员分数
	CmdForRedis("zcount" ,"zsw_zset", 0, 1000)					// 返回成员数量在分数范围内
	CmdForRedis("zcard" ,"zsw_zset")							// 成员数量
	CmdForRedis("zrevrank", "zsw_zset","zsw_key1")				// 获取成员排名， 从大到小, 排名第一是0
	CmdForRedis("zrank", "zsw_zset","zsw_key1")					// 获取成员排名， 从小到大,排名第一是0
	GetStringListFromRedis("zrange", "zsw_zset",0,1, "withscores")	// 获取 all 成员排名 从小到大   0,-1 是所有
	GetStringListFromRedis("zrange", "zsw_zset",0,-1)				// 获取 all 成员排名 从小到大   0,-1 是所有
	GetStringListFromRedis("zrevrange", "zsw_zset",0,2, "withscores")// 获取 all 成员排名 从大到小   0,-1 是所有
	GetStringListFromRedis("zrangebyscore", "zsw_zset",0,3000)		// 获取  成员区间排名	score
	GetStringListFromRedis("zrevrangebyscore", "zsw_zset",3000,0)		// 获取  成员区间排名 score
	GetStringListFromRedis("zrangebylex", "zsw_zset","(zsw_key1","[zsw_key4")// 获取  成员区间排名
	GetStringListFromRedis("zrevrangebylex", "zsw_zset","(zsw_key1","[zsw_key4")// 获取  成员区间排名
	CmdForRedis("zremRangeByLex", "zsw_zset","[zsw_key1","[zsw_key2")		// 移除按成员区间
	CmdForRedis("zremRangeByRank", "zsw_zset",0,0)				// 移除按排名 按排名， 返回移除成员数量
	CmdForRedis("zremRangeByScore", "zsw_zset",0,10000)			// 移除按排名 按分数， 返回移除成员数量

	//return false		// test open
*/