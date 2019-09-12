package zRedis

import (
	//"github.com/gomodule/redigo/redis"
	"github.com/garyburd/redigo/redis"
	"../log"
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
	return true
}



// 保存数据            dir 组信息 key value
func SaveStringToRedis(dir string, key string,value string)  {

	//data, _ := json.MarshalIndent(player, "", " ")
	//key := "BY_Player_UID_"+ strconv.Itoa(int( player.UserId))

	//_, err := RRedis.Do("hdel", dir, key)
	//fmt.Println("保存",dir, key,value)
	_, err := RRedis.Do("hset", dir, key,value)
	if err !=nil {
		log.PrintLogger("redis 保存的时候出错了:"+err.Error())
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
	//var key string
	//key = "BY_Player_UID_"+ strconv.Itoa(uid)
	//fmt.Println("",key)
	ret, err :=  RRedis.Do("hget",dir, key)
	//ret, err := RRedis.Do("hget", "ALL_Players", "BY_Player_UID_2027445")
	//fmt.Println(reflect.TypeOf(ret))

	if err !=nil {
		log.PrintLogger("redis 读取出错了:"+err.Error())
	}
	//fmt.Println(ret.(string))
	//var player Player.Player
	if ret!= nil {
		//fmt.Println("收到",string(ret.([]byte)))
		//if err := json.Unmarshal(ret.([]byte), &player); err != nil {
		//	log.Fatalf("JSON unmarshaling failed: %s", err)
		//}
		return string(ret.([]byte))
	}else{
		//fmt.Println("获取到数据为空")
		return ""
	}

}


//删除数据
func DelKeyToRedis(dir string,key string){
	_, err :=  RRedis.Do("hdel",dir, key)
	if err !=nil {
		log.PrintfLogger("redis 删除key %s 出错了:"+err.Error(), key)
	}
}





// 执行脚本，用于分布式，不用加锁的情况，因为脚本一次性执行的，类似存储过程
var AddScript = redis.NewScript(2,`
   local r = redis.call('hget',KEYS[1],KEYS[2])
   if r ~= nil then
		r = r + ARGV[1]
       redis.call('hset', KEYS[1],KEYS[2], r)
   end
   return r
`)
// 这里的参数num， 记住： num只是增量， 你只能要求人家增加多少， 具体完事之后是多少，会返回给你， 因为这个涉及到分布式多请求
func AddNumberToRedis(dir string,key string, num int) int{
	v, err := AddScript.Do(RRedis, dir,key, num)
	if err != nil {
		log.PrintLogger("AddNumberToRedis Error: " + err.Error())
		return 0
	}
	//fmt.Println("AddNumberToRedis",v , reflect.TypeOf(v))
	re := int(v.(int64))
	return re
}

