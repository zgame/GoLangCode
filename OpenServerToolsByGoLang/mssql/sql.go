package main

import (
	"fmt"
	"log"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"strconv"
)


// 创建房间的sql

func CreateRoomSql(RoomIndex int, ServerID int, ServerName string, ServerMachine string)  {
	fmt.Println("开始创建房间:", ServerID)
	list := GameRoomListSelect()[RoomIndex]
	f,err := ioutil.ReadFile("sql_create_room/"+ list.Name )
	if err != nil{
		fmt.Println("read file " + list.Name +  " Error")
	}
	sql := string(f)
	sql = fmt.Sprintf(sql,strconv.Itoa(ServerID), ServerName, strconv.Itoa(ServerID+10000), ServerMachine)
	//fmt.Println("sql:" ,sql)

	_ ,err = Engine.Query(sql)
	if err != nil{
		fmt.Println("数据库创建房间出现问题", list.Name)
		return
	}
	fmt.Println("------------------创建房间成功", list.Name, "  ID：", ServerID)
}



// 创建gameStore

func CreateGameStore(ServerID int,Android int)  {
	fmt.Println("开始创建GameStore",)

	f,err := ioutil.ReadFile("sql_game_store/game_store.sql")
	if err != nil{
		fmt.Println("read file  game_store.sql  Error")
	}
	sql := string(f)
	sql = fmt.Sprintf(sql,strconv.Itoa(ServerID), strconv.Itoa(Android))
	//fmt.Println("sql:" ,sql)

	_ ,err = Engine.Query(sql)
	if err != nil{
		fmt.Println("数据库创建game_store出错 :", ServerID)
		return
	}
	fmt.Println("--------------创建game store成功: " , ServerID)
}


// 创建机器人sql

func CreateAndroid(SqlIndex int,ServerID int,Android int)  {
	if SqlIndex == 0{
		fmt.Println("不需要创建机器人")
		return
	}
	fmt.Println("开始创建机器人",)
	list := SqlFileRobotListSelect()[SqlIndex]
	f,err := ioutil.ReadFile("sql_android_files/"+ list.Name)
	if err != nil{
		fmt.Println("read file  "+list.Name+"  Error")
	}
	sql := string(f)

	for i:=1;i<=Android;i++{
		sqlt := fmt.Sprintf(sql, strconv.Itoa(i),strconv.Itoa(ServerID),"%")
		fmt.Println("机器人",i)
		_ ,err = Engine.Query(sqlt)
		if err != nil{
			fmt.Println("数据库创建机器人出错 :", ServerID)
			return
		}
	}

	fmt.Println("--------------创建机器人成功: " , ServerID)
}














// 把列表数据保存到数据库中
func saveDataToSql(list []string) {
	engine := getDataBase()

	engine.Query("delete  from AaWhiteIPList") // 删除全部

	test_ip_list := make([]AaWhiteIPList, 0)
	for _, v := range list {
		test_ip_list = append(test_ip_list, AaWhiteIPList{Ip: v})
	}

	fmt.Printf("%v", test_ip_list)
	fmt.Println("")
	_, err := engine.Insert(test_ip_list) //插入多条，因为test_iii是结构数组
	if err != nil {
		fmt.Println("保存数据出错！", err)
		log.Fatal(err)
	}

}

// 从数据库中把数据读出来
func getSqlData() []AaWhiteIPList {

	var Engine *xorm.Engine
	Engine = getDataBase()

	// 从sql中获取数据
	has, _ := Engine.IsTableExist(new(AaWhiteIPList))
	fmt.Printf("是否存在这个表：", has)
	fmt.Println("")

	if has {
		fmt.Println("开始绑定数据结构...")
		err := Engine.Sync2(new(AaWhiteIPList))
		fmt.Println("绑定数据结构成功")

		var test_ip_list []AaWhiteIPList
		err = Engine.Find(&test_ip_list)

		if err != nil {
			fmt.Println("数据库查询出错！", err)
			log.Fatal(err)
			return nil
		}
		fmt.Printf("test_ip_list:%v", test_ip_list)
		fmt.Println("")

		return test_ip_list
	} else {
		fmt.Println("表不存在，那么创建它！")
		Engine.CreateTables(new(AaWhiteIPList))
		return nil
	}

}
