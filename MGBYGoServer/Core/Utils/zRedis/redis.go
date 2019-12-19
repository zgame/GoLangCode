package zRedis

import (
	"github.com/gomodule/redigo/redis"
	"../zLog"
	"fmt"
	"encoding/json"
)


var 	RRedis redis.Conn

// Redis数据库初始化
func InitRedis(address string, pwd string) (bool){
	re, err := redis.Dial("tcp", address)
	if err != nil {
		zLog.PrintLogger("Redis 服务器连接不上" + address + err.Error())
		return false
	}
	// 密码验证
	if _, err := re.Do("AUTH", pwd); err != nil {
		re.Close()
		zLog.PrintLogger("Redis 服务器密码不正确")
		return false
	}
	RRedis = re
	fmt.Println("redis 数据库 ok ！")

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
	return true

}

//-------------------------------------------通用函数-----------------------------------------------

// 通用函数， 第一个返回值是字符串， 第二个是数字
func CmdForRedis(cmd string, args ...interface{}) (string, int64) {
	ret, err := RRedis.Do(cmd, args...)
	if err != nil {
		zLog.PrintfLogger("=======redis  CmdForRedis ========= %s  %s   出错了: %s", cmd, args, err.Error())
		//panic("redis    出错 " + cmd  + "  " + err.Error())
	}
	re1,_ := redis.String(ret,err)
	re2,_ := redis.Int64(ret,err)

	if re1 != "" {
		fmt.Printf("redes do : %s  %v   result: %s \n", cmd, args, re1)
	}else{
		fmt.Printf("redes do : %s  %v   result: %d \n", cmd, args, re2)
	}

	return re1, re2
}

// 获取多个string返回值， json格式的数组形式
func GetStringListFromRedis(cmd string,args ...interface{}) string {
	ret, err := redis.Values(RRedis.Do(cmd, args...))
	if err != nil {
		zLog.PrintfLogger("=======redis  GetStringListFromRedis =========  出错了 %s %v : %s", cmd,args, err.Error())
		//panic("redis list 读取出错 " + cmd + "  " + err.Error())
	}
	if ret != nil {
		var str []string
		for _, v := range ret {
			//fmt.Println("",string(v.([]byte)))
			str = append(str, string(v.([]byte)))
		}
		data, _ := json.MarshalIndent(str, "", " ")
		fmt.Printf("redes do : %s  %v   result: %s \n", cmd, args, string(data))
		return string(data)
	} else {
		//fmt.Println("获取到数据为空")
		return ""
	}
}

//------------------------------------------- 脚本方式 -----------------------------------------------

// redis直接运行lua的脚本， 这个主要是用来进行分布式的统一性， 可以避免加分布式锁， 广泛用在处理跨服的活动上面，比如分布式抢红包，世界boss受伤，boss击杀， 主要是保证数值的增加，或者减少是分布式统一协调的
func RedisRunLuaScript(luaScript string, name string) int {

	var AddScript = redis.NewScript(0, luaScript)
	v, err := AddScript.Do(RRedis)
	if err != nil {
		zLog.PrintLogger(name + "  RedisRunLuaScript Error: " + err.Error())
		//panic(name + "  RedisRunLuaScript Error: " + err.Error())
		//os.Exit(0)
	}

	re := int(v.(int64))
	return re

}


//
//// 通用函数，适合返回string
//func CmdForRedisReturnString(cmd string , args ...interface{}) string {
//	ret, err := redis.String(RRedis.Do(cmd, args...))
//	if err != nil {
//		zLog.PrintfLogger("=======redis  CmdForRedisReturnString ========= %s  %s   出错了: %s", cmd, args, err.Error())
//		//panic("redis  出错 " + cmd + "  " + err.Error())
//	}
//	return ret
//}
//
//// 通用函数，适合返回数字
//func CmdForRedisReturnInt64(cmd string , args ...interface{}) int64 {
//	ret, err := redis.Int64(RRedis.Do(cmd, args...))
//	if err != nil {
//		zLog.PrintfLogger("=======redis  CmdForRedisReturnInt64 ========= %s  %s   出错了: %s", cmd, args, err.Error())
//		//panic("redis 出错 " + cmd + "  " + err.Error())
//	}
//	return ret
//}




//-------------------------------------------hash-----------------------------------------------

//// 保存数据            dir 组信息 key value
//func SaveStringToRedis(dir string, key string, value string) {
//	//startTime := ztimer.GetOsTimeMillisecond()
//
//	//data, _ := json.MarshalIndent(player, "", " ")
//	//key := "BY_Player_UID_"+ strconv.Itoa(int( player.UserId))
//
//	//_, err := RRedis.Do("hdel", dir, key)
//	//fmt.Println("保存",dir, key,value)
//
//	ztimer.CheckRunTimeCost(func() {
//		_, err := RRedis.Do("hset", dir, key, value)
//		if err != nil {
//			zLog.PrintfLogger("============redis=========== 保存 %s   时候出错了: %s", dir, err.Error())
//			panic("redis 保存出错 " + key + "  " + err.Error())
//		}
//	}, "SaveStringToRedis")
//
//	//if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
//	//	zLog.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ]SaveStringToRedis消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
//	//}
//	//if ret == '1'{
//	//	fmt.Println("save success", ret)
//	//} else {
//	//	fmt.Println("save failed", ret)
//	//}
//	//fmt.Println("redis save ",ret)
//}

// 获取数据
//func GetStringFromRedis(dir string, key string) string {
//
//	//fmt.Println("获取redis数据",dir, key)
//	//startTime := ztimer.GetOsTimeMillisecond()
//	//var key string
//	//key = "BY_Player_UID_"+ strconv.Itoa(uid)
//	//fmt.Println("",key)
//
//	ret, err := RRedis.Do("hget", dir, key)
//
//	//ret, err := RRedis.Do("hget", "ALL_Players", "BY_Player_UID_2027445")
//	//fmt.Println("=======redis======== 读取 ", dir, key , reflect.TypeOf(ret))
//
//	if err != nil {
//		zLog.PrintfLogger("=======redis========= 读取 %s  %s 出错了: %s", dir, key, err.Error())
//		panic("redis 读取出错 " + key + "  " + err.Error())
//	}
//	if ret != nil {
//		//fmt.Println("收到",string(ret.([]byte)))
//		//if err := json.Unmarshal(ret.([]byte), &player); err != nil {
//		//	zLog.Fatalf("JSON unmarshaling failed: %s", err)
//		//}
//
//		// 这里注释掉了
//		//if reflect.TypeOf(ret).Kind() == reflect.Int64 {
//		//	fmt.Println("ret int64 : ",  int(ret.(int64)))
//		//	return ""
//		//}
//		return string(ret.([]byte))
//	} else {
//		//fmt.Println("获取到数据为空")
//		return ""
//	}
//
//	//fmt.Println(ret.(string))
//	//var player UserModel.UserModel
//
//}

////删除数据
//func DelKeyToRedis(dir string, key string) {
//
//	_, err := RRedis.Do("hdel", dir, key)
//	if err != nil {
//		zLog.PrintfLogger("redis 删除key %s 出错了:"+err.Error(), key)
//		panic("redis 删除key  出错了:   " + key + "      " + err.Error())
//	}
//
//}

//// 获取是否存在key值，或者是否存在hash值， 如果不传hashKey，就只判断key
//func ExistKeyInRedis(key string, hashKey string) int {
//	var err error
//	var result interface{}
//
//	if hashKey == "" {
//		result, err = RRedis.Do("exists", key)
//	} else {
//		result, err = RRedis.Do("hexists", key, hashKey)
//	}
//	if err != nil {
//		zLog.PrintfLogger("redis 获取 key %s  hashKey %s 是否存在出错了:"+err.Error(), key, hashKey)
//		panic("redis 获取键值是否存在出错 ! key :   " + key + "   hashKey: " + hashKey + "   " + err.Error())
//	}
//
//	return int(result.(int64))
//}

//
//// 执行脚本，用于分布式，不用加锁的情况，因为脚本一次性执行的，类似存储过程
////---- KEYS[1] 是dir
////---- KEYS[2] 是key
////---- ARGV[1] 是参数
//var AddScript = redis.NewScript(2,`
//   local r = redis.call('hget',KEYS[1],KEYS[2])
//   if r ~= false then
//		r = r + ARGV[1]
//   else
//	    r = 1000000001
//   end
//   redis.call('hset', KEYS[1],KEYS[2], r)
//   return r
//`)
//
//// 这里的参数num， 记住： num只是增量， 你只能要求人家增加多少， 具体完事之后是多少，会返回给你， 因为这个涉及到分布式多请求
//func AddNumberToRedis(dir string,key string, num int) int{
//	startTime := ztimer.GetOsTimeMillisecond()
//	v, err := AddScript.Do(RRedis, dir,key, num)
//	if err != nil {
//		zLog.PrintLogger("AddNumberToRedis Error: " + err.Error())
//		panic("AddNumberToRedis Error: " + err.Error())
//		os.Exit(0)
//	}
//	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
//		zLog.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ] AddNumberToRedis 消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
//	}
//	//fmt.Println("AddNumberToRedis",v , reflect.TypeOf(v))
//	re := int(v.(int64))
//	//fmt.Println("========AddNumberToRedis===============",re)
//	return re
//}

//------------------------------------------- list -----------------------------------------------

// 增加数据
//func AddListFromRedis(dir string, value string) int {
//	ret, err := RRedis.Do("lpush", dir, value)
//	if err != nil {
//		zLog.PrintfLogger("=======redis  list ========= add %s  %s 出错了: %s", dir, err.Error())
//		panic("redis list 读取出错 " + dir + "  " + err.Error())
//	}
//	if ret != nil {
//		re := int(ret.(int64))
//		return re
//	} else {
//		//fmt.Println("获取到数据为空")
//		return 0
//	}
//}


//// 删除数据
//func DelListFromRedis(dir string, value string) int {
//	ret, err := RRedis.Do("lrem", dir, 0, value)
//	if err != nil {
//		zLog.PrintfLogger("=======redis  list del ========= 读取 %s  %s 出错了: %s", dir, err.Error())
//		panic("redis list del 出错 " + dir + "  " + err.Error())
//	}
//	if ret != nil {
//		re := int(ret.(int64))
//		return re
//	} else {
//		//fmt.Println("获取到数据为空")
//		return 0
//	}
//}
//
//// 删除数据最后一个
//func DelLastFromRedis(dir string) int {
//	_, err := RRedis.Do("rpop", dir)
//	if err != nil {
//		zLog.PrintfLogger("=======redis  list del rpop =========  %s  %s 出错了: %s", dir, err.Error())
//		panic("redis list del  rpop 出错 " + dir + "  " + err.Error())
//	}
//	//if ret!= nil {
//	//	return string(ret.(string))
//	//}else{
//	//	//fmt.Println("获取到数据为空")
//	//	return ""
//	//}
//	//if ret!= nil {
//	//	re := int(ret.(int64))
//	//	return re
//	//}else{
//	//	//fmt.Println("获取到数据为空")
//	return 0
//	//}
//}

//------------------------------------------- set -----------------------------------------------

//// 是否在集合中数据
//func GetSetFromRedis(dir string, value string) int {
//	ret, err := RRedis.Do("sismember", dir, value)
//	if err != nil {
//		zLog.PrintfLogger("=======redis  set ========= 读取 %s  %s 出错了: %s", dir, err.Error())
//		panic("redis set 读取出错 " + dir + "  " + err.Error())
//	}
//	if ret != nil {
//		re := int(ret.(int64))
//		return re
//	} else {
//		//fmt.Println("获取到数据为空")
//		return 0
//	}
//}
//
//// 删除数据
//func DelSetFromRedis(dir string, value string) int {
//	ret, err := RRedis.Do("srem", dir, value)
//	if err != nil {
//		zLog.PrintfLogger("=======redis  set del ========= 读取 %s  %s 出错了: %s", dir, err.Error())
//		panic("redis set del 出错 " + dir + "  " + err.Error())
//	}
//	if ret != nil {
//		re := int(ret.(int64))
//		return re
//	} else {
//		//fmt.Println("获取到数据为空")
//		return 0
//	}
//}
