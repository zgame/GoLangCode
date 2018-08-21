package main

import (
	"fmt"
	"os"
	"log"
)

func main()  {
	fmt.Println("begin TestLog ...")
	file, err := os.OpenFile("test.log",os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}

	logger := log.New(file, "", log.LstdFlags|log.Llongfile)
	log.Println("1.Println log with log.LstdFlags ...")
	logger.Println("1.Println log with log.LstdFlags ...")

	logger.SetFlags(log.LstdFlags)

	log.Println("2.Println log without log.LstdFlags ...")
	logger.Println("2.Println log without log.LstdFlags ...")

	//log.Panicln("3.std Panicln log without log.LstdFlags ...")
	//fmt.Println("3 Will this statement be execute ?")
	//logger.Panicln("3.Panicln log without log.LstdFlags ...")

	log.Println("4.Println log without log.LstdFlags ...")
	logger.Println("4.Println log without log.LstdFlags ...")

	log.Fatal("5.std Fatal log without log.LstdFlags ...")
	fmt.Println("5 Will this statement be execute ?")
	logger.Fatal("5.Fatal log without log.LstdFlags ...")
}
