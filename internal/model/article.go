package model

import "GyuBlog/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog-article"
}

type ArticleSwagger struct {
	list  []*Article
	Pager *app.Pager
}
