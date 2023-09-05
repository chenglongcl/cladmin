package menuservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"cladmin/util"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
)

type menu struct {
	ID             uint64
	ParentID       uint64
	Name           string
	URL            string
	Perms          string
	Type           int64
	Icon           string
	OrderNum       int64
	Enforcer       *casbin.Enforcer `inject:""`
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Menu = *menu

func NewMenuService(ctx *gin.Context, opts ...service.Option) Menu {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &menu{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Menu) Add() (*cladminmodel.SysMenu, *errno.Errno) {
	menuModel := &cladminmodel.SysMenu{
		ParentID: a.ParentID,
		Name:     a.Name,
		URL:      a.URL,
		Perms:    a.Perms,
		Type:     a.Type,
		Icon:     a.Icon,
		OrderNum: a.OrderNum,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysMenu.Create(menuModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return menuModel, nil
}

func (a Menu) Edit(menuModel *cladminmodel.SysMenu) *errno.Errno {
	err := cladminquery.Q.WithContext(a.ctx).SysMenu.Omit(field.AssociationFields).Save(menuModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return errNo
	}
	return nil
}

func (a Menu) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysMenu, *errno.Errno) {
	menuModel, err := cladminquery.Q.WithContext(a.ctx).SysMenu.
		Preload(cladminquery.Q.SysMenu.Roles).
		Select(fields...).Where(conditions...).Take()
	return menuModel, gormx.HandleError(err)
}

func (a Menu) InfoList(listParams *service.ListParams) ([]*cladminentity.MenuInfo, uint64, *errno.Errno) {
	menuModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, menuModel := range menuModels {
		ids = append(ids, menuModel.ID)
	}
	info := make([]*cladminentity.MenuInfo, 0)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	menuList := cladminentity.MenuList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.MenuInfo, len(menuModels)),
	}
	for _, menuModel := range menuModels {
		wg.Add(1)
		go func(menuModel *cladminmodel.SysMenu) {
			defer wg.Done()
			menuList.Lock.Lock()
			defer menuList.Lock.Unlock()
			menuList.IdMap[menuModel.ID] = &cladminentity.MenuInfo{
				ID:         menuModel.ID,
				ParentID:   menuModel.ParentID,
				ParentName: "",
				Name:       menuModel.Name,
				Url:        menuModel.URL,
				Perms:      menuModel.Perms,
				Type:       menuModel.Type,
				Icon:       menuModel.Icon,
				Open:       0,
				OrderNum:   menuModel.OrderNum,
				CreateTime: menuModel.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateTime: menuModel.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(menuModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, menuList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a Menu) List(listParams *service.ListParams) (result []*cladminmodel.SysMenu, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysMenu
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysMenu.WithContext(a.ctx)
		qc.ReplaceDB(qc.UnderlyingDB().Order(listParams.Options.CustomDBOrder))
	}
	base := qc.Select(listParams.Fields...).Where(listParams.Conditions...).Order(listParams.Orders...)
	for _, join := range listParams.Joins {
		base = base.Join(join.Table, join.On...)
	}
	if len(listParams.Groups) > 0 {
		base = base.Group(listParams.Groups...)
	}
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

func (a Menu) Delete(menuModel *cladminmodel.SysMenu) *errno.Errno {
	submenuModel, errNo := a.Get([]field.Expr{
		cladminquery.Q.SysMenu.ID,
	}, []gen.Condition{
		cladminquery.Q.SysMenu.ParentID.Eq(menuModel.ID),
	})
	if errNo != nil {
		return errNo
	}
	if submenuModel != nil && submenuModel.ID > 0 {
		return errno.ErrRecordHasChildren
	}
	_, err := cladminquery.Q.WithContext(a.ctx).SysMenu.Unscoped().Select(field.AssociationFields).Delete(menuModel)
	if err != nil {
		return gormx.HandleError(err)
	}
	return nil
}

func (a Menu) GetMenuNavByUserID(userID uint64) ([]*MenuTree, []string, *errno.Errno) {
	var (
		menuModels []*cladminmodel.SysMenu
		err        error
		errNo      *errno.Errno
	)
	menuModels, _, err = a.List(&service.ListParams{
		PS: util.PageSetting{},
		Options: struct {
			WithoutCount  bool
			Scenes        string
			CustomDBOrder string
			CustomFunc    func() interface{}
		}{WithoutCount: true},
		Fields: []field.Expr{
			cladminquery.Q.SysMenu.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if userID != 1 {
				conditions = append(conditions, cladminquery.Q.SysUserRole.UserID.Eq(userID))
			}
			return conditions
		}(), []gen.Condition{}...),
		Joins: func() []struct {
			Table schema.Tabler
			On    []field.Expr
		} {
			joins := make([]struct {
				Table schema.Tabler
				On    []field.Expr
			}, 0)
			if userID != 1 {
				joins = append(joins, []struct {
					Table schema.Tabler
					On    []field.Expr
				}{
					{
						Table: cladminquery.Q.SysRoleMenu,
						On: []field.Expr{
							cladminquery.Q.SysRoleMenu.MenuID.EqCol(cladminquery.Q.SysMenu.ID),
						},
					},
					{
						Table: cladminquery.Q.SysRole,
						On: []field.Expr{
							cladminquery.Q.SysRole.ID.EqCol(cladminquery.Q.SysRoleMenu.RoleID),
						},
					},
					{
						Table: cladminquery.Q.SysUserRole,
						On: []field.Expr{
							cladminquery.Q.SysUserRole.RoleID.EqCol(cladminquery.Q.SysRole.ID),
						},
					},
				}...)
			}
			return joins
		}(),
		Groups: func() []field.Expr {
			if userID != 1 {
				return []field.Expr{
					cladminquery.Q.SysMenu.ID,
				}
			}
			return nil
		}(),
		Orders: []field.Expr{
			cladminquery.Q.SysMenu.ParentID,
			cladminquery.Q.SysMenu.OrderNum,
			cladminquery.Q.SysMenu.ID,
		},
	})
	if errNo = gormx.HandleError(err); errNo != nil {
		return nil, nil, errNo
	}
	var (
		menuTrees MenuTrees
	)
	permissions := make([]string, 0)
	for _, menuModel := range menuModels {
		if menuModel.Type != 2 {
			menuTrees = append(menuTrees, &MenuTree{
				MenuID:     menuModel.ID,
				Name:       menuModel.Name,
				Open:       0,
				OrderNum:   menuModel.OrderNum,
				Icon:       menuModel.Icon,
				Url:        menuModel.URL,
				ParentID:   menuModel.ParentID,
				ParentName: "",
				Perms:      menuModel.Perms,
				Type:       menuModel.Type,
			})
		}
		if menuModel.Perms != "" {
			pSlice := strings.Split(menuModel.Perms, ",")
			permissions = append(permissions, pSlice...)
		}
	}
	list := menuTrees.ToTree()
	return list, permissions, nil
}

// MenuTree 菜单树
type MenuTree struct {
	MenuID     uint64       `json:"menuId"`
	Name       string       `json:"name"`
	Open       int64        `json:"open"`
	OrderNum   int64        `json:"orderNum"`
	Icon       string       `json:"icon"`
	Url        string       `json:"url"`
	ParentID   uint64       `json:"parentId"`
	ParentName string       `json:"parentName"`
	Perms      string       `json:"perms"`
	Type       int64        `json:"type"`
	List       *[]*MenuTree `json:"list,omitempty"`
}

// MenuTrees 菜单树列表
type MenuTrees []*MenuTree

// ForEach 遍历数据项
func (a MenuTrees) ForEach(fn func(*MenuTree, int)) MenuTrees {
	for i, item := range a {
		fn(item, i)
	}
	return a
}

// ToTree 转换为树形结构
func (a MenuTrees) ToTree() []*MenuTree {
	mi := make(map[uint64]*MenuTree)
	for _, item := range a {
		mi[item.MenuID] = item
	}
	var list []*MenuTree
	for _, item := range a {
		if item.ParentID == 0 {
			list = append(list, item)
			continue
		}
		if pItem, ok := mi[item.ParentID]; ok {
			if pItem.List == nil {
				var children []*MenuTree
				children = append(children, item)
				pItem.List = &children
				continue
			}
			*pItem.List = append(*pItem.List, item)
		}
	}
	return list
}
