package service

import (
	"fmt"
	"rango/models/origin"

	"github.com/sirupsen/logrus"
)

func Insert() int64 {
	var data []map[string]interface{}
	data = append(data, map[string]interface{}{
		"phone": 19488887777,
	}, map[string]interface{}{
		"phone": 15599990000,
	})
	affected, err := origin.UserTable().Insert(data)
	if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	id, _ := affected.LastInsertId()
	return id
}

func Select() bool {
	var User origin.User

	result, err := origin.UserTable().Select(&User)

	fmt.Print(User)

	if err != nil {
		fmt.Print(err)
		return false
	}
	return result
}

func Update() int64 {
	User := new(origin.User)

	affected, err := origin.UserTable().Db.Where("id = 2").Update(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Delete() int64 {
	User := new(origin.User)
	affected, err := origin.UserTable().Db.Where("id = 1").Delete(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Query() interface{} {
	sql := "select * from user where name = 'zhangsan'"
	affected, _ := origin.UserTable().Query(sql)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	return affected
}
