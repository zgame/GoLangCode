package aliPay

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web_gin/GlobalVar"
	"web_gin/MySql"
)

var (
	//appID = "2021002109699800"  // portia
	appID = "2016110200785785"  // sandbox appId
	url = "https://openapi.alipaydev.com/gateway.do"  //
	// RSA2(SHA256)
)
var client *alipay.AliPay


// 初始化
func Init() {

	// 去掉公钥，私钥的换行符号
	PublicKey := strings.ReplaceAll(GlobalVar.PublicKey,"\n","")
	PrivateKey := strings.ReplaceAll(GlobalVar.PrivateKey,"\n","")
	//PublicKey := GlobalVar.PublicKey
	//PrivateKey := GlobalVar.PrivateKey

	//fmt.Println("PublicKey   "+PublicKey)
	//fmt.Println("PrivateKey   "+PrivateKey)
	//l7xugGu42G7E69L1/iLbz2xSA02RAW/R6l4Ki251Fx4WZZI7fA/ED3gQQHrwEyoYN0XjpV/SlboAqodx7GbyftyYK0mxspXscfiRRUVYN1aKe0zxlqc64BbPDA/ChIVLEQJ7fFtnYUP1x7/idVbybADoa0e2XYweE4qapCZ30hCsYNRXqaCQjqeEbtlSgubtcQ+QWPvJxixN4gQV59aVqXP7+enlKm6Yxr8rAW7oXBBGhpADeUx56Sz+tqWIXDEVmxdsqjA1YAP+Ddia+5NDJGFpqqPmp+HpN3JX6eBaTG4wIFeMJdtcJhTgkLA/Loq05LAD8FtOx8MjsedaYNQi3g==
	//l7xugGu42G7E69L1/iLbz2xSA02RAW/R6l4Ki251Fx4WZZI7fA/ED3gQQHrwEyoYN0XjpV/SlboAqodx7GbyftyYK0mxspXscfiRRUVYN1aKe0zxlqc64BbPDA/ChIVLEQJ7fFtnYUP1x7/idVbybADoa0e2XYweE4qapCZ30hCsYNRXqaCQjqeEbtlSgubtcQ+QWPvJxixN4gQV59aVqXP7+enlKm6Yxr8rAW7oXBBGhpADeUx56Sz+tqWIXDEVmxdsqjA1YAP+Ddia+5NDJGFpqqPmp+HpN3JX6eBaTG4wIFeMJdtcJhTgkLA/Loq05LAD8FtOx8MjsedaYNQi3g==


	client = alipay.New(appID, PublicKey, PrivateKey, false)
}

// 拉起订单
func AliPayGetNo(c *gin.Context) {

	var p = alipay.AliPayTradeAppPay{}
	//p.NotifyURL = "http://xxx"
	//p.ReturnURL = "http://xxx"


	p.Subject = "itemID"
	p.OutTradeNo = "" + strconv.FormatInt(time.Now().UnixNano(),10)	// 后面增加渠道编号
	p.TotalAmount = "1.00"
	p.ProductCode = "QUICK_MSECURITY_PAY"		// 固定值
	p.GoodsType = "0"


	ppp := &GlobalVar.PayInfo{OpenId:"sdsd", ItemId:33}
	data, _ := json.MarshalIndent(ppp, "", " ")
	p.PassbackParams = string(data) 		// 扩展

	url, err := client.TradeAppPay(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)
	//c.JSON(200, gin.H{"aliPayUrl":url})
	c.String(200, fmt.Sprintf("{\"aliPayUrl\":\"%s\"}",url))
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
	var insertData MySql.Recharge
	insertData.Rmb = 1
	insertData.Payno = "sdfsd"
	result := MySql.InsertRechargeData(&insertData)
	if !result {
		c.JSON(200, "save database failed")
		return
	}

	//返回成功
	c.JSON(200, "success")
}
