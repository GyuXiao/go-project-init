// Package service
/**
  @author: zk.xiao
  @date: 2022/5/26
  @note:
**/
package service

import (
	"GyuBlog/global"
	"GyuBlog/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
