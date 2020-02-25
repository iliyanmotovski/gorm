package gorm

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type FakeGorm struct {
	mock.Mock
}

func (f *FakeGorm) Close() error {
	return f.Called().Error(0)
}

func (f *FakeGorm) DB() *sql.DB {
	return f.Called().Get(0).(*sql.DB)
}

func (f *FakeGorm) New() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) NewScope(value interface{}) *gorm.Scope {
	return f.Called(value).Get(0).(*gorm.Scope)
}

func (f *FakeGorm) CommonDB() gorm.SQLCommon {
	return f.Called().Get(0).(gorm.SQLCommon)
}

func (f *FakeGorm) Callback() *gorm.Callback {
	return f.Called().Get(0).(*gorm.Callback)
}

func (f *FakeGorm) SetLogger(log gorm.Logger) {
	f.Called(log)
}

func (f *FakeGorm) LogMode(enable bool) Gorm {
	return f.Called(enable).Get(0).(Gorm)
}

func (f *FakeGorm) SingularTable(enable bool) {
	f.Called(enable)
}

func (f *FakeGorm) Where(query interface{}, args ...interface{}) Gorm {
	return f.Called(query, args).Get(0).(Gorm)
}

func (f *FakeGorm) Or(query interface{}, args ...interface{}) Gorm {
	return f.Called(query, args).Get(0).(Gorm)
}

func (f *FakeGorm) Not(query interface{}, args ...interface{}) Gorm {
	return f.Called(query, args).Get(0).(Gorm)
}

func (f *FakeGorm) Limit(value int) Gorm {
	return f.Called(value).Get(0).(Gorm)
}

func (f *FakeGorm) Offset(value int) Gorm {
	return f.Called(value).Get(0).(Gorm)
}

func (f *FakeGorm) Order(value string, reorder ...bool) Gorm {
	return f.Called(value, reorder).Get(0).(Gorm)
}

func (f *FakeGorm) Select(query interface{}, args ...interface{}) Gorm {
	return f.Called(query, args).Get(0).(Gorm)
}

func (f *FakeGorm) Omit(columns ...string) Gorm {
	return f.Called(columns).Get(0).(Gorm)
}

func (f *FakeGorm) Group(query string) Gorm {
	return f.Called(query).Get(0).(Gorm)
}

func (f *FakeGorm) Having(query string, values ...interface{}) Gorm {
	return f.Called(query, values).Get(0).(Gorm)
}

func (f *FakeGorm) Joins(query string, args ...interface{}) Gorm {
	return f.Called(query, args).Get(0).(Gorm)
}

func (f *FakeGorm) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gorm {
	return f.Called(funcs).Get(0).(Gorm)
}

func (f *FakeGorm) Unscoped() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) Attrs(attrs ...interface{}) Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) Assign(attrs ...interface{}) Gorm {
	return f.Called(attrs).Get(0).(Gorm)
}

func (f *FakeGorm) First(out interface{}, where ...interface{}) Gorm {
	return f.Called(out, where).Get(0).(Gorm)
}

func (f *FakeGorm) Last(out interface{}, where ...interface{}) Gorm {
	return f.Called(out, where).Get(0).(Gorm)
}

func (f *FakeGorm) Find(out interface{}, where ...interface{}) Gorm {
	return f.Called(out, where).Get(0).(Gorm)
}

func (f *FakeGorm) Scan(dest interface{}) Gorm {
	return f.Called(dest).Get(0).(Gorm)
}

func (f *FakeGorm) Row() *sql.Row {
	return f.Called().Get(0).(*sql.Row)
}

func (f *FakeGorm) Rows() (*sql.Rows, error) {
	args := f.Called()
	return args.Get(0).(*sql.Rows), args.Error(1)
}

func (f *FakeGorm) ScanRows(rows *sql.Rows, result interface{}) error {
	return f.Called(rows, result).Error(0)
}

func (f *FakeGorm) Pluck(column string, value interface{}) Gorm {
	return f.Called(column, value).Get(0).(Gorm)
}

func (f *FakeGorm) Count(value interface{}) Gorm {
	return f.Called(value).Get(0).(Gorm)
}

func (f *FakeGorm) Related(value interface{}, foreignKeys ...string) Gorm {
	return f.Called(value, foreignKeys).Get(0).(Gorm)
}

func (f *FakeGorm) FirstOrInit(out interface{}, where ...interface{}) Gorm {
	return f.Called(out, where).Get(0).(Gorm)
}

func (f *FakeGorm) FirstOrCreate(out interface{}, where ...interface{}) Gorm {
	return f.Called(out, where).Get(0).(Gorm)
}

func (f *FakeGorm) Update(attrs ...interface{}) Gorm {
	return f.Called(attrs).Get(0).(Gorm)
}

func (f *FakeGorm) Updates(values interface{}, ignoreProtectedAttrs ...bool) Gorm {
	return f.Called(values, ignoreProtectedAttrs).Get(0).(Gorm)
}

func (f *FakeGorm) UpdateColumn(attrs ...interface{}) Gorm {
	return f.Called(attrs).Get(0).(Gorm)
}

func (f *FakeGorm) UpdateColumns(values interface{}) Gorm {
	return f.Called(values).Get(0).(Gorm)
}

func (f *FakeGorm) Save(value interface{}) Gorm {
	return f.Called(value).Get(0).(Gorm)
}

func (f *FakeGorm) Create(value interface{}) Gorm {
	return f.Called(value).Get(0).(Gorm)
}

func (f *FakeGorm) Delete(value interface{}, where ...interface{}) Gorm {
	return f.Called(value, where).Get(0).(Gorm)
}

func (f *FakeGorm) Raw(sql string, values ...interface{}) Gorm {
	return f.Called(sql, values).Get(0).(Gorm)
}

func (f *FakeGorm) Exec(sql string, values ...interface{}) Gorm {
	return f.Called(sql, values).Get(0).(Gorm)
}

func (f *FakeGorm) Model(value interface{}) Gorm {
	return f.Called(value).Get(0).(Gorm)
}

func (f *FakeGorm) Table(name string) Gorm {
	return f.Called(name).Get(0).(Gorm)
}

func (f *FakeGorm) Debug() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) Transaction(fc func(tx *gorm.DB) error) error {
	return f.Called(fc).Error(0)
}

func (f *FakeGorm) Begin() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) Commit() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) Rollback() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) RollbackUnlessCommitted() Gorm {
	return f.Called().Get(0).(Gorm)
}

func (f *FakeGorm) NewRecord(value interface{}) bool {
	return f.Called(value).Bool(0)
}

func (f *FakeGorm) RecordNotFound() bool {
	return f.Called().Bool(0)
}

func (f *FakeGorm) CreateTable(values ...interface{}) Gorm {
	return f.Called(values).Get(0).(Gorm)
}

func (f *FakeGorm) DropTable(values ...interface{}) Gorm {
	return f.Called(values).Get(0).(Gorm)
}

func (f *FakeGorm) DropTableIfExists(values ...interface{}) Gorm {
	return f.Called(values).Get(0).(Gorm)
}

func (f *FakeGorm) HasTable(value interface{}) bool {
	return f.Called(value).Bool(0)
}

func (f *FakeGorm) AutoMigrate(values ...interface{}) Gorm {
	return f.Called(values).Get(0).(Gorm)
}

func (f *FakeGorm) ModifyColumn(column string, typ string) Gorm {
	return f.Called(column, typ).Get(0).(Gorm)
}

func (f *FakeGorm) DropColumn(column string) Gorm {
	return f.Called(column).Get(0).(Gorm)
}

func (f *FakeGorm) AddIndex(indexName string, columns ...string) Gorm {
	return f.Called(indexName, columns).Get(0).(Gorm)
}

func (f *FakeGorm) AddUniqueIndex(indexName string, columns ...string) Gorm {
	return f.Called(indexName, columns).Get(0).(Gorm)
}

func (f *FakeGorm) RemoveIndex(indexName string) Gorm {
	return f.Called(indexName).Get(0).(Gorm)
}

func (f *FakeGorm) Association(column string) *gorm.Association {
	return f.Called(column).Get(0).(*gorm.Association)
}

func (f *FakeGorm) Preload(column string, conditions ...interface{}) Gorm {
	return f.Called(column, conditions).Get(0).(Gorm)
}

func (f *FakeGorm) Set(name string, value interface{}) Gorm {
	return f.Called(name, value).Get(0).(Gorm)
}

func (f *FakeGorm) InstantSet(name string, value interface{}) Gorm {
	return f.Called(name, value).Get(0).(Gorm)
}

func (f *FakeGorm) Get(name string) (interface{}, bool) {
	args := f.Called(name)
	return args.Get(0).(interface{}), args.Bool(1)
}

func (f *FakeGorm) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	f.Called(source, column, handler)
}

func (f *FakeGorm) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Gorm {
	return f.Called(field, dest, onDelete, onUpdate).Get(0).(Gorm)
}

func (f *FakeGorm) AddError(err error) error {
	return f.Called(err).Error(0)
}

func (f *FakeGorm) GetErrors() (errors []error) {
	return f.Called().Get(0).([]error)
}

func (f *FakeGorm) RowsAffected() int64 {
	return f.Called().Get(0).(int64)
}

func (f *FakeGorm) Error() error {
	return f.Called().Error(0)
}
