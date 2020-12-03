package Logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"web_gin/Action"
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
	// 判断是否重复购买
	for _, v := range arrDB {
		if v == itemId {
			return false
		}
	}
	return true
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
	startTime, _ := time.ParseInLocation("2006-01-02", start, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02", end, time.Local)
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
		Action.Error("OpenId不能为空", c)
		return 0, "", true
	}
	if ItemId == "" {
		Action.Error("ItemId不能为空", c)
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
		Action.Error("道具id不合法", c)
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
			Action.Error("道具重复购买", c)
			return 0, "", true
		}
	}
	return ItemPrice, string(myData), false
}
