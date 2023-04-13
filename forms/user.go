package forms

import "time"

type PasswordLoginForm struct {
	UserName  string `form:"user_name" json:"user_name"`
	Mobile    string `form:"mobile" json:"mobile" binding:"mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`       // 验证码id
}

type CreateForm struct {
	ID       string
	Password string
	UserName string
	HeadUrl  string
	Birthday *time.Time
	Address  string
	Desc     string
	Gender   string
	Role     int
	Mobile   string
}

type UpdateForm struct {
	ID       string
	Password string
	UserName string
	HeadUrl  string
	Birthday *time.Time
	Address  string
	Gender   string
	Mobile   string
}

type UserListForm struct {
	ID       string
	Password string
	NickName string
	HeadUrl  string
	Birthday string
	Address  string
	Gender   string
	mobile   string
	Current  int
	PageSize int
}
