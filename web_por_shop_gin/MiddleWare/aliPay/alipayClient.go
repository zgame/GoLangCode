package aliPay

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/encoding"
	"net/http"
	"web_gin/GlobalVar"
	"web_gin/MiddleWare/zLog"
)

// 客户端同步回调
func ClientGetSign(c *gin.Context) {
	zLog.PrintLogger("================ClientGetSign  客户端同步回调=================")
	var req *http.Request
	req = c.Request
	req.ParseForm()
	//fmt.Printf("req.PostForm :   %v \n",req.PostForm)
	//fmt.Println("----------------------------------------")
	//fmt.Printf("portia :   %s \n",c.PostForm("portia"))
	//fmt.Printf("%v \n",req.PostForm)
	//fmt.Println("----------------------------------------")

	str:=c.PostForm("portia")
	//fmt.Println(str)
	var back ClientCallBack
	var result ClientResult

	_ = json.Unmarshal([]byte(str), &back)
	//fmt.Println(back.ResultStatus)
	zLog.PrintLogger(back.Result)
	_ = json.Unmarshal([]byte(back.Result), &result)
	fmt.Println("result.AlipayTradeAppPayResponseS          ",result.AlipayTradeAppPayResponseS)
	//fmt.Println("result.AlipayTradeAppPayResponseS          ",result.Sign)
	//fmt.Println("result.AlipayTradeAppPayResponseS          ",result.SignType)
	//fmt.Println(result.TradeAppPayResponse.Timestamp)
	//fmt.Println(result.TradeAppPayResponse.OutTradeNo)
	////fmt.Println(back.Result.Alipay_trade_app_pay_response.Timestamp)
	//
	//ppp := result.TradeAppPayResponse
	//data, _ := json.Marshal(ppp)
	//zLog.PrintfLogger("result.TradeAppPayResponse json: ", string(data))

	// 组合验签字符串
	var resList map[string]string = make(map[string]string,0)
	json.Unmarshal([]byte(result.AlipayTradeAppPayResponseS)  ,&resList)

	fmt.Printf("组合验签字符串  %v ", resList)
	for k,v := range resList{

	}

	sign:= result.Sign
	ok, err := verifyData(str, sign, []byte(GlobalVar.AliPublicKey))	//验签
	fmt.Println(ok, err)
	if err != nil {
		c.JSON(200, "VerifySign failed")
		return
	}

	zLog.PrintLogger("============================================")
	c.JSON(200, "to be continue")
}

func test11()  {

	//{"code":"10000","msg":"Success","app_id":"2016110200785785","auth_app_id":"2016110200785785","charset":"utf-8","timestamp":"2020-11-19 18:25:19","out_trade_no":"1605781505114909200","total_amount":"0.01","trade_no":"2020111922001411520507219434","seller_id":"2088102181632824"}
}





func verifyData(data []byte,  sign string, key []byte) (ok bool, err error) {
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



//----------------------------------------------------------------------------
//         同步回调的数据格式
//----------------------------------------------------------------------------
type ClientCallBack struct {
	ResultStatus string  `json:"resultStatus"`
	Result string `json:"result"`
	Memo string `json:"memo"`
}
type ClientResult struct {
	//TradeAppPayResponse TradeAppPayResponse  `json:"alipay_trade_app_pay_response"`
	AlipayTradeAppPayResponseS string  `json:"alipay_trade_app_pay_response"`
	Sign string `json:"sign"`
	SignType string `json:"sign_type"`
}

//
type TradeAppPayResponse struct {
	Code string  `json:"code"`
	Msg string `json:"msg"`
	AppId string `json:"App_id"`
	OutTradeNo string `json:"Out_trade_no"`
	TradeNo string `json:"Trade_no"`
	TotalAmount string `json:"Total_amount"`
	SellerId string `json:"Seller_id"`
	Charset string `json:"Charset"`
	Timestamp string `json:"timestamp"`
	AuthAppId string `json:"auth_app_id"`
}