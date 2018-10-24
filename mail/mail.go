package main

import (
	"gopkg.in/gomail.v2"
	"fmt"
)

var mFrom  = "1966q.com"
var mTo  = "zm"
var mCc  = "z.com"
var mTitle  = "Hello15!"
var mText  = "Hello <b>Bob</b> and <i>Cora</i>!"


func main()  {
	fmt.Println("-----------")
	m := gomail.NewMessage()
	m.SetHeader("From", mFrom)
	m.SetHeader("To", mTo, mCc)
	m.SetHeader("Cc", mTo, mCc)
	//m.SetAddressHeader("Cc", mCc, "Dan")
	//m.SetAddressHeader("Cc", mTo, "Dan")
	m.SetHeader("Subject", mTitle)
	m.SetBody("text/html", mText)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 465, "196.com", "lzbhdd")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("send ok!")
}
