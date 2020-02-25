package gorm

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Gorm is an interface which DB implements
type Gorm interface {
	Close() error
	DB() *sql.DB
	New() Gorm
	NewScope(value interface{}) *gorm.Scope
	CommonDB() gorm.SQLCommon
	Callback() *gorm.Callback
	SetLogger(l gorm.Logger)
	LogMode(enable bool) Gorm
	SingularTable(enable bool)
	Where(query interface{}, args ...interface{}) Gorm
	Or(query interface{}, args ...interface{}) Gorm
	Not(query interface{}, args ...interface{}) Gorm
	Limit(value int) Gorm
	Offset(value int) Gorm
	Order(value string, reorder ...bool) Gorm
	Select(query interface{}, args ...interface{}) Gorm
	Omit(columns ...string) Gorm
	Group(query string) Gorm
	Having(query string, values ...interface{}) Gorm
	Joins(query string, args ...interface{}) Gorm
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gorm
	Unscoped() Gorm
	Attrs(attrs ...interface{}) Gorm
	Assign(attrs ...interface{}) Gorm
	First(out interface{}, where ...interface{}) Gorm
	Last(out interface{}, where ...interface{}) Gorm
	Find(out interface{}, where ...interface{}) Gorm
	Scan(dest interface{}) Gorm
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) Gorm
	Count(value interface{}) Gorm
	Related(value interface{}, foreignKeys ...string) Gorm
	FirstOrInit(out interface{}, where ...interface{}) Gorm
	FirstOrCreate(out interface{}, where ...interface{}) Gorm
	Update(attrs ...interface{}) Gorm
	Updates(values interface{}, ignoreProtectedAttrs ...bool) Gorm
	UpdateColumn(attrs ...interface{}) Gorm
	UpdateColumns(values interface{}) Gorm
	Save(value interface{}) Gorm
	Create(value interface{}) Gorm
	Delete(value interface{}, where ...interface{}) Gorm
	Raw(sql string, values ...interface{}) Gorm
	Exec(sql string, values ...interface{}) Gorm
	Model(value interface{}) Gorm
	Table(name string) Gorm
	Debug() Gorm
	Transaction(fc func(tx *gorm.DB) error) error
	Begin() Gorm
	Commit() Gorm
	Rollback() Gorm
	RollbackUnlessCommitted() Gorm
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(values ...interface{}) Gorm
	DropTable(values ...interface{}) Gorm
	DropTableIfExists(values ...interface{}) Gorm
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) Gorm
	ModifyColumn(column string, typ string) Gorm
	DropColumn(column string) Gorm
	AddIndex(indexName string, column ...string) Gorm
	AddUniqueIndex(indexName string, column ...string) Gorm
	RemoveIndex(indexName string) Gorm
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) Gorm
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) Gorm
	Set(name string, value interface{}) Gorm
	InstantSet(name string, value interface{}) Gorm
	Get(name string) (value interface{}, ok bool)
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	AddError(err error) error
	GetErrors() (errors []error)

	// extra
	Error() error
	RowsAffected() int64
}

type gormw struct {
	db *gorm.DB
}

// Open is a drop-in replacement for Open()
func Open(dialect string, args ...interface{}) (db Gorm, err error) {
	gormdb, err := gorm.Open(dialect, args...)
	return wrap(gormdb), err
}

// wrap wraps gorm.DB in an interface
func wrap(db *gorm.DB) Gorm {
	return &gormw{db}
}

func (w *gormw) Close() error {
	return w.db.Close()
}

func (w *gormw) DB() *sql.DB {
	return w.db.DB()
}

func (w *gormw) New() Gorm {
	return wrap(w.db.New())
}

func (w *gormw) NewScope(value interface{}) *gorm.Scope {
	return w.db.NewScope(value)
}

func (w *gormw) CommonDB() gorm.SQLCommon {
	return w.db.CommonDB()
}

func (w *gormw) Callback() *gorm.Callback {
	return w.db.Callback()
}

func (w *gormw) SetLogger(log gorm.Logger) {
	w.db.SetLogger(log)
}

func (w *gormw) LogMode(enable bool) Gorm {
	return wrap(w.db.LogMode(enable))
}

func (w *gormw) SingularTable(enable bool) {
	w.db.SingularTable(enable)
}

func (w *gormw) Where(query interface{}, args ...interface{}) Gorm {
	return wrap(w.db.Where(query, args...))
}

func (w *gormw) Or(query interface{}, args ...interface{}) Gorm {
	return wrap(w.db.Or(query, args...))
}

func (w *gormw) Not(query interface{}, args ...interface{}) Gorm {
	return wrap(w.db.Not(query, args...))
}

func (w *gormw) Limit(value int) Gorm {
	return wrap(w.db.Limit(value))
}

func (w *gormw) Offset(value int) Gorm {
	return wrap(w.db.Offset(value))
}

func (w *gormw) Order(value string, reorder ...bool) Gorm {
	return wrap(w.db.Order(value, reorder...))
}

func (w *gormw) Select(query interface{}, args ...interface{}) Gorm {
	return wrap(w.db.Select(query, args...))
}

func (w *gormw) Omit(columns ...string) Gorm {
	return wrap(w.db.Omit(columns...))
}

func (w *gormw) Group(query string) Gorm {
	return wrap(w.db.Group(query))
}

func (w *gormw) Having(query string, values ...interface{}) Gorm {
	return wrap(w.db.Having(query, values...))
}

func (w *gormw) Joins(query string, args ...interface{}) Gorm {
	return wrap(w.db.Joins(query, args...))
}

func (w *gormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gorm {
	return wrap(w.db.Scopes(funcs...))
}

func (w *gormw) Unscoped() Gorm {
	return wrap(w.db.Unscoped())
}

func (w *gormw) Attrs(attrs ...interface{}) Gorm {
	return wrap(w.db.Attrs(attrs...))
}

func (w *gormw) Assign(attrs ...interface{}) Gorm {
	return wrap(w.db.Assign(attrs...))
}

func (w *gormw) First(out interface{}, where ...interface{}) Gorm {
	return wrap(w.db.First(out, where...))
}

func (w *gormw) Last(out interface{}, where ...interface{}) Gorm {
	return wrap(w.db.Last(out, where...))
}

func (w *gormw) Find(out interface{}, where ...interface{}) Gorm {
	return wrap(w.db.Find(out, where...))
}

func (w *gormw) Scan(dest interface{}) Gorm {
	return wrap(w.db.Scan(dest))
}

func (w *gormw) Row() *sql.Row {
	return w.db.Row()
}

func (w *gormw) Rows() (*sql.Rows, error) {
	return w.db.Rows()
}

func (w *gormw) ScanRows(rows *sql.Rows, result interface{}) error {
	return w.db.ScanRows(rows, result)
}

func (w *gormw) Pluck(column string, value interface{}) Gorm {
	return wrap(w.db.Pluck(column, value))
}

func (w *gormw) Count(value interface{}) Gorm {
	return wrap(w.db.Count(value))
}

func (w *gormw) Related(value interface{}, foreignKeys ...string) Gorm {
	return wrap(w.db.Related(value, foreignKeys...))
}

func (w *gormw) FirstOrInit(out interface{}, where ...interface{}) Gorm {
	return wrap(w.db.FirstOrInit(out, where...))
}

func (w *gormw) FirstOrCreate(out interface{}, where ...interface{}) Gorm {
	return wrap(w.db.FirstOrCreate(out, where...))
}

func (w *gormw) Update(attrs ...interface{}) Gorm {
	return wrap(w.db.Update(attrs...))
}

func (w *gormw) Updates(values interface{}, ignoreProtectedAttrs ...bool) Gorm {
	return wrap(w.db.Updates(values, ignoreProtectedAttrs...))
}

func (w *gormw) UpdateColumn(attrs ...interface{}) Gorm {
	return wrap(w.db.UpdateColumn(attrs...))
}

func (w *gormw) UpdateColumns(values interface{}) Gorm {
	return wrap(w.db.UpdateColumns(values))
}

func (w *gormw) Save(value interface{}) Gorm {
	return wrap(w.db.Save(value))
}

func (w *gormw) Create(value interface{}) Gorm {
	return wrap(w.db.Create(value))
}

func (w *gormw) Delete(value interface{}, where ...interface{}) Gorm {
	return wrap(w.db.Delete(value, where...))
}

func (w *gormw) Raw(sql string, values ...interface{}) Gorm {
	return wrap(w.db.Raw(sql, values...))
}

func (w *gormw) Exec(sql string, values ...interface{}) Gorm {
	return wrap(w.db.Exec(sql, values...))
}

func (w *gormw) Model(value interface{}) Gorm {
	return wrap(w.db.Model(value))
}

func (w *gormw) Table(name string) Gorm {
	return wrap(w.db.Table(name))
}

func (w *gormw) Debug() Gorm {
	return wrap(w.db.Debug())
}

func (w *gormw) Transaction(fc func(tx *gorm.DB) error) error {
	return w.db.Transaction(fc)
}

func (w *gormw) Begin() Gorm {
	return wrap(w.db.Begin())
}

func (w *gormw) Commit() Gorm {
	return wrap(w.db.Commit())
}

func (w *gormw) Rollback() Gorm {
	return wrap(w.db.Rollback())
}

func (w *gormw) RollbackUnlessCommitted() Gorm {
	return wrap(w.db.RollbackUnlessCommitted())
}

func (w *gormw) NewRecord(value interface{}) bool {
	return w.db.NewRecord(value)
}

func (w *gormw) RecordNotFound() bool {
	return w.db.RecordNotFound()
}

func (w *gormw) CreateTable(values ...interface{}) Gorm {
	return wrap(w.db.CreateTable(values...))
}

func (w *gormw) DropTable(values ...interface{}) Gorm {
	return wrap(w.db.DropTable(values...))
}

func (w *gormw) DropTableIfExists(values ...interface{}) Gorm {
	return wrap(w.db.DropTableIfExists(values...))
}

func (w *gormw) HasTable(value interface{}) bool {
	return w.db.HasTable(value)
}

func (w *gormw) AutoMigrate(values ...interface{}) Gorm {
	return wrap(w.db.AutoMigrate(values...))
}

func (w *gormw) ModifyColumn(column string, typ string) Gorm {
	return wrap(w.db.ModifyColumn(column, typ))
}

func (w *gormw) DropColumn(column string) Gorm {
	return wrap(w.db.DropColumn(column))
}

func (w *gormw) AddIndex(indexName string, columns ...string) Gorm {
	return wrap(w.db.AddIndex(indexName, columns...))
}

func (w *gormw) AddUniqueIndex(indexName string, columns ...string) Gorm {
	return wrap(w.db.AddUniqueIndex(indexName, columns...))
}

func (w *gormw) RemoveIndex(indexName string) Gorm {
	return wrap(w.db.RemoveIndex(indexName))
}

func (w *gormw) Association(column string) *gorm.Association {
	return w.db.Association(column)
}

func (w *gormw) Preload(column string, conditions ...interface{}) Gorm {
	return wrap(w.db.Preload(column, conditions...))
}

func (w *gormw) Set(name string, value interface{}) Gorm {
	return wrap(w.db.Set(name, value))
}

func (w *gormw) InstantSet(name string, value interface{}) Gorm {
	return wrap(w.db.InstantSet(name, value))
}

func (w *gormw) Get(name string) (interface{}, bool) {
	return w.db.Get(name)
}

func (w *gormw) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	w.db.SetJoinTableHandler(source, column, handler)
}

func (w *gormw) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Gorm {
	return wrap(w.db.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (w *gormw) AddError(err error) error {
	return w.db.AddError(err)
}

func (w *gormw) GetErrors() (errors []error) {
	return w.db.GetErrors()
}

func (w *gormw) RowsAffected() int64 {
	return w.db.RowsAffected
}

func (w *gormw) Error() error {
	return w.db.Error
}
