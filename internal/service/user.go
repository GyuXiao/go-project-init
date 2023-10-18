package service

import (
	"GyuBlog/internal/model"
	"GyuBlog/pkg/jwt"
	"GyuBlog/pkg/snowflake"
	"GyuBlog/pkg/util"
)

type UserSignupRequest struct {
	UserName        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Gender          int    `json:"gender" binding:"oneof=0 1 2"` // 性别 未知：0 男：1 女：2
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UserLoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (svc *Service) Signup(p *UserSignupRequest) error {
	// 先判断待注册的用户的用户名是否已经存在
	err := svc.dao.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}
	// 通过雪花算法获取 userID
	userID, err := snowflake.GetID()
	if err != nil {
		return err
	}
	u := model.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: util.EncodeMd5([]byte(p.Password)),
		Email:    p.Email,
		Gender:   p.Gender,
	}
	// 注册用户
	return svc.dao.InsertUser(u)
}

func (svc *Service) Login(p *UserLoginRequest) (user *model.User, error error) {
	user = &model.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := svc.dao.Login(user); err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
