package models

type Tag struct {
	Id      int      `orm:"column(id);auto;pk"`
	TagName string   `orm:"column(tag_name);unique"`
	Order   int      `orm:"column(order)"`
	Musics  []*Music `orm:"reverse(many)"`
}

func (t *Tag) TableName() string {
	return TableName("tag")
}
