package converter

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"zhangda/go-tools/object"
)

func NewTableStruct() *TableStruct {
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

func (ts *TableStruct) Run() error {

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
	sqlStr := new(strings.Builder)
	sqlStr.WriteString("SELECT COLUMN_NAME,DATA_TYPE,COLUMN_TYPE,IS_NULLABLE,IFNULL(COLUMN_DEFAULT,\"''\"),TABLE_NAME,COLUMN_COMMENT ")
	sqlStr.WriteString("FROM information_schema.COLUMNS ")
	sqlStr.WriteString("WHERE table_schema = DATABASE() ")

	// 是否指定了具体的table
	if ts.table != "" {
		sqlStr.WriteString("AND TABLE_NAME = '" + ts.prefix + ts.table + "' ")
	}

	// sql排序
	sqlStr.WriteString("order by TABLE_NAME asc, ORDINAL_POSITION asc ")

	rows, err := ts.db.Query(sqlStr.String())
	if err != nil {
		log.Println("Error reading table information: ", err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		col := object.Column{}

		err = rows.Scan(&col.ColumnName, &col.Type, &col.ColumnType, &col.Nullable, &col.ColumnDefault, &col.TableName, &col.ColumnComment)

		if err != nil {
			log.Println(err.Error())
			return
		}

		col.Tag = col.ColumnName
		col.ColumnName = ts.camelCase(col.ColumnName, false)
		col.Type = typeForMysqlToGo[col.Type]
		jsonTag := col.Tag

		// 字段首字母本身大写, 是否需要删除tag
		if ts.config.RmTagIfUcFirst &&
			col.ColumnName[0:1] == strings.ToUpper(col.ColumnName[0:1]) {
			col.Tag = "-"

		} else {
			// 是否需要将tag转换成小写
			if ts.config.TagToLower {
				col.Tag = strings.ToLower(col.Tag)
				jsonTag = col.Tag
			}

			if !ts.config.JsonTagToHump {
				jsonTag = ts.camelCase(jsonTag, true)
			}
		}

		if ts.tagKey == "" {
			ts.tagKey = "xorm"
		}

		if ts.enableJsonTag {

			if col.Nullable == "NO" {
				col.Nullable = "notnull"
			}

			if len(col.ColumnDefault) == 0 {
				col.ColumnDefault = "''"
			}

			col.Tag = fmt.Sprintf("`%s:\"%s '%s' %s default(%s) comment('%s')\" json:\"%s\"`", ts.tagKey, col.ColumnType, col.Tag, col.Nullable, col.ColumnDefault, col.ColumnComment, jsonTag)

		} else {
			col.Tag = fmt.Sprintf("`%s:\"%s\"`", ts.tagKey, col.Tag)
		}

		if _, ok := tableColumns[col.TableName]; !ok {
			tableColumns[col.TableName] = []object.Column{}
		}

		tableColumns[col.TableName] = append(tableColumns[col.TableName], col)

	}

	return tableColumns, nil
}

func (ts *TableStruct) camelCase(str string, flag bool) string {
	// 是否有表前缀, 设置了就先去除表前缀
	if ts.prefix != "" {
		str = strings.Replace(str, ts.prefix, "", 1)
	}

	var text string

	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			// 字符长度大于1时
			if ts.config.UcFirstOnly == true {
				text += strings.ToUpper(p[0:1]) + strings.ToLower(p[1:])
			} else {
				text += strings.ToUpper(p[0:1]) + p[1:]
			}
		}
	}

	if flag {
		text = strings.ToLower(text[0:1]) + text[1:]
	}

	return text
}
