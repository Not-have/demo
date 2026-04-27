package service

import (
	"demoGo/db"
	"demoGo/models"
)

func AddUser(name string, age int, phone string) error {
	// 这里可以添加一些业务逻辑，比如验证输入数据的合法性等

	// 创建一个新的 User 实例
	user := models.User{
		Name:  name,
		Age:   age,
		Phone: phone,
	}

	// 将用户信息保存到数据库
	result := db.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
