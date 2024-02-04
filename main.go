package main

import (
	_ "music/routers"

	"music/models"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
