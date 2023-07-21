package initialization

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func SetupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:   global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:    600,
		MaxAge:     10,
		MaxBackups: 0,
		LocalTime:  true,
		Compress:   false,
	}, "", log.LstdFlags).WithCaller(2)

	fmt.Println("initialization.SetupLogger 日志初始化完成...")
	return nil
}