package ztimer

//----------------------------------------------------------------------------
// 定时器
//----------------------------------------------------------------------------


import (
	"time"
)

const (
	WarningTimeCost = 200 // 如果时间消耗超过 多少毫秒那么报警
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
			t := time.NewTimer(next.Sub(now)+1)	// 这里加一秒， 或者向上取整，是为了防止临界点问题， 防止偶尔59:59秒因为毫秒问题导致没有到整点
			<-t.C
			f()
		}
	}()
}

// 获取系统时间，精确到毫秒
func GetOsTimeMillisecond()  int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 用来检测运行时间的
func CheckRunTimeCost(f func(), msg string)  {
	startTime := GetOsTimeMillisecond()
	f()
	if GetOsTimeMillisecond()-startTime > WarningTimeCost {
		//zLog.PrintfLogger("----------!!!!!!!!!!!!!!!!!!!!!![ 警告 ]    %s     消耗时间: %d", msg, int(GetOsTimeMillisecond()-startTime))
	}
}
