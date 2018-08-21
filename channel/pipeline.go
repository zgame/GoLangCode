package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 1; x<100 ; x++ {
			naturals <- x
			//fmt.Println("ddddddddddddd",x)
		}
		defer close(naturals)
	}()
	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
		defer close(squares)
	}()
	// Printer (in main goroutine)
	for x := range squares {
		if x ==0 {
			break
		}
		fmt.Println(x)
	}
}
