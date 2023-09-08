package account

import (
	"rango/models/origin"
	"rango/tools/helper"

	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/doug-martin/goqu/v9"
)

type AccountService struct{}

func (the *AccountService) PutAccount(uid int64) (string, error) {
	// 查询一个没用过的账号
	accountData, err := origin.DeAccountTable().GetOne(goqu.Ex{"use_status": 0}, map[string]string{"id": "asc"})
	if err != nil {
		return "", err
	}

	// chatgpt key
	key := cast.ToString(accountData["key"])
	if key == "" {
		return "", errors.New("chatgpt key is empty")
	}

	err = AddDeAccountLog(uid, key)
	if err != nil {
		return "", errors.New("log error")
	}

	affect, _ := origin.DeAccountTable().Update(goqu.Ex{
		"use_status": 1,
	}, goqu.Ex{
		"id": helper.ToInt64(accountData, "id"),
	})
	if affect <= 0 {
		return "", errors.New("DeAccountTable update error")
	}

	return key, nil
}

// 增加账号下发记录
func AddDeAccountLog(uid int64, key string) (err error) {
	if uid <= 0 || key == "" {
		return errors.New("uid or key is empty")
	}

	// 增加日志
	_, affect, err := origin.DeAccountLogTable().Insert(map[string]any{
		"`uid`": uid,
		"`key`": key,
	})
	if err != nil {
		return
	}
	if affect < 1 {
		err = errors.New("data not insert")
		return
	}

	return nil
}
