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
	"web_gin/Logic"
	"web_gin/MiddleWare/zLog"
	"web_gin/MySql"
)

var (
	//appID = "2021002109699800"  // portia
	appID = "2016110200785785" // sandbox appId
)
var client *alipay.AliPay

// 初始化
func Init() {
	// 去掉公钥，私钥的换行符号
	PublicKey := strings.ReplaceAll(GlobalVar.AliPublicKey, "\n", "")
	PrivateKey := strings.ReplaceAll(GlobalVar.PrivateKey, "\n", "")

	client = alipay.New(appID, PublicKey, PrivateKey, true)
}

// 拉起订单
func GetPayInfo(c *gin.Context) {

	ItemPrice, myData, err := Logic.CheckItemId(c)
	if err {
		return
	}

	zLog.PrintfLogger("================Get ali PayInfo 拉起订单 ================= ")
	var p = alipay.AliPayTradeAppPay{}
	p.NotifyURL = GlobalVar.MyUrl + "portia_shop/alipay"
	p.Subject = "购买道具1分钱"
	p.OutTradeNo = "" + strconv.FormatInt(time.Now().UnixNano(), 10) // 后面增加渠道编号
	//p.TotalAmount = "0.01"
	p.TotalAmount = fmt.Sprintf("%.2f",ItemPrice)
	p.ProductCode = "QUICK_MSECURITY_PAY" // 固定值

	//fmt.Println("myData: ", string(myData))
	p.PassbackParams = myData // 扩展

	url, err2 := client.TradeAppPay(p)
	if err2 != nil {
		fmt.Println(err)
	}
	//zLog.PrintfLogger("生成的订单信息： %s" ,url)
	//zLog.PrintLogger("============================================")
	//c.JSON(200, gin.H{"aliPayUrl":url})
	c.String(200, fmt.Sprintf("{\"aliPayUrl\":\"%s\"}", url))
}


// 异步回调
func CallBack(c *gin.Context) {
	zLog.PrintLogger("================CallBack 异步回调=================")
	var req *http.Request
	req = c.Request
	c.Request.ParseForm()

	var notification, _ = client.GetTradeNotification(req)
	if notification != nil {
		zLog.PrintLogger("交易状态为:" + notification.TradeStatus)
	} else {
		fmt.Println("交易状态为: 验证失败")
		c.JSON(200, "VerifySign failed")
		return
	}

	// 验签成功， 解析我们自己的传输格式
	var pInfo GlobalVar.PayInfo
	err := json.Unmarshal([]byte(notification.PassbackParams), &pInfo)
	if err != nil {
		c.JSON(200, "PayInfo Error")
		return
	}
	// 然后保存数据库并发放道具
	Logic.SaveDataBase(&MySql.Recharge{Openid: pInfo.OpenId, Payno: notification.TradeNo, RechargeTime: notification.NotifyTime, Rmb: notification.TotalAmount, ItemId: pInfo.ItemId, Channel: "ali"})

	//返回成功
	c.JSON(200, "success")
}
