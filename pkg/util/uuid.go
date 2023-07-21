package util

import (
	"context"
	"github.com/blog-service/global"
	"github.com/blog-service/pkg/convert"
	"time"
)

var (
	UUID_BEGIN = 1000000
	UUID_KEY   = "global::uuid"
)

func Uuid(ctx context.Context) int64 {
	uuid, err := global.RedisDB.Exists(EncodeMD5(UUID_KEY)).Result()
	if err != nil {
		global.Logger.Error(ctx, "获取uuid 失败")
		return 0
	}

	if uuid > 0 {
		uuid = global.RedisDB.Incr(EncodeMD5(UUID_KEY)).Val()
		return uuid
	}

	uuid = convert.StrTo(global.RedisDB.Set(EncodeMD5(UUID_KEY), UUID_BEGIN, time.Hour*1000000).Val()).MustInt64()
	return uuid
}
