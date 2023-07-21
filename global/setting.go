package global

import (
	"github.com/blog-service/pkg/logger"
	"github.com/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettingS
	MongoDBSetting  *setting.MongoDBSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)
