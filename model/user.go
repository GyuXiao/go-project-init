package model

import (
	"GyuBlog/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID       uint64 `json:"user_id,string" gorm:"column:user_id"`
	UserName     string `json:"username" gorm:"column:username;index:username"`
	Password     string `json:"password" gorm:"column:password"`
	Email        string `json:"email" gorm:"column:email"`
	Gender       int    `json:"gender" gorm:"column:gender"`
	AccessToken  string
	RefreshToken string
}

func (u User) TableName() string {
	return "user"
}

func (u *User) Create() error {
	return DBEngine.Select("UserID", "UserName", "Password", "Email", "Gender").Create(&u).Error
}

// SelectUserByUsername
// 对于这个方法的返回值
// 1，如果在 User 表里找到了记录，返回 ErrorUserExit 业务码
// 2，如果在 User 表找不到记录，返回 nil
// 3，除以上的其他错误，都需要返回对应的错误
func SelectUserByUsername(username string) error {
	var user User
	result := DBEngine.Select("username").Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}
	return errcode.ErrorUserExit
}

func SelectUserIDAndPasswordByUsername(username string) (userID uint64, password string, err error) {
	var user User
	result := DBEngine.Select([]string{"user_id", "password"}).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return 0, "", result.Error
	}
	return user.UserID, user.Password, nil
}
