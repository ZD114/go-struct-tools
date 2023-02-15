package controller

import (
	"database/sql"
	"zhangda/go-tools/object"
)

type TableStruct struct {
	dsn            string
	savePath       string
	db             *sql.DB
	table          string
	prefix         string
	config         *object.TableConfig
	err            error
	realNameMethod string
	enableJsonTag  bool   // 是否添加json的tag, 默认不添加
	packageName    string // 生成struct的包名(默认为空的话, 则取名为: package model)
	tagKey         string // tag字段的key值,默认是orm
	dateToTime     bool   // 是否将 date相关字段转换为 time.Time,默认否
}

func (ts *TableStruct) Dsn(d string) *TableStruct {
	ts.dsn = d
	return ts
}

func (ts *TableStruct) TagKey(r string) *TableStruct {
	ts.tagKey = r
	return ts
}

func (ts *TableStruct) PackageName(r string) *TableStruct {
	ts.packageName = r
	return ts
}

func (ts *TableStruct) RealNameMethod(r string) *TableStruct {
	ts.realNameMethod = r
	return ts
}

func (ts *TableStruct) SavePath(p string) *TableStruct {
	ts.savePath = p
	return ts
}

func (ts *TableStruct) DB(d *sql.DB) *TableStruct {
	ts.db = d
	return ts
}

func (ts *TableStruct) Table(tab string) *TableStruct {
	ts.table = tab
	return ts
}

func (ts *TableStruct) Prefix(p string) *TableStruct {
	ts.prefix = p
	return ts
}

func (ts *TableStruct) EnableJsonTag(p bool) *TableStruct {
	ts.enableJsonTag = p
	return ts
}

func (ts *TableStruct) DateToTime(d bool) *TableStruct {
	ts.dateToTime = d
	return ts
}

func (ts *TableStruct) Config(c *object.TableConfig) *TableStruct {
	ts.config = c
	return ts
}
