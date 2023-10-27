// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package cladminquery

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"cladmin/dal/cladmindb/cladminmodel"
)

func newSysMenu(db *gorm.DB, opts ...gen.DOOption) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db, opts...)
	_sysMenu.sysMenuDo.UseModel(&cladminmodel.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewAsterisk(tableName)
	_sysMenu.ID = field.NewUint64(tableName, "id")
	_sysMenu.ParentID = field.NewUint64(tableName, "parent_id")
	_sysMenu.Name = field.NewString(tableName, "name")
	_sysMenu.URL = field.NewString(tableName, "url")
	_sysMenu.Perms = field.NewString(tableName, "perms")
	_sysMenu.Type = field.NewInt64(tableName, "type")
	_sysMenu.Icon = field.NewString(tableName, "icon")
	_sysMenu.OrderNum = field.NewInt64(tableName, "order_num")
	_sysMenu.IsTab = field.NewBool(tableName, "is_tab")
	_sysMenu.Status = field.NewBool(tableName, "status")
	_sysMenu.CreatedAt = field.NewTime(tableName, "created_at")
	_sysMenu.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysMenu.DeletedAt = field.NewField(tableName, "deleted_at")
	_sysMenu.Roles = sysMenuManyToManyRoles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Roles", "cladminmodel.SysRole"),
		Menus: struct {
			field.RelationField
			Roles struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Roles.Menus", "cladminmodel.SysMenu"),
			Roles: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Roles.Menus.Roles", "cladminmodel.SysRole"),
			},
		},
		Users: struct {
			field.RelationField
			Roles struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Roles.Users", "cladminmodel.SysUser"),
			Roles: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Roles.Users.Roles", "cladminmodel.SysRole"),
			},
		},
	}

	_sysMenu.fillFieldMap()

	return _sysMenu
}

type sysMenu struct {
	sysMenuDo sysMenuDo

	ALL       field.Asterisk
	ID        field.Uint64
	ParentID  field.Uint64
	Name      field.String
	URL       field.String
	Perms     field.String
	Type      field.Int64
	Icon      field.String
	OrderNum  field.Int64
	IsTab     field.Bool
	Status    field.Bool
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Roles     sysMenuManyToManyRoles

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.ParentID = field.NewUint64(table, "parent_id")
	s.Name = field.NewString(table, "name")
	s.URL = field.NewString(table, "url")
	s.Perms = field.NewString(table, "perms")
	s.Type = field.NewInt64(table, "type")
	s.Icon = field.NewString(table, "icon")
	s.OrderNum = field.NewInt64(table, "order_num")
	s.IsTab = field.NewBool(table, "is_tab")
	s.Status = field.NewBool(table, "status")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) WithContext(ctx context.Context) ISysMenuDo { return s.sysMenuDo.WithContext(ctx) }

func (s sysMenu) TableName() string { return s.sysMenuDo.TableName() }

func (s sysMenu) Alias() string { return s.sysMenuDo.Alias() }

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["name"] = s.Name
	s.fieldMap["url"] = s.URL
	s.fieldMap["perms"] = s.Perms
	s.fieldMap["type"] = s.Type
	s.fieldMap["icon"] = s.Icon
	s.fieldMap["order_num"] = s.OrderNum
	s.fieldMap["is_tab"] = s.IsTab
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt

}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysMenu) replaceDB(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuManyToManyRoles struct {
	db *gorm.DB

	field.RelationField

	Menus struct {
		field.RelationField
		Roles struct {
			field.RelationField
		}
	}
	Users struct {
		field.RelationField
		Roles struct {
			field.RelationField
		}
	}
}

func (a sysMenuManyToManyRoles) Where(conds ...field.Expr) *sysMenuManyToManyRoles {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a sysMenuManyToManyRoles) WithContext(ctx context.Context) *sysMenuManyToManyRoles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sysMenuManyToManyRoles) Session(session *gorm.Session) *sysMenuManyToManyRoles {
	a.db = a.db.Session(session)
	return &a
}

func (a sysMenuManyToManyRoles) Model(m *cladminmodel.SysMenu) *sysMenuManyToManyRolesTx {
	return &sysMenuManyToManyRolesTx{a.db.Model(m).Association(a.Name())}
}

type sysMenuManyToManyRolesTx struct{ tx *gorm.Association }

func (a sysMenuManyToManyRolesTx) Find() (result []*cladminmodel.SysRole, err error) {
	return result, a.tx.Find(&result)
}

func (a sysMenuManyToManyRolesTx) Append(values ...*cladminmodel.SysRole) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sysMenuManyToManyRolesTx) Replace(values ...*cladminmodel.SysRole) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sysMenuManyToManyRolesTx) Delete(values ...*cladminmodel.SysRole) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sysMenuManyToManyRolesTx) Clear() error {
	return a.tx.Clear()
}

func (a sysMenuManyToManyRolesTx) Count() int64 {
	return a.tx.Count()
}

type sysMenuDo struct{ gen.DO }

type ISysMenuDo interface {
	gen.SubQuery
	Debug() ISysMenuDo
	WithContext(ctx context.Context) ISysMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysMenuDo
	WriteDB() ISysMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysMenuDo
	Not(conds ...gen.Condition) ISysMenuDo
	Or(conds ...gen.Condition) ISysMenuDo
	Select(conds ...field.Expr) ISysMenuDo
	Where(conds ...gen.Condition) ISysMenuDo
	Order(conds ...field.Expr) ISysMenuDo
	Distinct(cols ...field.Expr) ISysMenuDo
	Omit(cols ...field.Expr) ISysMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	Group(cols ...field.Expr) ISysMenuDo
	Having(conds ...gen.Condition) ISysMenuDo
	Limit(limit int) ISysMenuDo
	Offset(offset int) ISysMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo
	Unscoped() ISysMenuDo
	Create(values ...*cladminmodel.SysMenu) error
	CreateInBatches(values []*cladminmodel.SysMenu, batchSize int) error
	Save(values ...*cladminmodel.SysMenu) error
	First() (*cladminmodel.SysMenu, error)
	Take() (*cladminmodel.SysMenu, error)
	Last() (*cladminmodel.SysMenu, error)
	Find() ([]*cladminmodel.SysMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysMenu, err error)
	FindInBatches(result *[]*cladminmodel.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*cladminmodel.SysMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysMenuDo
	Assign(attrs ...field.AssignExpr) ISysMenuDo
	Joins(fields ...field.RelationField) ISysMenuDo
	Preload(fields ...field.RelationField) ISysMenuDo
	FirstOrInit() (*cladminmodel.SysMenu, error)
	FirstOrCreate() (*cladminmodel.SysMenu, error)
	FindByPage(offset int, limit int) (result []*cladminmodel.SysMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysMenuDo) Debug() ISysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) ISysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() ISysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() ISysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Session(config *gorm.Session) ISysMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) ISysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) ISysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysMenuDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysMenuDo) Order(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) ISysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) ISysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() ISysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*cladminmodel.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*cladminmodel.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*cladminmodel.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*cladminmodel.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*cladminmodel.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*cladminmodel.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*cladminmodel.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*cladminmodel.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysMenu, err error) {
	buf := make([]*cladminmodel.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*cladminmodel.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*cladminmodel.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*cladminmodel.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*cladminmodel.SysMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysMenuDo) Delete(models ...*cladminmodel.SysMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
