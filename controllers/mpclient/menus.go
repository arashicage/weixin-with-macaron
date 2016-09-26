package mpclient

import (
	"log"

	"github.com/chanxuehong/wechat/mp/menu"
)

func InitMenu(d string) {

	var subButtons1 = make([]menu.Button, 5)
	subButtons1[0].SetAsViewButton("测试页面1", "http://" + d + ".cicp.net/html/index.html")
	subButtons1[1].SetAsViewButton("客服热线", "http://" + d + ".cicp.net/xxx")
	subButtons1[2].SetAsViewButton("weui", "http://" + d + ".net/weui/example/index.html")
	subButtons1[3].SetAsViewButton("index", "http://" + d + ".net/fpcy-html/index.html")
	subButtons1[4].SetAsViewButton("JSSDK", "https://inv-veri.chinatax.gov.cn/WebQuery/do/1100/yzmQuery")

	var mn menu.Menu
	mn.Buttons = make([]menu.Button, 3) //最多3个
//	mn.Buttons[0].SetAsViewButton("发票查验","https://inv-veri.chinatax.gov.cn/WebQuery/do/1100/yzmQuery")
	mn.Buttons[0].SetAsViewButton("发票查验", "http://" + d + ".cicp.net/x")
	mn.Buttons[1].SetAsScanCodeWaitMsgButton("扫码查票", "QRSCAN")
	mn.Buttons[2].SetAsSubMenuButton("查票助手", subButtons1)

	menuClient := (*menu.Client)(mpClient)
	if err := menuClient.CreateMenu(mn); err != nil {
		log.Println(err)
		return
	}
	log.Println("Menus OK!")

}