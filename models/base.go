package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

// Init registers the db driver, default database, and models in init.
func Init() {
	// need to register models in init
	orm.RegisterModel(new(Music), new(Tag), new(SearchHistory))

	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// need to register default database
	dbhost, _ := config.String("dbhost")
	dbport, _ := config.String("dbport")
	dbuser, _ := config.String("dbuser")
	dbpassword, _ := config.String("dbpassword")
	dbname, _ := config.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.DefaultTimeLoc = time.UTC
	orm.MaxIdleConnections(100)
	orm.MaxOpenConnections(100)
	orm.ConnMaxLifetime(time.Hour * 6)

	orm.RegisterDataBase("default", "mysql", dsn)
}

// 返回带前缀的表名
func TableName(str string) string {
	dbprefix, _ := config.String("dbprefix")
	return dbprefix + str
}
