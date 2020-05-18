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

func DealUserList(idStart int) {
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
	)
	fmt.Println(" --------------开始连接数据库-------------- ")
	platformDB := mssql.ConnectDB(userId, password, server, PlatformDBName)
	DataBaseBYDB := mssql.ConnectDB(userId, password, server, DataBaseBYDBName)
	logDB1 := mssql.ConnectDB(userId, password, server, logDBName1)
	logDB2 := mssql.ConnectDB(userId, password, server, logDBName2)

	//fmt.Println(" --------------开始查询充值列表--------------")
	sqlU := fmt.Sprintf("select  * from PlatformDB_202002.dbo.PPayCoinOrder_2020 with(nolock) where PayStatus=2 and SuccessTime >= 1578585600 and SuccessTime < 1578672000") // 一天
	//sqlU:= fmt.Sprintf( "select  * from PlatformDB_202002.dbo.PPayCoinOrder_2020 with(nolock) where PayStatus=2 and SuccessTime >= 1578585600 and SuccessTime < 1581264000") // 一个月
	_, rows, _ := mssql.Query(platformDB, sqlU)

	for rows.Next() { // 循环遍历
		var rechargeInfo RechargeList
		err := rows.Scan(&rechargeInfo.id,
			&rechargeInfo.orderNo,
			&rechargeInfo.UserId,
			&rechargeInfo.name,
			&rechargeInfo.payType,
			&rechargeInfo.payStatus,
			&rechargeInfo.kindId,
			&rechargeInfo.Money,
			&rechargeInfo.coin,
			&rechargeInfo.giftOnceCoin,
			&rechargeInfo.giftTotalCoin,
			&rechargeInfo.giftOnePayCoin,
			&rechargeInfo.createTime,
			&rechargeInfo.SuccessTime,
			&rechargeInfo.operateUserId,
			&rechargeInfo.remark,
			&rechargeInfo.IP,
			&rechargeInfo.sendStatus,
			&rechargeInfo.OnePay,
			&rechargeInfo.payGiftMoneyLimit,
			&rechargeInfo.crontabPayCount,
			&rechargeInfo.CrontabPayDate,
			&rechargeInfo.ClientKind,
			&rechargeInfo.AppstoreEnvironment,
			&rechargeInfo.ditchNumber,
			&rechargeInfo.coinType,
			&rechargeInfo.actionId,
			&rechargeInfo.userGameId,
			&rechargeInfo.gitPackageId,
			&rechargeInfo.appStoreProductId,
			&rechargeInfo.Diamond,
			&rechargeInfo.giftOnceDiamond,
			&rechargeInfo.giftTotalDiamond,
			&rechargeInfo.giftOnePayDiamond,
			&rechargeInfo.payDiamondMoneyLimit,
			&rechargeInfo.orderDitch,
			&rechargeInfo.channelId,
			&rechargeInfo.registerMachine,
			&rechargeInfo.otherMoney,
			&rechargeInfo.registerDate,
			&rechargeInfo.logonMachine,
			&rechargeInfo.vipLev,
			&rechargeInfo.receiptUserName,
			&rechargeInfo.itemId,
			&rechargeInfo.discountMoney,
			&rechargeInfo.payCount) // 赋值到结构体中
		if err != nil {
			zLog.PrintfLogger(" 遍历充值列表 id %d    , %s \n", rechargeInfo.id, err)
			continue
		}

		zLog.PrintfLogger(" --------------开始处理充值id : %d--------------", rechargeInfo.id)

		// -----------------------------获取单个充值行为------------------------
		dataTimeStr := time.Unix(int64(rechargeInfo.SuccessTime), 0).Format("2006-01-02 15:04:05")
		//fmt.Println("获取时间戳转日期时间", dataTimeStr)
		//fmt.Println("获取时间戳转日期", dataTimeStr[0:10])

		if rechargeInfo.gitPackageId > 0 {
			// 购买礼包

		} else {
			// 金币或者钻石
			if rechargeInfo.coin > 0 {
				// 金币
				getGold := rechargeInfo.coin + rechargeInfo.giftOnceCoin	// type =2
				emailGold:= rechargeInfo.giftOnePayCoin						// type =6

			} else if rechargeInfo.Diamond > 0 {
				//钻石
				getDiamond:= rechargeInfo.Diamond + rechargeInfo.giftOnceDiamond	// type =2
				emailDiamond:= rechargeInfo.giftOnePayDiamond					// type =6

			}
		}

	}

	mssql.CloseQuery(rows)
	mssql.CloseDB(platformDB)
	mssql.CloseDB(DataBaseBYDB)
	mssql.CloseDB(logDB1)
	mssql.CloseDB(logDB2)

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
