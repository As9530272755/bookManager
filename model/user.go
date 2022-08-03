package model

/*
 json:"username":定义 json 反向解析名字
 gorm:"not null"：定义字段在数据库中不能为空
 binding:"required"：定义用户在请求的时候不能传入空值
*/
// 定义 user 表结构
type User struct {
	Id       int    `json:"id" gorm:"primaryKey"` // 自定义主键
	Username string `json:"username" gorm:"not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Token    string `json:"token"`
}

// 自定义表名，因为默认 gorm 会在我们的表后面添加 s
func (User) TableName() string {
	return "user"
}
