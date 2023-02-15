package router

import (
	"fmt"
	"net/http"
	"zhangda/go-tools/object"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	server.Use(Recovery)
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	return server
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("router", r)

			c.JSON(http.StatusOK, object.FailMsg("系统内部错误"))
		}
	}()
	c.Next()
}
