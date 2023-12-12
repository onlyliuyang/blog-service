package middleware

import (
	"context"
	"github.com/blog-service/global"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	//return func(c *gin.Context) {
	//	rootContext, _ := global.Tracer.Extract(
	//		opentracing.HTTPHeaders,
	//		opentracing.HTTPHeadersCarrier(c.Request.Header),
	//	)
	//
	//	span := global.Tracer.StartSpan(c.Request.URL.Path, opentracing.ChildOf(rootContext))
	//	defer span.Finish()
	//	//加入日志追踪
	//	var traceID string
	//	var spanID string
	//	var spanContext = span.Context()
	//	switch spanContext.(type) {
	//	case jaeger.SpanContext:
	//		traceID = spanContext.(jaeger.SpanContext).TraceID().String()
	//		spanID = spanContext.(jaeger.SpanContext).SpanID().String()
	//		c.Set("X-Trace-ID", traceID)
	//		c.Set("X-Span-ID", spanID)
	//		c.Set("span", span)
	//		c.Set("tracer", global.Tracer)
	//	}
	//	//c.Request = c.Request.WithContext(spanContext)
	//	c.Next()
	//}

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
			c.Set("parentSpanContext", span.Context())
			c.Set("tracer", global.Tracer)
		}

		//返回span的spanContext
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

//
//span := tracer.StartSpan("span_root")
//ctx := opentracing.ContextWithSpan(context.Background(), span)
