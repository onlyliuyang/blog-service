package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	//获取对应的限流器的键值对名称
	Key(ctx *gin.Context) string
	//获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	//新增多个令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	LimiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        //自定义键值对名称
	FillInterval time.Duration //间隔多久时间放N个令牌
	Capacity     int64         //令牌桶容量
	Quantum      int64         //每次到达间隔时间后所放的具体令牌数量
}
