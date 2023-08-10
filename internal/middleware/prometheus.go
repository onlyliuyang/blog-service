package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/ca"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"sync"
	"time"
)

const (
	metricsPath = "metrics"
	faviconPath = "favicon.ico"
)

var (
	httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "request_seconds",
		Namespace:   "http_server",
		Subsystem:   "",
		Help:        "Histogram of response latency (seconds) of http handlers",
		ConstLabels: nil,
		Buckets:     nil,
	}, []string{"method", "code", "uri"})
)

func init() {
	prometheus.Register(httpHistogram)
}

// 定义采样路由
type handlerPath struct {
	sync.Map
}

// get获取path
func (hp *handlerPath) get(handler string) string {
	v, ok := hp.Load(handler)
	if !ok {
		return ""
	}
	return v.(string)
}

// set保存path到sync.Map
func (hp *handlerPath) set(ri gin.RouteInfo) {
	hp.Store(ri.Handler, ri.Path)
}

// GinPrometheus gin调用Prometheus的struct
type GinPrometheus struct {
	engine  *gin.Engine
	ignored map[string]bool
	pathMap *handlerPath
	updated bool
}

type Option func(ginPrometheus *GinPrometheus)

// Ignore添加忽略的路径
func Ignore(path ...string) Option {
	return func(ginPrometheus *GinPrometheus) {
		for _, p := range path {
			ginPrometheus.ignored[p] = true
		}
	}
}

// New gin prometheus
func NewGinPrometheus(e *gin.Engine, options ...Option) *GinPrometheus {
	if e == nil {
		return nil
	}

	gp := &GinPrometheus{
		engine: e,
		ignored: map[string]bool{
			metricsPath: true,
			faviconPath: true,
		},
		pathMap: &handlerPath{},
	}

	for _, o := range options {
		o(gp)
	}
	return gp
}

// updatePath 更新path
func (gp *GinPrometheus) updatePath() {
	gp.updated = true
	for _, ri := range gp.engine.Routes() {
		gp.pathMap.set(ri)
	}
}

func (gp *GinPrometheus) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !gp.updated {
			gp.updatePath()
		}

		//过滤请求
		if gp.ignored[ctx.Request.URL.String()] {
			ca.New()
			return
		}

		start := time.Now()
		ctx.Next()

		httpHistogram.WithLabelValues(
			ctx.Request.Method,
			strconv.Itoa(ctx.Writer.Status()),
			gp.pathMap.get(ctx.HandlerName()),
		).Observe(time.Since(start).Seconds())
	}
}
