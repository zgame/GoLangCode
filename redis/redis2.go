package main

import "fmt"
import (
	"github.com/garyburd/redigo/redis"
	"reflect"
	"time"
)

var RRedis redis.Conn

func Init() {
	re, err := redis.Dial("tcp", "172.16.140.123:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	if _, err := re.Do("AUTH", "Soonyo123"); err != nil {
		re.Close()
		fmt.Println("Redis 服务器密码不正确")
		return
	}
	RRedis = re
	//defer RRedis.Close()
}

func ggo()  {
	//res,err := RRedis.Do("hset","key","field","value")  //写
	//result,err := redis.Values(RRedis.Do("hgetall","key"))//读

	//res, err := RRedis.Do("hget", "ALL_Players", "BY_Player_UID_2027445")github.com/tengattack/gluasql  BY_AllPlayers_OpenId_Uid:74-D4-36-AD-09-c2
	//res, err := RRedis.Do("hset", " zzsw:123", "123", 111)
	//res, err := RRedis.Do("hget", " zzsw:123","123")
	//res, err := RRedis.Do("hget", "BY_AllPlayers_OpenId_Uid:74-D4-36-AD-00-01", "74-D4-36-AD-00-01")
	res, err := RRedis.Do("hget", "AllPlayers_UUID:BY_UUID", "BY_UUID")
	fmt.Println(reflect.TypeOf(res))
	if err != nil {
		fmt.Println("hget failed", err.Error())
	} else {
		if res!=nil {
			fmt.Printf("hget value :%s\n", res.([]byte))
		}else {
			fmt.Println("没有该记录")
		}

	}

	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	//RRedis.Close()
}

func main() {
	Init()

	for i:=0;i<10;i++{
		go func() {
			AddNumberToRedis("AllPlayers_UUID:BY_UUID","BY_UUID",1)
			AddNumberToRedis("AllPlayers_UUID:BY_UUID","BY_UUID",1)
			AddNumberToRedis("AllPlayers_UUID:BY_UUID","BY_UUID",1)
			ggo()
			AddNumberToRedis("AllPlayers_UUID:BY_UUID","BY_UUID",1)
			AddNumberToRedis("AllPlayers_UUID:BY_UUID","BY_UUID",1)
			ggo()

		}()
	}

	for  {
		time.Sleep(time.Second)
	}

}



// 执行脚本，用于分布式，不用加锁的情况，因为脚本一次性执行的，类似存储过程
var AddScript1 = redis.NewScript(2,`
   local r = redis.call('hget',KEYS[1],KEYS[2])
   if r ~= nil then
		r = r + ARGV[1]
       redis.call('hset', KEYS[1],KEYS[2], r)
   end
   return r
`)

// 这里的参数num， 记住： num只是增量， 你只能要求人家增加多少， 具体完事之后是多少，会返回给你， 因为这个涉及到分布式多请求
func AddNumberToRedis(dir string,key string, num int) int{

	v, err := AddScript1.Do(RRedis, dir,key, num)
	if err != nil {
		fmt.Println("AddNumberToRedis Error: " + err.Error())
		return 0
	}

	//fmt.Println("AddNumberToRedis",v , reflect.TypeOf(v))
	re := int(v.(int64))
	return re
}

