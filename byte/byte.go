package main

import "fmt"

func main() {
	var bb []byte
	bb = make([]byte,2)
	bb[0] =1

	bb = append(bb, byte(3))

	fmt.Printf("bb  %v   \n" ,bb)


	str := "21212"
	fmt.Printf("str  %v   \n" ,[]byte(str))
}
