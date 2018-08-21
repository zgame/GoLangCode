package main

import "fmt"

func main() {

	chaa := make(chan int)
	chaa2 := make(chan int)

	go test11(chaa)
	go test22(chaa, chaa2)

	select {

	case <-chaa2:
		fmt.Println("test2 over")
	}
	fmt.Println("88888888888888888")
	close(chaa)
	close(chaa2)

}

func test11(chaa chan int) {
	for i := 1; i < 100; i++ {

		fmt.Println("test1----", i)
	}
	chaa <- 1
}

func test22(chaa chan int, chaa2 chan int) {
	for {
		select {
		case <-chaa:
			fmt.Println("test1 over")

			chaa2 <- 1
			//return
			//default:							//打开default就不能阻塞了
		}
		fmt.Println("test22")
	}
}
