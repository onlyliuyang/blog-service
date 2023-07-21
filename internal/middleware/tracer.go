package middleware

import (
	"context"
	"fmt"
	"github.com/blog-service/global"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(span.Context()),
			)
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		}
		defer span.Finish()

		//加入日志追踪
		var traceID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String()
			spanID = spanContext.(jaeger.SpanContext).SpanID().String()
			c.Set("X-Trace-ID", traceID)
			c.Set("X-Span-ID", spanID)
		}

		fmt.Printf("span: %v, trace_id: %s, span_id: %s", span, c.GetString("X-Trace-ID"), c.GetString("X-Span-ID"))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
