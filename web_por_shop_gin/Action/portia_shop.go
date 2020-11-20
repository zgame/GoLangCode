package Action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web_gin/MySql"
)
// help
func GetShopHelp(c *gin.Context)  {
	c.JSON(200, gin.H{
		"/portia_shop/help": "帮助页面",
		//"/portia_shop/user": "查询用户页面，暂时没用",
		"/portia_shop/recharge?openid=***": "查询玩家一共充值多少",
		"/portia_shop/buy_list?openid=***": "查询玩家购买道具列表",
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

// 获取购买列表
func GetUserBuyList(c *gin.Context) {
	openId := c.Query("openid") // 获取get的参数
	if openId == "" {
		Error(fmt.Sprintf("输入参数openid为空 %s", openId),c)
		return
	}
	ShopList := MySql.GetUserShopList(openId)
	c.JSON(200, gin.H{"openid":  openId, "ShopList": ShopList})
}
