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
	today := time.Now().UTC().Truncate(24 * time.Hour) // 当天的起始时间
	tomorrow := today.Add(24 * time.Hour)
	var recordsJoinCategorys []forms.RecordsJoinCategory
	result := global.DB.Table(records.TableName()).Select("records.*,categories.type,categories.color,categories.icon,categories.name").Joins("left join categories on categories.id=records.category_id").Where("user_id=? and (records.record_date >= ? AND records.record_date < ?)", userId, today, tomorrow).Order("records.record_date desc").Find(&recordsJoinCategorys)
	if result.RowsAffected > 0 {
		return recordsJoinCategorys, true
	}
	return recordsJoinCategorys, false
}
