package zip

//----------------------------------------------------------------------------
// 压缩 和 解压缩
//----------------------------------------------------------------------------


import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/golang/snappy"
	lua "github.com/yuin/gopher-lua"
	"io"
)

var exports = map[string]lua.LGFunction{
	"encode": snappyZip,   //非
	"decode": snappyUnZip, //或
}

// ----------------------------------------------------------------------------

func zipLoader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

// ----------------------------------------------------------------------------

func LuaZipLoad(L *lua.LState) {
	L.PreloadModule("zip", zipLoader)
}


// snappy zip
func snappyZip(L *lua.LState) int {
	src := L.CheckString(1)
	in := snappy.Encode(nil, []byte(src))
	L.Push(lua.LString(string(in)))

	return 1
}

// snappy unzip
func snappyUnZip(L *lua.LState) int {
	src := L.CheckString(1)
	out,_:= snappy.Decode(nil, []byte(src))
	//fmt.Println(string(out))
	L.Push(lua.LString(string(out)))

	return 1
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
