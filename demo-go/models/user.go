package models

type User struct {
	Id    int    `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Age   int    `gorm:"column:age" json:"age"` // 数据库列名 + JSON 字段名
	Phone string `gorm:"column:phone" json:"phone"`
}

func (User) TableName() string {
	return "user_list"
}
