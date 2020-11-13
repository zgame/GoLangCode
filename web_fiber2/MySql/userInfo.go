package MySql

import (
	"fmt"
	"log"
)

// 根据用户表结构创建， 大写代表_
type Userinfo struct {
	Uid      int    `xorm:"int"`
	Openid   string `xorm:"varchar(30)"`
	Psw      string `xorm:"varchar(15)"`
	Mac      string `xorm:"varchar(30)"`
	Nickname string `xorm:"varchar(25)"`
}

// 查询数据
func GetUserInfoData() {
	fmt.Println("开始同步表结构")
	err := DataBaseEngine.Sync2(new(Userinfo)) //同步表跟结构
	if err != nil {
		fmt.Println("同步表结构出错！", err)
		log.Fatal(err)
		return
	}

	fmt.Println("开始获取单条数据")
	var dataUser Userinfo
	DataBaseEngine.Get(&dataUser) //获取单条数据
	println(dataUser.Uid)

	var DataUsers []Userinfo
	DataBaseEngine.Find(&DataUsers) //获取多条
	println(len(DataUsers))
	for i, v := range DataUsers {
		fmt.Println(i)
		fmt.Printf("%d:%d", i, v.Uid)
		fmt.Println("")
	}

	//println("-------------------------")

}

// 更新数据
func UpdateUserInfoData() {
	// update更新数据
	_, err := DataBaseEngine.Exec("update userinfo set uid = ? where uid = ?", 1112, 1111)
	if err != nil {
		fmt.Println("数据库更新出错！", err)
		log.Fatal(err)
		return
	}

}
// 插入单行数据
func InsertUserInfoData() {
	// insert 单条数据
	var insertData Userinfo
	insertData.Uid = 1111
	_, err := DataBaseEngine.Insert(&insertData)

	if err != nil {
		fmt.Println("数据库插入出错！", err)
		log.Fatal(err)
		return
	}
}
