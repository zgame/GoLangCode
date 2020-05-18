package main

import (
	"fmt"
)

var GoldKeys = "UserID,KindID,ServerID,ClientKind,ChangeScore,Score,Insure,OprAcc,ChangeReson,RecordTime,TableArea,ScoreIndb,InsureIndb,IsEmail,Type,SubType,Extend,iDitchId"
var DiamondKeys = "UserID,KindID,ServerID,ClientKind,ChangeDiamond,Diamond,OprAcc,ChangeReson,RecordTime,TableArea,DiamondIndb,IsEmail,Type,SubType,Extend,iDitchId"

// 获取

// insert语句
func GetInsertSql(dbName string, opType string, day1 string,  keys string, values string )  string{
	return fmt.Sprintf("insert into %s.dbo.%s_%s(%s) values(%s)", dbName, opType, day1 , keys, values)
}




//-------------------------------------金币---------------------------------------------

// 生成金币的充值语句
func GetGoldRechargeSql(rechargeInfo RechargeList, getGold int, lastAllGold int,dataTimeStr string,dbName string, day1 string, goldType int) string {
	subType := 1
	title := "充值金币赠送"
	if goldType == 6{	// 首充
		subType = 4
		title = "首充赠送"
	}
	goldValues := fmt.Sprintf("%d,%d,2273,%d,%d,%d,0,'系统赠送','%s','%s',161,%d,0,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,title,dataTimeStr,getGold+lastAllGold, goldType,subType, rechargeInfo.Money,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, GoldKeys, goldValues)
}
// 生成金币的增加语句
func GetGoldAddSql(rechargeInfo RechargeList, getGold int, lastAllGold int,dataTimeStr string,dbName string, day1 string) string {
	goldValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,0,'游戏操作','游戏写分','%s',40,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getGold, lastAllGold,dataTimeStr, rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, GoldKeys, goldValues)
}

// 生成金币减少的语句
func GetGoldReduceSql(rechargeInfo RechargeList, reduceGold int, lastAllGold int,dataTimeStr string,dbName string, day1 string) string {
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	goldValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,0,'游戏操作','游戏写分',%s,%d,0,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceGold, lastAllGold,reduceGoldTimeOff, table,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameScoreChangeRecord", day1, GoldKeys, goldValues)
}

//-------------------------------------钻石---------------------------------------------

// 生成钻石的充值语句
func GetDiamondRechargeSql(rechargeInfo RechargeList, getDiamond int, lastAllDiamond int,dataTimeStr string,dbName string, day1 string, DiamondType int) string {
	subType := 1
	title := "充值钻石赠送"
	if DiamondType == 6{ // 首充
		subType = 4
		title = "首充赠送"
	}
	DiamondValues := fmt.Sprintf("%d,%d,2273,%d,%d,%d,'系统赠送','%s','%s',37,%d,0,%d,%d,%d,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,title,dataTimeStr, getDiamond+lastAllDiamond, DiamondType,subType, rechargeInfo.Money,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}
// 生成钻石的增加语句
func GetDiamondAddSql(rechargeInfo RechargeList, getDiamond int, lastAllDiamond int,dataTimeStr string,dbName string, day1 string) string {
	DiamondValues := fmt.Sprintf("%d,%d,2240,%d,%d,%d,'游戏操作','个人奖池','%s',117,0,0,1,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, getDiamond, lastAllDiamond,dataTimeStr, rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}

// 生成钻石减少的语句
func GetDiamondReduceSql(rechargeInfo RechargeList, reduceDiamond int, lastAllDiamond int,dataTimeStr string,dbName string, day1 string) string {
	randTime:= ZRandomTo(20,60)
	table:= ZRandomTo(10,200)
	reduceGoldTimeOff := fmt.Sprintf("dateadd(ss,%d,'%s')",randTime,dataTimeStr)
	DiamondValues := fmt.Sprintf("%d,%d,2259,%d,%d,%d,'游戏操作','购买',%s,%d,0,0,4,1,10,%d", rechargeInfo.UserId, rechargeInfo.kindId, rechargeInfo.ClientKind, -reduceDiamond, lastAllDiamond,reduceGoldTimeOff, table,rechargeInfo.channelId)
	return GetInsertSql(dbName, "GameDiamondChangeRecord", day1, DiamondKeys, DiamondValues)
}
