package aliPay

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay"
	"net/http"
	"strings"
	"web_gin/GlobalVar"
)

var (
	appID = ""
	// RSA2(SHA256)
)
var client *alipay.AliPay


// 初始化
func Init() {

	PublicKey := strings.ReplaceAll(GlobalVar.PublicKey,"\n","")
	PrivateKey := strings.ReplaceAll(GlobalVar.PrivateKey,"\n","")
	client = alipay.New(appID, PublicKey, PrivateKey, false)
}


// 回调
func AliPayCallBack(c *gin.Context) {
	var req *http.Request
	req = c.Request
	ok, err := client.VerifySign(req.Form)	//验签
	fmt.Println(ok, err)
	if err != nil {
		c.JSON(200, "VerifySign failed")
		return
	}
	//查询是否重复


	// 判断金额是否一致，防止1分钱充值


	// 保存数据库


	//返回成功
	c.JSON(200, "success")
}
