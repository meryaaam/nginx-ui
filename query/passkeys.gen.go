// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/0xJacky/Nginx-UI/model"
)

func newPasskey(db *gorm.DB, opts ...gen.DOOption) passkey {
	_passkey := passkey{}

	_passkey.passkeyDo.UseDB(db, opts...)
	_passkey.passkeyDo.UseModel(&model.Passkey{})

	tableName := _passkey.passkeyDo.TableName()
	_passkey.ALL = field.NewAsterisk(tableName)
	_passkey.ID = field.NewUint64(tableName, "id")
	_passkey.CreatedAt = field.NewTime(tableName, "created_at")
	_passkey.UpdatedAt = field.NewTime(tableName, "updated_at")
	_passkey.DeletedAt = field.NewField(tableName, "deleted_at")
	_passkey.Name = field.NewString(tableName, "name")
	_passkey.UserID = field.NewUint64(tableName, "user_id")
	_passkey.RawID = field.NewString(tableName, "raw_id")
	_passkey.Credential = field.NewField(tableName, "credential")
	_passkey.LastUsedAt = field.NewInt64(tableName, "last_used_at")

	_passkey.fillFieldMap()

	return _passkey
}

type passkey struct {
	passkeyDo

	ALL        field.Asterisk
	ID         field.Uint64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Name       field.String
	UserID     field.Uint64
	RawID      field.String
	Credential field.Field
	LastUsedAt field.Int64

	fieldMap map[string]field.Expr
}

func (p passkey) Table(newTableName string) *passkey {
	p.passkeyDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p passkey) As(alias string) *passkey {
	p.passkeyDo.DO = *(p.passkeyDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *passkey) updateTableName(table string) *passkey {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Name = field.NewString(table, "name")
	p.UserID = field.NewUint64(table, "user_id")
	p.RawID = field.NewString(table, "raw_id")
	p.Credential = field.NewField(table, "credential")
	p.LastUsedAt = field.NewInt64(table, "last_used_at")

	p.fillFieldMap()

	return p
}

func (p *passkey) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *passkey) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["name"] = p.Name
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["raw_id"] = p.RawID
	p.fieldMap["credential"] = p.Credential
	p.fieldMap["last_used_at"] = p.LastUsedAt
}

func (p passkey) clone(db *gorm.DB) passkey {
	p.passkeyDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p passkey) replaceDB(db *gorm.DB) passkey {
	p.passkeyDo.ReplaceDB(db)
	return p
}

type passkeyDo struct{ gen.DO }

// FirstByID Where("id=@id")
func (p passkeyDo) FirstByID(id uint64) (result *model.Passkey, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DeleteByID update @@table set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=@id
func (p passkeyDo) DeleteByID(id uint64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("update passkeys set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p passkeyDo) Debug() *passkeyDo {
	return p.withDO(p.DO.Debug())
}

func (p passkeyDo) WithContext(ctx context.Context) *passkeyDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p passkeyDo) ReadDB() *passkeyDo {
	return p.Clauses(dbresolver.Read)
}

func (p passkeyDo) WriteDB() *passkeyDo {
	return p.Clauses(dbresolver.Write)
}

func (p passkeyDo) Session(config *gorm.Session) *passkeyDo {
	return p.withDO(p.DO.Session(config))
}

func (p passkeyDo) Clauses(conds ...clause.Expression) *passkeyDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p passkeyDo) Returning(value interface{}, columns ...string) *passkeyDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p passkeyDo) Not(conds ...gen.Condition) *passkeyDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p passkeyDo) Or(conds ...gen.Condition) *passkeyDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p passkeyDo) Select(conds ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p passkeyDo) Where(conds ...gen.Condition) *passkeyDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p passkeyDo) Order(conds ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p passkeyDo) Distinct(cols ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p passkeyDo) Omit(cols ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p passkeyDo) Join(table schema.Tabler, on ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p passkeyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p passkeyDo) RightJoin(table schema.Tabler, on ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p passkeyDo) Group(cols ...field.Expr) *passkeyDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p passkeyDo) Having(conds ...gen.Condition) *passkeyDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p passkeyDo) Limit(limit int) *passkeyDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p passkeyDo) Offset(offset int) *passkeyDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p passkeyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *passkeyDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p passkeyDo) Unscoped() *passkeyDo {
	return p.withDO(p.DO.Unscoped())
}

func (p passkeyDo) Create(values ...*model.Passkey) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p passkeyDo) CreateInBatches(values []*model.Passkey, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p passkeyDo) Save(values ...*model.Passkey) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p passkeyDo) First() (*model.Passkey, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passkey), nil
	}
}

func (p passkeyDo) Take() (*model.Passkey, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passkey), nil
	}
}

func (p passkeyDo) Last() (*model.Passkey, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passkey), nil
	}
}

func (p passkeyDo) Find() ([]*model.Passkey, error) {
	result, err := p.DO.Find()
	return result.([]*model.Passkey), err
}

func (p passkeyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Passkey, err error) {
	buf := make([]*model.Passkey, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p passkeyDo) FindInBatches(result *[]*model.Passkey, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p passkeyDo) Attrs(attrs ...field.AssignExpr) *passkeyDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p passkeyDo) Assign(attrs ...field.AssignExpr) *passkeyDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p passkeyDo) Joins(fields ...field.RelationField) *passkeyDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p passkeyDo) Preload(fields ...field.RelationField) *passkeyDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p passkeyDo) FirstOrInit() (*model.Passkey, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passkey), nil
	}
}

func (p passkeyDo) FirstOrCreate() (*model.Passkey, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passkey), nil
	}
}

func (p passkeyDo) FindByPage(offset int, limit int) (result []*model.Passkey, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p passkeyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p passkeyDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p passkeyDo) Delete(models ...*model.Passkey) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *passkeyDo) withDO(do gen.Dao) *passkeyDo {
	p.DO = *do.(*gen.DO)
	return p
}