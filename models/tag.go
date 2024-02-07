package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Tag struct {
	Id      int      `orm:"column(id);auto;pk"`
	TagName string   `orm:"column(tag_name);unique"`
	Order   int      `orm:"column(order)"`
	Musics  []*Music `orm:"reverse(many)"`
}

func (t *Tag) TableName() string {
	return TableName("tag")
}

func (t *Tag) GetMusicByTagName(tagName string, page int, pageSize int) ([]*Music, int, error) {
	o := orm.NewOrm()
	var musics []*Music
	_, err := o.QueryTable(new(Music)).Filter("Tags__Tag__TagName", tagName).Limit(pageSize).Offset((page - 1) * pageSize).All(&musics)
	if err != nil {
		return nil, 0, err
	}
	count, err := o.QueryTable(new(Music)).Filter("Tags__Tag__TagName", tagName).Count()
	if err != nil {
		return nil, 0, err
	}
	return musics, int(count), nil
}
