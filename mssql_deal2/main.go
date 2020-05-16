package main

import (
	"fmt"
	"runtime"
	"sync"
)


var Group = 1000 // 每次处理人数
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发
	fmt.Println("-----------------start--------------------------")

	for i := 0; i < 45; i++ {
		wg.Add(1)
		go DealUserList(i * Group)
	}
	wg.Wait()

	for {
		select {

		}
	}

	fmt.Println(" --------------end-------------- ")

}
