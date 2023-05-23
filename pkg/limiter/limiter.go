package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// 不同的接口需要的限流器采用的策略不完全一致，因此定义了一个通用的接口，主要是当前限流器所必要的方法

type LimiterIf interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimitBucketRule) LimiterIf
}

// 用于存储令牌桶与键值对名称的映射关系

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 用于存储令牌桶的一些相应规则属性
// Key: 自定义键值对名称
// FillInterval: 间隔时间
// Capacity: 令牌桶的容量
// Quantum: 每次到达间隔时间后放出的具体令牌数量

type LimitBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
