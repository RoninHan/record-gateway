package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/forms"
	"sample-go-service/models"
	"sample-go-service/utils"
	"time"
)

func CreateRecord(c *gin.Context) {
	RecordsForm := forms.RecordsForm{}
	if err := c.ShouldBind(&RecordsForm); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}
	userId := c.MustGet("userId").(string)
	
	record := models.Records{}
	record.ID = uuid.New().String()
	record.RecordDate = time.Now()
	record.Amount = RecordsForm.Amount
	record.CategoryId = RecordsForm.CategoryId
	record.UserId = userId
	record.Description = RecordsForm.Description

	result := dao.CreateRecord(record)
	if !result {
		Response.Err(c, 401, 401, "新增失败", "")
		return
	}
	Response.Success(c, 200, "success", "Create successfully")
}

func GetRecirds(c *gin.Context) {
	result, _ := dao.FindRecord()
	Response.Success(c, 200, "success", result)
}
