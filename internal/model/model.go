package model

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at" gorm:"column:updated_at"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local&timeout=1s",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		//db.Lo
	}

	//db.Callback().Create()
	db.Use(&OpentracingPlugin{})
	//db.Use(gormOpentracing.New())

	//span := opentracing.StartSpan("gormTracing unit test")
	//span := opentracing.SpanFromContext(c.Request.Context())
	//defer span.Finish()

	//timeoutCtx, _ := context.WithTimeout(context.Background(), 8*time.Second)

	//ctx := opentracing.ContextWithSpan(timeoutCtx, span)
	//db = db.WithContext(ctx)

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}
