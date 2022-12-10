package models

import (
	"xorm.io/xorm"
)

type BaseModel struct {
	Db *xorm.EngineGroup
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

func (the *BaseModel) Insert(data interface{}) (int64, error) {
	return the.Db.Insert(data)
}

func (the *BaseModel) Select(data interface{}) (bool, error) {
	return the.Db.Get(data)
}

func (the *BaseModel) Update() {
}

func (the *BaseModel) Delete() {
}

// 执行SQL查询
func (the *BaseModel) Query(sql string) (interface{}, error) {
	the.Db.ShowSQL(true)
	return the.Db.QueryInterface(sql)
}

// 执行SQL命令
func (the *BaseModel) Exec(sqlOrArgs ...interface{}) (interface{}, error) {
	the.Db.ShowSQL(true)
	return the.Db.Exec(sqlOrArgs...)
}
