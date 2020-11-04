package zRedis

import (
	"github.com/gomodule/redigo/redis"
	"ServerZgameTB/GlobalVar"
	//"github.com/garyburd/redigo/redis"
	"ServerZgameTB/Utils/log"
	"ServerZgameTB/Utils/ztimer"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

var RRedis redis.Conn

// Redis数据库初始化
func InitRedis(address string, pwd string) bool{
	re, err := redis.Dial("tcp", address)
	if err != nil {
		log.PrintLogger("Redis 服务器连接不上"+ address + err.Error())
		return false
	}
	// 密码验证
	if _, err := re.Do("AUTH", pwd); err != nil {
		re.Close()
		log.PrintLogger("Redis 服务器密码不正确")
		return false
	}
	RRedis = re
	fmt.Println("redis 数据库 ok ！")
	return true
}

//-------------------------------------------hash-----------------------------------------------

// 保存数据            dir 组信息 key value
func SaveStringToRedis(dir string, key string,value string)  {
	startTime := ztimer.GetOsTimeMillisecond()

	//data, _ := json.MarshalIndent(player, "", " ")
	//key := "BY_Player_UID_"+ strconv.Itoa(int( player.UserId))

	//_, err := RRedis.Do("hdel", dir, key)
	//fmt.Println("保存",dir, key,value)
	_, err := RRedis.Do("hset", dir, key,value)
	if err !=nil {
		log.PrintfLogger("============redis=========== 保存 %s   时候出错了: %s",dir,err.Error())
		panic("redis 保存出错 " + key + "  " + err.Error())
	}
	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		log.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ]SaveStringToRedis消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
	}
	//if ret == '1'{
	//	fmt.Println("save success", ret)
	//} else {
	//	fmt.Println("save failed", ret)
	//}
	//fmt.Println("redis save ",ret)
}



// 获取数据
func GetStringFromRedis(dir string,key string) string {

	//fmt.Println("获取redis数据",dir, key)
	startTime := ztimer.GetOsTimeMillisecond()
	//var key string
	//key = "BY_Player_UID_"+ strconv.Itoa(uid)
	//fmt.Println("",key)
	ret, err :=  RRedis.Do("hget",dir, key)

	//ret, err := RRedis.Do("hget", "ALL_Players", "BY_Player_UID_2027445")
	//fmt.Println("=======redis======== 读取 ", dir, key , reflect.TypeOf(ret))

	if err !=nil {
		log.PrintfLogger("=======redis========= 读取 %s  %s 出错了: %s", dir, key ,err.Error())
		panic("redis 读取出错 " + key + "  " + err.Error())
	}
	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		log.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ]GetStringFromRedis消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
	}
	//fmt.Println(ret.(string))
	//var player Player.Player
	if ret!= nil {
		//fmt.Println("收到",string(ret.([]byte)))
		//if err := json.Unmarshal(ret.([]byte), &player); err != nil {
		//	log.Fatalf("JSON unmarshaling failed: %s", err)
		//}
		if reflect.TypeOf(ret).Kind() == reflect.Int64 {
			fmt.Println("ret int64 : ",  int(ret.(int64)))
			return ""
		}
		return string(ret.([]byte))
	}else{
		//fmt.Println("获取到数据为空")
		return ""
	}

}


//删除数据
func DelKeyToRedis(dir string,key string){
	startTime := ztimer.GetOsTimeMillisecond()
	_, err :=  RRedis.Do("hdel",dir, key)
	if err !=nil {
		log.PrintfLogger("redis 删除key %s 出错了:"+err.Error(), key)
		panic("redis 删除key  出错了:   "+key+"      "+err.Error())
	}
	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		log.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ]DelKeyToRedis消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
	}
}


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
//		log.PrintLogger("AddNumberToRedis Error: " + err.Error())
//		panic("AddNumberToRedis Error: " + err.Error())
//		os.Exit(0)
//	}
//	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
//		log.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ] AddNumberToRedis 消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
//	}
//	//fmt.Println("AddNumberToRedis",v , reflect.TypeOf(v))
//	re := int(v.(int64))
//	//fmt.Println("========AddNumberToRedis===============",re)
//	return re
//}

//------------------------------------------- 脚本方式 -----------------------------------------------

// redis直接运行lua的脚本， 这个主要是用来进行分布式的统一性， 可以避免加分布式锁， 广泛用在处理跨服的活动上面，比如分布式抢红包，世界boss受伤，boss击杀， 主要是保证数值的增加，或者减少是分布式统一协调的
func RedisRunLuaScript(luaScript string ,name string)  int{
	startTime := ztimer.GetOsTimeMillisecond()
	var AddScript = redis.NewScript(0,luaScript)
	v, err := AddScript.Do(RRedis)
	if err != nil {
		log.PrintLogger(  name + "  RedisRunLuaScript Error: " + err.Error())
		panic(name + "  RedisRunLuaScript Error: " + err.Error())
		os.Exit(0)
	}
	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		log.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ] RedisRunLuaScript  %s 消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime),  name)
	}
	re := int(v.(int64))
	return re

}


//------------------------------------------- list -----------------------------------------------


// 增加数据
func AddListFromRedis(dir string, value string) int {
	ret, err :=   RRedis.Do("lpush",dir, value)
	if err !=nil {
		log.PrintfLogger("=======redis  list ========= add %s  %s 出错了: %s", dir, err.Error())
		panic("redis list 读取出错 " + dir + "  " + err.Error())
	}
	if ret!= nil {
		re := int(ret.(int64))
		return re
	}else{
		//fmt.Println("获取到数据为空")
		return 0
	}
}
// 获取数据
func GetListFromRedis(dir string) string {
	ret, err :=   redis.Values( RRedis.Do("lrange",dir, 0,-1))
	if err !=nil {
		log.PrintfLogger("=======redis  list ========= 读取 %s  %s 出错了: %s", dir, err.Error())
		panic("redis list 读取出错 " + dir + "  " + err.Error())
	}
	if ret!= nil {
		var str []string
		for _,v := range ret{
			//fmt.Println("",string(v.([]byte)))
			str = append(str, string(v.([]byte)))
		}
		data, _ := json.MarshalIndent(str, "", " ")
		return string(data)
	}else{
		//fmt.Println("获取到数据为空")
		return ""
	}
}

// 删除数据
func DelListFromRedis(dir string, value string) int {
	ret, err :=   RRedis.Do("lrem",dir, 0,value)
	if err !=nil {
		log.PrintfLogger("=======redis  list del ========= 读取 %s  %s 出错了: %s", dir, err.Error())
		panic("redis list del 出错 " + dir + "  " + err.Error())
	}
	if ret!= nil {
		re := int(ret.(int64))
		return re
	}else{
		//fmt.Println("获取到数据为空")
		return 0
	}
}

// 删除数据最后一个
func DelLastFromRedis(dir string) int {
	ret, err := RRedis.Do("rpop",dir)
	if err !=nil {
		log.PrintfLogger("=======redis  list del rpop =========  %s  %s 出错了: %s", dir, err.Error())
		panic("redis list del  rpop 出错 " + dir + "  " + err.Error())
	}
	//if ret!= nil {
	//	return string(ret.(string))
	//}else{
	//	//fmt.Println("获取到数据为空")
	//	return ""
	//}
	if ret!= nil {
		re := int(ret.(int64))
		return re
	}else{
		//fmt.Println("获取到数据为空")
		return 0
	}
}


//------------------------------------------- set -----------------------------------------------

// 是否在集合中数据
func GetSetFromRedis(dir string, value string) int {
	ret, err :=   RRedis.Do("sismember",dir, value)
	if err !=nil {
		log.PrintfLogger("=======redis  set ========= 读取 %s  %s 出错了: %s", dir, err.Error())
		panic("redis set 读取出错 " + dir + "  " + err.Error())
	}
	if ret!= nil {
		re := int(ret.(int64))
		return re
	}else{
		//fmt.Println("获取到数据为空")
		return 0
	}
}

// 删除数据
func DelSetFromRedis(dir string, value string) int {
	ret, err :=   RRedis.Do("srem",dir, value)
	if err !=nil {
		log.PrintfLogger("=======redis  set del ========= 读取 %s  %s 出错了: %s", dir, err.Error())
		panic("redis set del 出错 " + dir + "  " + err.Error())
	}
	if ret!= nil {
		re := int(ret.(int64))
		return re
	}else{
		//fmt.Println("获取到数据为空")
		return 0
	}
}
