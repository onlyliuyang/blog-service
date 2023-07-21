package service

import (
	"encoding/json"
	"errors"
	"github.com/blog-service/internal/dao"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
)

type AuthorCreateRequest struct {
	Name        string `form:"name" binding:"required"`
	Mobile      string `form:"mobile" binding:"required"`
	Email       string `form:"email" binding:"required,email"`
	HeadUrl     string `form:"head_url" binding:"required"`
	CountryCode int    `form:"country_code" binding:"required"`
	Status      int    `form:"status" binding:"required,oneof=1 2"`
}

type AuthorUpdateRequest struct {
	Id          int64  `form:"id" bind:"required"`
	Name        string `form:"name" binding:""`
	Mobile      string `form:"mobile" binding:""`
	Email       string `form:"email" binding:"email"`
	HeadUrl     string `form:"head_url" binding:""`
	CountryCode int    `form:"country_code" binding:""`
	Status      int    `form:"status" binding:"oneof=1 2"`
}

type AuthorResponse struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Mobile      string    `json:"mobile"`
	HeadUrl     string    `json:"head_url"`
	CountryCode int       `json:"country_code"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AuthorService struct {
	*Service
}

func (svc *AuthorService) CreateAuthor(ctx *gin.Context, params *AuthorCreateRequest) error {
	var err error
	var author dao.AuthorParams
	err = copier.Copy(&author, params)
	if err != nil {
		return err
	}
	err = svc.dao.CreateAuthor(ctx, &author)
	if err != nil {
		return err
	}
	return nil
}

func (svc *AuthorService) UpdateAuthor(ctx *gin.Context, params *AuthorUpdateRequest) error {
	var err error
	authorInfo, err := svc.dao.GetAuthor(ctx, params.Id)
	if err != nil {
		return err
	}

	if authorInfo.Id <= 0 {
		return errors.New(errcode.ErrorNotFoundAuthor.Msg())
	}

	var author dao.AuthorParams
	_ = copier.Copy(&author, params)
	err = svc.dao.UpdateAuthor(ctx, &author)
	return err
}

func (svc *AuthorService) ListAuthor(ctx *gin.Context) (list []*AuthorResponse, err error) {
	authorList, err := svc.dao.AuthorList(ctx, nil)
	if err != nil {
		return
	}

	bytes, _ := json.Marshal(authorList)
	_ = json.Unmarshal(bytes, &list)
	return
}
