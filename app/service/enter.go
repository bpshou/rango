package service

import (
	"rango/app/service/sms"
	"rango/app/service/user"
)

type ServiceGroup struct {
	UserService user.UserService
	SmsService  sms.SmsService
}

var ServiceGroupApp = new(ServiceGroup)
