package service

import (
	"encoding/json"
	"errors"
	_const "github.com/blog-service/const"
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

type UserAccountListRequest struct {
	Cursor string `form:"cursor"`
	Limit  int64  `form:"limit"`
}

type UserAccountListResponse struct {
	List       []*UserAccountResponse `json:"list"`
	NextCursor string                 `json:"nextCursor"`
}

type UserAccountResponse struct {
	Uid         int    `json:"uid"`
	Account     string `json:"account"`
	Mobile      string `json:"mobile"`
	CountryCode int    `json:"country_code"`
	Source      int    `json:"source"`
	State       int    `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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

func (svc *UserAccountService) GetUserInfo(ctx *gin.Context, userId int) (userInfo *model.UserAccount, err error) {
	userInfo, err = svc.dao.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}

	if userInfo.Uid <= 0 {
		return nil, errors.New("用户不存在")
	}

	if userInfo.State != _const.UserStateNormal {
		if errMsg, ok := _const.UserStateMap[userInfo.State]; ok {
			return userInfo, errors.New(errMsg)
		}
		return nil, errors.New("未知用户")
	}

	return
}

func (svc *UserAccountService) ImportUserList(ctx *gin.Context) {
	svc.dao.ImportUserList(ctx)
}

func (svc *UserAccountService) GetUserList(ctx *gin.Context, listRequest *UserAccountListRequest) (userListResponse UserAccountListResponse, err error) {
	userList, nextCursor, err := svc.dao.GetUserList(ctx, listRequest.Cursor, listRequest.Limit)
	if err != nil {
		return userListResponse, err
	}

	body, _ := json.Marshal(userList)
	_ = json.Unmarshal(body, &userListResponse.List)
	userListResponse.NextCursor = nextCursor

	return
}
