package dao

import (
	"github.com/blog-service/internal/model"
	"github.com/gin-gonic/gin"
)

func (d *Dao) RegisterUser(ctx *gin.Context, userAccount *model.UserAccount) error {
	var err error
	err = userAccount.Register(ctx, d.engine)
	return err
}
