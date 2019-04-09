package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

var SMTP_CONF = `
{
	"username":"tommyhey@163.com",
	"password":"ding123456",
	"host":"smtp.163.com:25",
	"fromAddress":"tommyhey@163.com",
	"sendTos":["dingximing@shhuzhong.com"]
}
`

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("smtp", SMTP_CONF)
	log.Critical("sendmail critical")
	time.Sleep(time.Second * 30)
}
