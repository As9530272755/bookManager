package model

// 自定义第三张表关联关系
/*
该结构体包含了用户与书籍关系
*/
type BookUser struct {
	// 下面两个元素分别是 User、Book 结构体的主键
	BookID int64 `gorm:"primaryKey"`
	UserID int64 `gorm:"primaryKey"`
}
