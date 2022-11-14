package main

import (
	_ "blog/routers"
	"blog/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
