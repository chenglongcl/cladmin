package service

import (
	"cladmin/model"
	"github.com/casbin/casbin"
)

type Role struct {
	Enforcer *casbin.Enforcer `inject:""`
}

// LoadAllPolicy 加载所有的角色策略
func (a *Role) LoadAllPolicy() error {
	roles, err := model.GetRolesAll()
	if err != nil {
		return err
	}
	for _, role := range roles {
		err = a.LoadPolicy(role.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadPolicy 加载角色权限策略
func (a *Role) LoadPolicy(id uint64) error {
	role, err := model.GetRole(id)
	if err != nil {
		return err
	}
	a.Enforcer.DeleteRole(role.RoleName)
	for _, menu := range role.Menu {
		if menu.Url == "" {
			continue
		}
		a.Enforcer.AddPermissionForUser(role.RoleName, menu.Url)
	}
	return nil
}
