package models

import (
	"database/sql"

	"github.com/gookit/goutil/maputil"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

type BaseModel struct {
	Db        *xorm.EngineGroup
	GormDb    *gorm.DB
	TableName string
	ShowSql   bool
}

// 实例化数据库
func NewDatabase(database string) BaseModel {
	// 使用数据库模型实例
	db := GetDatabaseXorm(database)
	// gorm 实例
	gormdb := GetDatabaseGorm(database)
	// 实例化结构体
	return BaseModel{
		Db:     db,
		GormDb: gormdb,
	}
}

/**
 * 批量插入
 *
 */
func (m *BaseModel) Insert(data ...interface{}) (lastId int64, rowsAffected int64, err error) {
	sql, _, err := NewSqlGen().InsertSql(m.TableName, data...)
	if err != nil {
		return
	}

	result, err := m.Exec(sql)
	if err != nil {
		logrus.Error("Insert Exec err", err.Error())
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		logrus.Error("Insert LastInsertId err", err.Error())
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

/**
 * 修改
 *
 */
func (m *BaseModel) Update(values interface{}, where goqu.Expression) (rowsAffected int64, err error) {
	sql, _, err := NewSqlGen().UpdateSql(m.TableName, values, where)
	if err != nil {
		return
	}

	result, err := m.Exec(sql)
	if err != nil {
		logrus.Error("Update Exec err", err.Error())
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

/**
 * 删除
 *
 */
func (m *BaseModel) Delete(where goqu.Expression) (rowsAffected int64, err error) {
	sql, _, err := NewSqlGen().DeleteSql(m.TableName, where)
	if err != nil {
		return
	}

	result, err := m.Exec(sql)
	if err != nil {
		logrus.Error("Delete Exec err", err.Error())
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

/**
 * 获取列表
 *
 */
func (m *BaseModel) GetList(where goqu.Expression, page int, size int, order map[string]string) (data []map[string]interface{}, err error) {
	sql, _, err := NewSqlGen().GetListSql(m.TableName, where, cast.ToUint(page), cast.ToUint(size), order)
	if err != nil {
		return
	}

	return m.GormQueryList(sql)
}

/**
 * 获取单个数据
 *
 */
func (m *BaseModel) GetOne(where goqu.Expression, order map[string]string) (data map[string]interface{}, err error) {
	sql, _, err := NewSqlGen().GetOneSql(m.TableName, where, order)
	if err != nil {
		return
	}

	return m.GormQuery(sql)
}

/**
 * 获取单个数据
 *
 */
func (m *BaseModel) GetCount(where goqu.Expression) (count int64, err error) {
	sql, _, err := NewSqlGen().GetCountSql(m.TableName, where)
	if err != nil {
		return
	}

	query, err := m.GormQuery(sql)
	if err != nil {
		logrus.Error("Query err", err.Error())
		return
	}

	count = cast.ToInt64(maputil.DeepGet(query, "count"))
	return
}

// ============================== 兼容封装 ==============================
/**
 * 执行SQL查询
 *
 */
func (m *BaseModel) Query(sqlOrArgs ...interface{}) ([]map[string]interface{}, error) {
	m.Db.ShowSQL(m.ShowSql)
	return m.Db.Engine.QueryInterface(sqlOrArgs...)
}

/**
 * 执行SQL命令
 *
 */
func (m *BaseModel) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	m.Db.ShowSQL(m.ShowSql)
	return m.Db.Engine.Exec(sqlOrArgs...)
}

/**
 * Gorm执行SQL查询 (map)
 *
 */
func (m *BaseModel) GormQuery(sql string, arg ...interface{}) (data map[string]interface{}, err error) {
	m.GormDb.Raw(sql, arg...).Scan(&data)
	return data, nil
}

/**
 * Gorm执行SQL查询 (列表)
 *
 */
func (m *BaseModel) GormQueryList(sql string, arg ...interface{}) (data []map[string]interface{}, err error) {
	m.GormDb.Raw(sql, arg...).Scan(&data)
	return data, nil
}

// ============================== 兼容封装 ==============================
