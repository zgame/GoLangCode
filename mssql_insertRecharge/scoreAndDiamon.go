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


// 获取

// insert语句
func GetInsertSql(dbName string, opType string, day1 string,  keys string, values string )  string{
	return fmt.Sprintf("insert into %s.dbo.%s_%s(%s) values(%s)", dbName, opType, day1 , keys, values)
}




//-------------------------------------金币---------------------------------------------

// 生成金币的充值语句
func GetScoreRechargeSql(rechargeInfo RechargeList, getGold int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, goldType int, subType int, title string, extend int) int {
	reason := "系统赠送"
	if title == "游戏写分"{
		reason = "游戏操作"
	}
	lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo)     // 获取玩家的历史金币数量
	table:= ZRandomTo(10,200)
	goldValues := fmt.Sprintf("%d,%d,2273,%d,%d,%d,0,'%s','%s','%s',%d,%d,0,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,reason,title,dataTimeStr,table,getGold+lastAllGold, goldType,subType, extend,rechargeInfo.channelId)
	addGoldSql:= GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)
	//zLog.PrintfLogger("插入充值金币语句 %s", addGoldSql)
	err, _ := mssql.Exec(dbNow, addGoldSql)
	if err != nil {
		zLog.PrintfLogger("GetScoreRechargeSql Exec Error %s ,sql: %s", err.Error(), addGoldSql)
	}
	return lastAllGold
}
// 生成金币的增加语句
func GetScoreAddSql(rechargeInfo RechargeList, getGold int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string) string {
	lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo)     // 获取玩家的历史金币数量
	table:= ZRandomTo(10,200)
	goldValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,0,'游戏操作','游戏写分','%s',%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,dataTimeStr, table,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)
}

// 生成金币减少的语句
func GetScoreReduceSql(rechargeInfo RechargeList, reduceGold int, dbNow *sql.DB, dataTimeStr string,dbName string, day1 string,lastAllGold int)  {
	//lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)     // 获取玩家的历史金币数量
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	goldValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,0,'游戏操作','游戏写分',%s,%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceGold, lastAllGold,reduceGoldTimeOff, table,rechargeInfo.channelId)
	reduceGoldSql := GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)

	//zLog.PrintfLogger("插入金币减少语句 %s", reduceGoldSql)
	err, _ := mssql.Exec(dbNow, reduceGoldSql)
	if err != nil {
		zLog.PrintfLogger("GetScoreReduceSql Exec Error %s ,sql: %s", err.Error(), reduceGoldSql)
	}
}

//-------------------------------------钻石---------------------------------------------

// 生成钻石的充值语句
func GetDiamondRechargeSql(rechargeInfo RechargeList, getDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, DiamondType int, subType int, title string, extend int) int {
	lastAllDiamond := GetHistoryDiamond(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo)                                               // 获取玩家的历史钻石数量
	DiamondValues := fmt.Sprintf("%d,%d,2273,%d,%d,%d,'系统赠送','%s','%s',37,%d,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,title,dataTimeStr, getDiamond+lastAllDiamond, DiamondType,subType, extend,rechargeInfo.channelId)
	addDiamondSql := GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
	//zLog.PrintfLogger("插入充值钻石语句 %s", addDiamondSql)
	err, _ := mssql.Exec(dbNow, addDiamondSql)
	if err != nil {
		zLog.PrintfLogger("GetDiamondRechargeSql Exec Error %s ,sql: %s", err.Error(), addDiamondSql)
	}
	return lastAllDiamond
}
// 生成钻石的增加语句
func GetDiamondAddSql(rechargeInfo RechargeList, getDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string) string {
	lastAllDiamond := GetHistoryDiamond(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo)
	DiamondValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,'游戏操作','个人奖池','%s',117,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,dataTimeStr, rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}

// 生成钻石减少的语句
func GetDiamondReduceSql(rechargeInfo RechargeList, reduceDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,lastAllDiamond int)  {
	//lastAllDiamond := GetHistoryDiamond(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	DiamondValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,'游戏操作','购买',%s,%d,0,0,4,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceDiamond, lastAllDiamond, reduceTimeOff, table,rechargeInfo.channelId)
	reduceGoldSql :=  GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
	//zLog.PrintfLogger("插入减少钻石语句 %s", reduceGoldSql)
	err, _ := mssql.Exec(dbNow, reduceGoldSql)
	if err != nil {
		zLog.PrintfLogger("GetDiamondReduceSql Exec Error %s ,sql: %s", err.Error(), reduceGoldSql)
	}
}

//-------------------------------------灵力---------------------------------------------

// 生成灵力的充值语句
func GetCoinRechargeSql(rechargeInfo RechargeList, getCoin int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, Type int, subType int, title string) int {
	table:= ZRandomTo(10,200)
	lastAllCoin := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo) // 获取玩家的历史灵力数量
	CoinValues := fmt.Sprintf("%d,%d,1967,%d,%d,%d,0,'游戏操作','%s','%s',%d, %d,0,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getCoin, lastAllCoin,title,dataTimeStr,table, getCoin+lastAllCoin, Type,subType, rechargeInfo.gitPackageId,rechargeInfo.channelId)
	AddCoinSql := GetInsertSql(dbName, "GameCoinChangeRecord", day1, CoinKeys, CoinValues)
	//zLog.PrintfLogger("增加灵力语句 %s", AddCoinSql)
	err, _ := mssql.Exec(dbNow, AddCoinSql)
	if err != nil {
		zLog.PrintfLogger("GetCoinRechargeSql Exec Error %s ,sql: %s", err.Error(), AddCoinSql)
	}
	return lastAllCoin
}

// 生成灵力减少的语句
func GetCoinReduceSql(rechargeInfo RechargeList, reduceCoin int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, lastAllCoin int)  {
	//lastAllCoin := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime) // 获取玩家的历史灵力数量
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	CoinValues := fmt.Sprintf("%d,%d,1259,%d,%d,%d,0,'游戏操作','游戏写分',%s,%d,0 , 0,0,1,6,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceCoin, lastAllCoin, reduceTimeOff, table,rechargeInfo.channelId)
	reduceCoinSql:= GetInsertSql(dbName, "GameCoinChangeRecord", day1, CoinKeys, CoinValues)
	//zLog.PrintfLogger("减少灵力语句 %s", reduceCoinSql)
	err, _ := mssql.Exec(dbNow, reduceCoinSql)
	if err != nil {
		zLog.PrintfLogger("GetCoinReduceSql Exec Error %s ,sql: %s", err.Error(), reduceCoinSql)
	}
}
//-------------------------------------道具---------------------------------------------

// 生成道具的充值语句
func GetItemRechargeSql(rechargeInfo RechargeList, itemId int, itemNum int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,  title string) int {
	lastAllItem := GetHistoryItem(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo) // 获取玩家的历史灵力数量
	ItemValues := fmt.Sprintf("%d,%d,1995,%d,%d,%d,'游戏操作','%s','%s',%d, 0,0,0,2,3,%d,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, itemId, itemNum,title,dataTimeStr,itemNum+lastAllItem,  rechargeInfo.gitPackageId,rechargeInfo.channelId)
	reduceItemSql:= GetInsertSql(dbName, "GameItemChangeRecord", day1, ItemKeys, ItemValues)
	//zLog.PrintfLogger("增加道具语句 %s", reduceItemSql)
	err, _ := mssql.Exec(dbNow, reduceItemSql)
	if err != nil {
		zLog.PrintfLogger("GetItemRechargeSql Exec Error %s ,sql: %s", err.Error(), reduceItemSql)
	}
	return lastAllItem
}

// 生成道具减少的语句
func GetItemReduceSql(rechargeInfo RechargeList, itemId int, itemNum int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,lastAllItem int )  {
	if itemId == 3102 {
		return
	}
	//lastAllItem := GetHistoryItem(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime, rechargeInfo) // 获取玩家的历史灵力数量
	title := "使用"
	switch itemId {
	case 151:
		title = "雪人大作战活动消耗"
	case 2007:
		title = "解锁炮"
	case 7003:
		title = "转换消耗道具"
	}
	randTime:= ZRandomTo(20,60)
	reduceTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	ItemValues := fmt.Sprintf("%d,%d,1201,%d,%d,%d,'游戏操作','%s',%s, %d,0,0,0,4,19,0,0,%d ", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, itemId, -itemNum, title,  reduceTimeOff,lastAllItem, rechargeInfo.channelId)
	reduceItemSql:= GetInsertSql(dbName, "GameItemChangeRecord", day1, ItemKeys, ItemValues)
	//zLog.PrintfLogger("减少道具语句 %s", reduceItemSql)
	err, _ := mssql.Exec(dbNow, reduceItemSql)
	if err != nil {
		zLog.PrintfLogger("GetItemReduceSql Exec Error %s ,sql: %s", err.Error(), reduceItemSql)
	}
}















