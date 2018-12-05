package main

import (
	"time"
	"fmt"
)

func main()  {
	t1 := time.Now()	//2018-04-22 14:27:28.7273956 +0800 CST m=+0.005000301
	fmt.Println(t1)

	t2 := t1.Unix()	//1524378448
	fmt.Println(t2)

	t3 := t1.UnixNano()	//1524378448727395600
	fmt.Println(t3)

	t4 := t1.String()	//2018-04-22 14:27:28.7273956 +0800 CST m=+0.005000301
	fmt.Println(t4)

	t5 := time.Millisecond		//1ms
	fmt.Println(t5)

	t6:=time.Now().UnixNano() / int64(time.Millisecond)	//1524378448805
	fmt.Println(t6)

	t7 := t1.Year()		//2018
	fmt.Println(t7)

	t8 := t1.Format("2006-01-02 15:04:05")		//2018-04-22 14:32:05
	fmt.Println(t8)



	t9 := time.Date(2017,2,4,5,7,8,0,time.Local)	//2017-02-04 05:07:08 +0800 CST
	fmt.Println(t9)



	t10 := time.Now().Add(time.Second)
	if time.Now().After(t10){
		 fmt.Println("到时间了")
	}else{
		fmt.Println("没到时间")
	}
	fmt.Println(t10)

	t11 := t1.Format("2006-01-02")		//2018-04-22 14:32:05
	fmt.Println(t11)


	t12,_,_ := t1.Clock()   //hour, min, sec
	fmt.Println(t12)


	for {
		tt := time.Now().UnixNano() / int64(time.Millisecond)
		fmt.Println("", tt)
		time.Sleep(time.Millisecond * 100)
	}




}
