package ztimer

import (
	"time"
)

// 计时器，用来定期检查配置的更新，包括后台控制的活动，开关，配置文件更新，用数据版本号来控制
func TimerCheckUpdate(f func(), timer time.Duration)  {
	go func() {
		tickerCheckUpdateData := time.NewTicker(time.Second * timer)
		defer tickerCheckUpdateData.Stop()

		for {
			select {
			case <-tickerCheckUpdateData.C:
				f()
			}
		}
	}()
}


// 夜里12点触发的计时器， 这里启动的时候也是要检查一次的
func TimerClock12(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			//f()
		}
	}()
}

