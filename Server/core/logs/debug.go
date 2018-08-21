package logs

import (
	"reflect"
	"runtime/debug"
)

// usage: 在goroutine开始时执行 defer PrintPanic()
func PrintPanic() {
	if r := recover(); r != nil {
		Critical("panic:%v", r)
		Critical("%s", debug.Stack())
	}
	//logs.GetLogger().Flush()
}

// usage: GoExec(func, param1, param2, ...)
func GoExec(f interface{}, params ...interface{}) {
	vf := reflect.ValueOf(f)
	vps := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		vps[i] = reflect.ValueOf(params[i])
	}

	go func() {
		defer PrintPanic()
		vf.Call(vps)
	}()
}
