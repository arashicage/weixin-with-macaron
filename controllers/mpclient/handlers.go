package mpclient

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
// "github.com/chanxuehong/wechat/util"

	"github.com/chanxuehong/wechat/mp/media"

	"github.com/chanxuehong/wechat/mp/menu"

//	"github.com/chanxuehong/wechat/mp/message/mass"
//	"github.com/chanxuehong/wechat/mp/message/mass/mass2users"

//	"github.com/chanxuehong/wechat/mp/message/template"

//	"encoding/json"

	"weixin-with-macaron/modules/crypt"

//	"net/url"

	"sort"
	"weixin-with-macaron/modules/biz"
)

// 被动接收消息（事件）推送，一个 URL 监听一个公众号的消息
// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {

	m := media.NewClient(AccessTokenServer, mpClient.HttpClient)
	info, err := m.UploadImage("qr.png")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(info.MediaId)
	//	QGyk0P63eRioWqUo2Py672IjnTaGwB2EQuL5LQ_brsRV8nAZWNA_HQZWVx6aAG2A

	// 简单起见，把用户发送过来的文本原样回复过去
	e := request.GetText(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	//	resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "<a href=\"http://712636.cicp.net/example/index.html\">"+"你好啊!"+text.Content+" </a>")
	resp := response.NewImage(e.FromUserName, e.ToUserName, e.CreateTime, "QGyk0P63eRioWqUo2Py672IjnTaGwB2EQuL5LQ_brsRV8nAZWNA_HQZWVx6aAG2A")

	mp.WriteRawResponse(w, r, resp) // 明文模式
	// mp.WriteAESResponse(w, r, resp) // 安全模式

}

func SubscribeMessageHandler(w http.ResponseWriter, r *mp.Request) {

	e := request.GetSubscribeEvent(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	//	resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "你好，欢迎关注长城软件纳税服务！")

	//	mp.WriteRawResponse(w, r, resp) // 明文模式
	// mp.WriteAESResponse(w, r, resp) // 安全模式


	a := []response.Article{
		{
			Title:       "您好，欢迎关注长城软件纳税服务！",
			Description: "您好，欢迎关注长城软件纳税服务！",
			PicURL:      "http://ips.chotee.com/wp-content/uploads/2015/iphone6s-ipad-pro/iphone6s_4k.jpg",
			URL:         "",
		},
		{
			Title:       "使用帮助",
			Description: "您好，学习如何使用本公众号提供的服务！",
			PicURL:      "http://ips.chotee.com/wp-content/uploads/2015/iphone6s-ipad-pro/iphone6s_4k.jpg",
			URL:         "www.baidu.com",
		},
		{
			Title:       "常见问题",
			Description: "您好，学习如何使用本公众号提供的服务！",
			PicURL:      "http://ips.chotee.com/wp-content/uploads/2015/iphone6s-ipad-pro/iphone6s_4k.jpg",
			URL:         "www.baidu.com",
		},
	}

	resp := response.NewNews(e.FromUserName, e.ToUserName, e.CreateTime, a)

	err := mp.WriteRawResponse(w, r, resp)
	if err != nil {
		log.Println(err.Error())
	}

	return

}

func ScanCodeWaitMsgEvent_Text(w http.ResponseWriter, r *mp.Request) {

	e := menu.GetScanCodeWaitMsgEvent(r.MixedMsg)

	log.Println(e.FromUserName, e.ToUserName, e.CreateTime, e.ScanCodeInfo.ScanResult, e.ScanCodeInfo.ScanType)

	s := strings.Split(e.ScanCodeInfo.ScanResult, ",")

	log.Println(s)

	if e.ScanCodeInfo.ScanType != "qrcode" || len(s) != 7 {
		log.Println("非法的二维码", e.FromUserName)

		resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "您扫描的不是有效的二维码。\n请您重新扫描！")
		err := mp.WriteRawResponse(w, r, resp) // 明文模式
		if err != nil {
			log.Println(err.Error())
		}

		return
	}

	resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "扫描信息为:" +
	"\n发票类型：" + s[1] +
	"\n发票代码：" + s[2] +
	"\n发票号码：" + s[3] +
	"\n开票金额：" + s[4] +
	"\n开票日期：" + s[5] +
	"\n购方识别号：" + s[6] +
	"\n\n查验结果为:一致")

	err := mp.WriteRawResponse(w, r, resp)
	if err != nil {
		log.Println(err.Error())
	}

}

func ScanCodeWaitMsgEvent_Templates(w http.ResponseWriter, r *mp.Request) {

	e := menu.GetScanCodeWaitMsgEvent(r.MixedMsg)

	log.Println(e.FromUserName, e.ToUserName, e.CreateTime, e.ScanCodeInfo.ScanResult, e.ScanCodeInfo.ScanType)

	s := strings.Split(e.ScanCodeInfo.ScanResult, ",")

	log.Println(s)

	if e.ScanCodeInfo.ScanType != "qrcode" || len(s) != 7 {
		log.Println("非法的二维码", e.FromUserName)

		resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "您扫描的不是有效的二维码。\n请您重新扫描！")
		err := mp.WriteRawResponse(w, r, resp) // 明文模式
		if err != nil {
			log.Println(err.Error())
		}

		return
	}

	resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "扫描信息为:" +
	"\n发票类型：" + s[1] +
	"\n发票代码：" + s[2] +
	"\n发票号码：" + s[3] +
	"\n开票金额：" + s[4] +
	"\n开票日期：" + s[5] +
	"\n购方识别号：" + s[6] +
	"\n\n查验结果为:一致")

	err := mp.WriteRawResponse(w, r, resp)
	if err != nil {
		log.Println(err.Error())
	}

}

func ScanCodeWaitMsgEvent_News(w http.ResponseWriter, r *mp.Request) {

	e := menu.GetScanCodeWaitMsgEvent(r.MixedMsg)

	log.Println(e.FromUserName, e.ToUserName, e.CreateTime, e.ScanCodeInfo.ScanResult, e.ScanCodeInfo.ScanType)

	s := strings.Split(e.ScanCodeInfo.ScanResult, ",")

	log.Println(s)

	if e.ScanCodeInfo.ScanType != "qrcode" || len(s) != 7 {
		log.Println("非法的二维码", e.FromUserName)

		// resp := response.NewText(e.FromUserName, e.ToUserName, e.CreateTime, "您扫描的不是有效的二维码。\n请您重新扫描！")
		// err := mp.WriteRawResponse(w, r, resp) // 明文模式
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		a := []response.Article{
			{
				Title:       "非法的二维码",
				Description: "您扫描的不是有效的二维码。请您重新扫描！",
				PicURL:      "http://ips.chotee.com/wp-content/uploads/2015/iphone6s-ipad-pro/iphone6s_4k.jpg",
				URL:         "",
			},
		}

		resp := response.NewNews(e.FromUserName, e.ToUserName, e.CreateTime, a)

		err := mp.WriteRawResponse(w, r, resp)
		if err != nil {
			log.Println(err.Error())
		}

		return
	}

	m := biz.CXFP(s[1] + ":" + s[2] + s[3], s[5], s[4])

	var keys []string
	content := s[2] + "," + s[3]

	for k, _ := range m {
		keys = append(keys, k)
	}

	// 按 key 排序输出
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(m[k])
		content = content + "," + m[k]
	}

	x1 := crypt.Base64Encode([]byte(content))
	x2 := crypt.Base64Encode(x1)

	a := []response.Article{
		{
			Title: "查验结果",
			Description: "扫描信息为:" +
			"\n发票类型：" + s[1] +
			"\n发票代码：" + s[2] +
			"\n发票号码：" + s[3] +
			"\n开票金额：" + s[4] +
			"\n开票日期：" + s[5][:4] + "-" + s[5][4:6] + "-" + s[5][6:],
			PicURL: "http://ips.chotee.com/wp-content/uploads/2015/iphone6s-ipad-pro/iphone6s_4k.jpg",
			URL:    "http://712636.cicp.net/z?result=" + string(x2),
		},
	}
	//	log.Println(a)

	resp := response.NewNews(e.FromUserName, e.ToUserName, e.CreateTime, a)

	err := mp.WriteRawResponse(w, r, resp)
	if err != nil {
		log.Println(err.Error())
	}

}

func verifyQRCode(s []string) bool {
	if len(s) != 7 {
		return false
	}
	// s[0]  版本
	// s[1]  发票类型
	// s[2]	 发票代码
	// s[3]  发票号码
	// s[4]	 开票金额
	// s[5]  开票日期
	// s[6]  购方识别号

	return true
}
