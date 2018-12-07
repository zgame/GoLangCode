package main

import "fmt"
import (
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strconv"
	"time"
	"sync"
)


func go2( RRedis redis.Conn)  {
	//res,err := RRedis.Do("hset","key","field","value")  //写
	//result,err := redis.Values(RRedis.Do("hgetall","key"))//读

	//res, err := RRedis.Do("hget", "ALL_Players", "BY_Player_UID_2027445")
	res, err := RRedis.Do("hset", "ALL_Players", "BY_Player_UID_2027445","ddddd")
	fmt.Println(reflect.TypeOf(res))
	if err != nil {
		fmt.Println("hget failed", err.Error())
	} else {
		//fmt.Printf("hget value :%s\n", res.([]byte))
		//fmt.Printf("hget value :%s\n", res.(string))
	}

	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
}

func zpop(c redis.Conn) (result int, err error) {
	dir := "ALL_Players"
	key:= "BY_Player_UID_2027445"

	defer func() {
		// Return connection to normal state on error.
		if err != nil {
			c.Do("DISCARD")
			fmt.Println("-----------DISCARD--------------")
		}
	}()

	// Loop until transaction is successful.
	for {
		if _, err := c.Do("WATCH",dir, key); err != nil {
			fmt.Println("-----------WATCH error--------------")
			return 0, err
		}

		//members, err := redis.Strings(c.Do("hget", dir,key))
		members, err := c.Do("hget", dir,key)
		if err != nil {
			fmt.Println("----------hget error--------------")
			return 0, err
		}
		//if len(members) != 1 {
		//	fmt.Println("-----------redis.ErrNil error--------------")
		//	return "", redis.ErrNil
		//}
		//fmt.Println("members",members)
		fmt.Printf("hget value :%s\n", members.([]byte))
		num,_ := strconv.Atoi(string(members.([]byte)))
		num ++

		c.Send("MULTI")
		c.Send("hset",dir, key, num)
		queued, err := c.Do("EXEC")
		if err != nil {
			fmt.Println("-----------EXEC error--------------")
			return 0, err
		}

		if queued != nil {
			result = num
			break
		}
	}

	return result, nil
}


func main() {
	re, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	//RRedis := re
	//go2(re)
	//dolua(re)

	for i:=0;i<100;i++{
		//go zpop(re)
		go dolua(re)
	}

	for{
		time.Sleep(time.Second * 1)
		//fmt.Println("---------------------------")
	}

}


var AddScript = redis.NewScript(2,`
   local r = redis.call('hget',KEYS[1],KEYS[2])
   if r ~= nil then
		r = r + ARGV[1]
       redis.call('hset', KEYS[1],KEYS[2], r)
   end
   return r
`)

//var AddScript = redis.NewScript(1, `
//    local r = redis.call('ZRANGE', KEYS[1], 0, 0)
//    if r ~= nil then
//        r = r[1]
//        redis.call('ZREM', KEYS[1], r)
//    end
//    return r
//`)

var Mutex sync.Mutex

func dolua(c redis.Conn) {
	Mutex.Lock()
	//v, err := redis.String(AddScript.Do(c, "zsw_zset",""))
	v, err := AddScript.Do(c, "ALL_Players","BY_Player_UID_2027445", 1)
	if err != nil {
		fmt.Println("dolua  error " , err.Error())
		return
	}
	Mutex.Unlock()
	fmt.Println("",v)
}
