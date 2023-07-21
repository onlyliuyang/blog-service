package initialization

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
)

func SetupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	fmt.Println("initialization.SetupDBEngine DB初始化完成...")
	return nil
}