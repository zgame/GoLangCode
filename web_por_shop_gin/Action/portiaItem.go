package Action

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"web_gin/GlobalVar"
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
	var arrItem []int = make([]int, 0)
	var ItemInfo *MySql.Shopmall
	ItemInfo = MySql.GetMallItemInfo(ItemId)
	if ItemInfo.Gift == "" || ItemInfo.Gift == "-1" {
		arrItem = append(arrItem, ItemId) // 发放单一道具
	} else {
		//	发放礼包道具
		list := GiftGetList(ItemInfo.Gift)
		for _, v := range list {
			item, _ := strconv.Atoi(v)
			arrItem = append(arrItem, item) // 发放多个道具的礼包
		}
	}

	ShopList := MySql.GetUserShopList(Openid)
	if ShopList == "" {
		// 说明是新增一条玩家数据
		// 转换成字符串
		str, _ := json.Marshal(arrItem)
		//str := fmt.Sprintf("[%d]", ItemId)
		uItem := &MySql.Useritem{Openid: Openid, ShopList: string(str)}
		if MySql.InsertUserItemData(uItem) {
			zLog.PrintLogger("道具发放成功")
			return string(str)
		} else {
			return ""
		}
	} else {
		// 取出老的数据
		var arrDB []int
		err := json.Unmarshal([]byte(ShopList), &arrDB)
		if err != nil {
			zLog.PrintfLogger("GiveItemToUser   Id: %s 发放道具解析已有数据出错 %s \n", Openid, err)
			return ""
		}
		// 加上老的数据
		arrItem = append(arrItem, arrDB...)

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
		if MySql.UpdateUserItemData(userItem, Openid) {
			zLog.PrintLogger("道具发放成功")
			return string(str)
		} else {
			return ""
		}
	}
}

// 获取礼包的道具列表
func GiftGetList(gift string) []string {
	list := strings.Split(gift, "#")
	return list
}

// 道具是否已经购买
func ItemCanBuy(Openid string, ItemId int) bool {
	var arrBuyItem []int = make([]int, 0) // 这里保存要购买的所有的道具
	var ItemInfo *MySql.Shopmall
	ItemInfo = MySql.GetMallItemInfo(ItemId)
	if ItemInfo.Gift == "" || ItemInfo.Gift == "-1" {
		arrBuyItem = append(arrBuyItem, ItemId) // 发放单一道具
		fmt.Println("单一道具")
	} else {
		// 礼包道具
		list := GiftGetList(ItemInfo.Gift)
		for _, v := range list {
			item, _ := strconv.Atoi(v)
			arrBuyItem = append(arrBuyItem, item)
		}
		fmt.Println("礼包道具")
	}

	ShopList := MySql.GetUserShopList(Openid)
	if ShopList == "" {
		return true // 没有玩家数据
	}
	var arrDB []int
	err := json.Unmarshal([]byte(ShopList), &arrDB)
	if err != nil {
		zLog.PrintfLogger("GiveItemToUser   Id: %s 已有道具解析数据出错 %s \n", Openid, err)
		return false
	}

	// 判断是否重复购买
	for _, item := range arrBuyItem {		//遍历礼包或单一道具
		have := false
		for _, v := range arrDB {			// 遍历已有列表
			if v == item {
				have = true //已经有了这个道具，那么break循环
				// 已经有了就break
				break
			}
		}
		if have == false {
			return true // 只要一个道具还没有，就可以购买了
		}
	}
	return false
}

// 获取道具的价格
func GetItemPrice(item *MySql.Shopmall) float64 {
	price := item.Price
	outPrice := item.Discountprice
	start := item.Starttime
	end := item.Endtime

	// 不在活动时间
	if start == "-1" || start == "" || end == "-1" || end == "" || outPrice == -1 || outPrice == 0 {
		return float64(price)
	}

	// 字符串变时间
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", start, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end, time.Local)
	if startTime.Before(time.Now()) && endTime.After(time.Now()) {
		// 活动时间内
		return outPrice
	}

	return float64(price)
}

//  对道具的判断
func CheckItemId(c *gin.Context) (float64, string, bool) {
	var debug bool
	var OpenId string
	var ItemId string
	OpenId = c.PostForm("OpenId") // 获取post的参数
	ItemId = c.PostForm("ItemId") // 获取post的参数
	if OpenId == "" {
		OpenId = c.Query("OpenId") // 获取get的参数
		ItemId = c.Query("ItemId") // 获取get的参数
	}
	if OpenId == "" {
		Error("OpenId不能为空", c)
		return 0, "", true
	}
	if ItemId == "" {
		Error("ItemId不能为空", c)
		return 0, "", true
	}
	itemId, _ := strconv.Atoi(ItemId)

	// 生成我们自己的信息传输
	pInfo := &GlobalVar.PayInfo{OpenId: OpenId, ItemId: itemId}
	myData, _ := json.Marshal(pInfo)

	var ItemInfo *MySql.Shopmall
	var ItemPrice float64 // 道具价格
	ItemInfo = MySql.GetMallItemInfo(itemId)

	// 道具是否合法
	if ItemInfo == nil {
		Error("道具id不合法", c)
		return 0, "", true
	} else {
		ItemPrice = GetItemPrice(ItemInfo) //获取道具价格

		//ItemPrice = 0.01
		////fmt.Println("测试阶段，价格", ItemPrice)
		//debug = true
	}

	// 增加道具是否购买重复的验证
	if !debug {
		if ItemCanBuy(OpenId, itemId) == false {
			Error("道具重复购买", c)
			return 0, "", true
		}
	}
	return ItemPrice, string(myData), false
}
