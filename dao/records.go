package dao

import (
	"fmt"
	"sample-go-service/global"
	"sample-go-service/models"
)

var records []models.Records

func CreateRecord(data models.Records) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("插入失败", err)
		return false
	}
	return true
}

func FindRecord() (*[]models.Records, bool) {
	result := global.DB.Find(&records)
	if result.RowsAffected > 0 {
		return &records, true
	}
	return &records, false
}
