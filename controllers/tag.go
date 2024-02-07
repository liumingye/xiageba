package controllers

import (
	"music/models"
	"strconv"
	"strings"
)

type TagController struct {
	baseController
}

// Yesterday 获取昨天的排名数据并设置页面显示。
func (c *TagController) Get() {
	// 获取音乐ID
	tagname := c.Ctx.Input.Param(":tagname")
	page := c.Ctx.Input.Param(":page")

	trimKeyword := strings.TrimSpace(tagname)
	if trimKeyword == "" {
		c.Abort("404")
	}

	if page == "" {
		page = "1"
	}

	// 将page转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Abort("500")
	}

	var pageSize = 30

	musics, total, _ := (&models.Tag{}).GetMusicByTagName(trimKeyword, pageInt, pageSize)

	c.Data["Musics"] = musics
	c.Data["Tagname"] = tagname
	c.Data["Page"] = page
	c.Data["Total"] = total

	c.SetPaginator(pageSize, total, "/t/"+tagname+"/%s")

	// 设置模板名称
	c.TplName = "tag.tpl"
}
