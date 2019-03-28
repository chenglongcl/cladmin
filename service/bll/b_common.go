package bll

import (
	"cladmin/service/menu_service"
	"cladmin/service/role_service"
	"cladmin/service/user_service"
)

type Common struct {
	UserAPI *user_service.User `inject:""`
	RoleAPI *role_service.Role `inject:""`
	MenuAPI *menu_service.Menu `inject:""`
}
