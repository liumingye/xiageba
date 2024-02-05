package template

import (
	"math/rand"
	"strconv"

	web "github.com/beego/beego/v2/server/web"
)

func random(min int, max int) string {
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func AddFuncMap() {
	web.AddFuncMap("random", random)
}
