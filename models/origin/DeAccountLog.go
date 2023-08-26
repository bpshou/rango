package origin

import (
	"rango/models"
)

type DeAccountLogModel struct {
	*models.BaseModel
}

// 实例化
func DeAccountLogTable() DeAccountLogModel {
	baseModel := models.NewDatabase("origin")
	baseModel.TableName = "de_account_log"
	return DeAccountLogModel{
		&baseModel,
	}
}
