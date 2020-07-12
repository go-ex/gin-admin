package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ex/gin-admin/routes"
)

/*
引导HTTP初始化
*/

// 监听HTTP
func RunHttpServer() {
	r := gin.Default()

	routes.LoadApiRoutes(r)

	_ = r.Run(":80")
}
