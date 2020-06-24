package main

import (
	"fmt"
	"runtime"
	"sync"
)


//------------------------------------------------------------------
// 处理新增充值用户的行为记录
//------------------------------------------------------------------

var Group = 10000 // 每次处理人数
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发
	fmt.Println("-----------------start--------------------------")

	for i := 0; i < 27; i++ {
		wg.Add(1)
		go DealUserList(i)
	}
	wg.Wait()

	fmt.Println("-----------------END--------------------------")
	fmt.Println("-----------------END--------------------------")
	fmt.Println("-----------------END--------------------------")
	fmt.Println("-----------------END--------------------------")

	for {
		select {

		}
	}

	fmt.Println(" --------------end-------------- ")

}
