package bll

import "cladmin/service"

type Common struct {
	UserAPI *service.User `inject:""`
	RoleAPI *service.Role `inject:""`
}
