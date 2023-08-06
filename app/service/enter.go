package service

import "rango/app/service/user"

type ServiceGroup struct {
	UserService user.UserService
}

var ServiceGroupApp = new(ServiceGroup)
