package models

import (
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type Music struct {
	Id     int    `orm:"column(id);pk"`
	Name   string `orm:"column(name)"`
	Singer string `orm:"column(singer)"`
	Pic    string `orm:"column(pic)"`
	Link1  string `orm:"column(link1)"`
	Link2  string `orm:"column(link2)"`
	Lyric  string `orm:"column(lyric)"`
	Tags   []*Tag `orm:"rel(m2m)"`
}

func (m *Music) TableName() string {
	return TableName("music")
}

func (t *Music) FuzzySearchMusic(keyword string, page int, pageSize int) ([]*Music, int, error) {
	o := orm.NewOrm()

	keyword = strings.ReplaceAll(keyword, "%", "\\%")
	keyword = strings.ReplaceAll(keyword, "_", "\\_")

	var musics []*Music
	// Use placeholders for parameters
	query := "SELECT * FROM music WHERE name LIKE ? OR singer LIKE ? OR id IN (SELECT music_id FROM music_tags WHERE tag_id IN (SELECT id FROM tag WHERE tag_name LIKE ?))"
	// Pass the parameters separately
	_, err := o.Raw(query+" LIMIT ?, ?", keyword+"%", keyword+"%", keyword+"%", (page-1)*pageSize, pageSize).QueryRows(&musics)
	if err != nil {
		return nil, 0, err
	}

	// Get total count for pagination
	var total int
	countQuery := "SELECT COUNT(*) FROM (" + query + ") AS total"
	// Pass the parameters separately
	err = o.Raw(countQuery, keyword+"%", keyword+"%", keyword+"%").QueryRow(&total)
	if err != nil {
		return nil, 0, err
	}

	return musics, total, nil
}
