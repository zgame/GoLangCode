package main

import (
	"database/sql"
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
func GetGiftPackageRechargeSql(rechargeInfo RechargeList, dbNow *sql.DB, dataTimeStr string, dbName string, day1 string, DataBaseBYDB *sql.DB ,  TestDb *sql.DB) {
	// 变量初始化
	getScore := 0   // 金币
	getDiamond := 0 // 钻石
	getCoin := 0    // 灵力
	title := ""     //原因
	Type := 2
	SubType := 3 // 子分类
	ItemArray := make([]GiftItem, 0)
	rate := ZRandomTo(3,10)		// 倍率

	switch rechargeInfo.gitPackageId {
	case 2: // 金币 钻石
		getScore = 100000
		getDiamond = 10
		title = "充值青铜礼包赠送"
		// 还差道具的
		ItemArray = AddItemArray(108, 1, ItemArray)
		ItemArray = AddItemArray(101, 10, ItemArray)

	case 3: // 金币 钻石
		getScore = 200000
		getDiamond = 20
		title = "充值白银礼包赠送"
		ItemArray = AddItemArray(109, 1, ItemArray)
		ItemArray = AddItemArray(120, 5, ItemArray)

	case 4: // 金币 钻石
		getScore = 300000
		getDiamond = 50
		title = "充值黄金礼包赠送"
		ItemArray = AddItemArray(110, 1, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)

	case 5: // 金币 钻石
		getScore = 500000
		getDiamond = 100
		title = "充值白金礼包赠送"
		ItemArray = AddItemArray(111, 1, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)

	case 101: //金币 钻石
		getScore = 800000
		getDiamond = 100
		title = "充值至尊礼包月卡赠送"

	case 901: //金币 钻石
		getScore = 150000 * rate
		title = "充值首充返利礼包赠送"

	case 1001: //金币 钻石
		getScore = 20000
		title = "充值挑战礼包赠送"
		ItemArray = AddItemArray(150, 1, ItemArray)

	case 3001: // 只有道具
		title = "充值红钻礼包赠送"
		ItemArray = AddItemArray(120, 10, ItemArray)

	case 3101: // 只有道具
	case 3102: // 只有道具
	case 3103: // 只有道具
	case 3104: // 只有道具
	case 3105: // 只有道具

	case 10001: // 只有道具
		getScore = 400000
		title = "充值满贯起航礼包赠送"

	case 10002: // 只有道具
		getScore = 120000
		getDiamond = 60
		title = "充值每日成长礼包赠送"

	case 10003: // 只有道具
		title = "充值满贯进阶礼包赠送"
		getScore = 2000000
		getDiamond = 200
		ItemArray = AddItemArray(120, 5, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)

	case 10004: // 只有道具
		title = "充值每日导弹达人礼包赠送"
		ItemArray = AddItemArray(108, 2, ItemArray)
		ItemArray = AddItemArray(109, 1, ItemArray)
		ItemArray = AddItemArray(111, 2, ItemArray)
		//Type = 3
		//SubType = 13
	case 10005: // 金币
		getScore = 300000
		title = "充值新起航礼包赠送"
		ItemArray = AddItemArray(2007, 180, ItemArray)
		ItemArray = AddItemArray(108, 1, ItemArray)

	case 10006: //  灵力
		//getCoin = 2400000
		title = "充值妖精场每日58元礼包1赠送"
		ItemArray = AddItemArray(111, 1, ItemArray)
		ItemArray = AddItemArray(120, 20, ItemArray)
		ItemArray = AddItemArray(7003, 20, ItemArray)
	case 10007: //  灵力
		//getCoin = 3000000
		title = "充值妖精场每日58元礼包2赠送"
		ItemArray = AddItemArray(111, 1, ItemArray)
		ItemArray = AddItemArray(120, 30, ItemArray)
		ItemArray = AddItemArray(7003, 30, ItemArray)
	case 10008: //  灵力
		//getCoin = 12000000
		title = "充值妖精场每日588元礼包1赠送"
		ItemArray = AddItemArray(7003, 300, ItemArray)
		ItemArray = AddItemArray(111, 5, ItemArray)
	case 10009: //  灵力
		//getCoin = 13000000
		title = "充值妖精场每日588元礼包2赠送"
		ItemArray = AddItemArray(7003, 350, ItemArray)
		ItemArray = AddItemArray(111, 6, ItemArray)
	case 10010: // 金币
		getScore = 120000
		title = "充值每日6元礼包1赠送"
		ItemArray = AddItemArray(2007, 30, ItemArray)
		ItemArray = AddItemArray(101, 5, ItemArray)
		ItemArray = AddItemArray(102, 2, ItemArray)
	case 10011: // 金币
		getScore = 150000
		title = "充值每日6元礼包2赠送"
		ItemArray = AddItemArray(2007, 50, ItemArray)
		ItemArray = AddItemArray(101, 7, ItemArray)
		ItemArray = AddItemArray(102, 3, ItemArray)
	case 10012: // 金币
		getScore = 180000
		title = "充值每日6元礼包3赠送"
		ItemArray = AddItemArray(2007, 100, ItemArray)
		ItemArray = AddItemArray(101, 10, ItemArray)
		ItemArray = AddItemArray(102, 5, ItemArray)
	case 10013: // 金币
		getScore = 330000
		title = "充值每日30元礼包1赠送"
		ItemArray = AddItemArray(2007, 150, ItemArray)
		ItemArray = AddItemArray(101, 20, ItemArray)
		ItemArray = AddItemArray(102, 10, ItemArray)
		ItemArray = AddItemArray(108, 1, ItemArray)
	case 10014: // 金币
		getScore = 390000
		title = "充值每日30元礼包2赠送"
		ItemArray = AddItemArray(2007, 180, ItemArray)
		ItemArray = AddItemArray(101, 30, ItemArray)
		ItemArray = AddItemArray(102, 15, ItemArray)
		ItemArray = AddItemArray(108, 1, ItemArray)
	case 10015: // 金币
		getScore = 450000
		title = "充值每日30元礼包3赠送"
		ItemArray = AddItemArray(2007, 210, ItemArray)
		ItemArray = AddItemArray(101, 40, ItemArray)
		ItemArray = AddItemArray(102, 20, ItemArray)
		ItemArray = AddItemArray(108, 2, ItemArray)





	case 10016: // 金币  灵力
		getScore = 1000000
		getCoin = 50000
		title = "充值每日98元礼包1赠送"
		ItemArray = AddItemArray(2007, 300, ItemArray)
		ItemArray = AddItemArray(120, 5, ItemArray)
		ItemArray = AddItemArray(109, 1, ItemArray)
	case 10017: // 金币  灵力
		getScore = 1100000
		getCoin = 60000
		title = "充值每日98元礼包2赠送"
		ItemArray = AddItemArray(2007, 350, ItemArray)
		ItemArray = AddItemArray(120, 7, ItemArray)
		ItemArray = AddItemArray(109, 1, ItemArray)
	case 10018: // 金币  灵力
		getScore = 1300000
		getCoin = 80000
		title = "充值每日98元礼包3赠送"
		ItemArray = AddItemArray(2007, 400, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)
		ItemArray = AddItemArray(109, 2, ItemArray)
	case 10019: // 金币  灵力
		getScore = 2000000
		getCoin = 100000
		title = "充值每日198元礼包1赠送"
		ItemArray = AddItemArray(2007, 500, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)
		ItemArray = AddItemArray(131, 2, ItemArray)
		ItemArray = AddItemArray(110, 1, ItemArray)
	case 10020: // 金币  灵力
		getScore = 2200000
		getCoin = 120000
		title = "充值每日198元礼包2赠送"
		ItemArray = AddItemArray(2007, 550, ItemArray)
		ItemArray = AddItemArray(120, 15, ItemArray)
		ItemArray = AddItemArray(131, 3, ItemArray)
		ItemArray = AddItemArray(110, 1, ItemArray)
	case 10021: // 金币  灵力
		getScore = 2500000
		getCoin = 150000
		title = "充值每日198元礼包3赠送"
		ItemArray = AddItemArray(2007, 600, ItemArray)
		ItemArray = AddItemArray(120, 20, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)
		ItemArray = AddItemArray(110, 2, ItemArray)
	case 10022: // 金币 钻石  灵力
		getScore = 6660000
		getDiamond = 666
		getCoin = 300000
		title = "充值每日648元礼包1赠送"
		ItemArray = AddItemArray(111, 1, ItemArray)
		ItemArray = AddItemArray(120, 20, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)
	case 10023: // 金币 钻石  灵力
		getScore = 7770000
		getDiamond = 777
		getCoin = 400000
		title = "充值每日648元礼包2赠送"
		ItemArray = AddItemArray(111, 1, ItemArray)
		ItemArray = AddItemArray(120, 30, ItemArray)
		ItemArray = AddItemArray(131, 6, ItemArray)
	case 10024: // 金币 钻石  灵力
		getScore = 8880000
		getDiamond = 888
		getCoin = 500000
		title = "充值每日648元礼包3赠送"
		ItemArray = AddItemArray(111, 2, ItemArray)
		ItemArray = AddItemArray(120, 50, ItemArray)
		ItemArray = AddItemArray(131, 8, ItemArray)

	case 11001: // 金币
		getScore = 180000
		title = "充值强化石礼包1赠送"
		ItemArray = AddItemArray(2007, 666, ItemArray)
		ItemArray = AddItemArray(101, 30, ItemArray)
	case 11002: // 金币
		getScore = 1080000
		title = "充值强化石礼包2赠送"
		ItemArray = AddItemArray(2007, 3888, ItemArray)
		ItemArray = AddItemArray(120, 10, ItemArray)
	case 11003: // 金币
		getScore = 3280000
		title = "充值强化石礼包3赠送"
		ItemArray = AddItemArray(2007, 13888, ItemArray)
		ItemArray = AddItemArray(131, 4, ItemArray)

	case 13001: // 金币
		getScore = 150000 * rate
		title = "充值金币超值礼包赠送"
	}

	if getScore > 0 {
		// 插入充值金币
		//zLog.PrintfLogger("-----------------礼包 %d 插入金币 %d", rechargeInfo.gitPackageId , getScore)
		lastAllScore:=GetScoreRechargeSql(rechargeInfo, getScore, dbNow, dataTimeStr, dbName, day1, Type, SubType, title, rechargeInfo.gitPackageId, DataBaseBYDB,  TestDb,0)
		GetScoreReduceSql(rechargeInfo, getScore, dbNow, dataTimeStr, dbName, day1,lastAllScore,0)
		//zLog.PrintfLogger("礼包 %d 插入充值金币语句 ", rechargeInfo.gitPackageId)
	}
	if getDiamond > 0 {
		// 插入充值钻石语句
		//zLog.PrintfLogger("-----------------礼包 %d 插入钻石 %d", rechargeInfo.gitPackageId , getDiamond)
		lastAllDiamond:=GetDiamondRechargeSql(rechargeInfo, getDiamond, dbNow, dataTimeStr, dbName, day1, Type, SubType, title, rechargeInfo.gitPackageId, DataBaseBYDB,  TestDb)
		GetDiamondReduceSql(rechargeInfo, getDiamond, dbNow, dataTimeStr, dbName, day1,lastAllDiamond)
		//zLog.PrintfLogger("礼包 %d 插入充值钻石语句 %s", rechargeInfo.gitPackageId, addDiamondSql)
	}
	if getCoin > 0 {
		// 插入灵力
		//zLog.PrintfLogger("-----------------礼包 %d 插入灵力 %d", rechargeInfo.gitPackageId , getCoin)
		lastAllCoin:=GetCoinRechargeSql(rechargeInfo, getCoin, dbNow, dataTimeStr, dbName, day1, Type, SubType, title, DataBaseBYDB,  TestDb)
		GetCoinReduceSql(rechargeInfo, getCoin, dbNow, dataTimeStr, dbName, day1,lastAllCoin)

	}
	for _, item := range ItemArray {
		// 插入道具
		//zLog.PrintfLogger("-----------------礼包 %d 插入道具 %d 数量 %d", rechargeInfo.gitPackageId , item.ItemId, item.ItemNum)
		lastAllItem:=GetItemRechargeSql(rechargeInfo, item.ItemId, item.ItemNum, dbNow, dataTimeStr, dbName, day1, title, DataBaseBYDB,  TestDb)
		GetItemReduceSql(rechargeInfo, item.ItemId, item.ItemNum, dbNow, dataTimeStr, dbName, day1,lastAllItem)
		if item.ItemId >=108 && item.ItemId <=111 {

			// 弹头产生金币消耗
			BombScore:= 0
			switch item.ItemId {
			case 108:
				BombScore = ZRandomTo(90000,110000)
			case 109:
				BombScore = ZRandomTo(430000,570000)
			case 110:
				BombScore = ZRandomTo(5500000,6500000)
			case 111:
				BombScore = ZRandomTo(13000000,18000000)
			}
			//zLog.PrintfLogger("-----------------因为包含弹头 %d，所以金币有所变化 %d", item.ItemId, BombScore)
			lastAllScore:= GetScoreRechargeSql(rechargeInfo, BombScore, dbNow, dataTimeStr, dbName, day1, 1, 1, "游戏写分", 10, DataBaseBYDB,  TestDb,5)
			GetScoreReduceSql(rechargeInfo, BombScore, dbNow, dataTimeStr, dbName, day1,lastAllScore,5)
		}
	}

}
