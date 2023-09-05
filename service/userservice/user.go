package userservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/auth"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"cladmin/util"
	"context"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"sync"
)

type user struct {
	ID             uint64
	Username       string
	Password       string
	Mobile         string
	Email          string
	Status         int32
	CreateUserID   uint64
	RoleIDList     []uint64
	Enforcer       *casbin.Enforcer `inject:""`
	serviceOptions *service.Options
	ctx            *gin.Context
}

type User = *user

func NewUserService(ctx *gin.Context, opts ...service.Option) User {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &user{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a User) Add() (*cladminmodel.SysUser, *errno.Errno) {
	var (
		checkUserModel *cladminmodel.SysUser
		roleModels     []*cladminmodel.SysRole
		err            error
	)
	//1.检查用户名是否存在
	checkUserModel, err = cladminquery.Q.WithContext(a.ctx).SysUser.Select(
		cladminquery.Q.SysUser.ID,
	).Where(
		cladminquery.Q.SysUser.Username.Eq(a.Username),
	).Take()
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	if checkUserModel != nil && checkUserModel.ID > 0 {
		return nil, errno.ErrUserExist
	}
	//2.查询目标角色
	roleModels, err = cladminquery.Q.WithContext(a.ctx).SysRole.Where(
		cladminquery.Q.SysRole.ID.In(a.RoleIDList...),
	).Find()
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	//3.创建用户
	password, _ := auth.Encrypt(a.Password)
	userModel := &cladminmodel.SysUser{
		Username:     a.Username,
		Password:     password,
		Email:        a.Email,
		Mobile:       a.Mobile,
		Status:       a.Status,
		CreateUserID: a.CreateUserID,
		Roles:        roleModels,
	}
	//3.1 跳过自动创建关联menu
	//见https://gorm.io/zh_CN/docs/associations.html 跳过自动创建、更新many2many
	err = cladminquery.Q.WithContext(a.ctx).SysUser.
		Omit(cladminquery.Q.SysUser.Roles.Field("*")).Create(userModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return userModel, nil
}

func (a User) Edit(userModel *cladminmodel.SysUser, roleIDs []uint64) *errno.Errno {
	var (
		checkUserModel *cladminmodel.SysUser
		roleModels     []*cladminmodel.SysRole
		err            error
	)
	if userModel == nil || userModel.ID == 0 {
		return errno.ErrValidation
	}
	//1.检查用户名是否存在
	checkUserModel, err = cladminquery.Q.WithContext(a.ctx).SysUser.Select(
		cladminquery.Q.SysUser.ID,
	).Where(
		cladminquery.Q.SysUser.Username.Eq(userModel.Username),
	).Take()
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	if checkUserModel != nil && checkUserModel.ID > 0 && checkUserModel.ID != userModel.ID {
		return errno.ErrUserExist
	}
	//2.查询目标角色
	roleModels, err = cladminquery.Q.WithContext(a.ctx).SysRole.Where(
		cladminquery.Q.SysRole.ID.In(roleIDs...),
	).Find()
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	//3.删除之前的用户&角色之间的引用
	err = cladminquery.Q.SysUser.Roles.Model(userModel).Delete(userModel.Roles...)
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	//4.更新用户
	userModel.Roles = roleModels
	err = cladminquery.Q.WithContext(a.ctx).SysUser.Save(userModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	return nil
}

func (a User) EditPersonal(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysUser.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a User) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysUser, *errno.Errno) {
	userModel, err := cladminquery.Q.WithContext(a.ctx).SysUser.
		Preload(cladminquery.Q.SysUser.Roles).Select(fields...).Where(conditions...).Take()
	return userModel, gormx.HandleError(err)
}

func (a User) InfoList(listParams *service.ListParams) ([]*cladminentity.UserInfo, uint64, *errno.Errno) {
	userModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, userModel := range userModels {
		ids = append(ids, userModel.ID)
	}
	info := make([]*cladminentity.UserInfo, 0)
	wg := sync.WaitGroup{}
	userList := cladminentity.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.UserInfo, len(userModels)),
	}
	finished := make(chan bool, 1)
	for _, userModel := range userModels {
		wg.Add(1)
		go func(user *cladminmodel.SysUser) {
			defer wg.Done()
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[user.ID] = &cladminentity.UserInfo{
				ID:           user.ID,
				Username:     user.Username,
				Mobile:       user.Mobile,
				Email:        user.Email,
				Status:       user.Status,
				CreateUserID: user.CreateUserID,
				CreateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateTime:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(userModel)
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
	return info, uint64(count), nil
}

func (a User) List(listParams *service.ListParams) (result []*cladminmodel.SysUser, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysUser
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysUser.WithContext(a.ctx)
		qc.ReplaceDB(qc.UnderlyingDB().Order(listParams.Options.CustomDBOrder))
	}
	base := qc.Select(listParams.Fields...).Where(listParams.Conditions...).Order(listParams.Orders...)
	offset, limit := util.MysqlPagination(listParams.PS)
	if !listParams.Options.WithoutCount {
		result, count, err = base.FindByPage(offset, limit)
	} else {
		if limit == -1 {
			result, err = base.Find()
		} else {
			result, err = base.Offset(offset).Limit(limit).Find()
		}
	}
	return
}

func (a User) Delete(userModel *cladminmodel.SysUser) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysUser.Unscoped().Select(field.AssociationFields).Delete(userModel)
	if err != nil {
		return gormx.HandleError(err)
	}
	return nil
}

// LoadAllPolicy 加载所有的用户策略
func (a User) LoadAllPolicy() error {
	var ctx context.Context
	if a.ctx == nil {
		ctx = context.Background()
	} else {
		ctx = a.ctx
	}
	userModels, err := cladminquery.Q.WithContext(ctx).SysUser.Preload(cladminquery.Q.SysUser.Roles).Find()
	if err != nil {
		return err
	}
	for _, userModel := range userModels {
		if len(userModel.Roles) != 0 {
			err = a.LoadPolicy(userModel.ID)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}

// LoadPolicy 加载用户权限策略
func (a User) LoadPolicy(id uint64) error {
	var ctx context.Context
	if a.ctx == nil {
		ctx = context.Background()
	} else {
		ctx = a.ctx
	}
	userModel, err := cladminquery.Q.WithContext(ctx).SysUser.
		Preload(cladminquery.Q.SysUser.Roles).
		Where(cladminquery.Q.SysUser.ID.Eq(id)).Take()
	if err != nil {
		return err
	}
	a.Enforcer.DeleteRolesForUser(userModel.Username)
	for _, role := range userModel.Roles {
		a.Enforcer.AddRoleForUser(userModel.Username, role.RoleName)
	}
	fmt.Println("更新角色权限关系", a.Enforcer.GetGroupingPolicy())
	return nil
}
