package service

import (
	"fmt"
	"rango/models"
	"rango/models/golang"
)

func Insert() int64 {
	User := new(golang.User)

	User.Name = "zhangsan"
	User.Age = 18
	User.CreateTime = "2022-05-20 12:12:12"

	affected, err := models.Db.Insert(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Select() bool {
	var User golang.User

	result, err := models.Db.Get(&User)

	fmt.Print(User)

	if err != nil {
		fmt.Print(err)
		return false
	}
	return result
}

func Update() int64 {
	User := new(golang.User)

	User.Name = "lisi"

	affected, err := models.Db.Where("id = 2").Update(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Delete() int64 {
	User := new(golang.User)
	affected, err := models.Db.Where("id = 1").Delete(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}
