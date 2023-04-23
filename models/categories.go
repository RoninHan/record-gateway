package models

type Categories struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
	Name  string `json:"name"`
}

func (Categories) TableName() string {
	return "categories"
}
