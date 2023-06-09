package models

import "time"

type Records struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	UserId      string    `json:"user_id"`
	CategoryId  string    `json:"category_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	RecordDate  time.Time `json:"record_date" gorm:"type:datetime"`
}

func (Records) TableName() string {
	return "records"
}
