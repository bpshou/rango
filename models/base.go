package models

import (
	"database/sql"
	"strings"

	"github.com/spf13/cast"

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
	sql, _, err := goqu.Insert(m.TableName).Rows(data...).ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

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
func (m *BaseModel) Update(values interface{}, where goqu.Expression) (rowsAffected int64, err error) {
	sql, _, err := goqu.Update(m.TableName).Set(values).Where(where).ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

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
func (m *BaseModel) Delete(where goqu.Expression) (rowsAffected int64, err error) {
	sql, _, err := goqu.Delete(m.TableName).Where(where).ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

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
func (m *BaseModel) GetList(where goqu.Expression, page int, size int, order map[string]string) (data []map[string]interface{}, err error) {
	model := goqu.Select("*").From(m.TableName).Where(where)
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

	sql, _, err := model.ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return m.Query(sql)
}

/**
 * 获取单个数据
 *
 */
func (m *BaseModel) GetOne(where goqu.Expression, order map[string]string) (data map[string]interface{}, err error) {
	model := goqu.Select("*").From(m.TableName).Where(where).Limit(1)
	if len(order) > 0 {
		for field, sc := range order {
			if sc == "desc" {
				model.Order(goqu.C(field).Desc())
			} else {
				model.Order(goqu.C(field).Asc())
			}
		}
	}

	sql, _, err := model.ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)

	queryList, err := m.Query(sql)
	if err != nil {
		logrus.Error("Query err", err.Error())
		return
	}

	if len(queryList) <= 0 {
		return
	}

	data = queryList[0]
	return
}

/**
 * 获取单个数据
 *
 */
func (m *BaseModel) GetCount(where goqu.Expression) (count int64, err error) {
	model := goqu.Select(goqu.COUNT("*").As("count")).From(m.TableName).Where(where)

	sql, _, err := model.ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)

	queryList, err := m.Query(sql)
	if err != nil {
		logrus.Error("Query err", err.Error())
		return
	}

	if len(queryList) <= 0 {
		return
	}

	count = cast.ToInt64(queryList[0]["count"])
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

// ============================== 兼容封装 ==============================
