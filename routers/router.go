package routers

import (
	"music/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.CtrlGet("/", (*controllers.MainController).Get)
	beego.CtrlGet("/music/:id([0-9]+)", (*controllers.MusicController).Get)
	beego.CtrlGet("/s/:text/page/:page([0-9]+)", (*controllers.SearchController).Get)
	beego.CtrlGet("/s/:keyword(.+)", (*controllers.SearchController).Get)
	beego.CtrlGet("/s/:keyword(.+)/:page([0-9]+)", (*controllers.SearchController).Get)
	beego.CtrlGet("/r/yesterday", (*controllers.RankController).Yesterday)
	beego.CtrlGet("/r/yesterday/:page([0-9]+)", (*controllers.RankController).Yesterday)
	beego.CtrlGet("/history", (*controllers.HistoryController).Get)
	beego.CtrlGet("/history/:page([0-9]+)", (*controllers.HistoryController).Get)
	beego.CtrlGet("/t/:tagname(.+)", (*controllers.TagController).Get)
	beego.CtrlGet("/t/:tagname(.+)/:page([0-9]+)", (*controllers.TagController).Get)
}
