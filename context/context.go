package main

import (
	"fmt"
	_ "net/http/pprof"
	"net/http"
	"sync"
	"time"
	"context"
	"log"
)

func main()  {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:8081",nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		fmt.Println("",i)
		wg.Done()
	}
	//wg.Wait()

	someHandler()
}


func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())				// 用取消函数来通知
	//ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 5) )		// 到时间自动停止
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Printf("done")
			return
		default:
			log.Printf("work")
		}
	}
}