package main

import (
	"os"
	"log"
)

func debugUid(uid int)  {
	file, _ := os.OpenFile("uid.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("UID:...",uid)
}
