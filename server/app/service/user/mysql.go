package user

import (
	"rango/models/origin"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
)

func (the *UserService) Insert() int64 {
	var data []map[string]interface{}

	data = append(data, map[string]interface{}{
		"phone": 13188881111,
	}, map[string]interface{}{
		"phone": 13188882222,
	})

	logrus.Debug(data)

	lastId, affected, err := origin.DeUserTable().Insert(data)

	logrus.Debug(lastId, affected, err)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	return lastId
}

func (the *UserService) Select() (list []map[string]interface{}) {

	data := goqu.Ex{
		"id": goqu.Op{"gt": 0},
	}

	list, err := origin.DeUserTable().GetList(data, 1, 2, map[string]string{})

	logrus.Debug(list)

	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return
}

func (the *UserService) Update() int64 {
	data := goqu.Record{
		"phone": "13188883333",
	}
	where := goqu.Ex{
		"id": goqu.Op{"eq": 2},
	}

	affected, err := origin.DeUserTable().Update(data, where)

	logrus.Debug(affected, err)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	return affected
}

func (the *UserService) Delete() int64 {
	where := goqu.Ex{
		"id": goqu.Op{"eq": 3},
	}

	affected, err := origin.DeUserTable().Delete(where)

	logrus.Debug(affected, err)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	return affected
}

func (the *UserService) GetOne() (one map[string]interface{}) {

	data := goqu.Ex{
		"id": goqu.Op{"eq": 1},
	}

	one, err := origin.DeUserTable().GetOne(data, map[string]string{})

	logrus.Debug(one)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	return
}

func (the *UserService) GetCount() (count int64) {

	data := goqu.Ex{
		"id": goqu.Op{"gte": 1},
	}

	count, err := origin.DeUserTable().GetCount(data)

	logrus.Debug(count)

	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return
}
