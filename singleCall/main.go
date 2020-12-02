package main


import (
	"fmt"
	"singleCall/ppp"
	"singleCall/singeCall2"
)


func main()  {

	str := fmt.Sprintf("%.2f",0.01 *2)
	println(str)

	fmt.Println("main.  main")
	singeCall2.Singe2()


	singeCall2.Main()

	//singeCall2
	fmt.Println("Test_v2",singeCall2.Test_v2)



	p := ppp.Ppp{1}
	fmt.Println("",p)


	fmt.Println("",ppp.Zsw_ppp)


	ppp.Show()

	t := Ttt{2}
fmt.Println("",t)
}
