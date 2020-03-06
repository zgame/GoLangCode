package main

import (
	_ "./routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "zsw1:zsw123@/by_statis_db?charset=utf8")
	//fmt.Println("-----init----------")
	//orm.RunSyncdb("default", false, true)
	//fmt.Println("------init---------")
}


func main() {
	//fmt.Println("---------------")
	//o := orm.NewOrm()
	//o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//
	//fmt.Println("---------------")
	//res, err := o.Raw("SELECT * FROM user").Exec()
	//if err == nil {
	//	num, _ := res.RowsAffected()
	//	fmt.Println("mysql row affected nums: ", num)
	//}
	//
	//
	//z1 := new(models.ZswT)
	//z1.Id = 2
	//z1.Name = "2223333"
	////o.Insert(z1)
	////o.Update(z1)
	//
	//z2 := new(models.ZswT)
	//z2.Name = "44"
	//o.Read(z2)
	//fmt.Println("z2",z2.Id)
	//
	//user := new(models.User)
	////user.Name = "zsw"
	//user.Id = 1
	//o.Read(user)
	//fmt.Println("------------")
	//fmt.Println("zsw:", user.Pwd)
	//
	//
	//fmt.Println("******************************")
	//var user2 []models.User
	//num,err := o.Raw("SELECT * FROM user ").QueryRows(&user2)
	//fmt.Printf("%v",user2[0].Name)
	//fmt.Println("")
	//fmt.Println("num",num)
	beego.Run()
}

