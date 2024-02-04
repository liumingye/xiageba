package controllers

import (
	"music/models"
	"music/modules/utils"
	"strconv"
	"strings"
)

type SearchController struct {
	baseController
}

func (c *SearchController) Get() {
	// 获取音乐ID
	keyword := c.Ctx.Input.Param(":keyword")
	page := c.Ctx.Input.Param(":page")

	if page == "" {
		page = "1"
	}

	// 将page转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Abort("404")
	}

	trimKeyword := strings.TrimSpace(keyword)
	if trimKeyword == "" {
		c.Abort("404")
	}

	var pageSize = 30

	musics, total, _ := models.FuzzySearchMusic(trimKeyword, pageInt, pageSize)

	c.Data["Musics"] = musics
	c.Data["Keyword"] = keyword
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Total"] = total

	c.SetPaginator(pageSize, total)

	// 设置模板名称
	c.TplName = "search.tpl"
}

func (c *SearchController) SetPaginator(per int, nums int) *utils.Paginator {
	p := utils.NewPaginator(c.Ctx, per, nums)
	c.Data["paginator"] = p
	return p
}
