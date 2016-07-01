package main

import (
	_ "web1/models"
	_ "web1/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

