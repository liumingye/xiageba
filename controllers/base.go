package controllers

import (
	"strconv"
	"strings"
	"time"

	"music/modules/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

const (
	DefaultPage  = "1"
	PageSize     = 30
	CacheTimeout = 1 * time.Minute
)

func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
}

func (c *baseController) SetPaginator(per int, nums int, linkFormat string) *utils.Paginator {
	p := utils.NewPaginator(c.Ctx, per, nums, linkFormat)
	c.Data["paginator"] = p
	return p
}

// getPage 返回输入参数中的页面。
//
// 它不接受任何参数并返回一个字符串。
func (c *baseController) getPage() (page string) {
	page = c.Ctx.Input.Param(":page")
	if page == "" {
		page = DefaultPage
	}
	return
}

// getPageInt 返回给定页面字符串的整数表示。
//
// page: 表示页面号的字符串。
// int: 页面号的整数表示。
func (c *baseController) getPageInt(page string) (pageInt int) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Abort("500")
	}
	return
}
