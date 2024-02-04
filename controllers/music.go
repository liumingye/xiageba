package controllers

import (
	"math/rand"
	"music/models"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm/hints"
)

type MusicController struct {
	baseController
}

func (c *MusicController) Get() {
	// 获取音乐ID
	id := c.Ctx.Input.Param(":id")

	// 将ID转换为整数
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Abort("404")
	}

	// 查询音乐信息
	music := models.Music{Id: idInt}
	err = c.o.Read(&music)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Music"] = &music

	// 加载音乐关联的标签
	c.o.LoadRelated(&music, "Tags", hints.OrderBy("order"))

	// 查询与音乐相关联的标签
	tagArray := make([]map[string]string, 0)
	for _, tag := range music.Tags {
		tagArray = append(tagArray, map[string]string{"name": tag.TagName, "rand": strconv.Itoa(rand.Intn(6) + 1)})
	}
	c.Data["Tags"] = tagArray

	// 获取下载链接
	var links []map[string]string
	appendLink := func(link string) {
		parts := strings.Split(link, "|")
		if len(parts) == 2 {
			links = append(links, map[string]string{"url": parts[0], "code": parts[1]})
		}
	}
	appendLink(music.Link1)
	appendLink(music.Link2)
	c.Data["Links"] = links

	// 分割歌词文本
	c.Data["Lryic"] = strings.Split(music.Lyric, "\n")

	// 设置模板名称
	c.TplName = "music.tpl"
}
