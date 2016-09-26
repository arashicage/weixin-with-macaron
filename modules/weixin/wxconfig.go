package weixin

import (
	"fmt"
	"time"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/jssdk"
)

type WxConfig struct {
	AppId     string    `json:"appid"`
	Timestamp string    `json:"timestamp"`
	NonceStr  string    `json:"nonceStr"`
	Signature string    `json:"signature"`
}

func (c WxConfig) String() string {

	s :=
	"wx.config({" + "\n" +
	"    debug: false," + "\n" +
	"    appId: '" + c.AppId + "'," + "\n" +
	"    timestamp: '" + c.Timestamp + "'," + "\n" +
	"    nonceStr: '" + c.NonceStr + "'," + "\n" +
	"    signature: '" + c.Signature + "'," + "\n" +
	"    jsApiList: [" + "\n" +
	"        'scanQRCode'," + "\n" +
	"        'hideOptionMenu'" + "\n" +
	"    ]" + "\n" +
	"});"

	return s;
}

func GetWxConfig(url string) WxConfig {

	AccessTokenServer := mp.NewDefaultAccessTokenServer("wxdfa36c0c9e2a6cd5", "d2a6db1bace397defb272c5c3faa3c6f", nil)
	mpClient := mp.NewClient(AccessTokenServer, nil)
	TicketServer := jssdk.NewDefaultTicketServer(mpClient)

	t, err := TicketServer.Ticket()
	if err != nil {
		fmt.Println(err.Error())
	}
	timestamp := time.Now().Unix()
	ts := fmt.Sprint(timestamp)

	signature := jssdk.WXConfigSign(t, "kcozwaYD9QXN17Tt", ts, url)

	c := WxConfig{
		AppId:     "wxdfa36c0c9e2a6cd5",
		Timestamp: ts,
		NonceStr:  "kcozwaYD9QXN17Tt",
		Signature: signature,
	}

	fmt.Println(c)

	return c

}
