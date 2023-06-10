package origin

import (
	"rango/models"
)

type UserModel struct {
	*models.BaseModel
}

type User struct {
	Id         int
	Phone      string
	CreateDate string
	UpdateDate string
}

// 实例化
func UserTable() UserModel {
	baseModel := models.NewDatabase("origin")
	baseModel.TableName = "user"
	return UserModel{
		&baseModel,
	}
}
