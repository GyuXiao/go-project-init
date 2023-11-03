package service

import (
	"GyuBlog/dao/mysql"
	"GyuBlog/global"
	"context"
)

type Service struct {
	ctx context.Context
	dao *mysql.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = mysql.New(global.DBEngine)
	return svc
}
