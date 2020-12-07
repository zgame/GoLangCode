package zMongoDB


// -----------------------------------------------------------------------------
// mongodb
// 启动服务器： cd D:\Program Files\MongoDBCollection\Server\3.4\bin
//				mongod.exe --dbpath "D:\Program Files\MongoDBCollection\Server\3.4\data\db"
// 启动客户端： cd D:\Program Files\MongoDBCollection\Server\3.4\bin
//				mongod.exe
// 客户端可视化工具：https://robomongo.org/download

// 命令：	show dbs
//			db				// 查看当前选择的数据库是哪一个
//			use  ***     // 有就切换， 没有就创建
//			show collections    		//集合列表

// -----------------------------------------------------------------------------


import (
"fmt"
"log"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string
}
var MongoSession * mgo.Session
var MongoDBCollection *mgo.Collection

func main() {
	MongoDBInit()
	var err error

	fmt.Println("----------------insert--------------")
	err = MongoDBCollection.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("----------------删除--------------")
	err = MongoDBCollection.Remove(bson.M{"name": "Ale"})
	err = MongoDBCollection.Remove(bson.M{"name": "Cla"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----------------更新--------------")
	err = MongoDBCollection.Update(bson.M{"name": "Cla"}, 	&Person{"Cla", "343333333333"})
	if err != nil {
		log.Fatal(err)
	}



	fmt.Println("---------------find-------------")
	result := Person{}
	err = MongoDBCollection.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)

	fmt.Println("---------------end-------------")
	EndMongoDB()
}
// MongoDB数据库初始化
func MongoDBInit()  {
	var err error
	MongoSession, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	MongoSession.SetMode(mgo.Monotonic, true)
	MongoDBCollection = MongoSession.DB("test").C("people")

}
// 关闭数据库
func EndMongoDB()  {
	MongoSession.Close()
}
