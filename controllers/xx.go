package controllers

import (
	"fmt"
	"strings"

	"gopkg.in/macaron.v1"

	"weixin-with-macaron/modules/crypt"
)

func Validataion2(ctx *macaron.Context) {

	result := ctx.Query("result")

	x1, err := crypt.Base64Decode([]byte(result))
	x2, err := crypt.Base64Decode(x1)

	if err != nil {
		fmt.Println(err.Error())
	}

	m := strings.Split(string(x2), ",")

	detail := fpxx_01{
		Fpdm:   m[0],
		Fphm:   m[1],
		Kprq:   m[1+2],
		Je:     m[10+2],
		Se:     m[11+2],
		Jshj:   m[12+2],
		Xfmc:   m[6+2],
		Xfsbh:  m[7+2],
		Xfdzdh: m[8+2],
		Xfyhzh: m[9+2],
		Gfmc:   m[2+2],
		Gfsbh:  m[3+2],
		Gfdzdh: m[4+2],
		Gfyhzh: m[5+2],
	}

	ctx.Data["info"] = detail
	ctx.Data["result"] = m[2]

	ctx.HTML(200, "r")
}
