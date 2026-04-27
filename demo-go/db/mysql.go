package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 当前项目使用 SQLite，本地数据库文件为 demo-01.db。
	var err error
	DB, err = gorm.Open(sqlite.Open("demo-01.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
}
