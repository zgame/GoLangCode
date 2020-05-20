package main

import (
	"./zLog"
	"database/sql"
)

// 礼包处理
func GetGiftPackageRechargeSql(rechargeInfo RechargeList, dbNow *sql.DB, dataTimeStr string, dbName string, day1 string) {
	// 变量初始化
	getScore := 0   // 金币
	getDiamond := 0 // 钻石
	getCoin := 0    // 灵力
	title := ""     //原因
	Type := 2
	SubType := 1 // 子分类
	switch rechargeInfo.gitPackageId {
	case 6:
		// 金币 钻石
		getScore = 200000
		getDiamond = 10
		title = "充值青铜礼包赠送"
		Type = 2
		SubType = 3
		// 还差道具的

	case 7:
		// 金币 钻石
		getScore = 400000
		getDiamond = 20
		title = "充值白银礼包赠送"
		Type = 2
		SubType = 3

	case 8:
		// 金币 钻石
		getScore = 1000000
		getDiamond = 50
		title = "充值黄金礼包赠送"
		Type = 2
		SubType = 3
	case 9:
		// 金币 钻石
		getScore = 2000000
		getDiamond = 100
		title = "充值白金礼包赠送"
		Type = 2
		SubType = 3
	case 101:
		//金币 钻石
		getScore = 2000000
		getDiamond = 100
		title = "充值白金礼包赠送"
		Type = 2
		SubType = 3

	case 3101, 3102, 3105:

	case 10004:
	case 10005:
	case 10006, 10007, 10008, 10009:
	case 10010, 10011, 10012, 10013, 10014, 10015:
	case 10016, 10017, 10018, 10019, 10020, 10021:
	case 10022, 10023, 10024:
	case 10125, 10126, 10127, 10128:
	case 10129, 10130:
	case 11001, 11002, 11003, 11004:
	case 13001, 13002, 13003:
	case 14001, 14002:
	case 15005, 15006, 15007, 15008:
	case 16005:

	}

	if getScore > 0 {
		// 插入充值金币
		addGoldSql := GetGoldRechargeSql(rechargeInfo, getScore, dbNow, dataTimeStr, dbName, day1, Type, SubType, title)
		zLog.PrintfLogger("礼包 %d 插入充值金币语句 %s", rechargeInfo.gitPackageId, addGoldSql)
	}
	if getDiamond > 0 {
		// 插入充值钻石语句
		addDiamondSql := GetDiamondRechargeSql(rechargeInfo, getDiamond, dbNow, dataTimeStr, dbName, day1, Type, SubType, title)
		zLog.PrintfLogger("礼包 %d 插入充值钻石语句 %s", rechargeInfo.gitPackageId, addDiamondSql)
	}
	if getCoin > 0 {
		// 插入灵力

	}

}
