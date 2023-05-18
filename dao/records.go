package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
	"time"
)

var records models.Records
var categories models.Categories

func CreateRecord(data models.Records) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("插入失败", err)
		return false
	}
	return true
}

func FindRecord(userId string) ([]forms.RecordsJoinCategory, bool) {
	var recordsJoinCategorys []forms.RecordsJoinCategory
	// Get the current month's start and end dates
	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1).Add(time.Hour*time.Duration(23) +
		time.Minute*time.Duration(59) + time.Second*time.Duration(59))

	result := global.DB.Table(records.TableName()).Select("records.*,categories.type,categories.color,categories.icon,categories.name").Joins("left join categories on categories.id=records.category_id").Where("user_id=? and (records.record_date  BETWEEN ? AND ?)", userId, firstDayOfMonth, lastDayOfMonth).Order("records.record_date desc").Find(&recordsJoinCategorys)
	if result.RowsAffected > 0 {
		return recordsJoinCategorys, true
	}
	return recordsJoinCategorys, false
}
