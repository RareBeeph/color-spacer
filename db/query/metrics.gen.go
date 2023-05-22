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

	"colorspacer/db/model"
)

func newMetric(db *gorm.DB, opts ...gen.DOOption) metric {
	_metric := metric{}

	_metric.metricDo.UseDB(db, opts...)
	_metric.metricDo.UseModel(&model.Metric{})

	tableName := _metric.metricDo.TableName()
	_metric.ALL = field.NewAsterisk(tableName)
	_metric.ID = field.NewUint(tableName, "id")
	_metric.CreatedAt = field.NewTime(tableName, "created_at")
	_metric.UpdatedAt = field.NewTime(tableName, "updated_at")
	_metric.DeletedAt = field.NewField(tableName, "deleted_at")
	_metric.RedSquared = field.NewFloat64(tableName, "red_squared")
	_metric.GreenSquared = field.NewFloat64(tableName, "green_squared")
	_metric.BlueSquared = field.NewFloat64(tableName, "blue_squared")
	_metric.RedDotGreen = field.NewFloat64(tableName, "red_dot_green")
	_metric.RedDotBlue = field.NewFloat64(tableName, "red_dot_blue")
	_metric.GreenDotBlue = field.NewFloat64(tableName, "green_dot_blue")

	_metric.fillFieldMap()

	return _metric
}

type metric struct {
	metricDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	RedSquared   field.Float64
	GreenSquared field.Float64
	BlueSquared  field.Float64
	RedDotGreen  field.Float64
	RedDotBlue   field.Float64
	GreenDotBlue field.Float64

	fieldMap map[string]field.Expr
}

func (m metric) Table(newTableName string) *metric {
	m.metricDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m metric) As(alias string) *metric {
	m.metricDo.DO = *(m.metricDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *metric) updateTableName(table string) *metric {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.RedSquared = field.NewFloat64(table, "red_squared")
	m.GreenSquared = field.NewFloat64(table, "green_squared")
	m.BlueSquared = field.NewFloat64(table, "blue_squared")
	m.RedDotGreen = field.NewFloat64(table, "red_dot_green")
	m.RedDotBlue = field.NewFloat64(table, "red_dot_blue")
	m.GreenDotBlue = field.NewFloat64(table, "green_dot_blue")

	m.fillFieldMap()

	return m
}

func (m *metric) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *metric) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 10)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["red_squared"] = m.RedSquared
	m.fieldMap["green_squared"] = m.GreenSquared
	m.fieldMap["blue_squared"] = m.BlueSquared
	m.fieldMap["red_dot_green"] = m.RedDotGreen
	m.fieldMap["red_dot_blue"] = m.RedDotBlue
	m.fieldMap["green_dot_blue"] = m.GreenDotBlue
}

func (m metric) clone(db *gorm.DB) metric {
	m.metricDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m metric) replaceDB(db *gorm.DB) metric {
	m.metricDo.ReplaceDB(db)
	return m
}

type metricDo struct{ gen.DO }

type IMetricDo interface {
	gen.SubQuery
	Debug() IMetricDo
	WithContext(ctx context.Context) IMetricDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMetricDo
	WriteDB() IMetricDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMetricDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMetricDo
	Not(conds ...gen.Condition) IMetricDo
	Or(conds ...gen.Condition) IMetricDo
	Select(conds ...field.Expr) IMetricDo
	Where(conds ...gen.Condition) IMetricDo
	Order(conds ...field.Expr) IMetricDo
	Distinct(cols ...field.Expr) IMetricDo
	Omit(cols ...field.Expr) IMetricDo
	Join(table schema.Tabler, on ...field.Expr) IMetricDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMetricDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMetricDo
	Group(cols ...field.Expr) IMetricDo
	Having(conds ...gen.Condition) IMetricDo
	Limit(limit int) IMetricDo
	Offset(offset int) IMetricDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMetricDo
	Unscoped() IMetricDo
	Create(values ...*model.Metric) error
	CreateInBatches(values []*model.Metric, batchSize int) error
	Save(values ...*model.Metric) error
	First() (*model.Metric, error)
	Take() (*model.Metric, error)
	Last() (*model.Metric, error)
	Find() ([]*model.Metric, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Metric, err error)
	FindInBatches(result *[]*model.Metric, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Metric) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMetricDo
	Assign(attrs ...field.AssignExpr) IMetricDo
	Joins(fields ...field.RelationField) IMetricDo
	Preload(fields ...field.RelationField) IMetricDo
	FirstOrInit() (*model.Metric, error)
	FirstOrCreate() (*model.Metric, error)
	FindByPage(offset int, limit int) (result []*model.Metric, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMetricDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m metricDo) Debug() IMetricDo {
	return m.withDO(m.DO.Debug())
}

func (m metricDo) WithContext(ctx context.Context) IMetricDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m metricDo) ReadDB() IMetricDo {
	return m.Clauses(dbresolver.Read)
}

func (m metricDo) WriteDB() IMetricDo {
	return m.Clauses(dbresolver.Write)
}

func (m metricDo) Session(config *gorm.Session) IMetricDo {
	return m.withDO(m.DO.Session(config))
}

func (m metricDo) Clauses(conds ...clause.Expression) IMetricDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m metricDo) Returning(value interface{}, columns ...string) IMetricDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m metricDo) Not(conds ...gen.Condition) IMetricDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m metricDo) Or(conds ...gen.Condition) IMetricDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m metricDo) Select(conds ...field.Expr) IMetricDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m metricDo) Where(conds ...gen.Condition) IMetricDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m metricDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMetricDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m metricDo) Order(conds ...field.Expr) IMetricDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m metricDo) Distinct(cols ...field.Expr) IMetricDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m metricDo) Omit(cols ...field.Expr) IMetricDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m metricDo) Join(table schema.Tabler, on ...field.Expr) IMetricDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m metricDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMetricDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m metricDo) RightJoin(table schema.Tabler, on ...field.Expr) IMetricDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m metricDo) Group(cols ...field.Expr) IMetricDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m metricDo) Having(conds ...gen.Condition) IMetricDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m metricDo) Limit(limit int) IMetricDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m metricDo) Offset(offset int) IMetricDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m metricDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMetricDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m metricDo) Unscoped() IMetricDo {
	return m.withDO(m.DO.Unscoped())
}

func (m metricDo) Create(values ...*model.Metric) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m metricDo) CreateInBatches(values []*model.Metric, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m metricDo) Save(values ...*model.Metric) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m metricDo) First() (*model.Metric, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metric), nil
	}
}

func (m metricDo) Take() (*model.Metric, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metric), nil
	}
}

func (m metricDo) Last() (*model.Metric, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metric), nil
	}
}

func (m metricDo) Find() ([]*model.Metric, error) {
	result, err := m.DO.Find()
	return result.([]*model.Metric), err
}

func (m metricDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Metric, err error) {
	buf := make([]*model.Metric, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m metricDo) FindInBatches(result *[]*model.Metric, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m metricDo) Attrs(attrs ...field.AssignExpr) IMetricDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m metricDo) Assign(attrs ...field.AssignExpr) IMetricDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m metricDo) Joins(fields ...field.RelationField) IMetricDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m metricDo) Preload(fields ...field.RelationField) IMetricDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m metricDo) FirstOrInit() (*model.Metric, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metric), nil
	}
}

func (m metricDo) FirstOrCreate() (*model.Metric, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metric), nil
	}
}

func (m metricDo) FindByPage(offset int, limit int) (result []*model.Metric, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m metricDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m metricDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m metricDo) Delete(models ...*model.Metric) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *metricDo) withDO(do gen.Dao) *metricDo {
	m.DO = *do.(*gen.DO)
	return m
}
