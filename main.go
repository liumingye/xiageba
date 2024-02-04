package main

import (
	"music/modules/cache"
	_ "music/routers"

	"music/models"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	cache.Init()
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
