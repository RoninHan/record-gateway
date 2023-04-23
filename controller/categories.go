package controller

import (
	"github.com/google/uuid"
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/forms"
	"sample-go-service/models"
	"sample-go-service/utils"

	"github.com/gin-gonic/gin"
)

func CreateCategories(c *gin.Context) {
	CategoriesForm := forms.CreateCategories{}
	if err := c.ShouldBind(&CategoriesForm); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}

	categories := models.Categories{}
	categories.ID = uuid.New().String()
	categories.Icon = CategoriesForm.Icon
	categories.Name = CategoriesForm.Name
	categories.Color = CategoriesForm.Color
	categories.Type = CategoriesForm.Type

	result := dao.CreateCategories(categories)
	if !result {
		Response.Err(c, 401, 401, "新增失败", "")
		return
	}
	Response.Success(c, 200, "success", "Create successfully")
}

func GetCategories(c *gin.Context) {
	result, _ := dao.FindCategories()
	Response.Success(c, 200, "success", result)
}
