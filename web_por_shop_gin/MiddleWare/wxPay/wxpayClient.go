package wxPay

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/encoding"
	"web_gin/Action"
	"web_gin/MySql"

	//"sort"
	//"strings"
	//"web_gin/GlobalVar"
	"web_gin/MiddleWare/zLog"
)

// 客户端同步回调
func ClientGetSign(c *gin.Context) {
	zLog.PrintLogger("================ClientGetSign  客户端同步回调=================")
	//var req *http.Request
	//req = c.Request
	//req.ParseForm()

	openId := c.PostForm("OpenId")
	if openId == ""{
		Action.Error("OpenId 不能为空",c)
		return
	}
	//fmt.Println(clientForm)
	//var back ClientCallBack
	//var aliMsg ClientResult
	//
	////  客服端用 portia字段发来整个返回字符串给我
	//_ = json.Unmarshal([]byte(clientForm), &back)
	//// 解析出来sign
	//_ = json.Unmarshal([]byte(back.Result), &aliMsg)




	//
	//var p = wxpay.OrderQueryParam{}
	//p.TransactionId = "测试充值"
	//p.OutTradeNo = GlobalVar.MyUrl + "portia_shop/wxpay"
	//
	//result, err2 := client.OrderQuery(p)
	//if err2 != nil {
	//	//fmt.Println("微信服务器返回错误："+err2.Error())
	//	Action.Error("微信查询服务器返回错误："+err2.Error(),c)
	//	return
	//}




	//zLog.PrintLogger("===================同步验签成功=========================")
	// 查找道具信息
	//tradeNo := mapMsg["trade_no"]
	////fmt.Println("tradeNo", tradeNo)
	//recharge := MySql.GetRechargeData(tradeNo)
	//if recharge == nil {
	//	Action.Error("没有找到已充值的订单信息，请稍后再试", c)
	//	return
	//}
	// 返回用户的道具列表
	//openId := recharge.Openid
	ShopList := MySql.GetUserShopList(openId)
	fmt.Println("玩家道具列表：",ShopList)
	c.JSON(200, gin.H{"openid": openId, "ShopList": ShopList})

}

//----------------------------------------------------------------------------
//         同步回调的数据格式
//----------------------------------------------------------------------------
type ClientCallBack struct {
	ResultStatus string `json:"resultStatus"`
	Result       string `json:"result"`
	Memo         string `json:"memo"`
}
type ClientResult struct {
	AlipayTradeAppPayResponse map[string]string `json:"alipay_trade_app_pay_response"`
	Sign                      string            `json:"sign"`
	SignType                  string            `json:"sign_type"`
}

func verifyData(data []byte, sign string, key []byte) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}
	err = encoding.VerifyPKCS1v15(data, signBytes, key, crypto.SHA256)

	if err != nil {
		return false, err
	}
	return true, nil
}
