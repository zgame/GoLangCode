package main

import "fmt"
import (
	"github.com/garyburd/redigo/redis"
	"reflect"
)

var RRedis redis.Conn

func Init() {
	re, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	RRedis = re
	//defer RRedis.Close()
}

func ggo()  {
	//res,err := RRedis.Do("hset","key","field","value")  //写
	//result,err := redis.Values(RRedis.Do("hgetall","key"))//读

	res, err := RRedis.Do("hget", "ALL_Players", "BY_Player_UID_2027445")
	fmt.Println(reflect.TypeOf(res))
	if err != nil {
		fmt.Println("hget failed", err.Error())
	} else {
		fmt.Printf("hget value :%s\n", res.([]byte))
	}

	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
}

func main() {
	Init()
	ggo()
}