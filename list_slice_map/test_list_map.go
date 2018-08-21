package main

import "fmt"

func main(){

	fmt.Println("-------------------------------数组---------------------")
	var list1 =[6]int{1,2,3,4,5,6}
	fmt.Printf("%v",list1)
	fmt.Println("")

	var list2 [6]int
	list2[5] =1
	list2 = [6]int{1,2,3,4,5}
	fmt.Printf("%v",list2)
	fmt.Println("")

	var aa = [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /*  第三行索引为 2 */
	}
	fmt.Printf("%v",aa)
	fmt.Println("")

	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}	//[]或者[...]都可以
	fmt.Printf("%v",balance)
	fmt.Println("")


	bb := []int{2,3,4}
	fmt.Printf("%v",bb)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----------------------slice------------------------------")




	//var slice1 []int
	slice1 := make([]int,1)
	slice1 = append(slice1, 1)
	slice1 =append(slice1, 2)
	slice1 =append(slice1, 3)
	slice1 =append(slice1, 4)
	slice1 =append(slice1, 5)
	fmt.Printf("%v",slice1)
	fmt.Println("")

	i:= len(slice1)-1
	slice1 = append(slice1[:i],slice1[i+1:]...)




	fmt.Printf("%v",slice1)
	fmt.Println("")


	slice1 = make([]int,0)
	slice1 = append(slice1, 1)
	fmt.Printf("%v",slice1)
	fmt.Println("")


	fmt.Println("---------------------------map -------------------------")
	map1 := make(map[string]int)
	map1["1"] = 1
	map1["22"] = 22
	map1["33"] = 33


	fmt.Printf("%v",map1)
	fmt.Println("")
	fmt.Println("1",map1["1"])
	fmt.Println("1",map1["2"])


	delete(map1,"1")

	fmt.Printf("%v",map1)
	fmt.Println("")

	map1 = make(map[string]int)

	map1["22"] = 22
	fmt.Printf("%v",map1)
	fmt.Println("")


}