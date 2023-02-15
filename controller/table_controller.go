package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"zhangda/go-tools/object"
)

func NewTable2Struct() *TableStruct {
	return &TableStruct{}
}

var typeForMysqlToGo = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func (ts *TableStruct) Run(ctx *gin.Context) error {

	if ts.config == nil {
		ts.config = new(object.TableConfig)
	}

	// 链接mysql, 获取db对象
	ts.dialMysql()
	if ts.err != nil {
		return ts.err
	}

	// 获取表和字段的schema
	_, err := ts.getColumns()
	if err != nil {
		return err
	}

	return nil

}

func (ts *TableStruct) dialMysql() {
	if ts.db == nil {
		if ts.dsn == "" {
			ts.err = errors.New("dsn数据库配置缺失")
			return
		}
		ts.db, ts.err = sql.Open("mysql", ts.dsn)
	}
	return
}

func (ts *TableStruct) getColumns(table ...string) (tableColumns map[string][]object.Column, err error) {

	// 根据设置,判断是否要把 date 相关字段替换为 string
	if ts.dateToTime == false {
		typeForMysqlToGo["date"] = "string"
		typeForMysqlToGo["datetime"] = "time.Time"
		typeForMysqlToGo["timestamp"] = "string"
		typeForMysqlToGo["time"] = "string"
	}

	tableColumns = make(map[string][]object.Column)

	// sql
	var sqlStr = `SELECT COLUMN_NAME,DATA_TYPE,COLUMN_TYPE,IS_NULLABLE,IFNULL(COLUMN_DEFAULT,"''"),TABLE_NAME,COLUMN_COMMENT
		FROM information_schema.COLUMNS 
		WHERE table_schema = DATABASE()`

	// 是否指定了具体的table
	if ts.table != "" {
		sqlStr += fmt.Sprintf(" AND TABLE_NAME = '%s'", ts.prefix+ts.table)
	}

	// sql排序
	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"

	return
}
