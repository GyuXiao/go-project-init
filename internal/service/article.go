// Package service
/**
  @author: zk.xiao
  @date: 2022/5/18
  @note:
**/
package service

type CountArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof= 0 1"`
}

type ArticleListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof= 0 1"`
}

type CreateArticleRequest struct {
	Name      string `form:"name" binding:"required, min=3, max=100"`
	CreatedBy string `form:"created_by" binding:"required, min=3, max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof= 0 1"`
}

type UpdateArticleRequest struct {
	ID         string `form:"id" binding:"required, gte=1"`
	Name       string `form:"name" binding:"min=3, max=100"`
	State      uint8  `form:"state" binding:"required, oneof= 0 1"`
	ModifiedBy string `form:"modified_by" binding:"required, min=3, max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}
