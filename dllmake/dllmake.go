package main

//import "C"

//export hello
func hello(a int,s string) int{
	println("hello:",s,a)
	return 0
	
}
//export Hello
func Hello(s string)int {
	println("Hello",s)
	return 1
}

//export Sum
func Sum(a int, b int, s string) int {
	println(s)
	return a + b
}

func main() {

	// Need a main function to make CGO compile package as C shared library
}
