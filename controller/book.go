package controller

import (
	"bookManager/dao/mysql"
	"bookManager/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加书籍
func AddBookHandler(c *gin.Context) {
	pbook := new(model.Book)
	if err := c.ShouldBind(pbook); err != nil {
		c.JSON(400, gin.H{"err msg": err.Error()})
		return
	}

	mysql.DB.Create(pbook)
	c.JSON(200, gin.H{"msg": "创建成功"})
}

// 查看数据列表
func GetBookHandler(c *gin.Context) {
	// 由于查询多本数据通过 [] 查询
	listBook := []model.Book{}
	mysql.DB.Find(&listBook)
	c.JSON(200, gin.H{"books": listBook})
}

// 查看指定书籍通过 http://10.0.0.134/book/4 这种 URL 获取 id 等于 4 的书籍
func GetBookDetailHandler(c *gin.Context) {
	// 通过 c.Param 获取 ID
	bookIdStr := c.Param("id")

	// 将通过 URL 获取的 id 转换为 int 类型，
	bookIdInt, err := strconv.Atoi(bookIdStr)
	if err != nil {
		log.Panic(err)
	}

	// 将 int 类型强制转换为
	book := model.Book{Id: int64(bookIdInt)}

	if rows := mysql.DB.Where(&book).First(&book).Row(); rows == nil {
		c.JSON(400, gin.H{"msg": "未能查询到该书籍"})
		return
	}

	c.JSON(200, gin.H{"msg": "已经查询到该书籍", "书籍信息": book})
}

// 修改书籍
func UpdateBookHandler(c *gin.Context) {
	// 获取 url 的 ID 字段
	Id := c.Param("id")
	// 转换为 int 类型
	oldBookId, _ := strconv.Atoi(Id)

	// 获取用户在 body 中传入的修改书籍信息
	Book := new(model.Book)

	// 校验数据
	if err := c.ShouldBind(Book); err != nil {
		if err := c.ShouldBind(Book); err != nil {
			c.JSON(400, gin.H{"err msg": err.Error()})
			return
		}
	}

	// 通过 id 过滤我们想要更新的书籍内容为用户在 body 字段传入的信息
	mysql.DB.Model(model.Book{Id: int64(oldBookId)}).Updates(&Book)
	c.JSON(200, gin.H{"msg": "更新", "更新后：": Book})
}
