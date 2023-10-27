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

func newSysDictData(db *gorm.DB, opts ...gen.DOOption) sysDictData {
	_sysDictData := sysDictData{}

	_sysDictData.sysDictDataDo.UseDB(db, opts...)
	_sysDictData.sysDictDataDo.UseModel(&cladminmodel.SysDictData{})

	tableName := _sysDictData.sysDictDataDo.TableName()
	_sysDictData.ALL = field.NewAsterisk(tableName)
	_sysDictData.ID = field.NewUint64(tableName, "id")
	_sysDictData.DictTypeID = field.NewUint64(tableName, "dict_type_id")
	_sysDictData.DictLabel = field.NewString(tableName, "dict_label")
	_sysDictData.DictValue = field.NewString(tableName, "dict_value")
	_sysDictData.Remark = field.NewString(tableName, "remark")
	_sysDictData.Sort = field.NewUint64(tableName, "sort")
	_sysDictData.CreatedAt = field.NewTime(tableName, "created_at")
	_sysDictData.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysDictData.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysDictData.fillFieldMap()

	return _sysDictData
}

type sysDictData struct {
	sysDictDataDo sysDictDataDo

	ALL        field.Asterisk
	ID         field.Uint64
	DictTypeID field.Uint64
	DictLabel  field.String
	DictValue  field.String
	Remark     field.String
	Sort       field.Uint64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (s sysDictData) Table(newTableName string) *sysDictData {
	s.sysDictDataDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictData) As(alias string) *sysDictData {
	s.sysDictDataDo.DO = *(s.sysDictDataDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictData) updateTableName(table string) *sysDictData {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.DictTypeID = field.NewUint64(table, "dict_type_id")
	s.DictLabel = field.NewString(table, "dict_label")
	s.DictValue = field.NewString(table, "dict_value")
	s.Remark = field.NewString(table, "remark")
	s.Sort = field.NewUint64(table, "sort")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysDictData) WithContext(ctx context.Context) ISysDictDataDo {
	return s.sysDictDataDo.WithContext(ctx)
}

func (s sysDictData) TableName() string { return s.sysDictDataDo.TableName() }

func (s sysDictData) Alias() string { return s.sysDictDataDo.Alias() }

func (s *sysDictData) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictData) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["dict_type_id"] = s.DictTypeID
	s.fieldMap["dict_label"] = s.DictLabel
	s.fieldMap["dict_value"] = s.DictValue
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysDictData) clone(db *gorm.DB) sysDictData {
	s.sysDictDataDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictData) replaceDB(db *gorm.DB) sysDictData {
	s.sysDictDataDo.ReplaceDB(db)
	return s
}

type sysDictDataDo struct{ gen.DO }

type ISysDictDataDo interface {
	gen.SubQuery
	Debug() ISysDictDataDo
	WithContext(ctx context.Context) ISysDictDataDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysDictDataDo
	WriteDB() ISysDictDataDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysDictDataDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysDictDataDo
	Not(conds ...gen.Condition) ISysDictDataDo
	Or(conds ...gen.Condition) ISysDictDataDo
	Select(conds ...field.Expr) ISysDictDataDo
	Where(conds ...gen.Condition) ISysDictDataDo
	Order(conds ...field.Expr) ISysDictDataDo
	Distinct(cols ...field.Expr) ISysDictDataDo
	Omit(cols ...field.Expr) ISysDictDataDo
	Join(table schema.Tabler, on ...field.Expr) ISysDictDataDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictDataDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysDictDataDo
	Group(cols ...field.Expr) ISysDictDataDo
	Having(conds ...gen.Condition) ISysDictDataDo
	Limit(limit int) ISysDictDataDo
	Offset(offset int) ISysDictDataDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictDataDo
	Unscoped() ISysDictDataDo
	Create(values ...*cladminmodel.SysDictData) error
	CreateInBatches(values []*cladminmodel.SysDictData, batchSize int) error
	Save(values ...*cladminmodel.SysDictData) error
	First() (*cladminmodel.SysDictData, error)
	Take() (*cladminmodel.SysDictData, error)
	Last() (*cladminmodel.SysDictData, error)
	Find() ([]*cladminmodel.SysDictData, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysDictData, err error)
	FindInBatches(result *[]*cladminmodel.SysDictData, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*cladminmodel.SysDictData) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysDictDataDo
	Assign(attrs ...field.AssignExpr) ISysDictDataDo
	Joins(fields ...field.RelationField) ISysDictDataDo
	Preload(fields ...field.RelationField) ISysDictDataDo
	FirstOrInit() (*cladminmodel.SysDictData, error)
	FirstOrCreate() (*cladminmodel.SysDictData, error)
	FindByPage(offset int, limit int) (result []*cladminmodel.SysDictData, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysDictDataDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysDictDataDo) Debug() ISysDictDataDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictDataDo) WithContext(ctx context.Context) ISysDictDataDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictDataDo) ReadDB() ISysDictDataDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictDataDo) WriteDB() ISysDictDataDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictDataDo) Session(config *gorm.Session) ISysDictDataDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictDataDo) Clauses(conds ...clause.Expression) ISysDictDataDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictDataDo) Returning(value interface{}, columns ...string) ISysDictDataDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictDataDo) Not(conds ...gen.Condition) ISysDictDataDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictDataDo) Or(conds ...gen.Condition) ISysDictDataDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictDataDo) Select(conds ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictDataDo) Where(conds ...gen.Condition) ISysDictDataDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictDataDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysDictDataDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDictDataDo) Order(conds ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictDataDo) Distinct(cols ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictDataDo) Omit(cols ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictDataDo) Join(table schema.Tabler, on ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictDataDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictDataDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictDataDo) Group(cols ...field.Expr) ISysDictDataDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictDataDo) Having(conds ...gen.Condition) ISysDictDataDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictDataDo) Limit(limit int) ISysDictDataDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictDataDo) Offset(offset int) ISysDictDataDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictDataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictDataDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictDataDo) Unscoped() ISysDictDataDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictDataDo) Create(values ...*cladminmodel.SysDictData) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictDataDo) CreateInBatches(values []*cladminmodel.SysDictData, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictDataDo) Save(values ...*cladminmodel.SysDictData) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictDataDo) First() (*cladminmodel.SysDictData, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDictData), nil
	}
}

func (s sysDictDataDo) Take() (*cladminmodel.SysDictData, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDictData), nil
	}
}

func (s sysDictDataDo) Last() (*cladminmodel.SysDictData, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDictData), nil
	}
}

func (s sysDictDataDo) Find() ([]*cladminmodel.SysDictData, error) {
	result, err := s.DO.Find()
	return result.([]*cladminmodel.SysDictData), err
}

func (s sysDictDataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysDictData, err error) {
	buf := make([]*cladminmodel.SysDictData, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictDataDo) FindInBatches(result *[]*cladminmodel.SysDictData, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictDataDo) Attrs(attrs ...field.AssignExpr) ISysDictDataDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictDataDo) Assign(attrs ...field.AssignExpr) ISysDictDataDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictDataDo) Joins(fields ...field.RelationField) ISysDictDataDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictDataDo) Preload(fields ...field.RelationField) ISysDictDataDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictDataDo) FirstOrInit() (*cladminmodel.SysDictData, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDictData), nil
	}
}

func (s sysDictDataDo) FirstOrCreate() (*cladminmodel.SysDictData, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysDictData), nil
	}
}

func (s sysDictDataDo) FindByPage(offset int, limit int) (result []*cladminmodel.SysDictData, count int64, err error) {
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

func (s sysDictDataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictDataDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictDataDo) Delete(models ...*cladminmodel.SysDictData) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictDataDo) withDO(do gen.Dao) *sysDictDataDo {
	s.DO = *do.(*gen.DO)
	return s
}