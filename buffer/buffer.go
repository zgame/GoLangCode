package main

import "fmt"

func main()  {
	str1 := "0123456"
	str2 := "789012345"
	str3 := "67890123456"

	b1 := []byte(str1)
	b2 := []byte(str2)
	b3 := []byte(str3)

	bufferEnd := make([]byte,len(b1)+len(b2)+len(b3))
	copy(bufferEnd, b1)
	copy(bufferEnd[len(b1):len(b1)+len(b2)], b2)
	copy(bufferEnd[len(b1)+len(b2):], b3)

	fmt.Println(string(bufferEnd))
}
