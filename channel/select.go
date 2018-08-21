package main

import "fmt"

func main() {

	chaa := make(chan int)
	tt := make(chan int)
	go test1(chaa)

	go test2(chaa,tt)

	select {
	case <-tt:
		fmt.Println("ddddddddddtttttttt")
	}
	fmt.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	close(tt)

}

func test1(chaa chan int)  {
	for i:=1;i<100;i++{

		fmt.Println("test----",i)
	}
	chaa <- 1
}

func test2(chaa chan int,tt chan int)  {
	for {
		select {
		case <-chaa:
			fmt.Println("over")
			close(chaa)
			tt<-1
			return
		default:

		}
		fmt.Printf("run")


	}
}