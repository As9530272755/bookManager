package controller

import (
	"bookManager/dao/mysql"
	"bookManager/model"
	"fmt"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// 注册功能
func RegisterHanlder(c *gin.Context) {
	user := new(model.User)

	// shouldBind 方法参数校验和参数绑定，获取 json 中复杂数据，这里用于用户注册
	if err := c.ShouldBind(user); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// 创建用户
	mysql.DB.Create(user)

	// 回显
	c.JSON(200, gin.H{"msg": "注册成功"})
}

// 登录功能
func LoginHanlder(c *gin.Context) {
	// 用于存放用户从 web 页面输入的数据信息
	user := new(model.User)

	// 对用户输入的数据进行参数校验
	if err := c.ShouldBind(user); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	// 判断当前输入的用户名和密码是否正确
	u := model.User{Username: user.Username, Password: user.Password}

	fmt.Println(u)

	// 通过 where().First().Row() ,对 u 这个实例结构体进行查找，如果未能查找到任何数据那么 rows = nil ，就直接响应客户端账号或密码错误
	if rows := mysql.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(403, gin.H{"msg": "用户名或密码错误！"})
		return
	}

	// 随机生成字符串做为 token
	token := uuid.New().String()

	// 将 token 写入到数据库
	mysql.DB.Model(&u).Update("token", token)

	c.JSON(200, gin.H{"msg": "登录成功", "token": token})

}

// 获取所有用户信息
func ListUserHandler(c *gin.Context) {
	listUser := []model.User{}

	mysql.DB.Find(&listUser)
	c.JSON(200, gin.H{"用户信息": listUser})
}

// 删除用户
func DeleteUserHandler(c *gin.Context) {
	ID := c.Param("id")
	userId, _ := strconv.Atoi(ID)

	mysql.DB.Where("id = ?", userId).Delete(&model.User{})
	c.JSON(200, gin.H{"msg": "用户删除成功！！"})
}
