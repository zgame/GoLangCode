package zRandom

import (
	"math/rand"
	"time"
	"fmt"
)

// 随机 [0,num)
func ZRandom(num int) int {
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	x := rand.Intn(num)
	return x
}

// 随机[ min, max)
func ZRandomTo(min int, max int) int {
	if min >= max || max == 0 {
		//fmt.Println("随机数格式不正确")
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// 随机[ min, max) float32
func ZRandomFloatTo(min float32, max float32) float32 {
	if min >= max || max <= 0 {
		fmt.Println("随机数格式不正确")
		return max
	}
	rand.Seed(time.Now().UnixNano())

	ma:= int(max*1000)
	mi:= int(min*1000)
	re := rand.Intn(ma-mi) + mi
	return float32(re/1000)
}

// 获取百分比方法， 比如10几率， 那么小于等于10，返回true
func ZRandomPercentRate(rate int) bool {
	rr := ZRandomTo(1,101)
	if rr <= rate{
		return true
	}else{
		return false
	}
}
