package controllers

import (
	"strings"

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

func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
}

func (c *baseController) SetPaginator(per int, nums int) *utils.Paginator {
	p := utils.NewPaginator(c.Ctx, per, nums)
	c.Data["paginator"] = p
	return p
}
