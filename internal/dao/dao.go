// Package dao
/**
  @author: zk.xiao
  @date: 2022/5/26
  @note:
**/
package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
