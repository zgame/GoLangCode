package main

import (
	"fmt"
	"runtime"
	"sync"
)

var RecordTimeDict = []string{
	"GameCoinChangeRecord_",
	"GameDiamondChangeRecord_",
	"GameItemChangeRecord_",
	"GameLotteryChangeRecord_",
	"GameScoreChangeRecord_",
	"HDBZExchangeInfo_",
	"HunGameChipRecord_"}

var Group = 1 // 每次处理人数
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置cpu的核的数量，从而实现高并发
	fmt.Println("-----------------start--------------------------")

	for i := 0; i < 1; i++ {
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
