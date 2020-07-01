package main

import (

	"fmt"
	"runtime"
	"sync"
)


//------------------------------------------------------------------
// 处理新增充值用户的行为记录
//------------------------------------------------------------------

//var Group = 1000 // 每次处理人数
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发
	fmt.Println("-----------------start--------------------------")
	fmt.Println("-----------------插入充值记录--------------------------")

	//TestDB = mssql.ConnectDB(userId, password, server, TestDBName)

	for i := 0; i < 32; i++ {
		wg.Add(1)
		go DealUserList(i )
	}
	wg.Wait()

	fmt.Println("-----------------插入充值记录--------------------------")
	fmt.Println("-----------------END--------------------------")
	fmt.Println("-----------------END--------------------------")
	fmt.Println("-----------------END--------------------------")


	for {
		select {

		}
	}

	//fmt.Println(" --------------end-------------- ")

}
