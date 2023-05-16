package forms

import (
	"sample-go-service/models"
)

type RecordsForm struct {
	CategoryId  string `json:"category_id" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RecordsJoinCategory struct {
	models.Records
	Type  string `json:"type" `
	Icon  string `json:"icon"`
	Color string `json:"color"`
	Name  string `json:"name"`
}
