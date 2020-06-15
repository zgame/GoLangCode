package main

import (
	"fmt"
	"time"
)

type UserList struct {
	id        int
	uid       int		// 充值用户uid
	initDate  string
	lastDate  string
	days      int
	uid2      int		// 免费用户uid
	initDate2 string
	lastDate2 string
	days2     int
	matchType int
	dayNum    int
}


// 获取玩家持续时间的日期列表
func getTimeList(start string, days int ) []string  {
	result := make([]string,0)
	//fmt.Println(" start " ,start)
	startTime,err := time.ParseInLocation("2006-01-02T00:00:00Z", start, time.Local)
	if err!=nil {
		fmt.Println("",err.Error())
	}
	//fmt.Println("startTime : ",startTime)
	for i:=0;i<days;i++{
		time := startTime.AddDate(0,0,i)
		timeString := time.Format("2006-01-02")
		//fmt.Println("", timeString )
		result = append(result, timeString)
	}
	return result
}
