package service

import "my-backend/service/user"

type servicePackage struct {
	UserService user.UserService
}

var ServicePackageApp = new(servicePackage)
