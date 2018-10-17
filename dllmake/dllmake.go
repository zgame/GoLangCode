package main

import "fmt"
//import "C"

//export hello
func hello(s string) {
	fmt.Println("",s)
	fmt.Println("hello From DLL: Bye!")
}
//export Hello
func Hello() {
	fmt.Println("Hello From DLL: Bye!")
}

//export Sum
func Sum(a int, b int) int {
	return a + b
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
