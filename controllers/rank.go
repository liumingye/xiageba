package controllers

import (
	"context"
	"music/models"
	"strconv"
	"time"

	"music/modules/cache"
)

type RankController struct {
	baseController
}

const (
	CacheKeyPrefixYesterdayRankPage = "yesterday_rank_page_"
	CacheKeyYesterdayRankTotal      = "yesterday_rank_total"
	CacheTimeout                    = 1 * time.Minute
	DefaultPage                     = "1"
	PageSize                        = 30
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
	c.SetPaginator(PageSize, total)
	c.TplName = "rank.tpl"
}

// getPage 返回输入参数中的页面。
//
// 它不接受任何参数并返回一个字符串。
func (c *RankController) getPage() string {
	page := c.Ctx.Input.Param(":page")
	if page == "" {
		page = DefaultPage
	}
	return page
}

// getPageInt 返回给定页面字符串的整数表示。
//
// page: 表示页面号的字符串。
// int: 页面号的整数表示。
func (c *RankController) getPageInt(page string) int {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Abort("500")
	}
	return pageInt
}

// getRankFromCache 从缓存中获取排名和总数。
//
// cacheRankKey, cacheTotalKey string.
// interface{}, int.
func (c *RankController) getRankFromCache(cacheRankKey, cacheTotalKey string) (rank interface{}, total int) {
	if cache.Bm == nil {
		return nil, 0
	}
	cacheRank, errRank := cache.Bm.Get(context.Background(), cacheRankKey)
	cacheTotal, errTotal := cache.Bm.Get(context.Background(), cacheTotalKey)
	if errRank != nil || errTotal != nil {
		return nil, 0
	}
	if cacheRank != nil && cacheTotal != nil {
		rank = cacheRank
		total = cacheTotal.(int)
	}
	return rank, total
}

// fetchAndCacheRank fetches the rank and caches it.
//
// Parameters:
//
//	pageInt int - the page number
//	cacheRankKey string - the cache key for the rank
//	cacheTotalKey string - the cache key for the total
//
// Returns:
//
//	interface{} - the rank
//	int - the total
func (c *RankController) fetchAndCacheRank(pageInt int, cacheRankKey, cacheTotalKey string) (interface{}, int) {
	rank, total, err := (&models.SearchHistory{}).GetSearchRank(pageInt, PageSize, time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		c.Abort("500")
	}
	go cache.Bm.Put(context.Background(), cacheRankKey, rank, CacheTimeout)
	go cache.Bm.Put(context.Background(), cacheTotalKey, total, CacheTimeout)
	return rank, total
}
