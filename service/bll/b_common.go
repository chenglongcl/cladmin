package bll

import (
	"cladmin/service/menuservice"
	"cladmin/service/roleservice"
	"cladmin/service/userservice"
)

type Common struct {
	UserAPI userservice.User `inject:""`
	RoleAPI roleservice.Role `inject:""`
	MenuAPI menuservice.Menu `inject:""`
}
