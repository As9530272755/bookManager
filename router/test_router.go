package router

import "github.com/gin-gonic/gin"

// 接收 *gin.Engine 参数，这样的话就可以在 LoadTestRouter 方法中实现对不同 URL 的处理
func LoadTestRouter(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "LoadTestRouter")
	})
}
