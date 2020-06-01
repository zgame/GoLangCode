package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"strings"
	"time"
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
	dataBaseArray := make([]RechargeList,0)

	fmt.Println(" --------------开始连接数据库-------------- ")
	platformDB := mssql.ConnectDB(userId, password, server, PlatformDBName)
	DataBaseBYDB := mssql.ConnectDB(userId, password, server, DataBaseBYDBName)
	logDB1 := mssql.ConnectDB(userId, password, server, logDBName1)
	logDB2 := mssql.ConnectDB(userId, password, server, logDBName2)
	TestDB := mssql.ConnectDB(userId, password, server, TestDBName)


	//fmt.Println(" --------------开始查询充值列表--------------")
	daySecond := 86400		// 一天秒数
	day110 := 1578585600	// 1月10号
	day1:= day110 + (daySecond * idStart)
	day2:= day1 + daySecond
	sqlU := fmt.Sprintf("select ID,UserID,PayStatus,KindID,Money,Coin,GiftOnceCoin,GiftOnePayCoin,SuccessTime,ClientKind,GiftPackageID,Diamond,GiftOnceDiamond,GiftOnePayDiamond,ChannelID  from testdb.dbo.PPayCoinOrder_2020 with(nolock) where PayStatus=2 and SuccessTime >= %d and SuccessTime < %d ",day1,day2) // 一天
	//sqlU:= fmt.Sprintf( "select  * from PlatformDB_202002.dbo.PPayCoinOrder_2020 with(nolock) where PayStatus=2 and SuccessTime >= 1578585600 and SuccessTime < 1581264000") // 一个月
	//fmt.Println("sql:",sqlU)
	_, rows, _ := mssql.Query(TestDB, sqlU)

	for rows.Next() { // 循环遍历
		var rechargeInfo RechargeList
		err := rows.Scan(&rechargeInfo.id, &rechargeInfo.UserId, &rechargeInfo.payStatus, &rechargeInfo.kindId, &rechargeInfo.Money,
			&rechargeInfo.coin, &rechargeInfo.giftOnceCoin, &rechargeInfo.giftOnePayCoin, &rechargeInfo.SuccessTime,
			&rechargeInfo.ClientKind, &rechargeInfo.gitPackageId,
			&rechargeInfo.Diamond, &rechargeInfo.giftOnceDiamond, &rechargeInfo.giftOnePayDiamond, &rechargeInfo.channelId) // 赋值到结构体中
		if err != nil {
			zLog.PrintfLogger(" 遍历充值列表 id %d    , %s \n", rechargeInfo.id, err)
			continue
		}
		dataBaseArray = append(dataBaseArray, rechargeInfo) //添加到列表

	}
	zLog.PrintfLogger(" --------------一共有 : %d  条数据--------------", len(dataBaseArray))
	for index,rechargeInfo := range dataBaseArray{
		zLog.PrintfLogger(" --------------开始处理充值index : %d     rechargeinfo.UserId: %d--------------", index,rechargeInfo.UserId)
		zLog.PrintfLogger(" --------------开始处理充值id : %d--------------", rechargeInfo.id)

		// -----------------------------获取单个充值行为------------------------
		dataTimeStr := time.Unix(int64(rechargeInfo.SuccessTime), 0).Format("2006-01-02 15:04:05")
		dayString := dataTimeStr[0:10]
		day1 := strings.Replace(dayString, "-", "", -1) // 去掉-， 整理成表的后缀
		dbNow, dbName := GetMonth(day1, logDB1, logDB2)
		//dbNow = dbNow
		//fmt.Println("获取时间戳转日期时间", dataTimeStr)
		//fmt.Println("获取时间戳转日期", dataTimeStr[0:10])

		if rechargeInfo.gitPackageId > 0 {
			// 购买礼包
			GetGiftPackageRechargeSql(  rechargeInfo, dbNow, dataTimeStr, dbName, day1 ,TestDB)
		} else {
			// 金币或者钻石
			if rechargeInfo.coin > 0 {
				// 金币
				getGold := rechargeInfo.coin + rechargeInfo.giftOnceCoin // type =2
				emailGold := rechargeInfo.giftOnePayCoin                 // type =6

				// 插入充值金币语句
				lastAllScore:=GetScoreRechargeSql(rechargeInfo, getGold, dbNow, dataTimeStr, dbName, day1, 2,1,"充值金币赠送", rechargeInfo.Money,TestDB)
				// 插入邮件赠送
				if emailGold > 0 {
					GetScoreRechargeSql(rechargeInfo, emailGold, dbNow, dataTimeStr, dbName, day1, 6,4,"首充赠送", rechargeInfo.Money,TestDB)
				}
				GetScoreReduceSql(rechargeInfo, getGold+emailGold, dbNow, dataTimeStr, dbName, day1,lastAllScore)

			} else if rechargeInfo.Diamond > 0 {
				//钻石
				getDiamond := rechargeInfo.Diamond + rechargeInfo.giftOnceDiamond // type =2
				emailDiamond := rechargeInfo.giftOnePayDiamond                    // type =6

				// 插入充值钻石语句
				lastAllDiamond:=GetDiamondRechargeSql(rechargeInfo, getDiamond, dbNow, dataTimeStr, dbName, day1, 2,1,"充值钻石赠送", rechargeInfo.Money,TestDB)
				// 插入邮件赠送
				if emailDiamond > 0 {
					GetDiamondRechargeSql(rechargeInfo, emailDiamond, dbNow, dataTimeStr, dbName, day1, 6,4,"首充赠送", rechargeInfo.Money,TestDB)
				}
				GetDiamondReduceSql(rechargeInfo, getDiamond+emailDiamond, dbNow, dataTimeStr, dbName, day1,lastAllDiamond)

			}
		}
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
	if strings.Contains(table1, "202001") {
		dbNow = logDB1
		dbName = "BY_LOG_202001"
	} else {
		dbNow = logDB2
		dbName = "BY_LOG_202002"
	}
	return dbNow, dbName
}
