package mysql

import (
	"GyuBlog/internal/model"
	"GyuBlog/pkg/util"
)

// CheckUserExist 检查指定用户名的用户是否存在
func (d *Dao) CheckUserExist(username string) error {
	u := model.User{UserName: username}
	return u.SelectUsersByName(d.engine, username)
}

// InsertUser
// 参数是不是太多了？
func (d *Dao) InsertUser(userID uint64, username string, password string, email string, gender int) error {
	u := model.User{
		UserID:   userID,
		UserName: username,
		Password: util.EncodeMd5([]byte(password)),
		Email:    email,
		Gender:   gender,
	}
	return u.Create(d.engine)
}
