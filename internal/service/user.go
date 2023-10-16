package service

import (
	"GyuBlog/pkg/snowflake"
)

type UserSignupRequest struct {
	UserName        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Gender          int    `json:"gender" binding:"oneof=0 1 2"` // 性别 未知：0 男：1 女：2
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

func (svc *Service) Signup(p *UserSignupRequest) error {
	err := svc.dao.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}
	userID, err := snowflake.GetID()
	if err != nil {
		return err
	}
	// 注册用户
	return svc.dao.InsertUser(userID, p.UserName, p.Password, p.Email, p.Gender)
}
