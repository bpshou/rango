package models

import (
	"database/sql"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

type BaseModel struct {
	Db        *xorm.EngineGroup
	TableName string
	ShowSql   bool
}

// 实例化数据库
func NewDatabase(database string) BaseModel {
	// 使用数据库模型实例
	db := GetDatabaseXorm(database)
	// 实例化结构体
	return BaseModel{
		Db: db,
	}
}

/**
 * 批量插入
 *
 */
func (m *BaseModel) Insert(data ...interface{}) (lastId int64, rowsAffected int64, err error) {
	sql, _, _ := goqu.Insert(m.TableName).Rows(data...).ToSQL()
	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	result, err := m.Exec(sql)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

/**
 * 修改
 *
 */
func (m *BaseModel) Update(data goqu.Record, where goqu.Ex) (rowsAffected int64, err error) {
	sql, _, _ := goqu.Update(m.TableName).Set(data).Where(where).ToSQL()
	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	result, err := m.Exec(sql)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

/**
 * 删除
 *
 */
func (m *BaseModel) Delete(where goqu.Ex) (rowsAffected int64, err error) {
	sql, _, _ := goqu.Delete(m.TableName).Where(where).ToSQL()
	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	result, err := m.Exec(sql)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

/**
 * 获取列表
 *
 */
func (m *BaseModel) GetList(data goqu.Ex, page int, size int, order map[string]string) ([]map[string]interface{}, error) {
	model := goqu.Select("*").From(m.TableName).Where(data)
	if page > 0 {
		model.Offset(uint(page))
	}
	if size > 0 {
		model.Limit(uint(size))
	}
	if len(order) > 0 {
		for field, sc := range order {
			if sc == "desc" {
				model.Order(goqu.C(field).Desc())
			} else {
				model.Order(goqu.C(field).Asc())
			}
		}
	}
	sql, _, _ := model.ToSQL()
	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return m.Db.QueryInterface(sql)
}

/**
 * 获取单个数据
 *
 */
func (m *BaseModel) GetOne(data goqu.Ex, page int, size int, order map[string]string) ([]map[string]interface{}, error) {
	model := goqu.Select("*").From(m.TableName).Where(data)
	if page > 0 {
		model.Offset(uint(page))
	}
	if size > 0 {
		model.Limit(uint(size))
	}
	if len(order) > 0 {
		for field, sc := range order {
			if sc == "desc" {
				model.Order(goqu.C(field).Desc())
			} else {
				model.Order(goqu.C(field).Asc())
			}
		}
	}
	sql, args, _ := model.ToSQL()
	return m.Db.QueryInterface(sql, args)
}

// ============================== 兼容封装 ==============================
/**
 * 执行SQL查询
 *
 */
func (m *BaseModel) Query(sqlOrArgs ...interface{}) ([]map[string]interface{}, error) {
	m.Db.ShowSQL(m.ShowSql)
	return m.Db.Engine.QueryInterface(sqlOrArgs)
}

/**
 * 执行SQL命令
 *
 */
func (m *BaseModel) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	m.Db.ShowSQL(m.ShowSql)
	return m.Db.Engine.Exec(sqlOrArgs...)
}

// ============================== 兼容封装 ==============================
