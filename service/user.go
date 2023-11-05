package service

import (
	"GyuBlog/model"
	"GyuBlog/pkg/errcode"
	"GyuBlog/pkg/jwt"
	"GyuBlog/pkg/snowflake"
	"GyuBlog/pkg/util"
	"github.com/jinzhu/gorm"
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
	err := model.SelectUserByUsername(p.UserName)
	// 如果用户已经存在或者查询数据时发现其他错误，则返回错误，不能继续注册操作
	if err != nil {
		return err
	}
	// 通过雪花算法获取 userID
	userID, snowErr := snowflake.GetID()
	if snowErr != nil {
		return snowErr
	}
	u := &model.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: util.EncodeMd5([]byte(p.Password)),
		Email:    p.Email,
		Gender:   p.Gender,
	}
	// 注册用户
	return u.Create()
}

func (svc *Service) Login(p *UserLoginRequest) (user *model.User, error error) {
	userID, password, err := model.SelectUserIDAndPasswordByUsername(p.UserName)
	// 数据库查询错误
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	// 数据库里没有记录
	if err == gorm.ErrRecordNotFound {
		return nil, errcode.ErrorUserNotExit
	}
	// 查到记录的密码错误
	if password != util.EncodeMd5([]byte(p.Password)) {
		return nil, errcode.ErrorUserPassword
	}

	// 数据查询成功
	user = &model.User{
		UserName: p.UserName,
		UserID:   userID,
	}
	accessToken, refreshToken, genError := jwt.GenToken(user.UserID, user.UserName)
	if genError != nil {
		return nil, genError
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
