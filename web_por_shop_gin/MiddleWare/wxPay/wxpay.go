package wxPay

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/wxpay"
	"net/http"
	"strconv"
	"time"
	"web_gin/Action"
	"web_gin/GlobalVar"
	"web_gin/Logic"
	"web_gin/MiddleWare/zLog"
	"web_gin/MySql"
)

var client * wxpay.WXPay
//var client2 * wxpay2.Client
var (
	appID = "wx31d6f1e3483da01e"
	apiKey = "d77a1fa0395f2b40b63639f9ec00c8da"
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
	var OpenId string
	var ItemId string
	OpenId = c.PostForm("OpenId") // 获取get的参数
	ItemId = c.PostForm("ItemId") // 获取get的参数
	if OpenId == "" {
		OpenId = "sdsd"
		ItemId = "33"
	}
	itemId, _ := strconv.Atoi(ItemId)
	// 增加道具是否购买重复的验证
	if false {
		if Logic.ItemCanBuy(OpenId, itemId) == false {
			zLog.PrintLogger("道具重复购买")
			Action.Error("道具重复购买", c)
			return
		}
	}
	// 生成我们自己的信息传输
	pInfo := &GlobalVar.PayInfo{OpenId: OpenId, ItemId: itemId}
	data, _ := json.Marshal(pInfo)


	fmt.Println("========== Get wxPay Info ==========")
	var p = wxpay.UnifiedOrderParam{}
	p.Body = "测试充值"
	p.NotifyURL = GlobalVar.MyUrl + "portia_shop/wxpay"
	p.TradeType = wxpay.K_TRADE_TYPE_APP
	p.SpbillCreateIP = c.ClientIP()
	p.TotalFee = 101   		// 单位1分钱
	p.OutTradeNo = "" + strconv.FormatInt(time.Now().UnixNano(), 10) // 后面增加渠道编号

	p.Attach = string(data) // 扩展


	result, err := client.UnifiedOrder(p)
	if err != nil {
		fmt.Println("微信服务器返回错误："+err.Error())
	}

	//rmb := strconv.FormatFloat(0.01*2,'E',-1,64)
	//rmb := fmt.Sprintf("%.2f", 0.22)
	//
	//fmt.Println(rmb)

	//url,_ := json.Marshal(result)
	//fmt.Printf("%s \n",url)
	//c.String(200, fmt.Sprintf("{\"wxPayUrl\":\"%s\"}", url))
	c.JSON(200, result)


	////// 统一下单
	//params := make(wxpay2.Params)
	//params.SetString("body", "test").
	//	SetString("out_trade_no", "436577857").
	//	SetInt64("total_fee", 1).
	//	SetString("spbill_create_ip", "127.0.0.1").
	//	SetString("notify_url", "http://notify.objcoding.com/notify").
	//	SetString("trade_type", "APP")
	//pp, _ := client2.UnifiedOrder(params)
	//jj,_ := json.MarshalIndent(pp,""," ")
	//fmt.Printf("jj %s \n",jj)


}



// 回调
func WxPayCallBack(c *gin.Context) {
	fmt.Println("-----------------------微信支付回调-----------------")
	var req *http.Request
	req = c.Request
	notification, err := client.GetTradeNotification(req)
	if notification != nil {
		fmt.Println("交易状态为:", notification.ResultCode)
	}
	if err != nil{
		c.XML(200, gin.H{"return_code":"Failed", "return_msg": "Not"})
		return
	}
	fmt.Println(notification, err)



	//验签成功， 解析我们自己的传输格式
	var pInfo GlobalVar.PayInfo
	err = json.Unmarshal([]byte(notification.Attach), &pInfo)
	if err != nil {
		c.JSON(200, "PayInfo Error")
		return
	}
	// 然后保存数据库并发放道具
	rmb := fmt.Sprintf("%.2f", float64(notification.TotalFee) * 0.01)
	Logic.SaveDataBase(&MySql.Recharge{Openid: pInfo.OpenId, Payno: notification.TransactionId, RechargeTime: notification.TimeEnd, Rmb: rmb, ItemId: pInfo.ItemId, Channel: "wx"})



	//返回成功
	c.XML(200, gin.H{"return_code":"SUCCESS", "return_msg": "OK"})
}

