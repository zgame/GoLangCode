package main

import (
	"gopkg.in/gomail.v2"
	"fmt"
	"github.com/go-ini/ini"
	log "github.com/jeanphorn/log4go"
)

//var mFrom  = "zhushw@soonyo.com"
//var mTo  = "zhushw@soonyo.com"
//var mCc  = "19665944@qq.com"
//var mTitle  = "GameServer Error"
//var mText  = "Hello <b>Bob</b> and <i>Cora</i>!"


func sendMail(msg string)  {
	// -----------读取配置文件-----------
	f, err := ini.Load("Setting.ini")
	if err != nil {
		fmt.Println("ini配置文件出错！", err)
		return
	}

	host := f.Section("Mail").Key("host").Value()
	port,_ := f.Section("Mail").Key("port").Int()
	username := f.Section("Mail").Key("username").Value()
	pwd := f.Section("Mail").Key("pwd").Value()

	mFrom := f.Section("Mail").Key("mFrom").Value()
	mTo := make([]string,0)
	for i:=1;i<=8;i++{
		str := fmt.Sprintf("mTo%d",i)
		mTot := f.Section("Mail").Key(str).Value()
		if mTot != ""{
			mTo = append(mTo,mTot)    //不空的地址组合成一个数组切片
		}
	}
	fmt.Println("mto: %v",mTo)

	// -----------邮件地址-----------
	m := gomail.NewMessage()
	m.SetHeader("From", mFrom)
	m.SetHeader("To", mTo...)		// 注意这里的写法
	//m.SetHeader("Cc", mTo, mCc)
	//m.SetAddressHeader("Cc", mCc, "Dan")
	//m.SetAddressHeader("Cc", mTo, "Dan")

	// -----------邮件内容-----------
	m.SetHeader("Subject", "GameServer Error---游戏服务器GM工具监控发现问题报警邮件")
	m.SetBody("text/html", msg)
	//m.Attach("/home/Alex/lolcat.jpg")

	// -----------发送-----------
	d := gomail.NewDialer(host, port, username, pwd)

	if err := d.DialAndSend(m); err != nil {
		log.LOGGER("Mail error").Info("发送邮件失败！")
	}

	fmt.Println("send ok!")
}
