package controllers

import (
	"html/template"
	"net/http"
	"runtime"

	"github.com/beego/beego/v2"
)

type ErrorController struct {
	baseController
}

func (c *ErrorController) showErr(errCode int, errContent string) {
	c.Data = map[interface{}]interface{}{
		"Error":        errCode,
		"Title":        http.StatusText(errCode),
		"Content":      template.HTML(errContent),
		"BeegoVersion": beego.VERSION,
		"GoVersion":    runtime.Version(),
	}
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error404() {
	c.showErr(404, "The page you have requested has flown the coop."+
		"<br>Perhaps you are here because:"+
		"<br>The page has moved"+
		"<br>The page no longer exists"+
		"<br>You were looking for your puppy and got lost"+
		"<br>You like 404 pages")
}
func (c *ErrorController) Error500() {
	c.showErr(500, "The page you have requested is down right now."+
		"<br>Please try again later and report the error to the website administrator")
}
