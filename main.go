package main

import (
	_ "mysql-monitor-dashboard/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

