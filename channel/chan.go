package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var num int32
var zcha chan int
var wg sync.WaitGroup

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发

	num =0
	countt := 20
	zcha := make(chan int)

	for i:=0;i<countt;i++{
		wg.Add(1)
		go test(i)
	}


	go func() {
		for i:=0;i<countt;i++{
			//<- zcha
			fmt.Println("---------------",zcha)
			temp := <- zcha
			fmt.Println("zcha:::::", temp)
		}

	}()


	wg.Wait()
	close(zcha)
	fmt.Println("over")

}

func test(index int)  {
	tmp := atomic.LoadInt32(&num)
	num = atomic.AddInt32(&tmp, 1)
	//num++
	fmt.Println("this is ", index, "    num:", num)
	wg.Done()
	zcha <- index

	//atomic.StoreInt32(&d, 666)

}
