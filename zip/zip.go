package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//gzip1()
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
	b := []byte("hello")
	w,err := zlib.NewWriterLevel(&in,zlib.DefaultCompression)
	if err!=nil {
		fmt.Println("err ", err.Error())
	}
	w.Write(b)
	w.Close()

	fmt.Println(string(b))
	fmt.Println(len(b))
	fmt.Printf(" %x   \n ",in.Bytes())
	fmt.Println(in.String())
	fmt.Println(in.Len())
	fmt.Println("---------------------------")


	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(out.String())
}

func gzip1()  {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// Setting the Header fields is optional.
	//zw.Name = "a-new-hope.txt"
	//zw.Comment = "an epic space opera by George Lucas"
	//zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("buf: %x   \n ",buf.Bytes())



	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	//println(zr)

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}


}