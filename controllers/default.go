package controllers

import (
	"context"
	"music/models"
	"music/modules/cache"
	"time"
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

	yesterdayRank, thisWeekRank, thisMonthRank, lastMonthRank := c.getRankFromCache()

	if yesterdayRank == nil || thisWeekRank == nil || thisMonthRank == nil || lastMonthRank == nil {
		yesterdayRank, thisWeekRank, thisMonthRank, lastMonthRank = c.fetchAndCacheRank()
	}

	c.Data["YesterdayRank"] = yesterdayRank
	c.Data["ThisWeekRank"] = thisWeekRank
	c.Data["ThisMonthRank"] = thisMonthRank
	c.Data["LastMonthRank"] = lastMonthRank
	c.Data["LatestSearchTerms"] = latestSearchTerms

	c.TplName = "index.tpl"
}

func (c *MainController) getRankFromCache() ([]*models.SearchRank, []*models.SearchRank, []*models.SearchRank, []*models.SearchRank) {
	if cache.Bm == nil {
		return nil, nil, nil, nil
	}
	yesterdayRank, err1 := cache.Bm.Get(context.Background(), "yesterdayRank")
	thisWeekRank, err2 := cache.Bm.Get(context.Background(), "thisWeekRank")
	thisMonthRank, err3 := cache.Bm.Get(context.Background(), "thisMonthRank")
	lastMonthRank, err4 := cache.Bm.Get(context.Background(), "lastMonthRank")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, nil, nil, nil
	}
	return yesterdayRank.([]*models.SearchRank), thisWeekRank.([]*models.SearchRank), thisMonthRank.([]*models.SearchRank), lastMonthRank.([]*models.SearchRank)
}

func (c *MainController) fetchAndCacheRank() ([]*models.SearchRank, []*models.SearchRank, []*models.SearchRank, []*models.SearchRank) {
	model := &models.SearchHistory{}
	f := model.GetSearchRank
	timeNow := time.Now()
	yesterdayRank, _, _ := f(1, 10, timeNow.AddDate(0, 0, -1), timeNow, false)
	thisWeekRank, _, _ := f(1, 10, timeNow.AddDate(0, 0, -7), timeNow, false)
	thisMonthRank, _, _ := f(1, 10, timeNow.AddDate(0, -1, 0), timeNow, false)
	lastMonthRank, _, _ := f(1, 10, timeNow.AddDate(0, -2, 0), timeNow.AddDate(0, -1, 0), false)
	go func() {
		cacheTimeout := 24 * time.Hour
		cache.Bm.Put(context.Background(), "yesterdayRank", yesterdayRank, cacheTimeout)
		cache.Bm.Put(context.Background(), "thisWeekRank", thisWeekRank, cacheTimeout)
		cache.Bm.Put(context.Background(), "thisMonthRank", thisMonthRank, cacheTimeout)
		cache.Bm.Put(context.Background(), "lastMonthRank", lastMonthRank, cacheTimeout)
	}()
	return yesterdayRank, thisWeekRank, thisMonthRank, lastMonthRank
}
