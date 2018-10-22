package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

//-------------------------------------------------------------------
// 在dll调用传递参数的时候，需要进行指针类型的转换
func IntPtr(n int) uintptr {
	return uintptr(n)
}
//func StrPtr(s string) uintptr {
//
//	ss,_:=syscall.UTF16PtrFromString(s)
//	p:=unsafe.Pointer(ss)
//	return uintptr(p)
//}

func StrPtr(s string) uintptr {
	//ss:=syscall.StringToUTF16Ptr(s)
	//ss:="sdfsdfsdf"
	p:=unsafe.Pointer(&s)
	return uintptr(p)
}



//-------------------------------------------------------------------

func Lib_add() {
	lib := syscall.NewLazyDLL("exportgo.dll")//exportgo  libcppmakedll  libcmakedll
	add := lib.NewProc("hello")

	_, _, err := add.Call(2,StrPtr("sdf"))
	if err != nil {
		//fmt.Println("lib.dll运算结果为:", ret)
	}
}

func DllTestDef_add() {
	DllTestDef, _ := syscall.LoadLibrary("exportgo.dll")
	defer syscall.FreeLibrary(DllTestDef)
	add, err := syscall.GetProcAddress(DllTestDef, "Sum")
	ret, _, err := syscall.Syscall(add,
		3,
		4,
		6,
		StrPtr("sum"))
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}

}

func DllTestDef_add2() {
	DllTestDef := syscall.MustLoadDLL("exportgo.dll")
	add := DllTestDef.MustFindProc("Hello")
	ret, _, err := add.Call(StrPtr("sdf"))
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}
}







//-------------------------------------------------------------------
func main() {
	fmt.Println("start ")

	Lib_add()
	//DllTestDef_add()
	//DllTestDef_add2()

}
