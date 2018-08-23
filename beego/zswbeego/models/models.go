package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int
	Name        string
	Pwd     	string
	Login_time       string
	Is_admin	int
	Is_dashboard	int
	Is_statis	int
	Is_edit	int
}

type ZswT struct {
	Id int
	Name string
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User),new(ZswT))
}
