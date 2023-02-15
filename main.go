package main

import (
	"fmt"
	"net/http"
	"zhangda/go-tools/config"
	"zhangda/go-tools/log"
	"zhangda/go-tools/router"
	converter "zhangda/go-tools/run"
)

func main() {
	nts := converter.NewTableStruct()

	err := nts.
		EnableJsonTag(config.GetConfig().EnableJsonTag).
		SavePath(config.GetConfig().SavePath).
		Dsn(config.GetConfig().MysqlUrl).
		Run()
	fmt.Println(err)

	log.InitLogger(config.GetConfig().Logging.File, config.GetConfig().Logging.Level, config.GetConfig().Spring.Application.Name, int(config.GetConfig().Server.Port))

	newRouter := router.InitRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.GetConfig().Server.Port),
		Handler:      newRouter,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Logger.Error("服务器启动异常", log.Any("serverError", err))
	}
}
