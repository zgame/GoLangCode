package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

type Header struct {
	var1 string
} 

func main()  {
	keepalive := "keepalive"
	keepalive1 := "coins"
	keepalive2 := "avg_bet"
	keepalive3 := "bc"
	keepalive4 := "time"
	keepalive5 := ":"
	keepalive6 := ","
	strBytes := []byte(keepalive)
	strBytes1 := []byte(keepalive1)
	strBytes2 := []byte(keepalive2)
	strBytes3 := []byte(keepalive3)
	strBytes4 := []byte(keepalive4)
	strBytes5 := []byte(keepalive5)
	strBytes6 := []byte(keepalive6)

	//enc := mahonia.NewDecoder("UTF-8")
	//output := enc.ConvertString(keepalive)
	//
	//fmt.Println("",output)
	//strBytes := []byte(output)
	//
	//
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint16(12))


	// 中间数值编码0d03 0a05 00
	// 中级数值编码03c0 3e88 000a 0700
	//    uint64   8084 1e00 0000 0000
	//    int64    8084 1e00 0000 0000
	//			   0012 7a00 0000 0000


	str := "6b656570616c697665"
	b, _ := hex.DecodeString(str)
	//encodedStr := hex.EncodeToString(b)
	fmt.Printf("@@@@--bytes-->%02x \n",b)
	var hh Header
	buf1:= bytes.NewBuffer(b)
	binary.Read(buf1,binary.LittleEndian, &hh)
	fmt.Println("-------------",hh.var1)
	//fmt.Printf("@@@@--string-->%s \n",encodedStr)



	fmt.Printf("%s   %x \n", keepalive,strBytes)
	fmt.Printf("%s   %x \n", keepalive1,strBytes1)
	fmt.Printf("%s   %x \n", keepalive2,strBytes2)
	fmt.Printf("%s   %x \n", keepalive3,strBytes3)
	fmt.Printf("%s   %x \n", keepalive4,strBytes4)
	fmt.Printf("%s   %x \n", keepalive5,strBytes5)
	fmt.Printf("%s   %x \n", keepalive6,strBytes6)
	fmt.Printf("%x \n", buffer)

	//-------------





}
