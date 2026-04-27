package service

import (
	"demoGo/db"
	"demoGo/models"
	"fmt"
)

func GetUserList() []models.User {
	var list []models.User

	result := db.DB.Find(&list)

	if result.Error != nil {
		fmt.Println("查询数据失败:", result.Error)
	}

	return list
}
