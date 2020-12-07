package log

import (
	"os"
	"log"
	"fmt"
	"time"
	"strconv"
	"runtime"
)

//--------------------------------------------------------------------------------------------------
// 错误日志处理
//--------------------------------------------------------------------------------------------------
var ShowLog = true
var ServerPort int
func CheckError(e error) bool{
	if e!=nil{
		file, _ := os.OpenFile("error.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		logger := log.New(file, "", log.LstdFlags|log.Llongfile)
		logger.Println("...error:...",e.Error())
		if ShowLog{
			fmt.Println("错误："+e.Error())
		}
		return true
	}
	return false
}

// 输出字符串日志，带显示出控制台
func PrintLogger(str string) {
	_logger(str)
	if ShowLog{
		fmt.Println("Log："+str)
	}
}
// 输出字符串日志， 不在控制台显示
func WriteLogger(str string) {
	_logger(str)
}

// 输出字符串日志， 不在控制台显示
func WritefLogger(format string, a ...interface{}) {
	str:= fmt.Sprintf(format,a...)
	_logger(str)
}

// 格式化字符串日志，带显示出控制台
func PrintfLogger(format string, a...interface{})  {
	str:= fmt.Sprintf(format,a...)
	_logger(str)
	if ShowLog{
		fmt.Println("Log："+str)
	}
}

// 内部函数
func _logger(str string)  {
	//go func() {
		// 判断一下目录， 如果没有目录就创建出来
		logDir := "Logs"
		exist, err := PathExists(logDir)
		if err != nil {
			fmt.Printf("get Logs/ dir error![%v]\n", err)
			return
		}
		if !exist {
			err := os.Mkdir(logDir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir %s failed![%v]\n", logDir, err)
			} else {
				fmt.Printf("mkdir %s success!\n", logDir)
			}
		}

		// 根据日期进行分文件
		t1:=time.Now()
		t11 := t1.Format("2006-01-02")

		file, _ := os.OpenFile(logDir+"/Logger_"+t11+"_" + strconv.Itoa(ServerPort) +".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		logger := log.New(file, "", log.LstdFlags)
		logger.Println("[Log:]", str)
	//}()

}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}




func GetSysMemInfo()  string{
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)

	str:= ""
	//str += "   Lookups:" + strconv.Itoa( int(memStat.Lookups))
	//str += "M   TotalAlloc:" + strconv.Itoa( int(memStat.TotalAlloc/1000000))//从服务开始运行至今分配器为分配的堆空间总和
	str += "  Sys:" + strconv.Itoa( int(memStat.Sys/1000000) )+ "M"
	//str += "M   Mallocs:" + strconv.Itoa( int(memStat.Mallocs))//服务malloc的次数
	//str += "次   Frees:" + strconv.Itoa( int(memStat.Frees))//服务回收的heap objects
	str += "   HeapAlloc:" + strconv.Itoa( int(memStat.HeapAlloc/1000000)) + "M"//服务分配的堆内存
	str += "   HeapSys:" + strconv.Itoa( int(memStat.HeapSys/1000000))+ "M"//系统分配的堆内存
	str += "   HeapIdle:" + strconv.Itoa( int(memStat.HeapIdle/1000000))+ "M"//申请但是为分配的堆内存，（或者回收了的堆内存）
	str += "   HeapInuse:" + strconv.Itoa( int(memStat.HeapInuse/1000000))+ "M"//正在使用的堆内存
	str += "   HeapReleased:" + strconv.Itoa( int(memStat.HeapReleased))+ "M"//返回给OS的堆内存，类似C/C++中的free。
	//str += "   HeapObjects:" + strconv.Itoa( int(memStat.HeapObjects))+ "个"//堆内存块申请的量
	str += "   StackInuse:" + strconv.Itoa( int(memStat.StackInuse/1000000)) + "M"//正在使用的栈
	str += "   StackSys:" + strconv.Itoa( int(memStat.StackSys/1000000)) + "M"//系统分配的作为运行栈的内存
	//str += "   NumGC:" + strconv.Itoa( int(memStat.NumGC))+ "次"////垃圾回收的内存大小
	//str += "   NumForcedGC:" + strconv.Itoa( int(memStat.NumForcedGC))
	//str += "   LastGC:" + strconv.Itoa( int(memStat.LastGC))
	return str

}
