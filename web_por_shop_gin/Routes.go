package main

import (
	"github.com/gin-gonic/gin"
	"web_gin/Action"
	"web_gin/MiddleWare/aliPay"
	"web_gin/MiddleWare/wxPay"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", Action.Ping)
	r.GET("/", Action.Welcome)

	r.GET("/get", Action.Get)
	//r.GET("/user/login", Action.Login)


	r.POST("/user/login", Action.Login)
	r.GET("/user/info", Action.Info)
	r.POST("/user/logout", Action.Logout)

	r.GET("/cookie", Action.Cookie)


	// ------------------ portia shop -------------------
	r.GET("/portia_shop/help", Action.GetShopHelp)
	r.GET("/portia_shop/user", Action.GetUserList)
	//r.GET("/portia_shop/recharge", Action.GetUserRechargeList)
	r.GET("/portia_shop/buy_list", Action.GetUserBuyList)
	r.GET("/portia_shop/mall_list", Action.GetUserMallList)

	// --------------------- ali wx---------------------------
	r.POST("/portia_shop/alipayget", aliPay.GetPayInfo)      //客户端获取订单信息
	r.POST("/portia_shop/alipaysign", aliPay.ClientGetSign) //客户端同步回调验证订单信息
	r.POST("/portia_shop/alipay", aliPay.CallBack)          // 支付宝异步回调

	r.GET("/portia_shop/wxpayget", wxPay.GetPayInfo)
	r.POST("/portia_shop/wxpayget", wxPay.GetPayInfo)
	r.POST("/portia_shop/wxpaysign", wxPay.ClientGetSign)
	r.POST("/portia_shop/wxpay", wxPay.WxPayCallBack)
}
