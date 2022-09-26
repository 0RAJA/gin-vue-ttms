package alipay_test

import (
	"log"
	"net/http"
	"strconv"

	"ttms/internal/global"
	pay "ttms/internal/pkg/alipay"

	_ "ttms/internal/setting"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/xid"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

var AliPayClient *pay.Client

func ExampleInit() {
	AliPayClient = pay.Init(pay.Config{
		KAppID:               global.Settings.AliPay.KAppID,
		KPrivateKey:          global.Settings.AliPay.KPrivateKey,
		IsProduction:         global.Settings.AliPay.IsProduction,
		AppPublicCertPath:    global.Settings.AliPay.AppPublicCertPath,
		AliPayRootCertPath:   global.Settings.AliPay.AliPayRootCertPath,
		AliPayPublicCertPath: global.Settings.AliPay.AliPayPublicCertPath,
		NotifyURL:            global.Settings.AliPay.NotifyURL,
		ReturnURL:            global.Settings.AliPay.ReturnURL,
	})
	var s = gin.Default()
	s.GET("/alipay", payUrl)
	s.GET("/callback", callback)
	s.POST("/notify", notify)
	if err := s.Run(":8080"); err != nil {
		panic(err)
	}
}

func payUrl(c *gin.Context) {
	orderID := strconv.FormatInt(xid.Next(), 10)
	url, err := AliPayClient.Pay(pay.Order{
		ID:          orderID,
		Subject:     "ttms购票:" + orderID,
		TotalAmount: 30,
		Code:        pay.LaptopWebPay,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "系统错误")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "校验失败")
		return
	}
	c.JSON(http.StatusOK, "支付成功:"+orderID)
}

func notify(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("支付成功:" + orderID)
	// 做自己的事
}
