package models

type User struct {
	Id   int    `gorm:"column:id" json:"id"` // 数据库列名 + JSON 字段名
	Name string `gorm:"column:username" json:"name"`
	Age  int    `gorm:"column:age" json:"age"`
}

func (User) TableName() string {
	return "user_list"
}
