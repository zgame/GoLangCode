package Logic

import (
	"encoding/json"
	"fmt"
	"web_gin/MiddleWare/zLog"
	"web_gin/MySql"
)

// 保存数据库
func SaveDataBase(insertData *MySql.Recharge) string {
	// 判断金额是否一致，防止1分钱充值

	// 保存数据库
	if MySql.InsertRechargeData(insertData) == false {
		return MySql.GetUserShopList(insertData.Openid)
	}
	// 发放道具
	return GiveItemToUser(insertData.Openid, insertData.ItemId)

}

//道具发放
func GiveItemToUser(Openid string, ItemId int) string {
	// 道具发放
	ShopList := MySql.GetUserShopList(Openid)
	if ShopList == "" {
		// 说明是新增一条玩家数据
		str := fmt.Sprintf("[%d]", ItemId)
		uItem := &MySql.Useritem{Openid: Openid, ShopList: str}
		if MySql.InsertUserItemData(uItem) {
			zLog.PrintLogger("道具发放成功")
			return str
		} else {
			return ""
		}
	} else {
		// 取出老的数据
		var arrItem []int = make([]int, 0)
		var arrDB []int
		err := json.Unmarshal([]byte(ShopList), &arrDB)
		if err != nil {
			zLog.PrintfLogger("GiveItemToUser   Id: %s 发放道具解析已有数据出错 %s \n", Openid, err)
			return ""
		}
		// 加上新的数据
		arrItem = append(arrItem, arrDB...)
		arrItem = append(arrItem, ItemId)
		//数据排重
		var strMap map[int]string = make(map[int]string, 0)
		for _, v := range arrItem {
			strMap[v] = "true"
		}
		var shopList []int = make([]int, 0)
		for k, _ := range strMap {
			shopList = append(shopList, k)
		}

		// 转换成字符串
		str, _ := json.Marshal(shopList)
		// 保存
		userItem := &MySql.Useritem{Openid: Openid, ShopList: string(str)}
		user := &MySql.Useritem{Openid: Openid}
		if MySql.UpdateUserItemData(userItem, user) {
			zLog.PrintLogger("道具发放成功")
			return string(str)
		} else {
			return ""
		}

	}
}

// 道具是否已经购买
func ItemCanBuy(Openid string, itemId int) bool {
	ShopList := MySql.GetUserShopList(Openid)
	var arrDB []int
	err := json.Unmarshal([]byte(ShopList), &arrDB)
	if err != nil {
		zLog.PrintfLogger("GiveItemToUser   Id: %s 已有道具解析数据出错 %s \n", Openid, err)
		return false
	}
	for _, v := range arrDB {
		if v == itemId {
			return false
		}
	}
	return true
}
