package xorm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.EngineGroup

// 参见文档 https://xorm.io/zh/docs/chapter-01/1.engine/
func EngineGroup() {
	conns := []string{
		"root:secret20$@192.168.5.5:3306/golang?charset=utf8mb4",
		"root:secret20$@192.168.5.5:3306/golang?charset=utf8mb4",
	}

	var err error
	engine, err = xorm.NewEngineGroup("mysql", conns)

	if err != nil {
		return
	}

	// results, err := engine.DBMetas()
	_, err = engine.Exec("select * from act_account where id = 1")

	fmt.Println(err.Error())
	if err != nil {
		return
	}
	// fmt.Println("%p", *results)
}
