package dao

import (
	"sample-go-service/forms"
	"sample-go-service/global"
)

var RecordsList []forms.RecordsList

func CreateRecord(data forms.RecordsForm) bool {
	result := global.DB.Create(data)
	return result.RowsAffected > 0
}

func FindRecord() (*[]forms.RecordsList, bool) {
	result := global.DB.Find(&RecordsList)
	if result.RowsAffected > 0 {
		return &RecordsList, true
	}
	return &RecordsList, false
}
