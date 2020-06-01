package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	"strconv"
)

//遗留分数
func GetHistoryScore( dbName string, day1 string ,dbNow *sql.DB,  rechargeInfo UserList, TestDB *sql.DB) (int,string) {
	return GetHistory(dbName,day1,dbNow,rechargeInfo.UserId,"GameScoreChangeRecord","Score", rechargeInfo.id,0,TestDB)
}

//遗留钻石
func GetHistoryDiamond( dbName string, day1 string ,dbNow *sql.DB,  rechargeInfo UserList, TestDB *sql.DB) (int,string) {
	return GetHistory(dbName,day1,dbNow,rechargeInfo.UserId,"GameDiamondChangeRecord","Diamond", rechargeInfo.id,0,TestDB)
}
//遗留灵力
func GetHistoryCoin( dbName string, day1 string ,dbNow *sql.DB,  rechargeInfo UserList, TestDB *sql.DB) (int,string) {
	return GetHistory(dbName,day1,dbNow,rechargeInfo.UserId,"GameCoinChangeRecord","Coin", rechargeInfo.id,0,TestDB)
}
//遗留道具
func GetHistoryItem( dbName string, day1 string ,dbNow *sql.DB,  rechargeInfo UserList, itemId int, TestDB *sql.DB) (int,string) {
	return GetHistory(dbName,day1,dbNow,rechargeInfo.UserId,"GameItemChangeRecord","ItemIndbNum", rechargeInfo.id,itemId,TestDB)
}


// 获取历史遗留
func GetHistory( dbName string, day1 string ,dbNow *sql.DB, userId int, tableNameT string, keyName string ,  rechargeId int, itemId int, TestDB *sql.DB) (int,string) {
	dayInt,_ := strconv.Atoi(day1)
	tableName := ""

	for i:=60;i>0;i-- {

		//num := dayInt - 20200210
		if dayInt==20200200{
			dayInt = 20200131	// 跳到1月份
		}
		if dayInt == 20200100{
			zLog.PrintfLogger(" dayInt 太往前了，已经要搜到12月份了, userid: %d  id :%d ", userId, rechargeId)
			InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId,TestDB,itemId)
			return 0,""
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
		sqlStr := fmt.Sprintf("select top(1)%s from %s where RecordTime = (select max(RecordTime) from %s where UserID = %d ) and UserID = %d %s", keyName,tableName, tableName, userId, userId,itemAdd)
		zLog.PrintfLogger("获取%s历史sql: %s ", keyName, sqlStr)

		_, rows, _ := mssql.Query(dbNow, sqlStr)
		for rows.Next() { // 循环遍历
			var result int
			var recordTime string
			err := rows.Scan(&result,&recordTime)
			if err != nil {
				zLog.PrintfLogger(" %s历史遗留表 rechargeId %d    , %s \n", keyName, result, err)
				continue
			}
			//if result >= 0 {
				zLog.PrintfLogger("userid : %d,   %s   id:%d 获取数量： %d", userId,  keyName, itemId, result)
				mssql.CloseQuery(rows)
				return result,recordTime
			//}
		}
		mssql.CloseQuery(rows)
	}
	return 0,""
}


// 把没有找到数据的玩家uid insert到表中
func InsertUserIdWhenCanNotFindOut(userId int, keyName string, id int,TestDB *sql.DB, itemId int)  {
	if TestDB == nil{
		return
	}

	sqlStr := fmt.Sprintf("insert into dbo.can_not_find_last_%s(userId,id) values (%d,%d)", keyName, userId, id)
	if itemId>0 {
		sqlStr = fmt.Sprintf("insert into dbo.can_not_find_last_%s(userId,id,itemId) values (%d,%d,%d)", keyName, userId, id, itemId)
	}
	err,_ :=mssql.Exec(TestDB, sqlStr)
	if err!= nil{
		zLog.PrintfLogger(" 没有找到数据的玩家 insert   uid: %d    id: %d ,   %s \n", userId,id, err)
	}
}
