package main

import (
	"./zLog"
	"database/sql"
	"fmt"
	"./mssql"
	"strconv"
)

//遗留分数
func GetHistoryScore( dbName string, day1 string ,dbNow *sql.DB, userId int, rechargeTime int) int {
	return GetHistory(dbName,day1,dbNow,userId,"GameScoreChangeRecord","Score",GetTimeFromInt(rechargeTime))
}

//遗留钻石
func GetHistoryDiamond( dbName string, day1 string ,dbNow *sql.DB, userId int, rechargeTime int) int {
	return GetHistory(dbName,day1,dbNow,userId,"GameDiamondChangeRecord","Diamond",GetTimeFromInt(rechargeTime))
}


// 获取历史遗留
func GetHistory( dbName string, day1 string ,dbNow *sql.DB, userId int, tableName string, keyName string , rechargeTime string) int {
	dayInt,_ := strconv.Atoi(day1)

	for i:=30;i>0;i-- {
		tableName := fmt.Sprintf("%s.dbo.%s_%d", dbName, tableName,dayInt)
		//num := dayInt - 20200210
		dayInt--
		if dayInt==20200200{
			dayInt = 20200131	// 跳到1月份
		}
		if dayInt == 20200100{
			zLog.PrintfLogger(" dayInt 太往前了，已经要搜到12月份了, userid: %d ", userId)
			InsertUserIdWhenCanNotFindOut(userId,keyName)
			return 0
			//dayInt = 20191230	// 跳到12月份
		}
		//if dayInt < 20200100{
		//	tableName = fmt.Sprintf("BY_LOG_201912.dbo.GameScoreChangeRecord_%d",  dayInt)
		//}
		//if dayInt == 20191200{
		//	zLog.PrintfLogger(" dayInt 太往前了，已经要搜到11月份了, userid: %d ", userId)
		//	return 0
		//}

		sql := fmt.Sprintf("select top(1)%s from %s where RecordTime = (select max(RecordTime) from %s where UserID = %d and RecordTime < '%s') and UserID = %d", keyName,tableName, tableName, userId,userId, rechargeTime)
		zLog.PrintfLogger("获取%s历史sql: %s ", keyName, sql)

		_, rows, _ := mssql.Query(dbNow, sql)
		for rows.Next() { // 循环遍历
			var score int
			err := rows.Scan(&score)
			if err != nil {
				zLog.PrintfLogger(" 遍历历史遗留表 id %d    , %s \n", score, err)
				continue
			}
			if score >0 {
				zLog.PrintfLogger("userid : %d, 获取数量： %d", userId, score)
				return score
			}
		}
	}
	return 0
}


// 把没有找到数据的玩家uid insert到表中
func InsertUserIdWhenCanNotFindOut(userId int, keyName string)  {
	sql := fmt.Sprintf("insert into dbo.can_not_find_last_%s(userId) values (%d)", keyName, userId)
	mssql.Exec(TestDB, sql)
}
