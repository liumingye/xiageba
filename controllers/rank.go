package controllers

import (
	"context"
	"music/models"
	"time"

	"music/modules/cache"
)

type RankController struct {
	baseController
}

const (
	CacheKeyPrefixYesterdayRankPage = "yesterday_rank_page_"
	CacheKeyYesterdayRankTotal      = "yesterday_rank_total"
	TitleYesterdaySearchRank        = "昨日搜索排行"
)

// Yesterday 获取昨天的排名数据并设置页面显示。
func (c *RankController) Yesterday() {
	page := c.getPage()
	pageInt := c.getPageInt(page)
	c.Data["Title"] = TitleYesterdaySearchRank

	cacheRankKey := CacheKeyPrefixYesterdayRankPage + page
	cacheTotalKey := CacheKeyYesterdayRankTotal

	rank, total := c.getRankFromCache(cacheRankKey, cacheTotalKey)
	if rank == nil {
		rank, total = c.fetchAndCacheRank(pageInt, cacheRankKey, cacheTotalKey)
	}

	c.Data["Data"] = rank
	c.SetPaginator(PageSize, total, "/r/%s")
	c.TplName = "rank.tpl"
}

// getRankFromCache 从缓存中获取排名和总数。
//
// cacheRankKey, cacheTotalKey string.
// []models.SearchRank, int.
func (c *RankController) getRankFromCache(cacheRankKey, cacheTotalKey string) (rank []models.SearchRank, total int) {
	if cache.Bm == nil {
		return nil, 0
	}
	cacheData, err := cache.Bm.GetMulti(context.Background(), []string{cacheRankKey, cacheTotalKey})
	if err != nil {
		return nil, 0
	}
	if len(cacheData) == 2 {
		rank, total = cacheData[0].([]models.SearchRank), cacheData[1].(int)
	}
	return
}

// fetchAndCacheRank fetches the rank and caches it.
//
// Parameters:
//
//	page int - the page number
//	cacheRankKey string - the cache key for the rank
//	cacheTotalKey string - the cache key for the total
//
// Returns:
//
//	[]models.SearchRank - the rank
//	int - the total
func (c *RankController) fetchAndCacheRank(page int, cacheRankKey string, cacheTotalKey string) (rank []models.SearchRank, total int) {
	rank, total, err := (&models.SearchHistory{}).GetSearchRank(page, PageSize, time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		c.Abort("500")
	}
	go func() {
		cache.Bm.Put(context.Background(), cacheRankKey, rank, CacheTimeout)
		cache.Bm.Put(context.Background(), cacheTotalKey, total, CacheTimeout)
	}()
	return
}
