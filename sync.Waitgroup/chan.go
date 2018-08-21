package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"sync"
)

var num int32
var wg sync.WaitGroup
func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发
	num =0


	for i:=0;i<100;i++{
		wg.Add(1)
		go test(i)
	}

	wg.Wait()

	

}

func test(index int)  {
	//tmp := atomic.LoadInt32(&num)
	//num = atomic.AddInt32(&tmp, 1)
	num ++
	fmt.Println("this is ", index, "    num:", num)
	wg.Done()

	//atomic.StoreInt32(&d, 666)

}
