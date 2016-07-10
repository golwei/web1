package class

import (
	"time"
)

//type Article struct {
//	Id      int
//	Title   string
//	Content string `orm:"type(text)`
//	Author  *User
//	Replys  int
//	Views   int
//	Tag     string
//	Time    time.Time `orm:"auto_now_add"`
//}

type Article struct {
	Id      int
	Title   string
	Content string `orm:"type(text)"`
	Author  *User  `orm:"rel(fk);size(30)"`

	NumReplys int
	NumViews  int

	Time time.Time `orm:"auto_now_add;type(datetime)"`

	Defunct bool
}
