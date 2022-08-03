package router

import (
	"bookManager/controller"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	// 注册 /register 路由
	r.POST("/register", controller.RegisterHanlder)

	// 注册 login 路由
	r.POST("/login", controller.LoginHanlder)

	r.GET("/listuser", controller.ListUserHandler)

	r.DELETE("/deluser/:id", controller.DeleteUserHandler)

	// 实现版本划分,这样在访问的时候就需要加上 /api/v1 的前缀
	v1 := r.Group("/api/v1")

	// 注册添加数据路由
	v1.POST("/book", controller.AddBookHandler)

	// 注册获取书籍路由
	v1.GET("/book", controller.GetBookHandler)

	// 注册查询单本书籍路由 /:id 获取 URL 中的 id 字段
	v1.GET("/book/:id", controller.GetBookDetailHandler)

	// 注册修改书籍路由
	v1.PUT("/book/:id", controller.UpdateBookHandler)
}
