package dao

import (
	"encoding/json"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AuthorParams struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	HeadUrl     string `json:"head_url"`
	CountryCode int    `json:"country_code"`
	Status      int    `json:"status"`
}

func (d *Dao) AuthorList(ctx *gin.Context, params *AuthorParams) (list []*model.Author, err error) {
	var author *model.Author
	_ = copier.Copy(&author, params)
	list, err = author.AuthorList(ctx, d.engine)
	return
}

func (d *Dao) AuthorCount(ctx *gin.Context, params *AuthorParams) (count int64, err error) {
	var author *model.Author
	_ = copier.Copy(&author, params)
	count, err = author.AuthorCount(ctx, d.engine)
	return
}

func (d *Dao) UpdateAuthor(ctx *gin.Context, params *AuthorParams) error {
	var author *model.Author
	updates := make(map[string]interface{})
	bytes, _ := json.Marshal(params)
	_ = json.Unmarshal(bytes, &updates)
	err := author.UpdateAuthor(ctx, d.engine, params.Id, updates)
	return err
}

func (d *Dao) DeleteAuthor(ctx *gin.Context, params *AuthorParams) error {
	var author *model.Author
	err := author.AuthorDelete(ctx, d.engine, params.Id)
	return err
}

func (d *Dao) CreateAuthor(ctx *gin.Context, params *AuthorParams) error {
	var err error
	var author model.Author
	author.Id = util.Uuid(ctx)
	err = copier.Copy(&author, params)
	if err != nil {
		return err
	}

	_, err = author.CreateAuthor(ctx, d.engine)
	return err
}

func (d *Dao) GetAuthor(ctx *gin.Context, id int64) (info *model.Author, err error) {
	var author model.Author
	info, err = author.AuthorById(ctx, d.engine, id)
	return
}
