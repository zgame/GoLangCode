package zip

//----------------------------------------------------------------------------
// 压缩 和 解压缩
//----------------------------------------------------------------------------


import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/golang/snappy"
	"io"
)


// snappy zip
func SnappyZip(strZip string) string {
	in := snappy.Encode(nil, []byte(strZip))
	//fmt.Printf("%x \n", in)
	return string(in)
}

// snappy unzip
func SnappyUnZip(strZip string) string  {
	out,_:= snappy.Decode(nil, []byte(strZip))
	//fmt.Println(string(out))
	return string(out)
}


// 压缩
func ZipWrite(strZip string) string {
	var in bytes.Buffer
	w,err := zlib.NewWriterLevel(&in,zlib.DefaultCompression)
	if err!=nil {
		fmt.Println("err ", err.Error())
	}
	w.Write([]byte(strZip))
	w.Close()

	return in.String()
}
// 解压缩
func ZipRead(strZip string) string{
	var in * bytes.Buffer
	in = bytes.NewBufferString(strZip)

	var out bytes.Buffer
	r, _ := zlib.NewReader(in)
	io.Copy(&out, r)
	//fmt.Println(out.String())
	return out.String()
}

//
//
//func zlib1()  {
//	var in bytes.Buffer
//	b := []byte("史蒂芬森东方闪电饭是钢十多个问个问asjflajdfja;ldfj;lajfd;lajdf;lajdf;lajsdfja;ldjf;lajdeowiueowiueroiwureoiqwuropiwueropiwueopwiueropiwuerwueopueroiwuerowieruoiewruoiewr题我有34537646857dfhgdjgjddfasdasfsaggnndfd/-/-/-/!@#!@%#^$&*%&^")
//	w,err := zlib.NewWriterLevel(&in,zlib.DefaultCompression)
//	if err!=nil {
//		fmt.Println("err ", err.Error())
//	}
//	w.Write(b)
//	w.Close()
//
//	fmt.Println(string(b))
//	fmt.Println(len(b))
//	fmt.Println(in.Bytes())
//	fmt.Println(in.Len())
//	fmt.Println("---------------------------")
//
//
//	var out bytes.Buffer
//	r, _ := zlib.NewReader(&in)
//	io.Copy(&out, r)
//	fmt.Println(out.String())
//}