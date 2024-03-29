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

func newSysArticle(db *gorm.DB, opts ...gen.DOOption) sysArticle {
	_sysArticle := sysArticle{}

	_sysArticle.sysArticleDo.UseDB(db, opts...)
	_sysArticle.sysArticleDo.UseModel(&cladminmodel.SysArticle{})

	tableName := _sysArticle.sysArticleDo.TableName()
	_sysArticle.ALL = field.NewAsterisk(tableName)
	_sysArticle.ID = field.NewUint64(tableName, "id")
	_sysArticle.UserID = field.NewUint64(tableName, "user_id")
	_sysArticle.CateID = field.NewUint64(tableName, "cate_id")
	_sysArticle.Title = field.NewString(tableName, "title")
	_sysArticle.Thumb = field.NewString(tableName, "thumb")
	_sysArticle.Content = field.NewString(tableName, "content")
	_sysArticle.ReleaseTime = field.NewString(tableName, "release_time")
	_sysArticle.CreatedAt = field.NewTime(tableName, "created_at")
	_sysArticle.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysArticle.DeletedAt = field.NewField(tableName, "deleted_at")

	_sysArticle.fillFieldMap()

	return _sysArticle
}

type sysArticle struct {
	sysArticleDo sysArticleDo

	ALL         field.Asterisk
	ID          field.Uint64
	UserID      field.Uint64
	CateID      field.Uint64
	Title       field.String
	Thumb       field.String
	Content     field.String
	ReleaseTime field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (s sysArticle) Table(newTableName string) *sysArticle {
	s.sysArticleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysArticle) As(alias string) *sysArticle {
	s.sysArticleDo.DO = *(s.sysArticleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysArticle) updateTableName(table string) *sysArticle {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.UserID = field.NewUint64(table, "user_id")
	s.CateID = field.NewUint64(table, "cate_id")
	s.Title = field.NewString(table, "title")
	s.Thumb = field.NewString(table, "thumb")
	s.Content = field.NewString(table, "content")
	s.ReleaseTime = field.NewString(table, "release_time")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysArticle) WithContext(ctx context.Context) ISysArticleDo {
	return s.sysArticleDo.WithContext(ctx)
}

func (s sysArticle) TableName() string { return s.sysArticleDo.TableName() }

func (s sysArticle) Alias() string { return s.sysArticleDo.Alias() }

func (s *sysArticle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysArticle) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["cate_id"] = s.CateID
	s.fieldMap["title"] = s.Title
	s.fieldMap["thumb"] = s.Thumb
	s.fieldMap["content"] = s.Content
	s.fieldMap["release_time"] = s.ReleaseTime
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysArticle) clone(db *gorm.DB) sysArticle {
	s.sysArticleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysArticle) replaceDB(db *gorm.DB) sysArticle {
	s.sysArticleDo.ReplaceDB(db)
	return s
}

type sysArticleDo struct{ gen.DO }

type ISysArticleDo interface {
	gen.SubQuery
	Debug() ISysArticleDo
	WithContext(ctx context.Context) ISysArticleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysArticleDo
	WriteDB() ISysArticleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysArticleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysArticleDo
	Not(conds ...gen.Condition) ISysArticleDo
	Or(conds ...gen.Condition) ISysArticleDo
	Select(conds ...field.Expr) ISysArticleDo
	Where(conds ...gen.Condition) ISysArticleDo
	Order(conds ...field.Expr) ISysArticleDo
	Distinct(cols ...field.Expr) ISysArticleDo
	Omit(cols ...field.Expr) ISysArticleDo
	Join(table schema.Tabler, on ...field.Expr) ISysArticleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysArticleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysArticleDo
	Group(cols ...field.Expr) ISysArticleDo
	Having(conds ...gen.Condition) ISysArticleDo
	Limit(limit int) ISysArticleDo
	Offset(offset int) ISysArticleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysArticleDo
	Unscoped() ISysArticleDo
	Create(values ...*cladminmodel.SysArticle) error
	CreateInBatches(values []*cladminmodel.SysArticle, batchSize int) error
	Save(values ...*cladminmodel.SysArticle) error
	First() (*cladminmodel.SysArticle, error)
	Take() (*cladminmodel.SysArticle, error)
	Last() (*cladminmodel.SysArticle, error)
	Find() ([]*cladminmodel.SysArticle, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysArticle, err error)
	FindInBatches(result *[]*cladminmodel.SysArticle, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*cladminmodel.SysArticle) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysArticleDo
	Assign(attrs ...field.AssignExpr) ISysArticleDo
	Joins(fields ...field.RelationField) ISysArticleDo
	Preload(fields ...field.RelationField) ISysArticleDo
	FirstOrInit() (*cladminmodel.SysArticle, error)
	FirstOrCreate() (*cladminmodel.SysArticle, error)
	FindByPage(offset int, limit int) (result []*cladminmodel.SysArticle, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysArticleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysArticleDo) Debug() ISysArticleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysArticleDo) WithContext(ctx context.Context) ISysArticleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysArticleDo) ReadDB() ISysArticleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysArticleDo) WriteDB() ISysArticleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysArticleDo) Session(config *gorm.Session) ISysArticleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysArticleDo) Clauses(conds ...clause.Expression) ISysArticleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysArticleDo) Returning(value interface{}, columns ...string) ISysArticleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysArticleDo) Not(conds ...gen.Condition) ISysArticleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysArticleDo) Or(conds ...gen.Condition) ISysArticleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysArticleDo) Select(conds ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysArticleDo) Where(conds ...gen.Condition) ISysArticleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysArticleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysArticleDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysArticleDo) Order(conds ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysArticleDo) Distinct(cols ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysArticleDo) Omit(cols ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysArticleDo) Join(table schema.Tabler, on ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysArticleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysArticleDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysArticleDo) Group(cols ...field.Expr) ISysArticleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysArticleDo) Having(conds ...gen.Condition) ISysArticleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysArticleDo) Limit(limit int) ISysArticleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysArticleDo) Offset(offset int) ISysArticleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysArticleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysArticleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysArticleDo) Unscoped() ISysArticleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysArticleDo) Create(values ...*cladminmodel.SysArticle) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysArticleDo) CreateInBatches(values []*cladminmodel.SysArticle, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysArticleDo) Save(values ...*cladminmodel.SysArticle) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysArticleDo) First() (*cladminmodel.SysArticle, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysArticle), nil
	}
}

func (s sysArticleDo) Take() (*cladminmodel.SysArticle, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysArticle), nil
	}
}

func (s sysArticleDo) Last() (*cladminmodel.SysArticle, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysArticle), nil
	}
}

func (s sysArticleDo) Find() ([]*cladminmodel.SysArticle, error) {
	result, err := s.DO.Find()
	return result.([]*cladminmodel.SysArticle), err
}

func (s sysArticleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*cladminmodel.SysArticle, err error) {
	buf := make([]*cladminmodel.SysArticle, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysArticleDo) FindInBatches(result *[]*cladminmodel.SysArticle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysArticleDo) Attrs(attrs ...field.AssignExpr) ISysArticleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysArticleDo) Assign(attrs ...field.AssignExpr) ISysArticleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysArticleDo) Joins(fields ...field.RelationField) ISysArticleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysArticleDo) Preload(fields ...field.RelationField) ISysArticleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysArticleDo) FirstOrInit() (*cladminmodel.SysArticle, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysArticle), nil
	}
}

func (s sysArticleDo) FirstOrCreate() (*cladminmodel.SysArticle, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*cladminmodel.SysArticle), nil
	}
}

func (s sysArticleDo) FindByPage(offset int, limit int) (result []*cladminmodel.SysArticle, count int64, err error) {
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

func (s sysArticleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysArticleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysArticleDo) Delete(models ...*cladminmodel.SysArticle) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysArticleDo) withDO(do gen.Dao) *sysArticleDo {
	s.DO = *do.(*gen.DO)
	return s
}
