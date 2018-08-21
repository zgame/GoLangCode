package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for{
		select {
		case t:= <-ticker.C:
			fmt.Println("",t)
		}
	}

}
