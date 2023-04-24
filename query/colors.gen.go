// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"colorspacer/model"
)

func newColor(db *gorm.DB, opts ...gen.DOOption) color {
	_color := color{}

	_color.colorDo.UseDB(db, opts...)
	_color.colorDo.UseModel(&model.Color{})

	tableName := _color.colorDo.TableName()
	_color.ALL = field.NewAsterisk(tableName)
	_color.ID = field.NewUint(tableName, "id")
	_color.CreatedAt = field.NewTime(tableName, "created_at")
	_color.UpdatedAt = field.NewTime(tableName, "updated_at")
	_color.DeletedAt = field.NewField(tableName, "deleted_at")
	_color.R = field.NewFloat64(tableName, "r")
	_color.G = field.NewFloat64(tableName, "g")
	_color.B = field.NewFloat64(tableName, "b")

	_color.fillFieldMap()

	return _color
}

type color struct {
	colorDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	R         field.Float64
	G         field.Float64
	B         field.Float64

	fieldMap map[string]field.Expr
}

func (c color) Table(newTableName string) *color {
	c.colorDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c color) As(alias string) *color {
	c.colorDo.DO = *(c.colorDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *color) updateTableName(table string) *color {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.R = field.NewFloat64(table, "r")
	c.G = field.NewFloat64(table, "g")
	c.B = field.NewFloat64(table, "b")

	c.fillFieldMap()

	return c
}

func (c *color) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *color) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["r"] = c.R
	c.fieldMap["g"] = c.G
	c.fieldMap["b"] = c.B
}

func (c color) clone(db *gorm.DB) color {
	c.colorDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c color) replaceDB(db *gorm.DB) color {
	c.colorDo.ReplaceDB(db)
	return c
}

type colorDo struct{ gen.DO }

type IColorDo interface {
	gen.SubQuery
	Debug() IColorDo
	WithContext(ctx context.Context) IColorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IColorDo
	WriteDB() IColorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IColorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IColorDo
	Not(conds ...gen.Condition) IColorDo
	Or(conds ...gen.Condition) IColorDo
	Select(conds ...field.Expr) IColorDo
	Where(conds ...gen.Condition) IColorDo
	Order(conds ...field.Expr) IColorDo
	Distinct(cols ...field.Expr) IColorDo
	Omit(cols ...field.Expr) IColorDo
	Join(table schema.Tabler, on ...field.Expr) IColorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IColorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IColorDo
	Group(cols ...field.Expr) IColorDo
	Having(conds ...gen.Condition) IColorDo
	Limit(limit int) IColorDo
	Offset(offset int) IColorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IColorDo
	Unscoped() IColorDo
	Create(values ...*model.Color) error
	CreateInBatches(values []*model.Color, batchSize int) error
	Save(values ...*model.Color) error
	First() (*model.Color, error)
	Take() (*model.Color, error)
	Last() (*model.Color, error)
	Find() ([]*model.Color, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Color, err error)
	FindInBatches(result *[]*model.Color, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Color) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IColorDo
	Assign(attrs ...field.AssignExpr) IColorDo
	Joins(fields ...field.RelationField) IColorDo
	Preload(fields ...field.RelationField) IColorDo
	FirstOrInit() (*model.Color, error)
	FirstOrCreate() (*model.Color, error)
	FindByPage(offset int, limit int) (result []*model.Color, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IColorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c colorDo) Debug() IColorDo {
	return c.withDO(c.DO.Debug())
}

func (c colorDo) WithContext(ctx context.Context) IColorDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c colorDo) ReadDB() IColorDo {
	return c.Clauses(dbresolver.Read)
}

func (c colorDo) WriteDB() IColorDo {
	return c.Clauses(dbresolver.Write)
}

func (c colorDo) Session(config *gorm.Session) IColorDo {
	return c.withDO(c.DO.Session(config))
}

func (c colorDo) Clauses(conds ...clause.Expression) IColorDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c colorDo) Returning(value interface{}, columns ...string) IColorDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c colorDo) Not(conds ...gen.Condition) IColorDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c colorDo) Or(conds ...gen.Condition) IColorDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c colorDo) Select(conds ...field.Expr) IColorDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c colorDo) Where(conds ...gen.Condition) IColorDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c colorDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IColorDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c colorDo) Order(conds ...field.Expr) IColorDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c colorDo) Distinct(cols ...field.Expr) IColorDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c colorDo) Omit(cols ...field.Expr) IColorDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c colorDo) Join(table schema.Tabler, on ...field.Expr) IColorDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c colorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IColorDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c colorDo) RightJoin(table schema.Tabler, on ...field.Expr) IColorDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c colorDo) Group(cols ...field.Expr) IColorDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c colorDo) Having(conds ...gen.Condition) IColorDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c colorDo) Limit(limit int) IColorDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c colorDo) Offset(offset int) IColorDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c colorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IColorDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c colorDo) Unscoped() IColorDo {
	return c.withDO(c.DO.Unscoped())
}

func (c colorDo) Create(values ...*model.Color) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c colorDo) CreateInBatches(values []*model.Color, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c colorDo) Save(values ...*model.Color) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c colorDo) First() (*model.Color, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Color), nil
	}
}

func (c colorDo) Take() (*model.Color, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Color), nil
	}
}

func (c colorDo) Last() (*model.Color, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Color), nil
	}
}

func (c colorDo) Find() ([]*model.Color, error) {
	result, err := c.DO.Find()
	return result.([]*model.Color), err
}

func (c colorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Color, err error) {
	buf := make([]*model.Color, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c colorDo) FindInBatches(result *[]*model.Color, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c colorDo) Attrs(attrs ...field.AssignExpr) IColorDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c colorDo) Assign(attrs ...field.AssignExpr) IColorDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c colorDo) Joins(fields ...field.RelationField) IColorDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c colorDo) Preload(fields ...field.RelationField) IColorDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c colorDo) FirstOrInit() (*model.Color, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Color), nil
	}
}

func (c colorDo) FirstOrCreate() (*model.Color, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Color), nil
	}
}

func (c colorDo) FindByPage(offset int, limit int) (result []*model.Color, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c colorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c colorDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c colorDo) Delete(models ...*model.Color) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *colorDo) withDO(do gen.Dao) *colorDo {
	c.DO = *do.(*gen.DO)
	return c
}
