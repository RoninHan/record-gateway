package forms

import "time"

type RecordsForm struct {
	UserId      string     `json:"user_id" binding:"required"`
	CategoryId  string     `json:"category_id" binding:"required"`
	Amount      float32    `json:"amount" binding:"required"`
	Description string     `json:"description" binding:"required"`
	RecordDate  *time.Time `json:"record_date" binding:"required"`
}
