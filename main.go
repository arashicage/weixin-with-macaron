package main

import (
	"gopkg.in/macaron.v1"
	"weixin-with-macaron/controllers/mpclient"

	"weixin-with-macaron/controllers"

//	"github.com/chanxuehong/wechat/mp/jssdk"
//	"github.com/chanxuehong/wechat/mp"
//	"fmt"
//	"time"

	"weixin-with-macaron/modules/weixin"
	"github.com/go-macaron/cache"
)

func main() {
	d := "712636"
//	d := "659770"
	mpclient.InitMenu(d)

	m := macaron.Classic()
	m.Use(cache.Cacher())

	//	默认为静态 dir 为 public 目录
	m.Use(macaron.Static("www"))
	//	设置模板目录,模板后缀
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
	}))

	//	微信入口
	m.Any("/wechat", mpclient.ServerFrontendx().ServeHTTP)

	//	手工输入页面
	m.Get("/x", func(ctx *macaron.Context,c1 cache.Cache) {

		s := c1.Get("signature")
		var c weixin.WxConfig

		if s == nil {
			//compute signature and cache it
			//val := weixin.GetWxConfig("http://" + d + ".cicp.net/x")
			val := weixin.GetWxConfig("http://" + d + ".cicp.net/x")
			c1.Put("signature", val, 7200)
			c = val
		} else {
			c = s.(weixin.WxConfig)
		}

		// 直接获取
		//c := weixin.GetWxConfig("http://" + d + ".cicp.net/x")
		//c := weixin.GetWxConfig("http://" + d + ".cicp.net/x")

		ctx.Data["debug"] = false
		ctx.Data["appid"] = c.AppId
		ctx.Data["timestamp"] = c.Timestamp
		ctx.Data["nonceStr"] = c.NonceStr
		ctx.Data["signature"] = c.Signature

		ctx.HTML(200, "index")
	})

	m.Get("/x1", func(ctx *macaron.Context) {

		c := weixin.GetWxConfig("http://712636.cicp.net/x1")

		ctx.Data["debug"] = false
		ctx.Data["appid"] = c.AppId
		ctx.Data["timestamp"] = c.Timestamp
		ctx.Data["nonceStr"] = c.NonceStr
		ctx.Data["signature"] = c.Signature

		ctx.HTML(200, "index")

	})

	m.Get("/xxx", func(ctx *macaron.Context) {
		ctx.HTML(200, "xxx")
	})

	//	结果页面
	m.Get("/r", controllers.Validataion)

	m.Get("/z", controllers.Validataion2)

	m.Get("/c", controllers.GetWxConfig)

	m.Get("/yzmQuery", controllers.RetriveCaptcha)

	m.Post("/query", controllers.VerifyParam)

	m.Run(80)

}
