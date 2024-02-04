package routers

import (
	"music/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/music/?:id", &controllers.MusicController{})
	beego.Router("/s/?:text/page/?:page", &controllers.SearchController{})
	beego.Router("/s/:keyword", &controllers.SearchController{})
	beego.Router("/s/:keyword/:page", &controllers.SearchController{})
}
