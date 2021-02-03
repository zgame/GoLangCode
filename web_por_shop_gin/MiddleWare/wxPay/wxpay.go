package wxPay

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/wxpay"
	"net/url"
	"web_gin/Action"
	"web_gin/MiddleWare/zLog"

	//wxpay2 "github.com/objcoding/wxpay"
	"net/http"
	"strconv"
	"time"
	"web_gin/GlobalVar"
	"web_gin/MySql"
)

var client * wxpay.WXPay
//var client2 * wxpay2.Client
var (
	appID = "wx31d6f1e3483da01e"
	apiKey = "eSf41sG9PMIlW6SheSf41sG9PMIlW6Sh"
	mchId = "1604563934"
	// RSA2(SHA256)
)


// 初始化
func Init() {
	//publicKey := strings.ReplaceAll(GlobalVar.AliPublicKey,"\n","")
	//privateKey := strings.ReplaceAll(GlobalVar.PrivateKey,"\n","")
	client = wxpay.New(appID, apiKey, mchId, true)
	//account1 := wxpay2.NewAccount(appID, mchId, apiKey, false)
	//client2 = wxpay2.NewClient(account1)
	//client2.SetSignType("HMACSHA256")
}


// 获取订单信息
func GetPayInfo(c *gin.Context)  {

	ItemPrice, myData, err := Action.CheckItemId(c)
	if err {
		return
	}


	fmt.Println("========== Get wxPay Info ==========")
	var p = wxpay.UnifiedOrderParam{}
	p.Body = "购买道具"
	p.NotifyURL = GlobalVar.MyUrl + "portia_shop/wxpay"
	p.TradeType = wxpay.K_TRADE_TYPE_APP
	p.SpbillCreateIP = c.ClientIP()
	p.TotalFee =  int(ItemPrice*100)   		// 单位1分钱
	p.OutTradeNo = "" + strconv.FormatInt(time.Now().UnixNano(), 10) // 后面增加渠道编号

	p.Attach = myData // 扩展


	result, err2 := client.UnifiedOrder(p)
	if err2 != nil {
		//fmt.Println("微信服务器返回错误："+err2.Error())
		Action.Error("微信服务器返回错误："+err2.Error(),c)
		return
	}


	// 下面获取一下微信给的信息，然后组成串，签名发给客户端
	var m = make(url.Values)
	m.Set("appid", appID)
	m.Set("partnerid", mchId)
	m.Set("prepayid", result.PrepayId)
	m.Set("noncestr", result.NonceStr)
	m.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	m.Set("package", "Sign=WXPay")
	var sign = wxpay.SignMD5(m, apiKey)
	m.Set("sign", sign)

	var re map[string]string = make(map[string]string,0)
	for k,v := range m {
		re[k]=v[0]
	}
	c.JSON(200, re)
}



// 回调
func WxPayCallBack(c *gin.Context) {
	fmt.Println("-----------------------微信支付回调-----------------")
	var req *http.Request
	req = c.Request
	notification, err := client.GetTradeNotification(req)
	if notification != nil {
		zLog.PrintLogger("交易状态为:"+ notification.ResultCode)
	}
	if err != nil{
		zLog.PrintLogger("交易状态为: 验证失败"+err.Error())
		c.XML(200, gin.H{"return_code":"Failed", "return_msg": "Not"})
		return
	}
	//fmt.Println(notification, err)
	//zLog.PrintLogger("===================异步验签成功=========================")
	//验签成功， 解析我们自己的传输格式
	var pInfo GlobalVar.PayInfo
	err = json.Unmarshal([]byte(notification.Attach), &pInfo)
	if err != nil {
		c.JSON(200, "PayInfo Error")
		return
	}
	// 然后保存数据库并发放道具
	rmb := fmt.Sprintf("%.2f", float64(notification.TotalFee) * 0.01)
	st, _ := time.Parse("20060102150405", notification.TimeEnd) //string转time
	ts := st.Format("2006-01-02 15:04:05") //time转string
	Action.SaveDataBase(&MySql.Recharge{Openid: pInfo.OpenId, Payno: notification.TransactionId, RechargeTime:ts , Rmb: rmb, ItemId: pInfo.ItemId, Channel: "wx"})


	//返回成功
	c.XML(200, gin.H{"return_code":"SUCCESS", "return_msg": "OK"})
}
