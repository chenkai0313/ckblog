package validate

import (
	"ckblog/models"
	"github.com/astaxie/beego/validation"
	"errors"
)

type UserValidate struct {

}

func (user UserValidate)RegisterValite(u models.User) error {
	valid := validation.Validation{}
	valid.Required(u.UserName, "user_name")
	valid.Required(u.Email, "emil")
	valid.Required(u.Password, "password")
	valid.Email(u.Email, "emil")
	valid.MinSize(u.Password, 5,"password")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return  errors.New(err.Key+" "+err.Message)
		}
	}

	return nil
}


func (user UserValidate)LoginValite(u models.User) error {
	valid := validation.Validation{}
	valid.Required(u.UserName, "user_name")
	valid.Required(u.Password, "password")
	valid.MinSize(u.Password, 5,"password")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return  errors.New(err.Key+" "+err.Message)
		}
	}

	return nil
}

