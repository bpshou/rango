package service

import (
	"rango/models/origin"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
)

func Insert() int64 {
	var data []map[string]interface{}

	data = append(data, map[string]interface{}{
		"phone": 13188881111,
	}, map[string]interface{}{
		"phone": 13188882222,
	})

	logrus.Debug(data)

	lastId, affected, err := origin.UserTable().Insert(data)

	logrus.Debug(lastId, affected, err)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	return lastId
}

func Select() (list []map[string]interface{}) {

	data := goqu.Ex{
		"id": goqu.Op{"eq": 1},
	}

	list, err := origin.UserTable().GetList(data, 0, 0, map[string]string{})

	logrus.Debug(list)

	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return
}

func Update() int64 {
	data := goqu.Record{
		"phone": "13188883333",
	}
	where := goqu.Ex{
		"id": goqu.Op{"eq": 2},
	}

	affected, err := origin.UserTable().Update(data, where)

	logrus.Debug(affected, err)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	return affected
}

func Delete() int64 {
	where := goqu.Ex{
		"id": goqu.Op{"eq": 3},
	}

	affected, err := origin.UserTable().Delete(where)

	logrus.Debug(affected, err)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	return affected
}
