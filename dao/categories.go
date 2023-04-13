package dao

import (
	"sample-go-service/forms"
	"sample-go-service/global"
)

var CategoriesForms []forms.CategoriesForm

func CreateCategories(data forms.CreateCategories) bool {
	result := global.DB.Create(data)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func FindCategories() (*[]forms.CategoriesForm, bool) {
	result := global.DB.Find(&CategoriesForms)
	if result.RowsAffected > 0 {
		return &CategoriesForms, true
	}
	return &CategoriesForms, false
}
