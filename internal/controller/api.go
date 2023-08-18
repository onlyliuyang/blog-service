package controller

import v1 "github.com/blog-service/internal/controller/v1"

var (
	ArticleController  = new(v1.Article)
	CategoryController = new(v1.Category)
	CommentController  = new(v1.Comment)
	AuthorController   = new(v1.Author)
)
