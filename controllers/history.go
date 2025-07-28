package controllers

import "music/models"

type HistoryController struct {
	baseController
}

// Yesterday 获取昨天的排名数据并设置页面显示。
func (c *HistoryController) Get() {
	// 设置缓存头：公共缓存，有效期60秒
	c.Ctx.Output.Header("Cache-Control", "public, max-age=60")

	page := c.getPage()
	pageInt := c.getPageInt(page)

	searchHistoryModel := &models.SearchHistory{}
	latestSearchTerms, total, err := searchHistoryModel.GetLatestSearchTerms(pageInt, PageSize)

	if err != nil {
		c.Abort("500")
	}

	c.Data["Data"] = latestSearchTerms
	c.SetPaginator(PageSize, total, "/history/%s")
	c.TplName = "history.tpl"
}
