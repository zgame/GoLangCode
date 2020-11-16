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
	r.GET("/portia_shop/recharge", Action.GetUserRechargeList)
	r.GET("/portia_shop/buy_list", Action.GetUserBuyList)

	// --------------------- ali wx---------------------------
	r.GET("/portia_shop/alipay", aliPay.AliPayCallBack)
	r.GET("/portia_shop/wxpay", wxPay.WxPayCallBack)
}
