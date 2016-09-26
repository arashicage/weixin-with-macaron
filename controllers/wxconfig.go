package controllers

import (
	"gopkg.in/macaron.v1"
	"encoding/json"

	"weixin-with-macaron/modules/weixin"
	"fmt"
)

func GetWxConfig(ctx *macaron.Context) {

	url := ctx.Query("url")

	c := weixin.GetWxConfig(url)

	js,err := json.Marshal(c)
	if err != nil{
		fmt.Println(err.Error())
	}

	ctx.Write(js)

}