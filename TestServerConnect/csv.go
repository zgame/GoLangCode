package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
)

func getGoldFishMap()  {
	GoldFishMap = make(map[int]int ,0)

	f, err := os.Open("mgby_fish.csv")
	if err != nil {
		fmt.Println("读取csv错误")
	}
	rows, err := csv.NewReader(f).ReadAll()		// rows是每一行
	f.Close()
	if err != nil {
		fmt.Println("读取行错误")
	}
	for _,v := range rows{
		fish_type,_ := strconv.Atoi(v[1])
		if fish_type == 3 || fish_type == 16  {
			fish_id,_  := strconv.Atoi(v[0])
			GoldFishMap[fish_id] = fish_type
		}
	}

	//for k,v :=range GoldFishMap{
	//	fmt.Println("k:",k,"  v:",v)
	//}
}