package main

import (
	"fmt"
	"zhangda/go-tools/config"
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
}
