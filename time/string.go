package main

import (
	"fmt"

)

func main() {
	str := "A0-B1-D4-34-"
	//str1:="00-00"

	fmt.Println("",str)



	i :=40000
	i1 := i/256
	i2 := i%256

	ss1 := fmt.Sprintf("%x",i1)
	if len(ss1) == 1{
		str += "0"
	}
	str += ss1



	//fmt.Println("")
	//fmt.Println("ss1",ss1)
	////str1 = strconv.Itoa(ss1)
	//
	str += "-"
	ss2:= fmt.Sprintf("%x",i2)
	if len(ss2) == 1{
		str += "0"
	}
	str += ss2


	//fmt.Println("")
	//fmt.Println("ss2",ss2)
	////str2 := strconv.Itoa(ss2)
	//
	//fmt.Println("", len(ss1))
	//fmt.Println("", len(ss2))
	//fmt.Println("",str1)
	fmt.Println("",str)

}
