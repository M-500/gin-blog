package controller

import (
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 16:29
//
func RegisterRouter() *gin.Engine {
	Router := gin.New()
	// 基本中间件使用
	Router.Use(middlewares.PaginationMiddleware())

	// use cors
	Router.Use(middlewares.CorsMiddleWare())
	naRouter := Router.Group("/api/v1/na")
	InitNaRouter(naRouter)

	//Router.Use(middlewares.JwtAuth())
	//adminRouter := Router.Group("/api/v1/cms")
	//controller.InitAdminRouter(adminRouter)

	//Router.GET("/debug/pprof/*name", gin.WrapF(pprof.Handler))
	return Router
}
