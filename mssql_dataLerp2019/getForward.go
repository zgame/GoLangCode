package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	"strconv"
)

//遗留分数
func GetForwardScore( dbName string, dayStart string ,dbNow *sql.DB,    rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameScoreChangeRecord","Score", "ChangeScore", rechargeInfo.id,0, gameDB)
}

//遗留钻石
func GetForwardDiamond( dbName string, dayStart string ,dbNow *sql.DB,  rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameDiamondChangeRecord","Diamond", "ChangeDiamond",rechargeInfo.id,0, gameDB)
}
//遗留灵力
func GetForwardCoin( dbName string, dayStart string ,dbNow *sql.DB,   rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameCoinChangeRecord","Coin", "ChangeCoin",rechargeInfo.id,0, gameDB)
}
//遗留Lottery
func GetForwardLottery( dbName string, dayStart string ,dbNow *sql.DB,   rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameLotteryChangeRecord","Lottery", "ChangeLottery",rechargeInfo.id,0, gameDB)
}
//遗留道具
func GetForwardItem( dbName string, dayStart string ,dbNow *sql.DB,  rechargeInfo UserList, itemId int,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameItemChangeRecord","ItemIndbNum", "ItemNum",rechargeInfo.id,itemId, gameDB)
}


// 获取前项遗留
func GetForward( dbName string, dayStart string , logDB *sql.DB, userId int, tableNameT string, keyName string , changeKey string, rechargeId int, itemId int,  gameDB04 *sql.DB) (int,string) {
	dayInt,_ := strconv.Atoi(dayStart)
	tableName := ""

	for i:=0;i<60;i++ {
		//num := dayInt - 20200210
		//if dayInt == 20200132{
		//	dayInt = 20200201	// 跳到2月份
		//}
		if dayInt == 20190430{
			//zLog.PrintfLogger(" dayInt 太往后了，已经要搜到5月份了, userId: %d  id :%d ", userId, rechargeId)
			//InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId,TestDB,itemId)
			endTime := "2019-04-30 00:00:00"
			forwardScore,forwardDiamond,forwardCoin := GetDataBaseBY(gameDB04, userId)
			forwardItem := 0
			if itemId>0 {
				forwardItem = GetDataBaseBYItem(gameDB04,userId, itemId)
			}
			forwardLottery := GetDataBaseBYLottery(gameDB04, userId)
			switch keyName {
			case "Lottery":
				return forwardLottery,endTime
			case "Score":
				//zLog.PrintfLogger(" score %d ", forwardScore)
				return forwardScore,endTime
			case "Diamond":
				//zLog.PrintfLogger(" Diamond %d ", forwardDiamond)
				return forwardDiamond,endTime
			case "Coin":
				//zLog.PrintfLogger(" Coin %d ", forwardCoin)
				return forwardCoin,endTime
			case "ItemIndbNum":
				//zLog.PrintfLogger(" ItemIndbNum %d ", forwardItem)
				return forwardItem,endTime
			}

			//zLog.PrintfLogger(" 没有找到 %d ", 0)
			//InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId,LogDB,itemId)
			return -1,""
			//dayInt = 20191230	// 跳到12月份
		}

		tableName = fmt.Sprintf("%s.dbo.%s_%d", dbName, tableNameT,dayInt)
		//if dayInt > 20200200{
		//	tableName = fmt.Sprintf("BY_LOG_202002.dbo.%s_%d",  tableNameT, dayInt)	// 如果是2月，就用2月的库
		//}
		dayInt++

		itemAdd :=""
		if itemId>0 {
			itemAdd = fmt.Sprintf(" and ItemID = %d ",itemId)
		}
		sqlStr := fmt.Sprintf("select top(1)%s,%s,RecordTime from %s where RecordTime = (select min(RecordTime) from %s where UserID = %d %s) and UserID = %d %s", keyName, changeKey, tableName, tableName, userId,itemAdd, userId,itemAdd)
		//zLog.PrintfLogger("获取%s未来sql: %s ", keyName, sqlStr)

		_, rows, _ := mssql.Query(logDB, sqlStr)
		for rows.Next() { // 循环遍历
			var result int
			var change int
			var timeS string
			err := rows.Scan(&result, &change, &timeS)
			if err != nil {
				zLog.PrintfLogger(" %s历史遗留表 结果 %d  报错： %s \n", keyName, result, err)
				continue
			}
			//if result >= 0 {
			result = result - change		// 这里要注意， 因为获取出来的数据是经过变化的， 那么要去掉这个新的变化才是上一个遗留值
				//zLog.PrintfLogger("userId : %d,   %s   id:%d 获取数量： %d", userId,  keyName, itemId, result)
				mssql.CloseQuery(rows)
				return result,timeS
			//}
		}
		mssql.CloseQuery(rows)
	}
	return 0,""
}

