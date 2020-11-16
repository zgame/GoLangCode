package aliPay

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay"
	"net/http"
	"strings"
)

var (
	appID = ""
	// RSA2(SHA256)
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAozFwYXqaaI1RSILHcdi4\ntNEMpTeyEBUZ5rooSXlUutURtUgeu8MwYASKE90G+lHsAUBSFeDM3IYKPImdKHt2\nk7mrDrPrSJCkLNd3pBxmbo9JUif9CKAy/t5FtEYrIpC9IlVtwn4StTmXmmeo9GVb\nnYwETo3llfpJEZ8viFYUcIrvv0LxesQrdQw8wfBAIh4nYk8whtnmPMO/CRwLns4B\neykw2IuM2btulvOspTXrbgKhnLPfDKlgfLj49jcYI9nEp9yMpgZeI6Dqp7MHZ7fv\nAiZiB8n4WBCQhmy4Y6iN6lA16X3mHn1KPRrBINrHD6cDCW7e4F4hrHnj3nmWXrJ1\ndQIDAQAB"
	privateKey = "MIIEowIBAAKCAQEAozFwYXqaaI1RSILHcdi4tNEMpTeyEBUZ5rooSXlUutURtUge\nu8MwYASKE90G+lHsAUBSFeDM3IYKPImdKHt2k7mrDrPrSJCkLNd3pBxmbo9JUif9\nCKAy/t5FtEYrIpC9IlVtwn4StTmXmmeo9GVbnYwETo3llfpJEZ8viFYUcIrvv0Lx\nesQrdQw8wfBAIh4nYk8whtnmPMO/CRwLns4Beykw2IuM2btulvOspTXrbgKhnLPf\nDKlgfLj49jcYI9nEp9yMpgZeI6Dqp7MHZ7fvAiZiB8n4WBCQhmy4Y6iN6lA16X3m\nHn1KPRrBINrHD6cDCW7e4F4hrHnj3nmWXrJ1dQIDAQABAoIBAHpl4EYccKcuJuLd\nw70tsQtdJ8DbTyAk03Jr+T9yUwx2NnvjBboKIcRCY1WWl180BnDBz089dimIFzFk\nfY0ZXMxbm2LBqxyX76r6SG+8JU+TBIksGOpZTSY/i8Q0RLH+IP0ZWeNgL6Pg+EYE\nrYHwa5B0rd5FKwcb26Xt4Pa+qUHmoF8UUCC+4YtmAP/UsbUy+XwCDwafXxOdJqJf\nc0Jw+1NooftyVLabzXi9XLsuw3+9Xtw346M5dTw6/0rF2G7wzbj4WilUXGFnLzwM\n623mWfwCDJdLzReU3tKP8PAg+OVl5Zxb6Zk5S5yFphM7u2YqjyfonPoJ4CXNfUTU\nHee8ltkCgYEAzqBEgBbwosuHImFyj3mppQNWQv0KUb0iYOYan5RwscV8Q4FRt2ne\nWcDSIjpZh13z5pbRVqYByp0q601OK7BM24IjCBtLe+FczYExTzyOxLMaO7RVINFn\nuvbZZR23WldiZ0WpDZgxyo9RmIZRBk2jfG+HKYb4bZkmbnnAGXeM1Z8CgYEAyjBK\nrvxwmU5Op789HJzOYzCnWlfE561zQPPxzLGAlPxu6XcpFWXx1INpSmAigmxnRqJP\n53Qt+9w9xyHzeruv1+Y0vxU2FAdlTp8C7NmE+YG+Pla+85HbbylmCb6dZb8+9HMY\np9/0CiSWPavPKdu4pZ5KoyNidocQEgAAsIQjVGsCgYEAynYPqNLRhzKWfwGtFxjH\nOYFDjPAUpHMGtJvDiooQwqAXWq3kPCvoS1m8jP1PrGxLCK7PAHA5YScPXvCon/Zn\n2M5zNQZJuGDiZhspDdLwsZwtIENbBoUpdvFZotKzTjpBmZ+QPlnar/guo50410xL\n3SoK7o3p7roaBjYWHN4fiVECgYBRY2Ec0VdODvyQf+XMt75IpVQohL4peGO1mL0T\n5bvZvUe0SRhLmc7f+coPe2VI1PQ5taqug9Di2oQvvZXyKM0e/nbrGFG9fECmhlG6\nH9FsUnLPS0HwcB1BwQtnDsjzJSnlYtNg+ECXOKUVzCxHMEBCwtZOlzbSeYnZhRDB\n/V7vYwKBgBHczQdF89dmeEbocYuxgcMeUqg03CW4cSEvhhGFk14dBwWt+ZX/Q8NW\n/i/8wmwkYEsQSO08mgsAU+AtPZ3rtwyu85HMOXLIjJHYenFf4HSwqyGJwQjyHNIz\nyn0RD5RDvbYBYPuQYNxNVkARyurytPSaYWJfd+9+yPGpYd+8ogdX"
)
var client *alipay.AliPay


// 初始化
func Init() {

	aliPublicKey = strings.ReplaceAll(aliPublicKey,"\n","")
	privateKey = strings.ReplaceAll(privateKey,"\n","")
	client = alipay.New(appID, aliPublicKey, privateKey, false)
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
