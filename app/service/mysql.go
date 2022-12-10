package service

import (
	"fmt"
	"rango/models/golang"
)

func Insert() int64 {
	User := new(golang.User)

	User.Name = "zhangsan"
	User.Age = 18
	User.CreateTime = "2022-05-20 12:12:12"

	affected, err := golang.UserTable().Insert(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Select() bool {
	var User golang.User

	result, err := golang.UserTable().Select(&User)

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

	affected, err := golang.UserTable().Db.Where("id = 2").Update(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Delete() int64 {
	User := new(golang.User)
	affected, err := golang.UserTable().Db.Where("id = 1").Delete(User)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	return affected
}

func Query() interface{} {
	sql := "select * from user where name = 'zhangsan'"
	affected, _ := golang.UserTable().Query(sql)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	return affected
}
