package main

import "fmt"

func main() {

	chaa := make(chan int)
	chaa2:= make(chan int)
	//chaa3:= make(chan int)
	go test1(chaa)
	go test2(chaa,chaa2)
	//go test3(chaa3)

	select {
	//case <-chaa3:
	//	fmt.Println("test3 over")
	case <-chaa2 :
		fmt.Println("test2 over")
	}
	fmt.Println("88888888888888888")
	close(chaa2)
	//close(chaa3)

}

func test1(chaa chan int)  {
	for i:=1;i<100;i++{

		fmt.Println("test1----",i)
	}
	chaa <- 1
}

func test2(chaa chan int, chaa2 chan int) {
	//	for {
	select {
	case <-chaa:
		fmt.Println("test1 over")
		close(chaa)
		chaa2 <- 1
		//return
	}
	//fmt.Println("test2 over")
	//	}
}
