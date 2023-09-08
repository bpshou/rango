package origin

import (
	"rango/models"
)

type DeAccountModel struct {
	*models.BaseModel
}

// 实例化
func DeAccountTable() DeAccountModel {
	baseModel := models.NewDatabase("origin")
	baseModel.TableName = "de_account"
	return DeAccountModel{
		&baseModel,
	}
}
