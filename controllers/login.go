package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "test/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	this.TplNames = "login.tpl"
}

func (this *LoginController) LoginCheck() {
	Email := this.GetString("email")
	Pass := this.GetString("pass")
	beego.Info(Email, Pass)
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	em := qs.Filter("email", "Email")
	//pas := qs.Filter("pass", "Pass")
	//beego.Info(em, pas)
	err := o.Read(&em)
	if err == orm.ErrNoRows {
		beego.Info("Wrong!")
	} else {
		beego.Info("Yes!")
	}

}
