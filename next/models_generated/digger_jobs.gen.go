// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newDiggerJob(db *gorm.DB, opts ...gen.DOOption) diggerJob {
	_diggerJob := diggerJob{}

	_diggerJob.diggerJobDo.UseDB(db, opts...)
	_diggerJob.diggerJobDo.UseModel(&model.DiggerJob{})

	tableName := _diggerJob.diggerJobDo.TableName()
	_diggerJob.ALL = field.NewAsterisk(tableName)
	_diggerJob.ID = field.NewString(tableName, "id")
	_diggerJob.CreatedAt = field.NewTime(tableName, "created_at")
	_diggerJob.UpdatedAt = field.NewTime(tableName, "updated_at")
	_diggerJob.DeletedAt = field.NewField(tableName, "deleted_at")
	_diggerJob.DiggerJobID = field.NewString(tableName, "digger_job_id")
	_diggerJob.Status = field.NewInt16(tableName, "status")
	_diggerJob.BatchID = field.NewString(tableName, "batch_id")
	_diggerJob.StatusUpdatedAt = field.NewTime(tableName, "status_updated_at")
	_diggerJob.DiggerJobSummaryID = field.NewString(tableName, "digger_job_summary_id")
	_diggerJob.WorkflowFile = field.NewString(tableName, "workflow_file")
	_diggerJob.WorkflowRunURL = field.NewString(tableName, "workflow_run_url")
	_diggerJob.PlanFootprint = field.NewField(tableName, "plan_footprint")
	_diggerJob.PrCommentURL = field.NewString(tableName, "pr_comment_url")
	_diggerJob.TerraformOutput = field.NewString(tableName, "terraform_output")

	_diggerJob.fillFieldMap()

	return _diggerJob
}

type diggerJob struct {
	diggerJobDo

	ALL                field.Asterisk
	ID                 field.String
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field
	DiggerJobID        field.String
	Status             field.Int16
	BatchID            field.String
	StatusUpdatedAt    field.Time
	DiggerJobSummaryID field.String
	WorkflowFile       field.String
	WorkflowRunURL     field.String
	PlanFootprint      field.Field
	PrCommentURL       field.String
	TerraformOutput    field.String

	fieldMap map[string]field.Expr
}

func (d diggerJob) Table(newTableName string) *diggerJob {
	d.diggerJobDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d diggerJob) As(alias string) *diggerJob {
	d.diggerJobDo.DO = *(d.diggerJobDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *diggerJob) updateTableName(table string) *diggerJob {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.DiggerJobID = field.NewString(table, "digger_job_id")
	d.Status = field.NewInt16(table, "status")
	d.BatchID = field.NewString(table, "batch_id")
	d.StatusUpdatedAt = field.NewTime(table, "status_updated_at")
	d.DiggerJobSummaryID = field.NewString(table, "digger_job_summary_id")
	d.WorkflowFile = field.NewString(table, "workflow_file")
	d.WorkflowRunURL = field.NewString(table, "workflow_run_url")
	d.PlanFootprint = field.NewField(table, "plan_footprint")
	d.PrCommentURL = field.NewString(table, "pr_comment_url")
	d.TerraformOutput = field.NewString(table, "terraform_output")

	d.fillFieldMap()

	return d
}

func (d *diggerJob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *diggerJob) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 14)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["digger_job_id"] = d.DiggerJobID
	d.fieldMap["status"] = d.Status
	d.fieldMap["batch_id"] = d.BatchID
	d.fieldMap["status_updated_at"] = d.StatusUpdatedAt
	d.fieldMap["digger_job_summary_id"] = d.DiggerJobSummaryID
	d.fieldMap["workflow_file"] = d.WorkflowFile
	d.fieldMap["workflow_run_url"] = d.WorkflowRunURL
	d.fieldMap["plan_footprint"] = d.PlanFootprint
	d.fieldMap["pr_comment_url"] = d.PrCommentURL
	d.fieldMap["terraform_output"] = d.TerraformOutput
}

func (d diggerJob) clone(db *gorm.DB) diggerJob {
	d.diggerJobDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d diggerJob) replaceDB(db *gorm.DB) diggerJob {
	d.diggerJobDo.ReplaceDB(db)
	return d
}

type diggerJobDo struct{ gen.DO }

type IDiggerJobDo interface {
	gen.SubQuery
	Debug() IDiggerJobDo
	WithContext(ctx context.Context) IDiggerJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDiggerJobDo
	WriteDB() IDiggerJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDiggerJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDiggerJobDo
	Not(conds ...gen.Condition) IDiggerJobDo
	Or(conds ...gen.Condition) IDiggerJobDo
	Select(conds ...field.Expr) IDiggerJobDo
	Where(conds ...gen.Condition) IDiggerJobDo
	Order(conds ...field.Expr) IDiggerJobDo
	Distinct(cols ...field.Expr) IDiggerJobDo
	Omit(cols ...field.Expr) IDiggerJobDo
	Join(table schema.Tabler, on ...field.Expr) IDiggerJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDiggerJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDiggerJobDo
	Group(cols ...field.Expr) IDiggerJobDo
	Having(conds ...gen.Condition) IDiggerJobDo
	Limit(limit int) IDiggerJobDo
	Offset(offset int) IDiggerJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDiggerJobDo
	Unscoped() IDiggerJobDo
	Create(values ...*model.DiggerJob) error
	CreateInBatches(values []*model.DiggerJob, batchSize int) error
	Save(values ...*model.DiggerJob) error
	First() (*model.DiggerJob, error)
	Take() (*model.DiggerJob, error)
	Last() (*model.DiggerJob, error)
	Find() ([]*model.DiggerJob, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DiggerJob, err error)
	FindInBatches(result *[]*model.DiggerJob, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DiggerJob) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDiggerJobDo
	Assign(attrs ...field.AssignExpr) IDiggerJobDo
	Joins(fields ...field.RelationField) IDiggerJobDo
	Preload(fields ...field.RelationField) IDiggerJobDo
	FirstOrInit() (*model.DiggerJob, error)
	FirstOrCreate() (*model.DiggerJob, error)
	FindByPage(offset int, limit int) (result []*model.DiggerJob, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDiggerJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d diggerJobDo) Debug() IDiggerJobDo {
	return d.withDO(d.DO.Debug())
}

func (d diggerJobDo) WithContext(ctx context.Context) IDiggerJobDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d diggerJobDo) ReadDB() IDiggerJobDo {
	return d.Clauses(dbresolver.Read)
}

func (d diggerJobDo) WriteDB() IDiggerJobDo {
	return d.Clauses(dbresolver.Write)
}

func (d diggerJobDo) Session(config *gorm.Session) IDiggerJobDo {
	return d.withDO(d.DO.Session(config))
}

func (d diggerJobDo) Clauses(conds ...clause.Expression) IDiggerJobDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d diggerJobDo) Returning(value interface{}, columns ...string) IDiggerJobDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d diggerJobDo) Not(conds ...gen.Condition) IDiggerJobDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d diggerJobDo) Or(conds ...gen.Condition) IDiggerJobDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d diggerJobDo) Select(conds ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d diggerJobDo) Where(conds ...gen.Condition) IDiggerJobDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d diggerJobDo) Order(conds ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d diggerJobDo) Distinct(cols ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d diggerJobDo) Omit(cols ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d diggerJobDo) Join(table schema.Tabler, on ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d diggerJobDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d diggerJobDo) RightJoin(table schema.Tabler, on ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d diggerJobDo) Group(cols ...field.Expr) IDiggerJobDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d diggerJobDo) Having(conds ...gen.Condition) IDiggerJobDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d diggerJobDo) Limit(limit int) IDiggerJobDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d diggerJobDo) Offset(offset int) IDiggerJobDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d diggerJobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDiggerJobDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d diggerJobDo) Unscoped() IDiggerJobDo {
	return d.withDO(d.DO.Unscoped())
}

func (d diggerJobDo) Create(values ...*model.DiggerJob) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d diggerJobDo) CreateInBatches(values []*model.DiggerJob, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d diggerJobDo) Save(values ...*model.DiggerJob) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d diggerJobDo) First() (*model.DiggerJob, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerJob), nil
	}
}

func (d diggerJobDo) Take() (*model.DiggerJob, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerJob), nil
	}
}

func (d diggerJobDo) Last() (*model.DiggerJob, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerJob), nil
	}
}

func (d diggerJobDo) Find() ([]*model.DiggerJob, error) {
	result, err := d.DO.Find()
	return result.([]*model.DiggerJob), err
}

func (d diggerJobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DiggerJob, err error) {
	buf := make([]*model.DiggerJob, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d diggerJobDo) FindInBatches(result *[]*model.DiggerJob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d diggerJobDo) Attrs(attrs ...field.AssignExpr) IDiggerJobDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d diggerJobDo) Assign(attrs ...field.AssignExpr) IDiggerJobDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d diggerJobDo) Joins(fields ...field.RelationField) IDiggerJobDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d diggerJobDo) Preload(fields ...field.RelationField) IDiggerJobDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d diggerJobDo) FirstOrInit() (*model.DiggerJob, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerJob), nil
	}
}

func (d diggerJobDo) FirstOrCreate() (*model.DiggerJob, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DiggerJob), nil
	}
}

func (d diggerJobDo) FindByPage(offset int, limit int) (result []*model.DiggerJob, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d diggerJobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d diggerJobDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d diggerJobDo) Delete(models ...*model.DiggerJob) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *diggerJobDo) withDO(do gen.Dao) *diggerJobDo {
	d.DO = *do.(*gen.DO)
	return d
}
