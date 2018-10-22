package main

import (
	"github.com/yuin/gopher-lua"
	"fmt"
)
var fileMethods = map[string]lua.LGFunction{
	"zprint": zprint,
	"zprint2": zprint2,

}


func main() {

	L := lua.NewState()
	defer L.Close()


	//mod := L.RegisterModule("zmodule", map[string]lua.LGFunction{}).(*lua.LTable)
	mt:= L.NewTypeMetatable("zsw")
	L.SetGlobal("ZZ", mt)							// 设定全局mudule
	L.SetField(mt, "new", L.NewFunction(zprint))		// 绑定new函数



	//mt.RawSetString("__index",mt)						// 设定__index
	L.SetFuncs(mt, fileMethods)								// 设定metaTable的函数列表





	//// 执行lua文件
	if err := L.DoFile("test_meta_table/testm.lua"); err != nil {
		fmt.Println("加载main.lua文件出错了！")
		fmt.Println(err.Error())
	}





}

func zprint(L *lua.LState)  int {
	println("zzzzzzzzzzzzzzz")
	return 1
}
func zprint2(L *lua.LState)  int {
	println("2222222")
	return 1
}