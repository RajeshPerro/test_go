package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int
	Fname string
	Lname string
	Email string
	Pass  string
}

func init() {
	orm.RegisterModel(new(User))

}
