package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 当前项目连接本地 MySQL 数据库 demo-01。
	var err error

	dsn := "root:admin@tcp(127.0.0.1:3306)/demo-01?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("连接数据库失败: %v", err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("获取底层数据库连接失败")
	}

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("连接数据库失败: %v", err))
	}
}
