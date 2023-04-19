package controller

import (
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/forms"
	"sample-go-service/utils"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
	RecordsForm := forms.RecordsForm{}
	if err := c.ShouldBind(&RecordsForm); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}

	result := dao.CreateRecord(RecordsForm)
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
