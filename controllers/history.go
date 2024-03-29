package controllers

import "music/models"

type HistoryController struct {
	baseController
}

// Yesterday 获取昨天的排名数据并设置页面显示。
func (c *HistoryController) Get() {

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
