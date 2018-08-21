package main

import (
	"time"
	"fmt"
)

func main() {
	for {
		select {
		case <-time.After(time.Second ):
			fmt.Println("sdfsdfsdfsdf")
		}
	}
	fmt.Println(" over")
}
