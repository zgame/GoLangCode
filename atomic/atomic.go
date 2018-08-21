package main

import (
	"fmt"
	"runtime"
	"time"
	"sync/atomic"
)

var num int32
func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发
	num =0
	for i:=0;i<100;i++{
		go Test(i)
	}

	for
	{
		select {
		case <- time.After(time.Second * 2):		// 2秒之后退出
		return
		}
		fmt.Printf("-")
	}

}


func Test(index int)  {
	tmp := atomic.LoadInt32(&num)
	num = atomic.AddInt32(&tmp, 1)
	fmt.Println("this is ", index, "    num:", num)


}
