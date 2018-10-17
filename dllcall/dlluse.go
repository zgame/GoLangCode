package main

import (
	"fmt"
	"unsafe"
	"syscall"
)

//-------------------------------------------------------------------
// 在dll调用传递参数的时候，需要进行指针类型的转换
func IntPtr(n int) uintptr {
	return uintptr(n)
}
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}


//-------------------------------------------------------------------

func Lib_add() {
	lib := syscall.NewLazyDLL("exportgo.dll")//exportgo  libcppmakedll
	add := lib.NewProc("hello")
	ret, _, err := add.Call( StrPtr("ssssssss"))
	if err != nil {
		fmt.Println("lib.dll运算结果为:", ret)
	}
}

func DllTestDef_add() {
	DllTestDef, _ := syscall.LoadLibrary("exportgo.dll")
	defer syscall.FreeLibrary(DllTestDef)
	add, err := syscall.GetProcAddress(DllTestDef, "Sum")
	ret, _, err := syscall.Syscall(add,
		2,
		IntPtr(1),
		IntPtr(27),
		0)
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}

}

func DllTestDef_add2() {
	DllTestDef := syscall.MustLoadDLL("exportgo.dll")
	add := DllTestDef.MustFindProc("Hello")
	ret, _, err := add.Call()
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}
}







//-------------------------------------------------------------------
func main() {
	fmt.Println("start ")

	//Lib_add()
	DllTestDef_add()
	//DllTestDef_add2()

}
