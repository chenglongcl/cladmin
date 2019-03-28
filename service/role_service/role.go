package role_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"github.com/casbin/casbin"
	"sync"
)

type Role struct {
	Id           uint64
	RoleName     string
	Remark       string
	CreateUserId uint64
	MenuIdList   []int64
	Enforcer     *casbin.Enforcer `inject:""`
}

func (a *Role) Add() (id uint64, errNo *errno.Errno) {
	data := map[string]interface{}{
		"role_name":      a.RoleName,
		"remark":         a.Remark,
		"create_user_id": a.CreateUserId,
		"menu_id_list":   a.MenuIdList,
	}
	if roleExist, _ := model.CheckRoleByRoleName(data["role_name"].(string)); roleExist {
		return 0, errno.ErrRoleExist
	}
	id, err := model.AddRole(data)
	if err != nil {
		return 0, errno.ErrDatabase
	}
	return id, nil
}

func (a *Role) Get() (*model.Role, *errno.Errno) {
	role, err := model.GetRole(a.Id)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return role, nil
}

func (a *Role) GetList(ps util.PageSetting) ([]*model.RoleInfo, uint64, *errno.Errno) {
	w := make(map[string]interface{})
	if a.RoleName != "" {
		w["role_name like"] = "%" + a.RoleName + "%"
	}
	roles, count, err := model.GetRoleList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []uint64
	for _, role := range roles {
		ids = append(ids, role.Id)
	}

	info := make([]*model.RoleInfo, 0)
	wg := sync.WaitGroup{}
	roleList := model.RoleList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.RoleInfo, len(roles)),
	}
	finished := make(chan bool, 1)

	for _, role := range roles {
		wg.Add(1)
		go func(role *model.Role) {
			defer wg.Done()
			roleList.Lock.Lock()
			defer roleList.Lock.Unlock()
			roleList.IdMap[role.Id] = &model.RoleInfo{
				Id:           role.Id,
				RoleName:     role.RoleName,
				Remark:       role.Remark,
				CreateUserId: role.CreateUserId,
				CreateTime:   role.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}(role)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, roleList.IdMap[id])
	}
	return info, count, nil
}

func (a *Role) Edit() *errno.Errno {
	data := map[string]interface{}{
		"id":           a.Id,
		"role_name":    a.RoleName,
		"remark":       a.Remark,
		"menu_id_list": a.MenuIdList,
	}
	if roleNameExist, _ := model.CheckRoleByRoleNameId(data["id"].(uint64), data["role_name"].(string));
		roleNameExist {
		return errno.ErrRoleNameExist
	}
	err := model.EditRole(data)
	if err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Role) Delete() *errno.Errno {
	err := model.DeleteRole(a.Id)
	if err != nil {
		return errno.ErrDatabase
	}
	return nil
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
