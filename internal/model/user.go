package model

import (
	"GyuBlog/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID   uint64 `json:"user_id,string" gorm:"column:user_id"`
	UserName string `json:"username" gorm:"column:username;index:username"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Gender   int    `json:"gender" gorm:"column:gender"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

// SelectUserByName
// 对于这个方法的返回值
// 1，如果在 User 表里找到了记录，返回 ErrorUserExit 业务码
// 2，如果在 User 表找不到记录，返回 nil
// 3，除以上的其他错误，都需要返回对应的错误
func (u User) SelectUserByName(db *gorm.DB, username string) error {
	var user User
	result := db.Select(username).Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}
	return errcode.ErrorUserExit
}
