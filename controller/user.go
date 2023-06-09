package controller

import (
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/models"
	"sample-go-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"sample-go-service/forms"
)

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}

	// if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
	// 	Response.Err(c, 400, 400, "otp failed", "")
	// 	return
	// }

	user, ok := dao.FindUserInfo(PasswordLoginForm.Email, PasswordLoginForm.PassWord)
	if !ok {
		Response.Err(c, 401, 401, "未注册该用户", "")
		return
	}
	token := utils.CreateToken(c, user.ID, user.UserName, user.Role)
	userinfoMap := HandleUserModelToMap(user)
	result := map[string]interface{}{"user": userinfoMap, "token": token}

	Response.Success(c, 200, "success", result)
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Current, UserListForm.PageSize)
	// 判断
	if (total + len(userlist)) == 0 {
		Response.Err(c, 400, 400, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	Response.Success(c, 200, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})
}

func CreateUser(c *gin.Context) {
	CreateForm := forms.CreateForm{}
	if err := c.ShouldBind(&CreateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	_, ok := dao.FindUserInfo(CreateForm.Email, CreateForm.Password)
	if ok {
		Response.Err(c, 401, 401, "存在相同的user", "")
		return
	}
	var user models.User
	user.ID = uuid.New().String()
	user.Address = CreateForm.Address
	user.Email = CreateForm.Email
	user.Password = CreateForm.Password
	user.Birthday = CreateForm.Birthday
	user.UserName = CreateForm.UserName
	user.Role = 1
	user.Gender = CreateForm.Gender
	user.HeadUrl = CreateForm.HeadUrl
	user.Mobile = CreateForm.Mobile

	result := dao.InsertUser(user)
	Response.Success(c, 200, "添加成功", result)
}

func UpdateUser(c *gin.Context) {
	UpdateForm := forms.UpdateForm{}
	if err := c.ShouldBind(&UpdateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	result := dao.UpdateUser(UpdateForm.ID, UpdateForm)
	Response.Success(c, 200, "更新成功", result)
}

func HandleUserModelToMap(user *models.User) map[string]interface{} {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userItemMap := map[string]interface{}{
		"id":       user.ID,
		"UserName": user.UserName,
		"HeadUrl":  user.HeadUrl,
		"email":    user.Email,
		"birthday": birthday,
		"address":  user.Address,
		"desc":     user.Desc,
		"gender":   user.Gender,
		"role":     user.Role,
		"mobile":   user.Mobile,
	}
	return userItemMap
}
