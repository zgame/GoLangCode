package main

import (
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
func GetScoreRechargeSql(rechargeInfo RechargeList, getGold int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, goldType int, subType int, title string, extend int) string {
	lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)     // 获取玩家的历史金币数量
	table:= ZRandomTo(10,200)
	goldValues := fmt.Sprintf("%d,%d,2273,%d,%d,%d,0,'系统赠送','%s','%s',%d,%d,0,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,title,dataTimeStr,table,getGold+lastAllGold, goldType,subType, extend,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)
}
// 生成金币的增加语句
func GetScoreAddSql(rechargeInfo RechargeList, getGold int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string) string {
	//lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)     // 获取玩家的历史金币数量
	table:= ZRandomTo(10,200)
	goldValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,0,'游戏操作','游戏写分','%s',%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,dataTimeStr, table,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)
}

// 生成金币减少的语句
func GetScoreReduceSql(rechargeInfo RechargeList, reduceGold int, dbNow *sql.DB, dataTimeStr string,dbName string, day1 string) string {
	//lastAllGold := GetHistoryScore(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)     // 获取玩家的历史金币数量
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	goldValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,0,'游戏操作','游戏写分',%s,%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceGold, lastAllGold,reduceGoldTimeOff, table,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, ScoreKeys, goldValues)
}

//-------------------------------------钻石---------------------------------------------

// 生成钻石的充值语句
func GetDiamondRechargeSql(rechargeInfo RechargeList, getDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, DiamondType int, subType int, title string, extend int) string {
	lastAllDiamond := GetHistoryDiamond(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime)                                               // 获取玩家的历史钻石数量
	DiamondValues := fmt.Sprintf("%d,%d,2273,%d,%d,%d,'系统赠送','%s','%s',37,%d,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,title,dataTimeStr, getDiamond+lastAllDiamond, DiamondType,subType, extend,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}
// 生成钻石的增加语句
func GetDiamondAddSql(rechargeInfo RechargeList, getDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string) string {
	DiamondValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,'游戏操作','个人奖池','%s',117,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,dataTimeStr, rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}

// 生成钻石减少的语句
func GetDiamondReduceSql(rechargeInfo RechargeList, reduceDiamond int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string) string {
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	DiamondValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,'游戏操作','购买',%s,%d,0,0,4,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceDiamond, lastAllDiamond,reduceGoldTimeOff, table,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}

//-------------------------------------灵力---------------------------------------------

// 生成灵力的充值语句
func GetCoinRechargeSql(rechargeInfo RechargeList, getCoin int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string, Type int, subType int, title string) string {
	table:= ZRandomTo(10,200)
	lastAllCoin := GetHistoryCoin(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime) // 获取玩家的历史灵力数量
	CoinValues := fmt.Sprintf("%d,%d,1967,%d,%d,%d,0,'游戏操作','%s','%s',%d, %d,0,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getCoin, lastAllCoin,title,dataTimeStr,table, getCoin+lastAllCoin, Type,subType, rechargeInfo.gitPackageId,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameCoinChangeRecord", day1, CoinKeys, CoinValues)
}
//-------------------------------------道具---------------------------------------------

// 生成道具的充值语句
func GetItemRechargeSql(rechargeInfo RechargeList, itemId int, itemNum int, dbNow *sql.DB,dataTimeStr string,dbName string, day1 string,  title string) string {
	table:= ZRandomTo(10,200)
	lastAllItem := GetHistoryItem(dbName, day1,dbNow,rechargeInfo.UserId,rechargeInfo.SuccessTime) // 获取玩家的历史灵力数量
	ItemValues := fmt.Sprintf("%d,%d,1995,%d,%d,%d,0,'游戏操作','%s','%s',%d, %d,0,0,0,2,3,%d,0,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, itemId, itemNum,title,dataTimeStr,table, itemNum+lastAllItem,  rechargeInfo.gitPackageId,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameItemChangeRecord", day1, ItemKeys, ItemValues)
}


















