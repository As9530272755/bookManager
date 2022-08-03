package mysql

import (
	"bookManager/model"
	"log"

	// 由于该 package 是 mysql 所以为了避免冲突这里定义别名为 gmysql
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局 DB ，以便其他 package 函数之间调用
var DB *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(10.0.0.134:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("数据库连接失败：", err)
	}

	// 将实例化之后的 db 赋值给 DB
	DB = db

	// 创建表
	if err := DB.AutoMigrate(&model.User{}, &model.Book{}); err != nil {
		log.Panic(err)
	}
}
