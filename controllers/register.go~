package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/models"
)

type RegisterController struct {
	beego.Controller
}

func (m *RegisterController) Register() {
	m.TplNames = "register.tpl"
	// FirstName := m.GetStrings("name")
	// beego.Info(FirstName)

}

func (this *RegisterController) DataCatch() {
	if this.Ctx.Request.Method == "POST" {
		// Fname := r.GetString("name")
		// Lname := r.GetString("lname")
		// Email := r.GetString("email")
		// Pass := r.GetString("pass")
		// beego.Info(Fname, Lname, Email, Pass)
		//o.Using("default")
		user := new(models.User)
		user.Fname = this.GetString("name")
		user.Lname = this.GetString("lname")
		user.Email = this.GetString("email")
		user.Pass = this.GetString("pass")

		o := orm.NewOrm()
		o.Insert(user)
		this.Redirect("/", 302)
		return
	} else {
		this.TplNames = "register.tpl"
	}
}
