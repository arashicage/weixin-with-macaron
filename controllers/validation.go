package controllers

import (
	"gopkg.in/macaron.v1"
	"weixin-with-macaron/modules/biz"
)

type fpxx_01 struct {
	Fpdm   string
	Fphm   string
	Kprq   string
	Je     string
	Se     string
	Jshj   string
	Xfmc   string
	Xfsbh  string
	Xfdzdh string
	Xfyhzh string
	Gfmc   string
	Gfsbh  string
	Gfdzdh string
	Gfyhzh string
}

func Validataion(ctx *macaron.Context) {

	fplx := ctx.Query("fplx")
	fpdm := ctx.Query("fpdm")
	fphm := ctx.Query("fphm")
	kprq := ctx.Query("kprq")
	kpje := ctx.Query("kpje")

	key := fplx + ":" + fpdm + fphm

	m := biz.CXFP(key, kprq, kpje)

	detail := fpxx_01{
		Fpdm:   fpdm,
		Fphm:   fphm,
		Kprq:   m["01"],
		Je:     m["10"],
		Se:     m["11"],
		Jshj:   m["12"],
		Xfmc:   m["06"],
		Xfsbh:  m["07"],
		Xfdzdh: m["08"],
		Xfyhzh: m["09"],
		Gfmc:   m["02"],
		Gfsbh:  m["03"],
		Gfdzdh: m["04"],
		Gfyhzh: m["05"],
	}

	ctx.Data["info"] = detail
	ctx.Data["result"] = m["00"]

	ctx.HTML(200, "r")
}
