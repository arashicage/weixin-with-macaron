package mpclient

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/request"
)

var (
	//	个人测试号信息
	// appID          = "wx5c115be8b5f2ae77"
	// appsecret      = "d4624c36b6795d1d99dcf0547af5443d"
	// Token          = "wechat"
	// URL            = "http://712636.cicp.net/wechat"
	// EncodingAESKey = ""
	// oriId          = "gh_f3eb08b22b15"

	// 长城软件纳税服务信息
	appID          = "wxdfa36c0c9e2a6cd5"
	appsecret      = "d2a6db1bace397defb272c5c3faa3c6f"
	Token          = "gwssi"
	URL            = "http://712636.cicp.net/wechat"
	EncodingAESKey = "sq8t8vqmK69wfoICA3aFeTaZ2rvDTMb4j8BuqTJ4w9s"
	oriId          = "gh_32f799deb7f9"
)

// 主动调用微信 api，mp 子包里的 Client 基本都是这样的调用方式
// 一个应用只能有一个实例
var AccessTokenServer = mp.NewDefaultAccessTokenServer(appID, appsecret, nil)
var mpClient = mp.NewClient(AccessTokenServer, nil)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

func ServerFrontendx() *mp.ServerFrontend {

	messageServeMux := mp.NewMessageServeMux()
	// messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler)           // 注册文本处理 Handler
	messageServeMux.EventHandleFunc(request.EventTypeSubscribe, SubscribeMessageHandler) // 注册关注本公众号 Handler
	//	messageServeMux.EventHandleFunc(menu.EventTypeScanCodeWaitMsg, ScanCodeWaitMsgEvent_Text) // 扫描二维码 Handler,返回文本消息
	messageServeMux.EventHandleFunc(menu.EventTypeScanCodeWaitMsg, ScanCodeWaitMsgEvent_News) // 扫描二维码 Handler,返回 News消息
	//	messageServeMux.EventHandleFunc(menu.EventTypeScanCodeWaitMsg, ScanCodeWaitMsgEvent_Templates) // 扫描二维码 Handler,返回 Templates 消息

	// 下面函数的几个参数设置成你自己的参数: oriId, token, appId, aesKey
	mpServer := mp.NewDefaultServer(oriId, Token, appID, nil, messageServeMux)

	mpServerFrontend :=
		mp.NewServerFrontend(mpServer, mp.ErrorHandlerFunc(ErrorHandler), nil)

	return mpServerFrontend
}
