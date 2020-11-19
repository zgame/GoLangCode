package wxPay

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/wxpay"
	"net/http"
	"strings"
	"web_gin/GlobalVar"
)

var client * wxpay.WXPay
var (
	appID = ""
	// RSA2(SHA256)
)


// 初始化
func Init() {
	publicKey := strings.ReplaceAll(GlobalVar.AliPublicKey,"\n","")
	privateKey := strings.ReplaceAll(GlobalVar.PrivateKey,"\n","")
	client = wxpay.New(appID, publicKey, privateKey, false)
}


// 回调
func WxPayCallBack(c *gin.Context) {
	var req *http.Request
	req = c.Request
	result, err := client.GetTradeNotification(req)
	if result != nil {
		fmt.Println("交易状态为:", result.ResultCode)
	}
	if err != nil{
		c.XML(200, gin.H{"return_code":"Failed", "return_msg": "Not"})
		return
	}
	fmt.Println(result, err)

	//查询是否重复


	// 判断金额是否一致，防止1分钱充值

	// 保存数据库

	//返回成功
	c.XML(200, gin.H{"return_code":"SUCCESS", "return_msg": "OK"})
}

