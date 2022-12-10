package golang

import (
	"rango/models"
)

type UserModel struct {
	*models.BaseModel
}

type User struct {
	Id         int
	Name       string
	Age        int
	CreateTime string
}

// 实例化
func UserTable() UserModel {
	BaseModel := models.NewDatabase("golang")
	return UserModel{
		&BaseModel,
	}
}

func (the User) TableName() string {
	return "user"
}
