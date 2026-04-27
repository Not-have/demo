package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 当前项目使用 SQLite，本地数据库文件为 demo-01.db。
	var err error

	dsn := "root:admin@tcp(127.0.0.1:3306)/demo-01?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败")
	}
}
