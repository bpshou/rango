package golang

import (
	"rango/models"
)

type UserModel struct {
	models.BaseModel
}

type User struct {
	Id         int
	Name       string
	Age        int
	CreateTime string
}

func (this User) TableName() string {
	return "user"
}
