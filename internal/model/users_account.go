package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// UserAccount 模型对应于 "users_account" 表
type UserAccount struct {
	UID         int       `gorm:"column:uid;primaryKey;autoIncrement" json:"uid"`
	Account     string    `gorm:"column:account;uniqueIndex:un_account;not null" json:"account"`
	Mobile      string    `gorm:"column:mobile;uniqueIndex:un_phone;not null" json:"mobile"`
	Password    string    `gorm:"column:password" json:"password"`
	CountryCode uint16    `gorm:"column:country_code;not null;default:86" json:"country_code"`
	Source      int       `gorm:"column:source;not null;default:1" json:"source"`
	State       uint8     `gorm:"column:state;not null;default:1" json:"state"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;type:datetime(0);autoUpdateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;type:datetime(0);autoUpdateTime on update current_timestamp" json:"updated_at"`
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
