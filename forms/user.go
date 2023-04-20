package forms

import "time"

type PasswordLoginForm struct {
	Email    string `form:"email" json:"email"`
	PassWord string `form:"password" json:"password"`
}

type CreateForm struct {
	Email    string     `form:"email" json:"email" binding:"required"`
	Password string     `form:"password" json:"password" binding:"required"`
	UserName string     `form:"user_name" json:"user_name" `
	HeadUrl  string     `form:"head_url" json:"head_url" `
	Birthday *time.Time `form:"birthday" json:"birthday"`
	Address  string     `form:"address" json:"address"`
	Gender   string     `form:"gender" json:"gender"`
	Mobile   string     `form:"mobile" json:"mobile"`
}

type UpdateForm struct {
	ID       string
	Email    string
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
	Email    string
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
