package main

import (
	"github.com/asmcos/requests"
	"github.com/go-ini/ini"
	"log"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	cfgs, err := ini.Load(pwd + "\\config.ini")
	data := requests.Datas{
		"userId":          cfgs.Section("user").Key("userId").Value(),
		"password":        cfgs.Section("user").Key("password").Value(),
		"queryString":     cfgs.Section("user").Key("queryString").Value(),
		"operatorPwd":     "",
		"service":         "",
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
		log.Println(json["message"])
	}
}
