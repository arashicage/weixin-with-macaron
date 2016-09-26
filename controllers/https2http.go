package controllers

import (
	"fmt"
	"gopkg.in/macaron.v1"

	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RetriveCaptcha(ctx *macaron.Context) {
	targetUrl :=ctx.Query("url")
	fmt.Println("===="+targetUrl)

	var resp *http.Response
	var err error
	var data []byte
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, //忽略证书签名
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp, err = client.PostForm(targetUrl, nil)
	defer  resp.Body.Close()

	if data, err = ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("%s\n", data)
	}

}

func VerifyParam(ctx *macaron.Context) {
	targetUrl :=ctx.Data["url"].(string)

	var resp *http.Response
	var err error
	var data []byte
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, //忽略证书签名
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	param := url.Values{//map[string][]string
		"fpdm": {"1100152130"},
		"fphm": {"00920868"},
		"kprq": {"20150916"},
		"fpje": {"6618.74"},
		"fplx": {"01"},
		"yzm": {"1111"},
		"yzmSj": {"11"},
		"loginSj": {"22"},
		"token": {"33"},
		"username": {"44"},
		"index": {"55"},
	};

	resp, err = client.PostForm(targetUrl,param)
	defer  resp.Body.Close()

	if data, err = ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("%s\n", data)
	}
}