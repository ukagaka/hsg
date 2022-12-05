package User

import (
	"hxsg/internal/model"
	"hxsg/internal/model/Db"
)

type UserModel struct {
	model.Module
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) CheckUserName(name string) {
	Db.UserObjs.CheckUserName(name)
}
