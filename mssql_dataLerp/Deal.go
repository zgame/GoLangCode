package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"strconv"
	"strings"
)

var (
	userId     = "dbuser"
	password   = "CEDFE2CDA7DB84AC"
	server     = "172.16.140.89"
	logDBName1 = "BY_LOG_202001"
	logDBName2 = "BY_LOG_202002"

	//userIdReadOnly     = "dbuser_ro"
	//passwordReadOnly   = "35A20E7966ECDC93"
	PlatformDBName   = "PlatformDB_202001"
	DataBaseBYDBName = "DataBaseBY_202002"
	TestDBName = "testdb"

	//TestDB *sql.DB
)


func DealUserList(idStart int) {
	dataBaseArray := make([]UserList,0)

	fmt.Println(" --------------开始连接数据库-------------- ")
	platformDB := mssql.ConnectDB(userId, password, server, PlatformDBName)
	DataBaseBYDB := mssql.ConnectDB(userId, password, server, DataBaseBYDBName)
	logDB1 := mssql.ConnectDB(userId, password, server, logDBName1)
	logDB2 := mssql.ConnectDB(userId, password, server, logDBName2)
	TestDB := mssql.ConnectDB(userId, password, server, TestDBName)




	//fmt.Println(" --------------开始查询充值列表--------------")
	//daySecond := 86400		// 一天秒数
	//day110 := 1578585600	// 1月10号
	day1:= Group * idStart
	day2:= Group * (idStart + 1)
	sqlU := fmt.Sprintf("select *  from testdb.dbo.a1_user_chongzhi_id with(nolock) where id >= %d and id < %d ",day1,day2) // 一天
	//sqlU:= fmt.Sprintf( "select  * from PlatformDB_202002.dbo.PPayCoinOrder_2020 with(nolock) where PayStatus=2 and SuccessTime >= 1578585600 and SuccessTime < 1581264000") // 一个月
	//fmt.Println("sql:",sqlU)
	_, rows, _ := mssql.Query(TestDB, sqlU)

	for rows.Next() { // 循环遍历
		var Info UserList
		err := rows.Scan(&Info.id, &Info.UserId, &Info.registerDate, &Info.lastLoginDate,&Info.kindId,&Info.ClientKind,&Info.channelId) // 赋值到结构体中
		if err != nil {
			zLog.PrintfLogger(" 遍历玩家列表 id %d    , %s \n", Info.id, err)
			continue
		}
		dataBaseArray = append(dataBaseArray, Info) //添加到列表

	}
	zLog.PrintfLogger(" --------------一共有 : %d  条数据--------------", len(dataBaseArray))
	for index, userInfo := range dataBaseArray{
		zLog.PrintfLogger(" --------------开始处理index : %d     UserId: %d--------------", index, userInfo.UserId)
		zLog.PrintfLogger(" --------------开始处理id : %d--------------", userInfo.id)

		// -----------------------------获取单个 user 行为------------------------
		//dataTimeStr := time.Unix(int64(userInfo.SuccessTime), 0).Format("2006-01-02 15:04:05")
		dataTimeStr := userInfo.registerDate
		dayString := dataTimeStr[0:10]
		dayStart := strings.Replace(dayString, "-", "", -1) // 去掉-， 整理成表的后缀
		dataTimeStr2 := userInfo.lastLoginDate
		dayString2 := dataTimeStr2[0:10]
		day2 := strings.Replace(dayString2, "-", "", -1) // 去掉-， 整理成表的后缀
		dbNow, dbName := GetMonth(dayStart, logDB1, logDB2)
		//dbNow = dbNow
		//fmt.Println("获取时间戳转日期时间", dataTimeStr)
		//fmt.Println("获取时间戳转日期", dataTimeStr[0:10])
		day1N,_ :=strconv.Atoi(dayStart) // 注册日期
		day2N,_ :=strconv.Atoi(day2)     // 流失日期

		changeScore :=0
		if day1N < 20200110 {
			// 玩家的注册日期在1月10日之前， 用日志库金额， 要缝合1月10号或者之后的首次记录
			dayStart = "20200110"
			dbNow, dbName = GetMonth(dayStart, logDB1, logDB2)
			forwardScore, recordTime := GetForwardScore(dbName, dayStart,dbNow,  userInfo) // 获取玩家的最终金币数量
			//fmt.Println("forwardScore",forwardScore)
			//fmt.Println("recordTime",recordTime)
			dayStart = "20200109"
			lastAllScore,_ := GetHistoryScore(dbName, dayStart,dbNow,  userInfo,nil) // 获取玩家的历史金币数量
			//fmt.Println("lastAllScore",lastAllScore)
			changeScore = forwardScore - lastAllScore
			//fmt.Println("差额", changeScore)
			DealScore(changeScore,userInfo,logDB1,recordTime,logDB2, lastAllScore, -1)		// 进行修改， 在最终的地方，提前一点，进行缝合
		}else {
			//  新增注册用户，起始是0  ，不用缝合

		}
		if day2N > 20200210{
			// 玩家在2月10日依然留存， 用日志库金额
			dayStart = "20200210"
			dbNow, dbName = GetMonth(dayStart, logDB1, logDB2)
			forwardScore, _ := GetForwardScore(dbName, dayStart,dbNow,  userInfo) // 获取玩家的最终金币数量
			//fmt.Println("forwardScore",forwardScore)
			//fmt.Println("recordTime",recordTime)
			dayStart = "20200209"
			lastAllScore,recordTime := GetHistoryScore(dbName, dayStart,dbNow,  userInfo,nil) // 获取玩家的历史金币数量
			//fmt.Println("lastAllScore",lastAllScore)
			changeScore = forwardScore - lastAllScore
			fmt.Println("差额", changeScore)
			DealScore(changeScore,userInfo,logDB1,recordTime,logDB2, lastAllScore,1)		// 进行修改，在最后的记录上，增加一点时间进行缝合

		}else{
			// 玩家在2月10日前流失了 ， 玩家的结束金额应该是游戏库金额
			//dayStart = "20200210"
			//dbNow, dbName = GetMonth(dayStart, logDB1, logDB2)
			//forwardScore, _ := GetForwardScore(dbName, dayStart,dbNow,  userInfo) // 获取玩家的最终金币数量
			//fmt.Println("forwardScore",forwardScore)
			//fmt.Println("recordTime",recordTime)

			dayStart = day2		// 流失的时候
			lastAllScore,recordTime := GetHistoryScore(dbName, dayStart ,dbNow,  userInfo,nil)     // 获取玩家的历史金币数量
			//fmt.Println("lastAllScore",lastAllScore)
			changeScore = forwardScore - lastAllScore
			fmt.Println("差额", changeScore)
			DealScore(changeScore,userInfo,logDB1,recordTime,logDB2, lastAllScore,1)		// 进行修改，在最后的记录上，增加一点时间进行缝合
		}




		//if userInfo.gitPackageId > 0 {
		//	// 购买礼包
		//	GetGiftPackageRechargeSql(userInfo, dbNow, dataTimeStr, dbName, dayStart ,TestDB)
		//} else {
		//	// 金币或者钻石
		//	if userInfo.coin > 0 {
		//		// 金币
		//		getGold := userInfo.coin + userInfo.giftOnceCoin // type =2
		//		emailGold := userInfo.giftOnePayCoin             // type =6
		//
		//		// 插入充值金币语句
		//		lastAllScore:=GetScoreRechargeSql(userInfo, getGold, dbNow, dataTimeStr, dbName, dayStart, 2,1,"充值金币赠送", userInfo.Money,TestDB)
		//		// 插入邮件赠送
		//		if emailGold > 0 {
		//			GetScoreRechargeSql(userInfo, emailGold, dbNow, dataTimeStr, dbName, dayStart, 6,4,"首充赠送", userInfo.Money,TestDB)
		//		}
		//		GetScoreReduceSql(userInfo, getGold+emailGold, dbNow, dataTimeStr, dbName, dayStart,lastAllScore)
		//
		//	} else if userInfo.Diamond > 0 {
		//		//钻石
		//		getDiamond := userInfo.Diamond + userInfo.giftOnceDiamond // type =2
		//		emailDiamond := userInfo.giftOnePayDiamond                // type =6
		//
		//		// 插入充值钻石语句
		//		lastAllDiamond:=GetDiamondRechargeSql(userInfo, getDiamond, dbNow, dataTimeStr, dbName, dayStart, 2,1,"充值钻石赠送", userInfo.Money,TestDB)
		//		// 插入邮件赠送
		//		if emailDiamond > 0 {
		//			GetDiamondRechargeSql(userInfo, emailDiamond, dbNow, dataTimeStr, dbName, dayStart, 6,4,"首充赠送", userInfo.Money,TestDB)
		//		}
		//		GetDiamondReduceSql(userInfo, getDiamond+emailDiamond, dbNow, dataTimeStr, dbName, dayStart,lastAllDiamond)
		//
		//
		//	}
		//}
	}

	mssql.CloseQuery(rows)

	mssql.CloseDB(platformDB)
	mssql.CloseDB(DataBaseBYDB)
	mssql.CloseDB(logDB1)
	mssql.CloseDB(logDB2)
	mssql.CloseDB(TestDB)

	wg.Done()
}

// 获取当前的月份
func GetMonth(table1 string, logDB1 *sql.DB, logDB2 *sql.DB) (*sql.DB, string) {
	var dbNow *sql.DB
	var dbName string
	if strings.Contains(table1, "202002") {
		dbNow = logDB2
		dbName = "BY_LOG_202002"
	} else {
		dbNow = logDB1
		dbName = "BY_LOG_202001"
	}
	return dbNow, dbName
}


// 做平金币
func DealScore(score int,userInfo UserList, logDB1 *sql.DB, dataTimeStr string, logDB2  *sql.DB, lastAllScore int , addTime int)  {
	dayString := dataTimeStr[0:10]
	day1 := strings.Replace(dayString, "-", "", -1) // 去掉-， 整理成表的后缀
	dbNow, dbName := GetMonth(day1, logDB1, logDB2)
	if score>0 {
		GetScoreAddSql(userInfo , score, dbNow ,dataTimeStr ,dbName , day1 ,lastAllScore,addTime)
	}
	if score <0 {
		GetScoreReduceSql(userInfo, score, dbNow, dataTimeStr, dbName, day1,lastAllScore,addTime)
	}

}
