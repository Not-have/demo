package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// 这里可以放置数据库连接的初始化代码，例如连接 MySQL 数据库。

	dsn := "user:pass@tcp(127.0.0.1:3306)/demo-01?charset=utf8mb4&parseTime=True&loc=Local"

	_, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
}
