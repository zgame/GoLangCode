package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"			// 注意这里
	"time"
	"fmt"
)

func Counter() {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 200)
		counter++
	}
	//wg.Done()
}

func main() {
	flag.Parse()

	//远程获取pprof数据打开浏览器http://localhost:8080/debug/pprof/
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	//var wg sync.WaitGroup
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//	go Counter()
	//}
	//wg.Wait()

	//// sleep 10 mins, 在程序退出之前可以查看性能参数.
	//time.Sleep(60 * time.Second)

	for  {
		fmt.Println("打开浏览器http://localhost:8080/debug/pprof/")
		time.Sleep(time.Second *2)
	}
}