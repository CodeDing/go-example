package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("smtp", `{"username":"tommyhey@163.com","password":"xxxxxx","host":"smtp.163.com:25","fromAddress":"tommyhey@163.com","sendTos":["dingximing@shhuzhong.com"]}`)
	log.Critical("sendmail critical")
	time.Sleep(time.Second * 30)
}
