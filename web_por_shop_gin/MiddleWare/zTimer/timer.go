package zTimer

import "time"

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
