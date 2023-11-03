package service

import (
	"GyuBlog/dao/mysql"
	"GyuBlog/model"
	"context"
)

type Service struct {
	ctx context.Context
	dao *mysql.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	if svc.dao == nil {
		svc.dao = mysql.New(model.DBEngine)
	}
	return svc
}
