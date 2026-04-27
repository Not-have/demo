package models

type User struct {
	Id    int    `gorm:"column:id"` // 数据库列名 + JSON 字段名
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Phone string `gorm:"column:phone"`
}

func (User) TableName() string {
	return "user_list"
}
