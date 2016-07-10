package class

import (
	"time"

	"github.com/astaxie/beego/orm"
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


  7 type Article struct {
  8         Id      int
  9         Title   string
 10         Content string `orm:"type(text)"`
 11         Author  *User  `orm:"rel(fk);size(30)"`
 12 
 13         NumReplys int
 14         NumViews  int
 15 
 16 
 17         Time time.Time `orm:"auto_now_add;type(datetime)"`
 18 
 19         Defunct bool
 20 }
 21 
