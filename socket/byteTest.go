package main

import "fmt"

func main()  {
	str := "0welcome!"
	var bb []byte
	bb = []byte(str)

	bb[0] = 0

	fmt.Println("",bb)
	fmt.Printf("%c",bb[1])
	fmt.Println("")
	fmt.Printf("%d",bb)
	fmt.Println("")
	fmt.Printf("%x",bb)
	fmt.Println("")
	fmt.Printf("%s",bb)
}
