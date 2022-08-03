package model

// 定义 Book 表结构
type Book struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	BookName string `json:"bookname" binding:"required"`
	Desc     string `json:"desc" binding:"required"`

	// 与 user 表进行多对多关联，一本书可以被多人借阅
	Users []User `gorm:"many2many:book_user""` // book_user 表示第三张关联表名为 book_user
}

func (Book) TableName() string {
	return "book"
}
