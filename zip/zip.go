package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func main() {
	zlib1()
	//writer()

}

func writer()  {
	var in bytes.Buffer
	b := []byte("史蒂芬森东方闪电饭是钢十多个问个问题我有34537646857dfhgdjgjddfasdasfsaggnndfd/-/-/-/!@#!@%#^$&*%&^")
	w := zlib.NewWriter(&in)
	w.Write(b)
	w.Close()

	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(out.String())
}


func zlib1()  {
	var in bytes.Buffer
	b := []byte("史蒂芬森东方闪电饭是钢十多个问个问asjflajdfja;ldfj;lajfd;lajdf;lajdf;lajsdfja;ldjf;lajdeowiueowiueroiwureoiqwuropiwueropiwueopwiueropiwuerwueopueroiwuerowieruoiewruoiewr题我有34537646857dfhgdjgjddfasdasfsaggnndfd/-/-/-/!@#!@%#^$&*%&^")
	w,err := zlib.NewWriterLevel(&in,zlib.DefaultCompression)
	if err!=nil {
		fmt.Println("err ", err.Error())
	}
	w.Write(b)
	w.Close()

	fmt.Println(string(b))
	fmt.Println(len(b))
	fmt.Println(in.Bytes())
	fmt.Println(in.Len())
	fmt.Println("---------------------------")


	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(out.String())
}