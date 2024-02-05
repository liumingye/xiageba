package controllers

import (
	"music/models"
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

	keyword = strings.ReplaceAll(keyword, "%", "")

	trimKeyword := strings.TrimSpace(keyword)
	if trimKeyword == "" {
		c.Abort("404")
	}

	if page == "" {
		// 保存搜索词
		go (&models.SearchHistory{}).AddSearchHistory(keyword)
		page = "1"
	}

	// 将page转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Abort("500")
	}

	var pageSize = 30

	musics, total, _ := (&models.Music{}).FuzzySearchMusic(trimKeyword, pageInt, pageSize)

	c.Data["Musics"] = musics
	c.Data["Keyword"] = keyword
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Total"] = total

	c.SetPaginator(pageSize, total, "/s/"+keyword+"/%s")

	// 设置模板名称
	c.TplName = "search.tpl"
}
