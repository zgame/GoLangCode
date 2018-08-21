package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

var pool *redis.Pool
func init1() {
	pool = newPool("127.0.0.1:6379", 0)
}

func newPool(host string, db int) *redis.Pool {
	return &redis.Pool {
		MaxIdle: 50,
		MaxActive: 100,
		Dial: func() (redis.Conn, error) {
			options := redis.DialDatabase(db)
			//redis.DialPassword("zsw123")
			c, err := redis.Dial("tcp", host, options)
			if err != nil {
				panic(err.Error())
			}
			// 密码验证
			//if _, err := c.Do("AUTH", "zsw123"); err != nil {
			//	c.Close()
			//	return nil, err
			//}
			return c, err
		},
	}
}



func dododogo() string {
	c := pool.Get()
	defer c.Close()


	// 有序数组zset
	ret, err := redis.Strings(c.Do("zrange", "zsw_zset","0","220"))
	fmt.Println(ret)
	ret1, err := redis.Strings(c.Do("zadd", "zsw_zset","0.1","ss0.1"))
	fmt.Println(ret1)
	ret1, err = redis.Strings(c.Do("zRangeByScore", "zsw_zset","0","220"))
	fmt.Println(ret1)


	// 列表，后面添加的在前面
	ret, err = redis.Strings(c.Do("lpush", "list_zsw", "newwwwwwwww"))
	fmt.Println(ret)

	ret, err = redis.Strings(c.Do("lrange", "list_zsw","0","10"))
	fmt.Println(ret)


	// map key value
	ret, err = redis.Strings(c.Do("hset", "zsw_map", "new","new11"))
	fmt.Println(ret)

	ret, err = redis.Strings(c.Do("hgetall", "zsw_map"))
	fmt.Println(ret)

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")
	ret, err = redis.Strings(c.Do("hget", "zsw_map", "new"))
	fmt.Println(ret)




	// string set get
	n, err := c.Do("set", "key", "value1")
	fmt.Println("n:",n)
	n, err = redis.String(c.Do("get", "key"))
	fmt.Println("n:",n)
	fmt.Println("--------------------------------------------------")


	// 管道pipline
	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	c.Receive() // reply from SET
	v, err := redis.String(c.Receive()) // reply from GET
	fmt.Println("v:",v)




	//存储过程
	c.Send("MULTI")
	c.Send("INCR", "foo")
	c.Send("INCR", "bar")
	r, err := c.Do("EXEC")
	fmt.Println(r) // prints [1, 1]



	if err !=nil{
		fmt.Println("error:",err)
		return "failed"
	}

	return "Success"
}



func main(){
	init1()
	strf := dododogo()
	fmt.Println("",strf)
}
