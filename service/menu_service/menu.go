package menu_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"github.com/casbin/casbin"
)

type Menu struct {
	Id       uint64
	ParentId uint64
	Name     string
	Url      string
	Perms    string
	Type     int64
	Icon     string
	OrderNum int64

	Enforcer *casbin.Enforcer `inject:""`
}

func (a *Menu) Add() *errno.Errno {
	data := map[string]interface{}{
		"parent_id": a.ParentId,
		"name":      a.Name,
		"url":       a.Url,
		"perms":     a.Perms,
		"type":      a.Type,
		"icon":      a.Icon,
		"order_num": a.OrderNum,
	}
	if err := model.AddMenu(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Menu) Edit() ([]uint64, *errno.Errno) {
	data := map[string]interface{}{
		"id":        a.Id,
		"parent_id": a.ParentId,
		"name":      a.Name,
		"url":       a.Url,
		"perms":     a.Perms,
		"type":      a.Type,
		"icon":      a.Icon,
		"order_num": a.OrderNum,
	}
	if err := model.EditMenu(data); err != nil {
		return nil, errno.ErrDatabase
	}
	roleList := model.EditMenuGetRoles(a.Id)
	return roleList, nil
}
