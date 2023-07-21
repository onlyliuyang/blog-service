package initialization

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/pkg/setting"
	"strings"
	"time"
)

func SetupSetting(config string) error {
	set, err := setting.NewSetting("yaml", strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("MongoDB", &global.MongoDBSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	fmt.Println("initialization.SetupSetting 配置初始化完成...")

	return nil
}