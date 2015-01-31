package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "test/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", "root:root123@/go_test?charset=utf8")

	name := "default"
	force := false
	verbose := true

	orm.Debug = true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("Error!")
	}
}

func main() {

	beego.Run()
}
