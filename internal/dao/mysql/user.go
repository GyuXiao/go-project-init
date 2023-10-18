package mysql

import (
	"GyuBlog/internal/model"
	"GyuBlog/pkg/errcode"
	"GyuBlog/pkg/util"
	"github.com/jinzhu/gorm"
)

// CheckUserExist 检查指定用户名的用户是否存在
func (d *Dao) CheckUserExist(username string) error {
	u := model.User{UserName: username}
	return u.SelectUserByName(d.engine, username)
}

func (d *Dao) InsertUser(u model.User) error {
	return u.Create(d.engine)
}

func (d *Dao) Login(u *model.User) error {
	originPassword := u.Password
	// 通过用户名查找 userID 和 password
	userID, password, err := u.SelectUserIDAndPasswordByUsername(d.engine)
	// 数据库查询错误
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// 数据库里没有记录
	if err == gorm.ErrRecordNotFound {
		return errcode.ErrorUserNotExit
	}
	// 查到记录的密码错误
	if password != util.EncodeMd5([]byte(originPassword)) {
		return errcode.ErrorUserPassword
	}
	// 成功查到后记录 userID 然后返回
	u.UserID = userID
	return nil
}
