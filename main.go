package main

import (
	"bookManager/dao/mysql"
	"bookManager/router"
	"fmt"
)

func main() {
	// 初始化 mysql 数据库
	mysql.InitMysql()
	fmt.Println("测试连接数据", mysql.DB)

	// 1.将实例化 router 服务的方法拆分到 router 文件下
	r := router.Init_router()

	// 2.启动
	r.Run()
}
