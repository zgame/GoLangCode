package aliPay

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/encoding"
	"net/http"
	"strings"
	"web_gin/Action"
	"web_gin/GlobalVar"
	"web_gin/MySql"

	//"sort"
	//"strings"
	//"web_gin/GlobalVar"
	"web_gin/MiddleWare/zLog"
)

// 客户端同步回调
func ClientGetSign(c *gin.Context) {
	zLog.PrintLogger("================ClientGetSign  客户端同步回调=================")
	var req *http.Request
	req = c.Request
	req.ParseForm()

	clientForm := c.PostForm("portia")
	//fmt.Println(clientForm)
	var back ClientCallBack
	var aliMsg ClientResult

	//  客服端用 portia字段发来整个阿里字符串给我
	_ = json.Unmarshal([]byte(clientForm), &back)
	// 解析出来sign
	_ = json.Unmarshal([]byte(back.Result), &aliMsg)

	mapMsg := aliMsg.AlipayTradeAppPayResponse // 这里是阿里支付信息的map
	//fmt.Printf("aliMsg.AlipayTradeAppPayResponse        %v \n  ", mapMsg)

	// 用截取字符串的方式来获取关键的支付信息字符串用来验签
	end := strings.Index(back.Result, `,"sign":"`)
	msg := back.Result[33:end]

	// 开始验签
	_, err := verifyData([]byte(msg), aliMsg.Sign, encoding.FormatPublicKey(GlobalVar.AliPublicKey)) //验签
	//fmt.Println(ok, err)
	if err != nil {
		Action.Error(err.Error(),c)
		return
	}
	zLog.PrintLogger("===================同步验签成功=========================")
	// 查找道具信息
	tradeNo := mapMsg["trade_no"]
	//fmt.Println("tradeNo", tradeNo)
	recharge := MySql.GetRechargeData(tradeNo)
	if recharge == nil {
		Action.Error("没有找到已充值的订单信息，请稍后再试", c)
		return
	}
	// 返回用户的道具列表
	openId := recharge.Openid
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
