package Action

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"web_gin/MiddleWare/zLog"
	"web_gin/MySql"
)
// help
func GetShopHelp(c *gin.Context)  {
	c.JSON(200, gin.H{
		"https://shop.portia.xyz:8097/portia_shop/help": "帮助页面",
		//"/portia_shop/user": "查询用户页面，暂时没用",
		//"/portia_shop/recharge?openid=***": "查询玩家一共充值多少",
		"https://shop.portia.xyz:8097/portia_shop/buy_list?openid=***": "查询玩家购买道具列表",
		"https://shop.portia.xyz:8097/portia_shop/mall_list": "查询商城道具列表",
		"https://shop.portia.xyz:8097/portia_shop/control_list": "查询商城控制列表",
		"https://shop.portia.xyz:8097/portia_shop/zswpay": "测试用充值接口，上线去掉，不能对外",
	})
}


// test
func GetUserList(c *gin.Context) {
	openId := c.Query("openid") // 获取get的参数
	if openId == "" {
		Error(fmt.Sprintf("输入参数openid为空 %s", openId),c)
		return
	}

	user := MySql.GetUserInfoData(openId)
	if user == nil {
		Error(fmt.Sprintf("输入参数openid : %s 找不到", openId),c)
		return
	}


	c.JSON(200, gin.H{"Uid": user.Uid})
}

// test获取充值数据
func GetUserRechargeList(c *gin.Context) {
	openId := c.Query("openid") // 获取get的参数
	if openId == "" {
		Error(fmt.Sprintf("输入参数openid为空 %s", openId),c)
		return
	}

	result := MySql.GetRechargeData(openId)
	//MySql.UpdateAllItems(openId)

	c.JSON(200, gin.H{"rmb": result.Rmb})
}

// 获取已购买列表
func GetUserBuyList(c *gin.Context) {
	//ItemCanBuy("zsw222", 12)

	openId := c.Query("openid") // 获取get的参数
	if openId == "" {
		Error(fmt.Sprintf("输入参数openid为空 %s", openId),c)
		return
	}
	ShopList := MySql.GetUserShopList(openId)
	c.JSON(200, gin.H{"openid":  openId, "ShopList": ShopList})
}

// 获取商城列表
func GetUserMallList(c *gin.Context) {
	mallList := MySql.GetMallInfoData()
	list,err := json.Marshal(mallList)
	if err!= nil{
		zLog.PrintfLogger("获取商城列表错误 %s \n",err.Error())
	}
	result := string(list)
	//fmt.Println(result)


	// zip 压缩
	//var in bytes.Buffer
	//w,err := zlib.NewWriterLevel(&in,zlib.DefaultCompression)
	//_,err = w.Write([]byte(result))
	//err =   w.Close()
	//if err!=nil {
	//	zLog.PrintfLogger("压缩错误 %s \n ", err.Error())
	//}

	c.JSON(200, gin.H{"MallList": result, "bytes": ""})
}

// 获取商城控制列表
func GetControlMallList(c *gin.Context) {
	controlList := MySql.GetMallControlData()
	list,err := json.Marshal(controlList)
	if err!= nil{
		zLog.PrintfLogger("获取商城控制列表错误 %s \n",err.Error())
	}
	result := string(list)
	//fmt.Println(result)

	c.JSON(200, gin.H{"ControlList": result, "bytes": ""})
}



// 默认充值成功的接口
func ZswPay(c * gin.Context)  {
	OpenId := c.PostForm("OpenId") // 获取post的参数
	ItemId ,_ := strconv.Atoi(c.PostForm("ItemId"))
	_, _, err := CheckItemId(c)
	if err {
		return
	}
	// 然后保存数据库并发放道具
	SaveDataBase(&MySql.Recharge{Openid: OpenId, Payno: "假的订单，测试用", RechargeTime: time.Now().String(), Rmb: "0", ItemId: ItemId, Channel: "测试"})

	// 返回用户的道具列表
	ShopList := MySql.GetUserShopList(OpenId)
	fmt.Println("玩家道具列表：",ShopList)
	c.JSON(200, gin.H{"openid": OpenId, "ShopList": ShopList})
}