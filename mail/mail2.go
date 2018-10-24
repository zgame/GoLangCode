package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	fmt.Println("------------")
	// 邮箱地址
	UserEmail := "196om"
	// 端口号，:25也行
	Mail_Smtp_Port := ":587"
	//邮箱的授权码，去邮箱自己获取
	Mail_Password := "lzbubhdd"
	// 此处填写SMTP服务器
	Mail_Smtp_Host := "smtp.qq.com"
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)
	to := []string{"19m"}
	nickname := "发送人名称"
	user := UserEmail

	subject := "标题"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "邮件内容."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(Mail_Smtp_Host+Mail_Smtp_Port, auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
	fmt.Println("ok")
}
