package object

import (
	"xorm.io/xorm"

	"zhangda/go-tools/config"
)

var _dbs map[string]*xorm.Engine

func init() {
	_dbs = make(map[string]*xorm.Engine, len(config.GetConfig().Spring.DataSource))

	for k, v := range config.GetConfig().Spring.DataSource {
		if engine, err := xorm.NewEngine("mysql", v.Url); err != nil {
			panic("连接数据库失败, error=" + err.Error())
		} else {
			engine.SetMaxOpenConns(v.MaxOpenConn)
			engine.SetMaxIdleConns(v.MaxIdleConn)

			_dbs[k] = engine
		}
	}
}

func GetDatabases() map[string]*xorm.Engine {
	return _dbs
}

func GetDemandDB() *xorm.Engine {
	return _dbs["demand"]
}
