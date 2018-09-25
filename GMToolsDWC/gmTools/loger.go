package main

import (
	"os"
	"log"
	"runtime"
)

func loger(e string) {
	if e!=""{
		file, _ := os.OpenFile("Log.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		logger := log.New(file, "", log.LstdFlags|log.Llongfile)
		logger.Println(e)
	}
}

func logerDump() {

	file, _ := os.OpenFile("Log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	logger := log.New(file, "", log.LstdFlags|log.Llongfile)

	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	logger.Println(string(buf[:n]))

}

