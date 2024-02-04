package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Music struct {
	Id     int    `orm:"column(id);pk"`
	Name   string `orm:"column(name)"`
	Singer string `orm:"column(singer)"`
	Pic    string `orm:"column(pic)"`
	Link1  string `orm:"column(link1)"`
	Link2  string `orm:"column(link2)"`
	Tag    string `orm:"column(tag)"`
	Lyric  string `orm:"column(lyric)"`
	Tags   []*Tag `orm:"rel(m2m)"`
}

func (m *Music) TableName() string {
	return TableName("music")
}

func FuzzySearchMusic(keyword string, page int, pageSize int) ([]*Music, int, error) {
	o := orm.NewOrm()

	var musics []*Music
	query := "SELECT * FROM music WHERE name LIKE ? OR singer LIKE ? OR id IN (SELECT music_id FROM music_tags WHERE tag_id IN (SELECT id FROM tag WHERE tag_name LIKE ?))"
	_, err := o.Raw(query+" LIMIT ?, ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", (page-1)*pageSize, pageSize).QueryRows(&musics)
	if err != nil {
		return nil, 0, err
	}

	// Get total count for pagination
	var total int
	err = o.Raw("SELECT COUNT(*) FROM ("+query+") AS total", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").QueryRow(&total)
	if err != nil {
		return nil, 0, err
	}

	return musics, total, nil
}
