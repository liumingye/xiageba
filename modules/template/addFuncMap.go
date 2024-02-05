package template

import (
	"math/rand"
	"strconv"

	web "github.com/beego/beego/v2/server/web"
)

func random(min int, max int) string {
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func plus(a int, b int) int {
	return a + b
}

func AddFuncMap() {
	web.AddFuncMap("plus", plus)
	web.AddFuncMap("random", random)
}
