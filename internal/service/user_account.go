package service

import (
	"encoding/json"
	"github.com/blog-service/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAccountRegisterRequest struct {
	Account     string `form:"account" binding:"required"`
	Mobile      string `form:"mobile" binding:"required"`
	Password    string `form:"password" binding:"required"`
	CountryCode int    `form:"country_code" binding:"required"`
	Source      int    `form:"source" binding:"required"`
	State       int    `form:"state" binding:"required"`
}

type UserAccountService struct {
	Service
}

func (svc *UserAccountService) Resister(ctx *gin.Context, registerRequest *UserAccountRegisterRequest) error {
	var userAccount *model.UserAccount
	byteRegister, err := json.Marshal(registerRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteRegister, &userAccount)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userAccount.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userAccount.Password = string(hashedPassword)
	err = svc.dao.RegisterUser(ctx, userAccount)
	if err != nil {
		return err
	}
	return nil
}
