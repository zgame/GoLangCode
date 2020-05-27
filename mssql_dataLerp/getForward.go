package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	"strconv"
)

//遗留分数
func GetForwardScore( dbName string, dayStart string ,dbNow *sql.DB,    rechargeInfo UserList) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameScoreChangeRecord","Score", rechargeInfo.id,0)
}

//遗留钻石
func GetForwardDiamond( dbName string, dayStart string ,dbNow *sql.DB,  rechargeInfo UserList) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameDiamondChangeRecord","Diamond", rechargeInfo.id,0)
}
//遗留灵力
func GetForwardCoin( dbName string, dayStart string ,dbNow *sql.DB,   rechargeInfo UserList) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameCoinChangeRecord","Coin", rechargeInfo.id,0)
}
//遗留道具
func GetForwardItem( dbName string, dayStart string ,dbNow *sql.DB,  rechargeInfo UserList, itemId int) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameItemChangeRecord","ItemIndbNum", rechargeInfo.id,itemId)
}


// 获取历史遗留
func GetForward( dbName string, dayStart string ,dbNow *sql.DB, userId int, tableNameT string, keyName string , rechargeId int, itemId int) (int,string) {
	dayInt,_ := strconv.Atoi(dayStart)
	tableName := ""

	for i:=0;i<60;i++ {

		//num := dayInt - 20200210
		if dayInt == 20200132{
			dayInt = 20200201	// 跳到2月份
		}
		if dayInt == 20200230{
			zLog.PrintfLogger(" dayInt 太往后了，已经要搜到3月份了, userId: %d  id :%d ", userId, rechargeId)
			//InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId,TestDB,itemId)
			return 0,""
			//dayInt = 20191230	// 跳到12月份
		}

		tableName = fmt.Sprintf("%s.dbo.%s_%d", dbName, tableNameT,dayInt)
		if dayInt > 20200200{
			tableName = fmt.Sprintf("BY_LOG_202002.dbo.%s_%d",  tableNameT, dayInt)	// 如果是2月，就用2月的库
		}
		dayInt++

		itemAdd :=""
		if itemId>0 {
			itemAdd = fmt.Sprintf(" and ItemID = %d ",itemId)
		}
		sqlStr := fmt.Sprintf("select top(1)%s,RecordTime from %s where RecordTime = (select min(RecordTime) from %s where UserID = %d ) and UserID = %d %s", keyName,tableName, tableName, userId, userId,itemAdd)
		zLog.PrintfLogger("获取%s未来sql: %s ", keyName, sqlStr)

		_, rows, _ := mssql.Query(dbNow, sqlStr)
		for rows.Next() { // 循环遍历
			var result int
			var timeS string
			err := rows.Scan(&result,&timeS)
			if err != nil {
				zLog.PrintfLogger(" %s历史遗留表 rechargeId %d    , %s \n", keyName, result, err)
				continue
			}
			if result >= 0 {
				zLog.PrintfLogger("userId : %d,   %s   id:%d 获取数量： %d", userId,  keyName, itemId, result)
				mssql.CloseQuery(rows)
				return result,timeS
			}
		}
		mssql.CloseQuery(rows)
	}
	return 0,""
}

