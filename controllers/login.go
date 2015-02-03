package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	this.TplNames = "login.tpl"
}

func (this *LoginController) LoginCheck() {
	// Email := this.GetString("email")
	// Pass := this.GetString("pass")
	// beego.Info(Email, Pass)

	u := new(models.User)
	u.Email = this.GetString("email")
	u.Pass = this.GetString("pass")

	// beego.Info(u.Fname)

	o := orm.NewOrm()
	err := o.Read(u, "Email", "Pass")
	//qs := o.QueryTable("user")
	//em := qs.Filter("email", "Email")
	//pas := qs.Filter("pass", "Pass")
	//beego.Info(em, pas)

	if err == orm.ErrNoRows {
		beego.Info("No")
		this.Redirect("/login", 302)
		return
	} else {
		beego.Info("Yes")
		this.Redirect("/", 302)
		return
	}
	//beego.Info("fgfghfgh")

}
