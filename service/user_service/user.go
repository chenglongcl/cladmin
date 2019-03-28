package user_service

import (
	"cladmin/model"
	"cladmin/pkg/auth"
	"cladmin/pkg/errno"
	"cladmin/util"
	"fmt"
	"github.com/casbin/casbin"
	"sync"
)

type User struct {
	Id           uint64
	Username     string
	Password     string
	Mobile       string
	Email        string
	Status       int64
	CreateUserId uint64
	RoleId       []int64
	Enforcer     *casbin.Enforcer `inject:""`
}

func (a *User) Add() (id uint64, errNo *errno.Errno) {
	if userExist, _ := model.CheckUserByUsername(a.Username); userExist {
		return 0, errno.ErrUserExist
	}
	password, _ := auth.Encrypt(a.Password)
	data := map[string]interface{}{
		"username":       a.Username,
		"password":       password,
		"mobile":         a.Mobile,
		"email":          a.Email,
		"status":         a.Status,
		"create_user_id": a.CreateUserId,
		"role_id":        a.RoleId,
	}
	id, err := model.AddUser(data)
	if err != nil {
		return 0, errno.ErrDatabase
	}
	return id, nil
}

func (a *User) Get() (user *model.User, errNo *errno.Errno) {
	user, err := model.GetUser(a.Id)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return user, nil
}
func (a *User) GetList(ps util.PageSetting) ([]*model.UserInfo, uint64, *errno.Errno) {
	info := make([]*model.UserInfo, 0)
	w := make(map[string]interface{})
	if a.Username != "" {
		w["username like"] = "%" + a.Username + "%"
	}
	users, count, err := model.GetUserList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []uint64
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	finished := make(chan bool, 1)

	for _, u := range users {
		wg.Add(1)
		go func(u *model.User) {
			defer wg.Done()
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				Mobile:    u.Mobile,
				Email:     u.Email,
				Status:    u.Status,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, userList.IdMap[id])
	}
	return info, count, nil
}
func (a *User) Edit() *errno.Errno {
	if userExist, _ := model.CheckUserById(a.Id); !userExist {
		return errno.ErrNotUserExist
	}
	var password string
	if a.Password != "" {
		password, _ = auth.Encrypt(a.Password)
	}
	data := map[string]interface{}{
		"id":       a.Id,
		"password": password,
		"mobile":   a.Mobile,
		"email":    a.Email,
		"status":   a.Status,
		"role_id":  a.RoleId,
	}
	err := model.EditUser(data)
	if err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *User) Delete() *errno.Errno {
	err := model.DeleteUser(a.Id)
	if err != nil {
		return errno.ErrDatabase
	}
	return nil
}

// LoadAllPolicy 加载所有的用户策略
func (a *User) LoadAllPolicy() error {
	users, err := model.GetUsersAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if len(user.Role) != 0 {
			err = a.LoadPolicy(user.Id)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}

// LoadPolicy 加载用户权限策略
func (a *User) LoadPolicy(id uint64) error {
	user, err := model.GetUser(id)
	if err != nil {
		return err
	}
	a.Enforcer.DeleteRolesForUser(user.Username)
	for _, ro := range user.Role {
		a.Enforcer.AddRoleForUser(user.Username, ro.RoleName)
	}
	fmt.Println("更新角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}
