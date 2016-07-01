package class

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     string `orm:"pk;size(30)"`
	Nick   string `orm:"size(30)"`
	Info   string `orm:"null;size(140)"`
	Hobby  string `orm:"null"`
	Email  string `orm:"unique;size(50)"`
	Avator string `orm:"null;size(100)"`
	Url    string `orm:"null;size(100)"`

	Followers int
	Following int

	Regtime time.Time `orm:"auto_now_add"`

	Password string `orm:"size(100)"`
	Private  int
}


func Testmodel() {

	u := User{
		Id:    "jike",
		Nick:  "jike",
		Email: "123@q.com",
	}

	o := orm.NewOrm()
	_, err := o.Insert(&u)

	u1 := User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	u2 := User{Nick: "jike"}
	err = o.Read(&u2, "nick")
	fmt.Println(u2)

	u2.Nick = "jike2"
	_, err = o.Update(&u2)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	_, err = o.Delete(&u)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	if err != nil {
		fmt.Println(err)
	}

}

