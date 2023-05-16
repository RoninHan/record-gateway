package controller

import (
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/forms"
	"sample-go-service/models"
	"sample-go-service/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	float, _ := strconv.ParseFloat(RecordsForm.Amount, 32)
	record.Amount = float
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

func GetTodayRecirds(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	recirds, _ := dao.FindRecord(userId)
	user, _ := dao.GetUser(userId)
	token := utils.CreateToken(c, user.ID, user.UserName, user.Role)
	result := map[string]interface{}{"data": recirds, "token": token}

	Response.Success(c, 200, "success", result)
}
