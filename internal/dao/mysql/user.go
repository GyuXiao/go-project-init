package mysql

import (
	"GyuBlog/internal/model"
)

// CheckUserExist 检查指定用户名的用户是否存在
func (d *Dao) CheckUserExist(username string) error {
	u := model.User{UserName: username}
	return u.SelectUserByName(d.engine, username)
}

func (d *Dao) InsertUser(u model.User) error {
	return u.Create(d.engine)
}
