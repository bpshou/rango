package origin

import (
	"rango/models"
)

type DeUserModel struct {
	*models.BaseModel
}

type DeUser struct {
	Id         int
	Phone      string
	CreateTime string
	UpdateTime string
}

// 实例化
func DeUserTable() DeUserModel {
	baseModel := models.NewDatabase("origin")
	baseModel.TableName = "de_user"
	return DeUserModel{
		&baseModel,
	}
}
