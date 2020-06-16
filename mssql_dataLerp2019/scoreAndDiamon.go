package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
)

var ScoreKeys = "UserID,KindID,ServerID,ClientKind,ChangeScore,Score,Insure,OprAcc,ChangeReson,RecordTime,TableArea,ScoreIndb,InsureIndb,IsEmail,Type,SubType,Extend,iDitchId"
var DiamondKeys = "UserID,KindID,ServerID,ClientKind,ChangeDiamond,Diamond,OprAcc,ChangeReson,RecordTime,TableArea,DiamondIndb,IsEmail,Type,SubType,Extend,iDitchId"
var ItemKeys = "UserID,KindID,ServerID,ClientKind,ItemID,ItemNum,OprAcc,ChangeReson,RecordTime,ItemIndbNum, GetScore,MasterID, IsEmail,Type,SubType,Extend, IsBigMG, iDitchId"
var CoinKeys = "UserID,KindID,ServerID,ClientKind,ChangeCoin,Coin,Insure,OprAcc,ChangeReson,RecordTime,TableArea, CoinIndb , InsureIndb, IsEmail,Type,SubType,Extend, iDitchId"
var LotteryKeys = "UserID,KindID,ServerID,ClientKind,ChangeLottery,Lottery,OprAcc,ChangeReson,RecordTime,TableArea, LotteryIndb ,  IsEmail,Type,SubType,Extend, iDitchId"




// insert语句
func GetInsertSql(dbName string, opType string, day1 string,  keys string, values string )  string{
	return fmt.Sprintf("insert into %s.dbo.%s_%s(%s) values(%s)", dbName, opType, day1 , keys, values)
}


//-------------------------------------金币---------------------------------------------


//生成金币的增加语句
func GetScoreAddSql(rechargeInfo UserList, getGold int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,lastAllGold int, addTime int)  {
	table:= ZRandomTo(10,200)
	randTime:= 1//ZRandomTo(20,60)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	goldValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,0,'游戏操作','游戏写分',%s,%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,reduceGoldTimeOff, table,rechargeInfo.channelId)
	addGoldSql:=  GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)
	//zLog.PrintfLogger("插入增加金币语句 %s", addGoldSql)
	err, _ := mssql.Exec(dbNow, addGoldSql)
	if err != nil {
		zLog.PrintfLogger("GetScoreAddSql Exec Error %s ,sql: %s", err.Error(), addGoldSql)
	}
}

// 生成金币减少的语句
func GetScoreReduceSql(rechargeInfo UserList, reduceGold int, dbNow *sql.DB, dataTimeStr string,dbName string, day1 string,lastAllGold int, addTime int)  {
	//reduceGoldPart1 := int( reduceGold / 3)
	//reduceGoldPart2 := reduceGold - reduceGoldPart1
	//lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)     // 获取玩家的历史金币数量
	randTime:= 2//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	// --------------------第一次减少------------------------------
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	goldValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,0,'游戏操作','游戏写分',%s,%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, reduceGold, lastAllGold,reduceGoldTimeOff, table,rechargeInfo.channelId)
	reduceGoldSql := GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)

	//zLog.PrintfLogger("插入金币减少语句 %s", reduceGoldSql)
	err, _ := mssql.Exec(dbNow, reduceGoldSql)
	if err != nil {
		zLog.PrintfLogger("GetScoreReduceSql Exec Error %s ,sql: %s", err.Error(), reduceGoldSql)
	}

}

//-------------------------------------钻石---------------------------------------------


// 生成钻石的增加语句
func GetDiamondAddSql(rechargeInfo UserList, getDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,lastAllDiamond int, addTime int)  {
	//lastAllDiamond := GetHistoryDiamond(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo)
	randTime:= 1//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	DiamondValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,'游戏操作','个人奖池',%s,%d,0,0,3,45,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,reduceTimeOff, table,rechargeInfo.channelId)
	addDiamondSql := GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)

	//zLog.PrintfLogger("插入增加钻石语句 %s", addDiamondSql)
	err, _ := mssql.Exec(dbNow, addDiamondSql)
	if err != nil {
		zLog.PrintfLogger("GetDiamondAddSql Exec Error %s ,sql: %s", err.Error(), addDiamondSql)
	}
}

// 生成钻石减少的语句
func GetDiamondReduceSql(rechargeInfo UserList, reduceDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,lastAllDiamond int, addTime int)  {
	//lastAllDiamond := GetHistoryDiamond(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)
	randTime:= 2//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	DiamondValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,'游戏操作','购买',%s,%d,0,0,4,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, reduceDiamond, lastAllDiamond, reduceTimeOff, table,rechargeInfo.channelId)
	reduceGoldSql :=  GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
	//zLog.PrintfLogger("插入减少钻石语句 %s", reduceGoldSql)
	err, _ := mssql.Exec(dbNow, reduceGoldSql)
	if err != nil {
		zLog.PrintfLogger("GetDiamondReduceSql Exec Error %s ,sql: %s", err.Error(), reduceGoldSql)
	}
}

//-------------------------------------灵力---------------------------------------------


//生成灵力的增加语句
func GetCoinAddSql(rechargeInfo UserList, getCoin int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, lastAllCoin int, addTime int)  {
	randTime:= 1//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	//lastAllCoin := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo,TestDB) // 获取玩家的历史灵力数量
	CoinValues := fmt.Sprintf("%d,%d,1967,%d,%d,%d,0,'游戏操作','杀分',%s,%d, 0,0,0,1,6,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getCoin, lastAllCoin,reduceTimeOff,table,  rechargeInfo.channelId)
	AddCoinSql := GetInsertSql(dbName, "GameCoinChangeRecord", day1, CoinKeys, CoinValues)
	//zLog.PrintfLogger("增加灵力语句 %s", AddCoinSql)
	err, _ := mssql.Exec(dbNow, AddCoinSql)
	if err != nil {
		zLog.PrintfLogger("GetCoinRechargeSql Exec Error %s ,sql: %s", err.Error(), AddCoinSql)
	}

}

// 生成灵力减少的语句
func GetCoinReduceSql(rechargeInfo UserList, reduceCoin int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, lastAllCoin int, addTime int)  {
	//lastAllCoin := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime) // 获取玩家的历史灵力数量
	randTime:= 2//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	CoinValues := fmt.Sprintf("%d,%d,1259,%d,%d,%d,0,'游戏操作','杀分',%s,%d,0 , 0,0,1,6,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, reduceCoin, lastAllCoin, reduceTimeOff, table,rechargeInfo.channelId)
	reduceCoinSql:= GetInsertSql(dbName, "GameCoinChangeRecord", day1, CoinKeys, CoinValues)
	//zLog.PrintfLogger("减少灵力语句 %s", reduceCoinSql)
	err, _ := mssql.Exec(dbNow, reduceCoinSql)
	if err != nil {
		zLog.PrintfLogger("GetCoinReduceSql Exec Error %s ,sql: %s", err.Error(), reduceCoinSql)
	}
}
//-------------------------------------道具---------------------------------------------

//生成道具的增加语句
func GetItemAddSql(rechargeInfo UserList, itemId int, itemNum int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string ,lastAllItem int, addTime int)  {
	//lastAllItem := GetHistoryItem(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo,itemId,TestDB) // 获取玩家的历史灵力数量
	randTime:= 1//ZRandomTo(20,60)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)

	title := ""
	typeI := 0
	subType := 0
	extend := 0

	switch itemId {
	case 108, 109, 110, 111:
		title = "寻宝抽奖"
		typeI = 3
		subType = 7
		extend = 10
	case 101, 102, 2007:
		title = "捕获"
		typeI = 1
		subType = 17
		extend = 10
	case 103, 120, 7003:
		title = "购买"
		typeI = 3
		subType = 19
		extend = 300
	case 131:
		title = "网页活动"
		typeI = 6
		subType = 17
		extend = 0
	case 151:
		title = "增加：排行榜奖励"
		typeI = 3
		subType = 2
		extend = 0
	case 150:
		title = "赠送"
		typeI = 7
		subType = 3
		extend = 1001
	}

	ItemValues := fmt.Sprintf("%d,%d,1995,%d,%d,%d,'游戏操作','%s',%s,%d, 0,0,0,%d,%d,%d,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, itemId, itemNum,title,reduceTimeOff,lastAllItem,  typeI, subType,extend,rechargeInfo.channelId)
	reduceItemSql:= GetInsertSql(dbName, "GameItemChangeRecord", day1, ItemKeys, ItemValues)
	//zLog.PrintfLogger("增加道具语句 %s", reduceItemSql)
	err, _ := mssql.Exec(dbNow, reduceItemSql)
	if err != nil {
		zLog.PrintfLogger("GetItemRechargeSql Exec Error %s ,sql: %s", err.Error(), reduceItemSql)
	}

}

// 生成道具减少的语句
func GetItemReduceSql(rechargeInfo UserList, itemId int, itemNum int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,lastAllItem int , addTime int)  {
	randTime:= 2//ZRandomTo(20,60)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)

	if itemId == 3102 {
		return
	}
	//lastAllItem := GetHistoryItem(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo) // 获取玩家的历史灵力数量
	title := "使用"
	switch itemId {
	case 150:
		title = "扣除：竞技场报名"
	case 151:
		title = "雪人大作战活动消耗"
	case 2007:
		title = "解锁炮"
	case 7003:
		title = "转换消耗道具"
	}

	ItemValues := fmt.Sprintf("%d,%d,1201,%d,%d,%d,'游戏操作','%s',%s, %d,0,0,0,4,19,0,0,%d ", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, itemId, itemNum, title,  reduceTimeOff, lastAllItem, rechargeInfo.channelId)
	reduceItemSql:= GetInsertSql(dbName, "GameItemChangeRecord", day1, ItemKeys, ItemValues)
	//zLog.PrintfLogger("减少道具语句 %s", reduceItemSql)
	err, _ := mssql.Exec(dbNow, reduceItemSql)
	if err != nil {
		zLog.PrintfLogger("GetItemReduceSql Exec Error %s ,sql: %s", err.Error(), reduceItemSql)
	}
}





//-------------------------------------奖券---------------------------------------------


//生成奖券的增加语句
func GetLotteryAddSql(rechargeInfo UserList, getLottery int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, lastAllLottery int, addTime int)  {
	randTime:= 1//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	//lastAllLottery := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo,TestDB) // 获取玩家的历史灵力数量
	LotteryValues := fmt.Sprintf("%d,%d,1044,%d,%d,%d,'游戏操作','捕获',%s,%d, 0,0,1,17,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getLottery, lastAllLottery,reduceTimeOff,table,  rechargeInfo.channelId)
	AddLotterySql := GetInsertSql(dbName, "GameLotteryChangeRecord", day1, LotteryKeys, LotteryValues)
	//zLog.PrintfLogger("增加奖券语句 %s", AddLotterySql)
	err, _ := mssql.Exec(dbNow, AddLotterySql)
	if err != nil {
		zLog.PrintfLogger("GetLotteryRechargeSql Exec Error %s ,sql: %s", err.Error(), AddLotterySql)
	}

}

// 生成奖券减少的语句
func GetLotteryReduceSql(rechargeInfo UserList, reduceLottery int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, lastAllLottery int, addTime int)  {
	//lastAllLottery := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime) // 获取玩家的历史灵力数量
	randTime:= 2//ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	if addTime < 0 {
		randTime = -randTime
	}
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	LotteryValues := fmt.Sprintf("%d,%d,1045,%d,%d,%d,'游戏操作','幸运抽奖(奖券)扣除奖券',%s,%d,0 , 0,4,3,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, reduceLottery, lastAllLottery, reduceTimeOff, table,rechargeInfo.channelId)
	reduceLotterySql:= GetInsertSql(dbName, "GameLotteryChangeRecord", day1, LotteryKeys, LotteryValues)

	//zLog.PrintfLogger("减少奖券语句 %s", reduceCoinSql)
	err, _ := mssql.Exec(dbNow, reduceLotterySql)
	if err != nil {
		zLog.PrintfLogger("GetLotteryReduceSql Exec Error %s ,sql: %s", err.Error(), reduceLotterySql)
	}
}




