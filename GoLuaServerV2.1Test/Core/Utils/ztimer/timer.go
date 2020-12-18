package ztimer

import (
	"time"
)


// 计时器，毫秒计时
func TimerMillisecondCheckUpdate(f func(), timer time.Duration)  {
	go func() {
		tickerCheckUpdateData := time.NewTicker(time.Millisecond * timer)
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
func TimerClock0(f func()) {
	TimerClock(f,0)
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
			next = time.Date(next.Year(), next.Month(), next.Day(), clock, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}

func GetOsTimeMillisecond()  int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
}