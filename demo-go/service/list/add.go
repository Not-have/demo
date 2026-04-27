package service

import (
	"demoGo/db"
	"demoGo/models"
	"fmt"
)

func GetUserList() []models.User {
	var list []models.User
	err := db.DB.Find(&list).Error

	if err != nil {
		fmt.Println("Error fetching user list:", err)
		return nil
	}

	return list
}
