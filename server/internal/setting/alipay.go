package setting

import (
	"ttms/internal/global"
	pay "ttms/internal/pkg/alipay"
)

type alipay struct {
}

func (alipay) Init() {
	global.AliPayClient = pay.Init(pay.Config{
		KAppID:               global.Settings.AliPay.KAppID,
		KPrivateKey:          global.Settings.AliPay.KPrivateKey,
		IsProduction:         global.Settings.AliPay.IsProduction,
		AppPublicCertPath:    global.Settings.AliPay.AppPublicCertPath,
		AliPayRootCertPath:   global.Settings.AliPay.AliPayRootCertPath,
		AliPayPublicCertPath: global.Settings.AliPay.AliPayPublicCertPath,
		NotifyURL:            global.Settings.AliPay.NotifyURL,
		ReturnURL:            global.Settings.AliPay.ReturnURL,
	})
}
