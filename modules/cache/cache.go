package cache

import (
	"github.com/beego/beego/v2/client/cache"
)

var (
	Bm cache.Cache
)

// Init initialize the cache instance
func Init() {
	Bm, _ = cache.NewCache("file", `{"CachePath":"cache","FileSuffix":".bin","DirectoryLevel":"2","EmbedExpiry":"0"}`)
}
