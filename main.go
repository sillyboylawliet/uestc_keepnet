package main

import (
	"github.com/asmcos/requests"
	"github.com/go-ini/ini"
	"log"
)

func main() {
	cfgs, err := ini.Load("config.ini")
	data := requests.Datas{
		"userId":          cfgs.Section("user").Key("userId").Value(),
		"password":        cfgs.Section("user").Key("password").Value(),
		"service":         "",
		"queryString":     "userip=100.66.21.167&wlanacname=&nasip=171.88.130.251&wlanparameter=48-0e-ec-2b-d5-68&url=http://123.123.123.123/&userlocation=ethtrunk/2:511.807",
		"operatorPwd":     "",
		"operatorUserId":  "",
		"validcode":       "",
		"passwordEncrypt": "true",
	}
	resp, _ := requests.Post("http://172.25.249.64/eportal/InterFace.do?method=login", data)
	var json map[string]interface{}
	err = resp.Json(&json)
	if err != nil {
		log.Println("error")
		return
	}
	if json["result"] == "success" {
		log.Println("keep net success")
	} else {
		log.Println("error in response")
	}
}
