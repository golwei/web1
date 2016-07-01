package models

import (
	. "web1/models/class"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "a.db")

	orm.RegisterModel(new(User))
	_ = orm.RunSyncdb("default", false, true)

}
