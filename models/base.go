package models

import (
	"database/sql"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"xorm.io/xorm"
)

type BaseModel struct {
	Db        *xorm.EngineGroup
	TableName string
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

func (m *BaseModel) Insert(data ...interface{}) (sql.Result, error) {
	sql, _, _ := goqu.Insert(m.TableName).Rows(data...).ToSQL()
	sql = strings.ReplaceAll(sql, "\"", "")
	m.Db.ShowSQL(true)
	return m.Db.Engine.Exec(sql)
}

func (m *BaseModel) Select(data interface{}) (bool, error) {
	return m.Db.Get(data)
}

func (m *BaseModel) Update() {
}

func (m *BaseModel) Delete() {
}

// 执行SQL查询
func (m *BaseModel) Query(sql string) ([]map[string]interface{}, error) {
	m.Db.ShowSQL(true)
	return m.Db.Engine.QueryInterface(sql)
}

// 执行SQL命令
func (m *BaseModel) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	m.Db.ShowSQL(true)
	return m.Db.Engine.Exec(sqlOrArgs...)
}
