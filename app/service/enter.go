package service

import (
	"rango/app/service/account"
	"rango/app/service/sms"
	"rango/app/service/user"
)

type ServiceGroup struct {
	UserService    user.UserService
	SmsService     sms.SmsService
	AccountService account.AccountService
}

var ServiceGroupApp = new(ServiceGroup)
