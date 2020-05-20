package main

import (
	"./zLog"
	"database/sql"
	"fmt"
)

type GiftItem struct {
	ItemId  int
	ItemNum int
}

//添加道具列表
func AddItemArray(id int, num int, array []GiftItem) []GiftItem {
	ItemTemp := GiftItem{id, num}
	return append(array, ItemTemp)
}

// 礼包处理
func GetGiftPackageRechargeSql(rechargeInfo RechargeList, dbNow *sql.DB, dataTimeStr string, dbName string, day1 string) {
	// 变量初始化
	getScore := 0   // 金币
	getDiamond := 0 // 钻石
	getCoin := 0    // 灵力
	title := ""     //原因
	Type := 2
	SubType := 3 // 子分类
	ItemArray := make([]GiftItem, 0)
	rate := ZRandomTo(3,5)		// 倍率

	switch rechargeInfo.gitPackageId {
	case 6: // 金币 钻石
		getScore = 200000
		getDiamond = 10
		title = "充值青铜礼包赠送"
		// 还差道具的
		ItemArray = AddItemArray(108, 1, ItemArray)
		ItemArray = AddItemArray(101, 10, ItemArray)

	case 7: // 金币 钻石
		getScore = 400000
		getDiamond = 20
		title = "充值白银礼包赠送"
		ItemArray = AddItemArray(109, 1, ItemArray)
		ItemArray = AddItemArray(120, 5, ItemArray)

	case 8: // 金币 钻石
		getScore = 1000000
		getDiamond = 50
		title = "充值黄金礼包赠送"
		ItemArray = AddItemArray(110, 1, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)

	case 9: // 金币 钻石
		getScore = 2000000
		getDiamond = 100
		title = "充值白金礼包赠送"
		ItemArray = AddItemArray(111, 1, ItemArray)
		ItemArray = AddItemArray(120, 20, ItemArray)

	case 101: //金币 钻石
		getScore = 800000
		getDiamond = 100
		title = "充值至尊礼包月卡赠送"

	case 3101: // 只有道具
	case 3102: // 只有道具
	case 3105: // 只有道具

	case 10004: // 只有道具
		title = "充值礼包特殊奖励加赠"
		ItemArray = AddItemArray(108, 2, ItemArray)
		ItemArray = AddItemArray(109, 2, ItemArray)
		ItemArray = AddItemArray(110, 1, ItemArray)
		Type = 3
		SubType = 13
	case 10005: // 金币
		getScore = 600000
		title = "充值新起航礼包赠送"
		ItemArray = AddItemArray(2007, 300, ItemArray)
		ItemArray = AddItemArray(3102, 18, ItemArray)

	case 10006: //  灵力
		getCoin = 2400000
		title = "充值妖精场每日98元礼包1赠送"
		ItemArray = AddItemArray(120, 20, ItemArray)
		ItemArray = AddItemArray(7003, 20, ItemArray)
	case 10007: //  灵力
		getCoin = 3000000
		title = "充值妖精场每日98元礼包2赠送"
		ItemArray = AddItemArray(120, 30, ItemArray)
		ItemArray = AddItemArray(7003, 30, ItemArray)
	case 10008: //  灵力
		getCoin = 12000000
		title = "充值妖精场每日648元礼包1赠送"
		ItemArray = AddItemArray(7003, 300, ItemArray)
	case 10009: //  灵力
		getCoin = 13000000
		title = "充值妖精场每日648元礼包2赠送"
		ItemArray = AddItemArray(7003, 350, ItemArray)
	case 10010: // 金币
		getScore = 360000
		title = "充值每日6元礼包1赠送"
		ItemArray = AddItemArray(2007, 30, ItemArray)
		ItemArray = AddItemArray(101, 5, ItemArray)
		ItemArray = AddItemArray(102, 2, ItemArray)
	case 10011: // 金币
		getScore = 400000
		title = "充值每日6元礼包2赠送"
		ItemArray = AddItemArray(2007, 50, ItemArray)
		ItemArray = AddItemArray(101, 7, ItemArray)
		ItemArray = AddItemArray(102, 3, ItemArray)
	case 10012: // 金币
		getScore = 480000
		title = "充值每日6元礼包3赠送"
		ItemArray = AddItemArray(2007, 100, ItemArray)
		ItemArray = AddItemArray(101, 10, ItemArray)
		ItemArray = AddItemArray(102, 5, ItemArray)
	case 10013: // 金币
		getScore = 1320000
		title = "充值每日30元礼包1赠送"
		ItemArray = AddItemArray(2007, 180, ItemArray)
		ItemArray = AddItemArray(101, 20, ItemArray)
		ItemArray = AddItemArray(102, 10, ItemArray)
	case 10014: // 金币
		getScore = 1540000
		title = "充值每日30元礼包2赠送"
		ItemArray = AddItemArray(2007, 240, ItemArray)
		ItemArray = AddItemArray(101, 30, ItemArray)
		ItemArray = AddItemArray(102, 15, ItemArray)
	case 10015: // 金币
		getScore = 1760000
		title = "充值每日30元礼包3赠送"
		ItemArray = AddItemArray(2007, 300, ItemArray)
		ItemArray = AddItemArray(101, 40, ItemArray)
		ItemArray = AddItemArray(102, 20, ItemArray)
	case 10016: // 金币  灵力
		getScore = 4000000
		getCoin = 100000
		title = "充值每日98元礼包1赠送"
		ItemArray = AddItemArray(2007, 400, ItemArray)
		ItemArray = AddItemArray(120, 5, ItemArray)
		ItemArray = AddItemArray(109, 1, ItemArray)
	case 10017: // 金币  灵力
		getScore = 4400000
		getCoin = 120000
		title = "充值每日98元礼包2赠送"
		ItemArray = AddItemArray(2007, 500, ItemArray)
		ItemArray = AddItemArray(120, 7, ItemArray)
		ItemArray = AddItemArray(109, 1, ItemArray)
	case 10018: // 金币  灵力
		getScore = 4800000
		getCoin = 150000
		title = "充值每日98元礼包3赠送"
		ItemArray = AddItemArray(2007, 600, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)
		ItemArray = AddItemArray(109, 2, ItemArray)
	case 10019: // 金币  灵力
		getScore = 7600000
		getCoin = 200000
		title = "充值每日198元礼包1赠送"
		ItemArray = AddItemArray(2007, 666, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)
		ItemArray = AddItemArray(131, 2, ItemArray)
	case 10020: // 金币  灵力
		getScore = 8000000
		getCoin = 250000
		title = "充值每日198元礼包2赠送"
		ItemArray = AddItemArray(2007, 777, ItemArray)
		ItemArray = AddItemArray(120, 15, ItemArray)
		ItemArray = AddItemArray(131, 3, ItemArray)
	case 10021: // 金币  灵力
		getScore = 8400000
		getCoin = 300000
		title = "充值每日198元礼包3赠送"
		ItemArray = AddItemArray(2007, 888, ItemArray)
		ItemArray = AddItemArray(120, 20, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)
	case 10022: // 金币 钻石  灵力
		getScore = 24000000
		getDiamond = 1000
		getCoin = 500000
		title = "充值每日648元礼包1赠送"
		ItemArray = AddItemArray(110, 1, ItemArray)
		ItemArray = AddItemArray(120, 20, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)
	case 10023: // 金币 钻石  灵力
		getScore = 26000000
		getDiamond = 1100
		getCoin = 600000
		title = "充值每日648元礼包2赠送"
		ItemArray = AddItemArray(110, 1, ItemArray)
		ItemArray = AddItemArray(120, 30, ItemArray)
		ItemArray = AddItemArray(131, 6, ItemArray)
	case 10024: // 金币 钻石  灵力
		getScore = 30000000
		getDiamond = 1200
		getCoin = 800000
		title = "充值每日648元礼包3赠送"
		ItemArray = AddItemArray(110, 2, ItemArray)
		ItemArray = AddItemArray(120, 50, ItemArray)
		ItemArray = AddItemArray(131, 8, ItemArray)
	case 10125: // 金币
		getScore = 1200000
		title = "充值海盗宝藏（首次）赠送"
		ItemArray = AddItemArray(101, 15, ItemArray)
	case 10126: // 金币
		getScore = 600000
		title = "充值海盗宝藏赠送"
		ItemArray = AddItemArray(101, 15, ItemArray)
	case 10127: // 金币
		getScore = 2400000
		title = "充值海盗王宝藏（首次）赠送"
		ItemArray = AddItemArray(101, 35, ItemArray)
	case 10128: // 金币
		getScore = 1200000
		title = "充值海盗王宝藏赠送"
		ItemArray = AddItemArray(101, 35, ItemArray)
	case 10129: //   灵力
		getCoin = 3000000
		title = "充值超值礼包赠送"
		ItemArray = AddItemArray(7003, 50, ItemArray)
	case 10130: //   灵力
		getCoin = 13000000
		title = "充值豪华超值礼包赠送"
		ItemArray = AddItemArray(7003, 350, ItemArray)
	case 11001: // 金币
		getScore = 600000
		title = "充值初级强化石礼包赠送"
		ItemArray = AddItemArray(2007, 1288, ItemArray)
		ItemArray = AddItemArray(101, 30, ItemArray)
	case 11002: // 金币
		getScore = 4000000
		title = "充值中级强化石礼包赠送"
		ItemArray = AddItemArray(2007, 7500, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)
	case 11003: // 金币
		getScore = 16000000
		title = "充值高级强化石礼包赠送"
		ItemArray = AddItemArray(2007, 28000, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)
	case 11004: // 金币
		getScore = 400000
		title = "充值6元超值礼包赠送"
		ItemArray = AddItemArray(2007, 1000, ItemArray)
		ItemArray = AddItemArray(101, 20, ItemArray)
		ItemArray = AddItemArray(102, 10, ItemArray)
	case 13001: // 金币
		getScore = 300000 * rate
		title = "充值初级超值金币礼包赠送"
	case 13002: // 金币
		getScore = 1000000 * rate
		title = "充值中级超值金币礼包赠送"
	case 13003: // 金币
		getScore = 4800000 * rate
		title = "充值高级超值金币礼包赠送"
	case 14001: // 金币
		getScore = 8000000
		title = "充值追击海神礼包赠送"
		ItemArray = AddItemArray(120, 28, ItemArray)
	case 14002: // 金币
	case 15005: // 金币
	case 15006: // 金币
	case 15007: // 金币
	case 15008: // 金币
	case 16005: // 只有道具

	}

	if getScore > 0 {
		// 插入充值金币
		addScoreSql := GetScoreRechargeSql(rechargeInfo, getScore, dbNow, dataTimeStr, dbName, day1, Type, SubType, title, rechargeInfo.gitPackageId)
		zLog.PrintfLogger("礼包 %d 插入充值金币语句 %s", rechargeInfo.gitPackageId, addScoreSql)
	}
	if getDiamond > 0 {
		// 插入充值钻石语句
		addDiamondSql := GetDiamondRechargeSql(rechargeInfo, getDiamond, dbNow, dataTimeStr, dbName, day1, Type, SubType, title, rechargeInfo.gitPackageId)
		zLog.PrintfLogger("礼包 %d 插入充值钻石语句 %s", rechargeInfo.gitPackageId, addDiamondSql)
	}
	if getCoin > 0 {
		// 插入灵力
		addCoinSql := GetCoinRechargeSql(rechargeInfo, getCoin, dbNow, dataTimeStr, dbName, day1, Type, SubType, title)
		zLog.PrintfLogger("礼包 %d 插入充值灵力语句 %s", rechargeInfo.gitPackageId, addCoinSql)
	}
	for _, item := range ItemArray {
		fmt.Println(item.ItemId, "道具", item.ItemNum)
		addItemSql := GetItemRechargeSql(rechargeInfo, item.ItemId, item.ItemNum, dbNow, dataTimeStr, dbName, day1, title)
		zLog.PrintfLogger("礼包 %d 插入充值道具语句 %s", rechargeInfo.gitPackageId, addItemSql)
	}

}
