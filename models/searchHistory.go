package models

import (
	"errors"
	"time"

	"music/modules/utils"

	"github.com/beego/beego/v2/client/orm"
)

type SearchHistory struct {
	Id              int       `orm:"column(id);pk"`
	SearchTerm      string    `orm:"column(search_term)"`
	SearchTimestamp time.Time `orm:"column(search_timestamp)"`
}

func (t *SearchHistory) TableName() string {
	return TableName("search_history")
}

// 添加搜索历史记录
func (t *SearchHistory) AddSearchHistory(searchTerm string) error {
	if !isValidSearchTerm(searchTerm) {
		return errors.New("invalid search term")
	}
	o := orm.NewOrm()
	searchHistory := &SearchHistory{
		SearchTerm:      searchTerm,
		SearchTimestamp: time.Now(),
	}
	_, err := o.Insert(searchHistory)
	return err
}

func isValidSearchTerm(searchTerm string) bool {
	return utils.HanCounter(searchTerm) <= 50
}

type SearchRank struct {
	SearchTerm  string `orm:"column(search_term)"`
	SearchCount int    `orm:"column(search_count)"`
}

// GetSearchRank 根据提供的页面、页面大小、开始日期和结束日期检索搜索排名。
// 返回SearchRank数组、整数和错误。
func (t *SearchHistory) GetSearchRank(page, pageSize int, startDate, endDate time.Time, hasTotal bool) ([]*SearchRank, int, error) {
	o := orm.NewOrm()
	var searchRanks []*SearchRank

	startString := startDate.Format("2006-01-02")
	endString := endDate.Format("2006-01-02")

	// Use a predefined SQL query with named parameters
	query := `SELECT search_term, COUNT(*) AS search_count
            FROM search_history
            WHERE search_timestamp >= ? AND search_timestamp < ?
            GROUP BY search_term
            ORDER BY search_count DESC
            LIMIT ?, ?`

	// Execute the query and scan results directly into the SearchRank struct
	_, err := o.Raw(query, startString, endString, (page-1)*pageSize, pageSize).QueryRows(&searchRanks)
	if err != nil {
		return nil, 0, err
	}

	// Count total number of search terms
	var totalCount int = 0
	if hasTotal {
		countQuery := `SELECT COUNT(DISTINCT search_term)
						FROM search_history
						WHERE search_timestamp >= ? AND search_timestamp < ?`
		err = o.Raw(countQuery, startString, endString).QueryRow(&totalCount)
		if err != nil {
			return nil, 0, err
		}
	}

	// Return the results along with the total count
	return searchRanks, totalCount, nil
}

type LatestSearchTerms struct {
	SearchTerm string `json:"search_term"`
	SecondsAgo string `json:"seconds_ago"`
}

// 获取最新搜索词
func (t *SearchHistory) GetLatestSearchTerms(page int, pageSize int) ([]LatestSearchTerms, int, error) {
	o := orm.NewOrm()
	var searchTerms []SearchHistory

	// 执行查询
	_, err := o.QueryTable("search_history").OrderBy("-search_timestamp").Offset((page - 1) * pageSize).Limit(pageSize).All(&searchTerms)
	if err != nil {
		return nil, 0, err
	}

	// 将结果转换为所需的类型
	var result []LatestSearchTerms
	for _, term := range searchTerms {
		secondsAgo := time.Since(term.SearchTimestamp).Seconds()
		secondsString := utils.FormatSecondsAgo(secondsAgo)
		result = append(result, LatestSearchTerms{
			SearchTerm: term.SearchTerm,
			SecondsAgo: secondsString,
		})
	}

	// Count total number of search terms
	totalCount, err := o.QueryTable("search_history").Count()
	if err != nil {
		return nil, 0, err
	}

	return result, int(totalCount), nil
}
