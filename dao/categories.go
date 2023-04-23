package dao

import (
	"fmt"
	"sample-go-service/global"
	"sample-go-service/models"
)

var categoriesModal []models.Categories

func CreateCategories(data models.Categories) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("插入失败", err)
		return false
	}
	return true
}

func FindCategories() (*[]models.Categories, bool) {
	result := global.DB.Find(&categoriesModal)
	if result.RowsAffected > 0 {
		return &categoriesModal, true
	}
	return &categoriesModal, false
}
