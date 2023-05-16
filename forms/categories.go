package forms

type CreateCategories struct {
	Type  string `json:"type" binding:"required"`
	Icon  string `json:"icon" binding:"required"`
	Color string `json:"color" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type FindCategories struct {
	Type string `json:"type" binding:"required"`
}

type CategoriesForm struct {
	ID    string
	Type  string
	Icon  string
	Color string
	Name  string
}
