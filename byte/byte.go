package main

import (
	"fmt"
	"math"
	"encoding/binary"
)

func main() {
	var bb []byte
	bb = make([]byte,2)
	bb[0] =1

	bb = append(bb, byte(3))
	// byte数组可以append到末尾
	fmt.Printf("byte slice append test  %v   \n" ,bb)


	str := "21212"
	// string 跟[]byte之间的转换
	fmt.Printf("string to []byte  %v   \n" ,[]byte(str))



	// float 转 []byte
	fmt.Printf("float to []byte  %v\n" ,Float32ToByte(1))


	// uint32 转 []byte
	fmt.Printf("uint32 to []byte %v\n", Int32ToBytes(6565))


	var ii int32
	ii = -1
	i2:=uint32(ii)
	i3:=int32(i2)
	fmt.Println("",ii)
	fmt.Println("",i2)
	fmt.Println("",i3)


}


func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}


func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

