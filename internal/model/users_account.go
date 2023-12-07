package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// UserAccount 模型对应于 "users_account" 表
type UserAccount struct {
	Uid         int       `gorm:"column:uid;primaryKey;autoIncrement" json:"uid,omitempty" redis:"uid"`
	Account     string    `gorm:"column:account;uniqueIndex:un_account;not null" json:"account,omitempty" redis:"account"`
	Mobile      string    `gorm:"column:mobile;uniqueIndex:un_phone;not null" json:"mobile,omitempty" redis:"mobile"`
	Password    string    `gorm:"column:password" json:"password,omitempty" redis:"password"`
	CountryCode int       `gorm:"column:country_code;not null;default:86" json:"country_code,omitempty" redis:"country_code"`
	Source      int       `gorm:"column:source;not null;default:1" json:"source,omitempty" redis:"source"`
	State       int       `gorm:"column:state;not null;default:1" json:"state,omitempty" redis:"state"`
	CreatedAt   time.Time `gorm:"column:created_at;" json:"created_at,omitempty" redis:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty" redis:"updated_at"`
}

// TableName 返回模型对应的数据库表名
func (u *UserAccount) TableName() string {
	return "bs_users_account"
}

//用户注册
func (u *UserAccount) Register(ctx *gin.Context, db *gorm.DB) error {
	var err error
	err = db.WithContext(ctx).Table(u.TableName()).Create(u).Error
	return err
}

func (u *UserAccount) GetUserInfo(ctx *gin.Context, db *gorm.DB, userId int) (userInfo UserAccount, err error) {
	err = db.WithContext(ctx).Table(u.TableName()).Where("uid = ?", userId).Find(&userInfo).Error
	return
}

func (u *UserAccount) GetListByPage(ctx *gin.Context, db *gorm.DB, limit, offset int) (userList []*UserAccount, err error) {
	err = db.WithContext(ctx).Table(u.TableName()).Order("uid desc").Offset(offset).Limit(limit).Find(&userList).Error
	return
}
