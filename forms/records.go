package forms

import "time"

type RecordsForm struct {
	UserId      string
	CategoryId  string
	Amount      float32
	Description string
	RecordDate  *time.Time
}

type RecordsList struct {
	ID          uint
	UserId      string
	CategoryId  string
	Amount      float32
	Description string
	RecordDate  *time.Time
}
