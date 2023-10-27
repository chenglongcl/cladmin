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

func newSysDept(db *gorm.DB, opts ...gen.DOOption) sysDept {
	_sysDept := sysDept{}

	_sysDept.sysDeptDo.UseDB(db, opts...)
	_sysDept.sysDeptDo.UseModel(&cladminmodel.SysDept{})

	tableName := _sysDept.sysDeptDo.TableName()
	_sysDept.ALL = field.NewAsterisk(tableName)
	_sysDept.ID = field.NewUint64(tableName, "id")
	_sysDept.ParentID = field.NewUint64(tableName, "parent_id")
	_sysDept.Name = field.NewString(tableName, "name")
	_sysDept.Sort = field.NewInt32(tableName, "sort")
	_sysDept.CreatedAt = field.NewTime(tableName, "created_at")
	_sysDept.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysDept.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysDept.fillFieldMap()

	return _sysDept
}

type sysDept struct {
	sysDeptDo sysDeptDo

	ALL       field.Asterisk
	ID        field.Uint64
	ParentID  field.Uint64
	Name      field.String
	Sort      field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s sysDept) Table(newTableName string) *sysDept {
	s.sysDeptDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDept) As(alias string) *sysDept {
	s.sysDeptDo.DO = *(s.sysDeptDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDept) updateTableName(table string) *sysDept {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.ParentID = field.NewUint64(table, "parent_id")
	s.Name = field.NewString(table, "name")
	s.Sort = field.NewInt32(table, "sort")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysDept) WithContext(ctx context.Context) ISysDeptDo { return s.sysDeptDo.WithContext(ctx) }

func (s sysDept) TableName() string { return s.sysDeptDo.TableName() }

func (s sysDept) Alias() string { return s.sysDeptDo.Alias() }

func (s *sysDept) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDept) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["name"] = s.Name
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysDept) clone(db *gorm.DB) sysDept {
	s.sysDeptDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDept) replaceDB(db *gorm.DB) sysDept {
	s.sysDeptDo.ReplaceDB(db)
	return s
}

type sysDeptDo struct{ gen.DO }

type ISysDeptDo interface {
	gen.SubQuery
	Debug() ISysDeptDo
	WithContext(ctx context.Context) ISysDeptDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysDeptDo
	WriteDB() ISysDeptDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysDeptDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysDeptDo
	Not(conds ...gen.Condition) ISysDeptDo
	Or(conds ...gen.Condition) ISysDeptDo
	Select(conds ...field.Expr) ISysDeptDo
	Where(conds ...gen.Condition) ISysDeptDo
	Order(conds ...field.Expr) ISysDeptDo
	Distinct(cols ...field.Expr) ISysDeptDo
	Omit(cols ...field.Expr) ISysDeptDo
	Join(table schema.Tabler, on ...field.Expr) ISysDeptDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysDeptDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysDeptDo
	Group(cols ...field.Expr) ISysDeptDo
	Having(conds ...gen.Condition) ISysDeptDo
	Limit(limit int) ISysDeptDo
	Offset(offset int) ISysDeptDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDeptDo
	Unscoped() ISysDeptDo
	Create(values ...*cladminmodel.SysDept) error
	CreateInBatches(values []*cladminmodel.SysDept, batchSize int) error
	Save(values ...*cladminmodel.SysDept) error
	First() (*cladminmodel.SysDept, error)
	Take() (*cladminmodel.SysDept, error)
	Last() (*cladminmodel.SysDept, error)
	Find() ([]*cladminmodel.SysDept, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysDept, err error)
	FindInBatches(result *[]*cladminmodel.SysDept, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*cladminmodel.SysDept) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysDeptDo
	Assign(attrs ...field.AssignExpr) ISysDeptDo
	Joins(fields ...field.RelationField) ISysDeptDo
	Preload(fields ...field.RelationField) ISysDeptDo
	FirstOrInit() (*cladminmodel.SysDept, error)
	FirstOrCreate() (*cladminmodel.SysDept, error)
	FindByPage(offset int, limit int) (result []*cladminmodel.SysDept, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysDeptDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysDeptDo) Debug() ISysDeptDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDeptDo) WithContext(ctx context.Context) ISysDeptDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDeptDo) ReadDB() ISysDeptDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDeptDo) WriteDB() ISysDeptDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDeptDo) Session(config *gorm.Session) ISysDeptDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDeptDo) Clauses(conds ...clause.Expression) ISysDeptDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDeptDo) Returning(value interface{}, columns ...string) ISysDeptDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDeptDo) Not(conds ...gen.Condition) ISysDeptDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDeptDo) Or(conds ...gen.Condition) ISysDeptDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDeptDo) Select(conds ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDeptDo) Where(conds ...gen.Condition) ISysDeptDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDeptDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysDeptDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDeptDo) Order(conds ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDeptDo) Distinct(cols ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDeptDo) Omit(cols ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDeptDo) Join(table schema.Tabler, on ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDeptDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDeptDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDeptDo) Group(cols ...field.Expr) ISysDeptDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDeptDo) Having(conds ...gen.Condition) ISysDeptDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDeptDo) Limit(limit int) ISysDeptDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDeptDo) Offset(offset int) ISysDeptDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDeptDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDeptDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDeptDo) Unscoped() ISysDeptDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDeptDo) Create(values ...*cladminmodel.SysDept) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDeptDo) CreateInBatches(values []*cladminmodel.SysDept, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDeptDo) Save(values ...*cladminmodel.SysDept) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDeptDo) First() (*cladminmodel.SysDept, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDept), nil
	}
}

func (s sysDeptDo) Take() (*cladminmodel.SysDept, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDept), nil
	}
}

func (s sysDeptDo) Last() (*cladminmodel.SysDept, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDept), nil
	}
}

func (s sysDeptDo) Find() ([]*cladminmodel.SysDept, error) {
	result, err := s.DO.Find()
	return result.([]*cladminmodel.SysDept), err
}

func (s sysDeptDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysDept, err error) {
	buf := make([]*cladminmodel.SysDept, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDeptDo) FindInBatches(result *[]*cladminmodel.SysDept, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDeptDo) Attrs(attrs ...field.AssignExpr) ISysDeptDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDeptDo) Assign(attrs ...field.AssignExpr) ISysDeptDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDeptDo) Joins(fields ...field.RelationField) ISysDeptDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDeptDo) Preload(fields ...field.RelationField) ISysDeptDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDeptDo) FirstOrInit() (*cladminmodel.SysDept, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDept), nil
	}
}

func (s sysDeptDo) FirstOrCreate() (*cladminmodel.SysDept, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDept), nil
	}
}

func (s sysDeptDo) FindByPage(offset int, limit int) (result []*cladminmodel.SysDept, count int64, err error) {
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

func (s sysDeptDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDeptDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDeptDo) Delete(models ...*cladminmodel.SysDept) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDeptDo) withDO(do gen.Dao) *sysDeptDo {
	s.DO = *do.(*gen.DO)
	return s
}
