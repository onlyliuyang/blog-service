package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

const DEFAULT_CONFIGS = "configs/app/"

type Setting struct {
	vip *viper.Viper
}

func NewSetting(configType string, configs ...string) (*Setting, error) {
	vip := viper.New()
	vip.SetConfigName(getConfigName())
	vip.SetConfigType(configType)
	vip.AddConfigPath(DEFAULT_CONFIGS)
	if len(configs) > 0 {
		for _, config := range configs {
			if config != "" {
				vip.AddConfigPath(config)
			}
		}
	}

	err := vip.ReadInConfig()
	if err != nil {
		//判断是不是因为找不到配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件")
		}
		return nil, err
	}

	s := &Setting{
		vip: vip,
	}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vip.WatchConfig()
		s.vip.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("Reoload all Section")
			_ = s.ReloadAllSection()
		})
	}()
}

func getConfigName() string {
	goEnv := os.Getenv("go_env")
	switch goEnv {
	case "dev":
		return "config_dev"
	case "test":
		return "config_test"
	case "preview":
		return "config_preview"
	default:
		return "config_release"
	}
}
