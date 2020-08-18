package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
)

//遗留分数
func GetForwardScore( dbName string, dayStart string ,dbNow *sql.DB,    rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameScoreChangeRecord","Score", "ChangeScore", 0, gameDB)
}

//遗留钻石
func GetForwardDiamond( dbName string, dayStart string ,dbNow *sql.DB,  rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameDiamondChangeRecord","Diamond", "ChangeDiamond",0, gameDB)
}
//遗留灵力
func GetForwardCoin( dbName string, dayStart string ,dbNow *sql.DB,   rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameCoinChangeRecord","Coin", "ChangeCoin",0, gameDB)
}
//遗留道具
func GetForwardItem( dbName string, dayStart string ,dbNow *sql.DB,  rechargeInfo UserList, itemId int,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameItemChangeRecord","ItemIndbNum", "ItemNum",itemId, gameDB)
}
//遗留Lottery
func GetForwardLottery( dbName string, dayStart string ,dbNow *sql.DB,   rechargeInfo UserList,  gameDB *sql.DB) (int,string) {
	return GetForward(dbName,dayStart,dbNow,rechargeInfo.UserId,"GameLotteryChangeRecord","Lottery", "ChangeLottery",0, gameDB)
}

// 获取历史遗留
func GetForward( dbName string, dayStart string ,dbNow *sql.DB, userId int, tableNameT string, keyName string , changeKey string,  itemId int, gameDB02 *sql.DB) (int,string) {
	dayInt := 20200201
	var recordTime string
	//recordTime := "2020-02-01 04:11:55"
	//dayInt,_ := strconv.Atoi(dayStart)
	tableName := ""

	for i:=0;i<60;i++ {

		if dayInt==20200201{
			recordTime = "and RecordTime > '2020-02-01 04:12:26' "
		}else{
			recordTime = ""
		}

		//num := dayInt - 20200210
		//if dayInt == 20200132{
		//	dayInt = 20200201	// 跳到2月份
		//}
		if dayInt == 20200230{
			//endTime := "2020-03-01 00:00:00"
			//forwardScore,forwardDiamond,forwardCoin := GetDataBaseBY(gameDB02, userId)
			//forwardItem := 0
			//if itemId>0 {
			//	forwardItem = GetDataBaseBYItem(gameDB02,userId, itemId)
			//}
			//forwardLottery := GetDataBaseBYLottery(gameDB02, userId)
			//switch keyName {
			//case "Lottery":
			//	return forwardLottery,endTime
			//case "Score":
			//	//zLog.PrintfLogger(" score %d ", forwardScore)
			//	return forwardScore,endTime
			//case "Diamond":
			//	//zLog.PrintfLogger(" Diamond %d ", forwardDiamond)
			//	return forwardDiamond,endTime
			//case "Coin":
			//	//zLog.PrintfLogger(" Coin %d ", forwardCoin)
			//	return forwardCoin,endTime
			//case "ItemIndbNum":
			//	//zLog.PrintfLogger(" ItemIndbNum %d ", forwardItem)
			//	return forwardItem,endTime
			//}

			//zLog.PrintfLogger(" 没有找到 %d ", 0)
			//InsertUserIdWhenCanNotFindOut(userId,keyName, rechargeId,LogDB,itemId)
			return -1,"2020-02-30"
			//dayInt = 20191230	// 跳到12月份
		}

		//tableName = fmt.Sprintf("%s.dbo.%s_%d", dbName, tableNameT,dayInt)
		//if dayInt > 20200200{
			tableName = fmt.Sprintf("BY_LOG_202002.dbo.%s_%d",  tableNameT, dayInt)	// 如果是2月，就用2月的库
		//}
		dayInt++

		itemAdd :=""
		if itemId>0 {
			itemAdd = fmt.Sprintf(" and ItemID = %d ",itemId)
		}
		sqlStr := fmt.Sprintf("select top(1)%s,%s,RecordTime from %s where RecordTime = (select min(RecordTime) from %s where UserID = %d %s %s) and UserID = %d %s", keyName, changeKey, tableName, tableName, userId, itemAdd, recordTime, userId,itemAdd)
		//zLog.PrintfLogger("forward 获取%s未来sql: %s ", keyName, sqlStr)

		_, rows, _ := mssql.Query(dbNow, sqlStr)
		for rows.Next() { // 循环遍历
			var result int
			var change int
			var timeS string
			err := rows.Scan(&result, &change, &timeS)
			if err != nil {
				zLog.PrintfLogger(" %s历史遗留表 rechargeId %d    , %s \n", keyName, result, err)
				continue
			}
			//if result >= 0 {
			result = result - change		// 这里要注意， 因为获取出来的数据是经过变化的， 那么要去掉这个新的变化才是上一个遗留值
				//zLog.PrintfLogger("forward 往后查找  userId : %d,   %s   id:%d 获取数量： %d  time: %s", userId,  keyName, itemId, result, timeS)
				mssql.CloseQuery(rows)
				return result,timeS
			//}
		}
		mssql.CloseQuery(rows)
	}
	return 0,""
}

