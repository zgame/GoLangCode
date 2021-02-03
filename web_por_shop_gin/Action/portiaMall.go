package Action

import (
	"fmt"
	"time"
	"web_gin/MiddleWare/zLog"
	"web_gin/MySql"
)

// 处理一下商城的过期时间
func ResetMallTime(){
	mallList := MySql.GetMallInfoData()
	for _,element := range mallList{
		if element.Endtime == "" || element.Endtime == "-1"{
			continue
		}
		endTime,err := time.ParseInLocation("2006-01-02 15:04:05", element.Endtime, time.Local)
		if err!= nil {
			zLog.PrintLogger("商城过期时间有问题  "+err.Error() + "   " + element.Endtime)
			return
		}

		if time.Now().After(endTime){
			fmt.Println("过期了， 那么处理一下")
			element.Starttime = "-1"
			element.Discountprice = -1
			element.Endtime = "-1"
			MySql.UpdateMallInfoData(&element,element.Id)
		}
	}
}
