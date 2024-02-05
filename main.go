package main

import (
	"music/controllers"
	_ "music/models"
	_ "music/modules/cache"
	"music/modules/template"
	_ "music/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	template.AddFuncMap()
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	beego.Run()
}
