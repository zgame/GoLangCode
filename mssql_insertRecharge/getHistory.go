package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	"strconv"
)

//遗留分数
func GetHistoryScore( dbName string, day1 string ,dbNow *sql.DB, userId int, rechargeTime int, rechargeInfo RechargeList, TestDB *sql.DB) int {
	return GetHistory(dbName,day1,dbNow,userId,"GameScoreChangeRecord","Score",GetTimeFromInt(rechargeTime), rechargeInfo.id,0,TestDB)
}

//遗留钻石
func GetHistoryDiamond( dbName string, day1 string ,dbNow *sql.DB, userId int, rechargeTime int,rechargeInfo RechargeList, TestDB *sql.DB) int {
	return GetHistory(dbName,day1,dbNow,userId,"GameDiamondChangeRecord","Diamond",GetTimeFromInt(rechargeTime), rechargeInfo.id,0,TestDB)
}
//遗留灵力
func GetHistoryCoin( dbName string, day1 string ,dbNow *sql.DB, userId int, rechargeTime int,rechargeInfo RechargeList, TestDB *sql.DB) int {
	return GetHistory(dbName,day1,dbNow,userId,"GameCoinChangeRecord","Coin",GetTimeFromInt(rechargeTime), rechargeInfo.id,0,TestDB)
}
//遗留道具
func GetHistoryItem( dbName string, day1 string ,dbNow *sql.DB, userId int, rechargeTime int,rechargeInfo RechargeList, itemId int, TestDB *sql.DB) int {
	return GetHistory(dbName,day1,dbNow,userId,"GameItemChangeRecord","ItemIndbNum",GetTimeFromInt(rechargeTime), rechargeInfo.id,itemId,TestDB)
}


// 获取历史遗留
func GetHistory( dbName string, day1 string ,dbNow *sql.DB, userId int, tableNameT string, keyName string , rechargeTime string , rechargeId int, itemId int, TestDB *sql.DB) int {
	dayInt,_ := strconv.Atoi(day1)
	tableName := ""

	for i:=60;i>0;i-- {

		//num := dayInt - 20200210
		if dayInt==20200200{
			dayInt = 20200131	// 跳到1月份
		}
		if dayInt == 20200100{
			zLog.PrintfLogger(" dayInt 太往前了，已经要搜到12月份了, userid: %d  id :%d ", userId, rechargeId)
			InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId,TestDB)
			return 0
			//dayInt = 20191230	// 跳到12月份
		}
		tableName = fmt.Sprintf("%s.dbo.%s_%d", dbName, tableNameT,dayInt)
		if dayInt < 20200200{
			tableName = fmt.Sprintf("BY_LOG_202001.dbo.%s_%d",  tableNameT, dayInt)
		}
		//if dayInt < 20200100{
		//	tableName = fmt.Sprintf("BY_LOG_201912.dbo.%s_%d",  tableNameT, dayInt)
		//}
		//if dayInt == 20191200{
		//	zLog.PrintfLogger(" dayInt 太往前了，已经要搜到11月份了, userid: %d  id :%d", userId,rechargeId)
		//	InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId)
		//	return 0
		//}
		dayInt--

		itemAdd :=""
		if itemId>0 {
			itemAdd = fmt.Sprintf(" and ItemID = %d ",itemId)
		}
		sql := fmt.Sprintf("select top(1)%s from %s where RecordTime = (select max(RecordTime) from %s where UserID = %d and RecordTime < '%s') and UserID = %d %s", keyName,tableName, tableName, userId, rechargeTime,userId,itemAdd)
		//zLog.PrintfLogger("获取%s历史sql: %s ", keyName, sql)

		_, rows, _ := mssql.Query(dbNow, sql)
		for rows.Next() { // 循环遍历
			var result int
			err := rows.Scan(&result)
			if err != nil {
				zLog.PrintfLogger(" %s历史遗留表 rechargeId %d    , %s \n", keyName, result, err)
				continue
			}
			if result >= 0 {
				zLog.PrintfLogger("userid : %d,   %s   id:%d 获取数量： %d", userId,  keyName, itemId, result)
				mssql.CloseQuery(rows)
				return result
			}
		}
		mssql.CloseQuery(rows)
	}
	return 0
}


// 把没有找到数据的玩家uid insert到表中
func InsertUserIdWhenCanNotFindOut(userId int, keyName string, id int,TestDB *sql.DB)  {
	sql := fmt.Sprintf("insert into dbo.can_not_find_last_%s(userId,id) values (%d,%d)", keyName, userId, id)
	err,_ :=mssql.Exec(TestDB, sql)
	if err!= nil{
		zLog.PrintfLogger(" 没有找到数据的玩家 insert   uid: %d    id: %d ,   %s \n", userId,id, err)
	}
}
