package controllers

import (
	"music/models"
)

type MainController struct {
	baseController
}

func (c *MainController) Get() {
	// 获取最新搜索词
	searchHistoryModel := &models.SearchHistory{}
	latestSearchTerms, _, err := searchHistoryModel.GetLatestSearchTerms(1, 16)

	if err != nil {
		c.Abort("500")
	}

	c.Data["LatestSearchTerms"] = latestSearchTerms

	c.TplName = "index.tpl"
}
