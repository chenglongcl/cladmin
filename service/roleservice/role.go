package roleservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"cladmin/util"
	"context"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"sync"
)

type role struct {
	ID             uint64
	RoleName       string
	Remark         string
	CreateUserID   uint64
	MenuIDList     []uint64
	Enforcer       *casbin.Enforcer `inject:""`
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Role = *role

func NewRoleService(ctx *gin.Context, opts ...service.Option) Role {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &role{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Role) Add() (*cladminmodel.SysRole, *errno.Errno) {
	var (
		checkRoleModel *cladminmodel.SysRole
		menuModels     []*cladminmodel.SysMenu
		err            error
	)
	//1.检查角色名是否存在
	checkRoleModel, err = cladminquery.Q.WithContext(a.ctx).SysRole.Select(
		cladminquery.Q.SysRole.ID,
	).Where(
		cladminquery.Q.SysRole.RoleName.Eq(a.RoleName),
	).Take()
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	if checkRoleModel != nil && checkRoleModel.ID > 0 {
		return nil, errno.ErrRoleExist
	}
	//2.查询目标菜单
	menuModels, err = cladminquery.Q.WithContext(a.ctx).SysMenu.Where(
		cladminquery.Q.SysMenu.ID.In(a.MenuIDList...),
	).Find()
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	//3.创建角色
	menuIDList, _ := jsoniter.MarshalToString(a.MenuIDList)
	roleModel := &cladminmodel.SysRole{
		RoleName:     a.RoleName,
		Remark:       a.Remark,
		MenuIDList:   menuIDList,
		CreateUserID: a.CreateUserID,
		Menus:        menuModels,
	}
	//3.1 跳过自动创建关联menu
	//见https://gorm.io/zh_CN/docs/associations.html 跳过自动创建、更新many2many
	err = cladminquery.Q.WithContext(a.ctx).SysRole.
		Omit(cladminquery.Q.SysRole.Menus.Field("*")).Create(roleModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return roleModel, nil
}

func (a Role) Edit(roleModel *cladminmodel.SysRole, menuIDs []uint64) *errno.Errno {
	var (
		checkRoleModel *cladminmodel.SysRole
		menuModels     []*cladminmodel.SysMenu
		err            error
	)
	if roleModel == nil || roleModel.ID == 0 {
		return errno.ErrValidation
	}
	//1.检查角色名是否存在
	checkRoleModel, err = cladminquery.Q.WithContext(a.ctx).SysRole.Select(
		cladminquery.Q.SysRole.ID,
	).Where(
		cladminquery.Q.SysRole.RoleName.Eq(roleModel.RoleName),
	).Take()
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	if checkRoleModel != nil && checkRoleModel.ID > 0 && checkRoleModel.ID != roleModel.ID {
		return errno.ErrRoleExist
	}
	//2.查询目标菜单
	menuModels, err = cladminquery.Q.WithContext(a.ctx).SysMenu.Where(
		cladminquery.Q.SysMenu.ID.In(menuIDs...),
	).Find()
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	//3.删除之前的角色&菜单之间的引用
	err = cladminquery.Q.SysRole.Menus.Model(roleModel).Delete(roleModel.Menus...)
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	//4.更新数据
	roleModel.Menus = menuModels
	err = cladminquery.Q.WithContext(a.ctx).SysRole.Save(roleModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	return nil
}

func (a Role) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysRole, *errno.Errno) {
	roleModel, err := cladminquery.Q.WithContext(a.ctx).SysRole.
		Preload(cladminquery.Q.SysRole.Menus).Select(fields...).Where(conditions...).Take()
	return roleModel, gormx.HandleError(err)
}

func (a Role) InfoList(listParams *service.ListParams) ([]*cladminentity.RoleInfo, uint64, *errno.Errno) {
	roleModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, roleModel := range roleModels {
		ids = append(ids, roleModel.ID)
	}
	info := make([]*cladminentity.RoleInfo, 0)
	wg := sync.WaitGroup{}
	roleList := cladminentity.RoleList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.RoleInfo, len(roleModels)),
	}
	finished := make(chan bool, 1)
	for _, roleModel := range roleModels {
		wg.Add(1)
		go func(role *cladminmodel.SysRole) {
			defer wg.Done()
			roleList.Lock.Lock()
			defer roleList.Lock.Unlock()
			var menuIdList []int64
			_ = jsoniter.UnmarshalFromString(role.MenuIDList, &menuIdList)
			roleList.IdMap[role.ID] = &cladminentity.RoleInfo{
				ID:           role.ID,
				RoleName:     role.RoleName,
				Remark:       role.Remark,
				MenuIDList:   menuIdList,
				CreateUserID: role.CreateUserID,
				CreateTime:   role.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}(roleModel)
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
	return info, uint64(count), nil
}

func (a Role) List(listParams *service.ListParams) (result []*cladminmodel.SysRole, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysRole
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysRole.WithContext(a.ctx)
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

func (a Role) Delete(roleModel *cladminmodel.SysRole) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysRole.Unscoped().Select(field.AssociationFields).Delete(roleModel)
	if err != nil {
		return gormx.HandleError(err)
	}
	return nil
}

// LoadAllPolicy 加载所有的角色策略
func (a Role) LoadAllPolicy() error {
	var ctx context.Context
	if a.ctx == nil {
		ctx = context.Background()
	} else {
		ctx = a.ctx
	}
	roleModels, err := cladminquery.Q.WithContext(ctx).SysRole.Preload(cladminquery.Q.SysRole.Menus).Find()
	if err != nil {
		return err
	}
	for _, roleModel := range roleModels {
		err = a.LoadPolicy(roleModel.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadPolicy 加载角色权限策略
func (a Role) LoadPolicy(id uint64) error {
	var ctx context.Context
	if a.ctx == nil {
		ctx = context.Background()
	} else {
		ctx = a.ctx
	}
	roleModel, err := cladminquery.Q.WithContext(ctx).SysRole.
		Preload(cladminquery.Q.SysRole.Menus).
		Where(cladminquery.Q.SysRole.ID.Eq(id)).Take()
	if err != nil {
		return err
	}
	a.Enforcer.DeletePermissionsForUser(roleModel.RoleName)
	for _, menu := range roleModel.Menus {
		if menu.URL == "" {
			continue
		}
		a.Enforcer.AddPermissionForUser(roleModel.RoleName, menu.URL)
	}
	return nil
}
