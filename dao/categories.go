package dao

import (
	"fmt"
	"sample-go-service/forms"
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

func FindCategories(data forms.FindCategories) (*[]models.Categories, bool) {
	result := global.DB.Where(&models.Categories{Type: data.Type}).Find(&categoriesModal)
	if result.RowsAffected > 0 {
		return &categoriesModal, true
	}
	return &categoriesModal, false
}
