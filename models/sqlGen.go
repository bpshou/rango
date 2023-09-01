package models

import (
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
)

type SqlGen struct {
}

func NewSqlGen() *SqlGen {
	return &SqlGen{}
}

/**
 * 批量插入
 *
 */
func (m *SqlGen) InsertSql(tableName string, data ...interface{}) (sql string, args []interface{}, err error) {
	sql, args, err = goqu.Insert(tableName).Rows(data...).ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return
}

/**
 * 修改
 *
 */
func (m *SqlGen) UpdateSql(tableName string, values interface{}, where goqu.Expression) (sql string, args []interface{}, err error) {
	sql, args, err = goqu.Update(tableName).Set(values).Where(where).ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return
}

/**
 * 删除
 *
 */
func (m *SqlGen) DeleteSql(tableName string, where goqu.Expression) (sql string, args []interface{}, err error) {
	sql, args, err = goqu.Delete(tableName).Where(where).ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return
}

/**
 * 获取列表
 *
 */
func (m *SqlGen) GetListSql(tableName string, where goqu.Expression, page uint, size uint, order map[string]string) (sql string, args []interface{}, err error) {
	model := goqu.Select("*").From(tableName).Where(where)
	if page > 0 {
		model = model.Offset((page - 1) * size)
	}
	if size > 0 {
		model = model.Limit(size)
	}
	if len(order) > 0 {
		for field, sc := range order {
			if sc == "desc" {
				model = model.Order(goqu.C(field).Desc())
			} else {
				model = model.Order(goqu.C(field).Asc())
			}
		}
	}

	sql, args, err = model.ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return
}

/**
 * 获取单个数据
 *
 */
func (m *SqlGen) GetOneSql(tableName string, where goqu.Expression, order map[string]string) (sql string, args []interface{}, err error) {
	model := goqu.Select("*").From(tableName).Where(where).Limit(1)
	if len(order) > 0 {
		for field, sc := range order {
			if sc == "desc" {
				model = model.Order(goqu.C(field).Desc())
			} else {
				model = model.Order(goqu.C(field).Asc())
			}
		}
	}

	sql, args, err = model.ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return
}

/**
 * 获取单个数据
 *
 */
func (m *SqlGen) GetCountSql(tableName string, where goqu.Expression) (sql string, args []interface{}, err error) {
	model := goqu.Select(goqu.COUNT("*").As("count")).From(tableName).Where(where)

	sql, args, err = model.ToSQL()
	if err != nil {
		logrus.Error("sql err", err.Error())
		return
	}

	sql = strings.ReplaceAll(sql, "\"", "")
	logrus.Debug(sql)
	return
}
