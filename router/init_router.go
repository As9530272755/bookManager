package router

import (
	"github.com/gin-gonic/gin"
)

/*
加载其他路由文件中的路由
*/

// 这个方法作用加载或者初始化其他文件中的路由
func Init_router() *gin.Engine {
	r := gin.Default()

	// 传递 gin.Engine ，实现注册
	LoadTestRouter(r)

	// 注册 api 路由
	LoadApiRouter(r)
	return r
}
