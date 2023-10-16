// Package service
/**
  @author: zk.xiao
  @date: 2022/5/26
  @note:
**/
package service

import (
	"GyuBlog/global"
	"GyuBlog/internal/dao/mysql"
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
