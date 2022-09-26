package setting

type group struct {
	Dao        mDao
	Log        log
	Maker      maker
	Snowflake  sf
	Page       page
	Auto       auto
	Worker     worker
	MangerFunc mangerFunc
	LoadQR     loadQR
	LoadManger loadManger
	AliPay     alipay
}

var Group = new(group)

func AllInit() {
	Group.Dao.Init()
	Group.Snowflake.Init()
	Group.Maker.Init()
	Group.Log.Init()
	Group.Page.Init()
	Group.Worker.Init()
	Group.MangerFunc.Init()
	Group.LoadQR.Init()
	Group.LoadManger.Init()
	Group.AliPay.Init()
}
