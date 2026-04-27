package service

import (
	"demoGo/db"
	"demoGo/models"
	"fmt"
)

func GetUserList() []models.User {
	var list []models.User

	result := db.DB.Where("age LIKE?", "%22%").Find(&list)

	if result.Error != nil {
		fmt.Println("查询数据失败:", result.Error)
	}

	return list
}
