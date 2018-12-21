package zRedis

import (
	"github.com/gomodule/redigo/redis"
	//"github.com/garyburd/redigo/redis"
	"../log"
	"../ztimer"
	"../../GlobalVar"
	"fmt"
	"reflect"
	"os"
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



// 执行脚本，用于分布式，不用加锁的情况，因为脚本一次性执行的，类似存储过程
var AddScript = redis.NewScript(2,`
   local r = redis.call('hget',KEYS[1],KEYS[2])
   if r ~= false then		
		r = r + ARGV[1]			
   else			
	    r = 1000000001	
   end
   redis.call('hset', KEYS[1],KEYS[2], r)
   return r
`)

// 这里的参数num， 记住： num只是增量， 你只能要求人家增加多少， 具体完事之后是多少，会返回给你， 因为这个涉及到分布式多请求
func AddNumberToRedis(dir string,key string, num int) int{
	startTime := ztimer.GetOsTimeMillisecond()
	v, err := AddScript.Do(RRedis, dir,key, num)
	if err != nil {
		log.PrintLogger("AddNumberToRedis Error: " + err.Error())
		panic("AddNumberToRedis Error: " + err.Error())
		os.Exit(0)
	}
	if ztimer.GetOsTimeMillisecond()-startTime > GlobalVar.WarningTimeCost {
		log.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ] AddNumberToRedis 消耗时间: %d", int(ztimer.GetOsTimeMillisecond()-startTime))
	}
	//fmt.Println("AddNumberToRedis",v , reflect.TypeOf(v))
	re := int(v.(int64))
	//fmt.Println("========AddNumberToRedis===============",re)
	return re
}

