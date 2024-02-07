package controllers

import (
	"context"
	"music/models"
	"music/modules/cache"
	"strconv"
	"strings"
	"time"
)

type SearchController struct {
	baseController
}

func (c *SearchController) getMusicsFromCache(keyword string, page int) (musics []*models.Music, total int) {
	if cache.Bm == nil {
		return nil, 0
	}
	cacheData, err := cache.Bm.GetMulti(context.Background(), []string{"search_" + keyword + "_" + strconv.Itoa(page), "search_total_" + keyword})
	if err != nil {
		return nil, 0
	}
	if len(cacheData) == 2 {
		musics, total = cacheData[0].([]*models.Music), cacheData[1].(int)
	}
	return
}

func (c *SearchController) fetchAndCache(keyword string, page int) (musics []*models.Music, total int) {
	musics, total, err := (&models.Music{}).FuzzySearchMusic(keyword, page, 30)
	if err != nil {
		c.Abort("500")
	}
	go func() {
		cache.Bm.Put(context.Background(), "search_"+keyword+"_"+strconv.Itoa(page), musics, 24*time.Hour)
		cache.Bm.Put(context.Background(), "search_total_"+keyword, total, 24*time.Hour)
	}()
	return
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

	musics, total := c.getMusicsFromCache(keyword, pageInt)
	if musics == nil {
		musics, total = c.fetchAndCache(keyword, pageInt)
	}

	c.Data["Musics"] = musics
	c.Data["Keyword"] = keyword
	c.Data["Page"] = page
	c.Data["Total"] = total

	c.SetPaginator(pageSize, total, "/s/"+keyword+"/%s")

	// 设置模板名称
	c.TplName = "search.tpl"
}
