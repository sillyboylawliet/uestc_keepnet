package main

import (
	"github.com/asmcos/requests"
	"github.com/go-ini/ini"
	"log"
	"net/url"
	"os"
)

func testNet() (bool, string) {
	req := requests.Requests()
	req.Debug = 0
	req.SetTimeout(1)
	resp, err := req.Get("http://123.123.123.123")
	if err == nil {
		return false, resp.R.Request.URL.RawQuery
	}
	return true, ""
}

func main() {
	status, params := testNet()
	if status {
		return
	}
	pwd, _ := os.Getwd()
	cfgs, err := ini.Load(pwd + "/config.ini")
	headers := requests.Header{
		"Referer":    "http://172.25.249.64/eportal/index.jsp?" + params,
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}
	data := requests.Datas{
		"userId":          cfgs.Section("user").Key("userId").Value(),
		"password":        cfgs.Section("user").Key("password").Value(),
		"queryString":     url.QueryEscape(params),
		"operatorPwd":     "",
		"service":         "",
		"operatorUserId":  "",
		"validcode":       "",
		"passwordEncrypt": "false",
	}
	resp, _ := requests.Post("http://172.25.249.64/eportal/InterFace.do?method=login", headers, data)
	var json map[string]interface{}
	err = resp.Json(&json)
	if err != nil {
		log.Println("error")
		return
	}
	if json["result"] == "success" {
		log.Println("keep net success")
	} else {
		log.Println(json["message"])
	}
}
