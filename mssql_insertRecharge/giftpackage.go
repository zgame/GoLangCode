package main

import (
	"./zLog"
)

// 礼包处理
func GetGiftPackageRechargeSql(rechargeInfo RechargeList, getGold int, lastAllGold int, dataTimeStr string, dbName string, day1 string, goldType int) {
	switch rechargeInfo.gitPackageId {
	case 6, 7, 8, 9:
		// 金币 钻石
		addGoldSql := GetGoldRechargeSql(rechargeInfo, getGold, lastAllGold, dataTimeStr, dbName, day1, 2)
		zLog.PrintfLogger("插入充值金币语句 %s", addGoldSql)

	case 101:

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

}
