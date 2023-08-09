package authmodel

import (
	"github.com/hthai2201/dw-go-23/exercises/06/module/user/usermodel"
)

type LoginUser struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (LoginUser) TableName() string {
	return usermodel.User{}.TableName()
}
