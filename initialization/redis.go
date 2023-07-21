package initialization

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/go-redis/redis"
)

func SetupRedis() error  {
	var err error
	global.RedisDB = redis.NewClient(&redis.Options{
		Addr: global.RedisSetting.Host,
		Password: global.RedisSetting.Password,
		DB: global.RedisSetting.DB,
	})
	_, err = global.RedisDB.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("initialization.SetupRedis Redis初始化完成...")
	return nil
}