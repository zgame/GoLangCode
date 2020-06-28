package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	//TimerClock0(func() {
	//	fmt.Println("ddddddddddd")
	//})
	TimerClock2(func() {
		fmt.Println("",time.Now())
	})

	for{
		select {
		//case t:= <-ticker.C:
		//	fmt.Println("",t)
		}
	}

}

// 到时间触发的计时器
func TimerClock2(f func()) {
	go func() {
		for {
			var next time.Time
			now := time.Now()
			clockNow,_,_ := now.Clock()		//现在几点了
			mini:= now.Minute()
			//// 计算下一个时间点
			//if clockNow >= clock{
			//	// 现在的时间已经过了，那么就等明天吧
			//	next = now.Add(time.Hour * 24)
			//}else {
				next = now
			//}
			next = time.Date(next.Year(), next.Month(), next.Day(), clockNow, mini+1,0 , 0, next.Location())
			fmt.Println(" next ",next)
			fmt.Println("sub",next.Sub(now))
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}

func TimerClock0(f func()) {
	TimerClock(f,19)
}

// 到时间触发的计时器
func TimerClock(f func(),clock int) {
	go func() {
		for {
			var next time.Time
			now := time.Now()
			clockNow,_,_ := now.Clock()		//现在几点了
			// 计算下一个时间点
			if clockNow >= clock{
				// 现在的时间已经过了，那么就等明天吧
				next = now.Add(time.Hour * 24)
			}else {
				next = now
			}
			next = time.Date(next.Year(), next.Month(), next.Day(), clock, 29, 0, 0, next.Location())
			fmt.Println(" next ",next)
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}