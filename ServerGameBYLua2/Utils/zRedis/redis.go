package zRedis

import (
	//"github.com/gomodule/redigo/redis"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"../log"
)

var RRedis redis.Conn

// Redis数据库初始化
func InitRedis(address string, pwd string) bool{
	re, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println("Redis 服务器连接不上"+ address,err.Error())
		return false
	}
	// 密码验证
	if _, err := re.Do("AUTH", pwd); err != nil {
		re.Close()
		fmt.Println("Redis 服务器密码不正确")
		return false
	}
	RRedis = re
	return true
}

//func newPool(host string, db int) *redis.Pool {
//	return &redis.Pool {
//		MaxIdle: 50,
//		MaxActive: 100,
//		Dial: func() (redis.Conn, error) {
//			options := redis.DialDatabase(db)
//			//redis.DialPassword("zsw123")
//			c, err := redis.Dial("tcp", host, options)
//			if err != nil {
//				panic(err.Error())
//			}
//			// 密码验证
//			//if _, err := c.Do("AUTH", "zsw123"); err != nil {
//			//	c.Close()
//			//	return nil, err
//			//}
//			return c, err
//		},
//	}
//}


// 保存数据            dir 组信息 key value
func SaveStringToRedis(dir string, key string,value string)  {

	//data, _ := json.MarshalIndent(player, "", " ")
	//key := "BY_Player_UID_"+ strconv.Itoa(int( player.UserId))

	_, err := RRedis.Do("hdel", dir, key)
	_, err = RRedis.Do("hset", dir, key,value)
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
		fmt.Println("redis 读取出错了:", err)
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


// 获取数据
func DelKeyToRedis(dir string,key string){
	_, err :=  RRedis.Do("hget",dir, key)
	if err !=nil {
		fmt.Println("redis 删除key出错了:", err, key)
	}
}







//
//
//
//
//func dododogo() string {
//	c := RRedis.Get()
//	defer c.Close()
//
//
//	// 有序数组zset
//	ret, err := redis.Strings(c.Do("zrange", "zsw_zset","0","220"))
//	fmt.Println(ret)
//	ret1, err := redis.Strings(c.Do("zadd", "zsw_zset","0.1","ss0.1"))
//	fmt.Println(ret1)
//	ret1, err = redis.Strings(c.Do("zRangeByScore", "zsw_zset","0","220"))
//	fmt.Println(ret1)
//
//
//	// 列表，后面添加的在前面
//	ret, err = redis.Strings(c.Do("lpush", "list_zsw", "newwwwwwwww"))
//	fmt.Println(ret)
//
//	ret, err = redis.Strings(c.Do("lrange", "list_zsw","0","10"))
//	fmt.Println(ret)
//
//
//	// map key value
//	ret, err = redis.Strings(c.Do("hset", "zsw_map", "new","new11"))
//	fmt.Println(ret)
//
//	ret, err = redis.Strings(c.Do("hgetall", "zsw_map"))
//	fmt.Println(ret)
//
//
//
//	// string set get
//	n, err := c.Do("set", "key", "value1")
//	fmt.Println("n:",n)
//	n, err = redis.String(c.Do("get", "key"))
//	fmt.Println("n:",n)
//	fmt.Println("--------------------------------------------------")
//
//
//	// 管道pipline
//	c.Send("SET", "foo", "bar")
//	c.Send("GET", "foo")
//	c.Flush()
//	c.Receive() // reply from SET
//	v, err := redis.String(c.Receive()) // reply from GET
//	fmt.Println("v:",v)
//
//
//
//
//	//存储过程
//	c.Send("MULTI")
//	c.Send("INCR", "foo")
//	c.Send("INCR", "bar")
//	r, err := c.Do("EXEC")
//	fmt.Println(r) // prints [1, 1]
//
//
//
//	if err !=nil{
//		fmt.Println("error:",err)
//		return "failed"
//	}
//
//	return "Success"
//}
//
//
//func main(){
//	InitRedis()
//	strf := dododogo()
//	fmt.Println("",strf)
//}
