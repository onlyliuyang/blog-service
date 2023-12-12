package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	opentracingLog "github.com/opentracing/opentracing-go/log"
	"net/http"
	"strings"
	"time"
)

func Post(ctx *gin.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodPost, path, params, headers)
}

func Delete(ctx *gin.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodDelete, path, params, headers)
}

func Get(ctx *gin.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodGet, path, params, headers)
}

func Put(ctx *gin.Context, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	return doRequest(ctx, http.MethodPut, path, params, headers)
}

func doRequest(ctx *gin.Context, method string, path string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	tracer, _ := ctx.Get("tracer")
	parentSpanContext, _ := ctx.Get("parentSpanContext")

	bytes, _ := json.Marshal(params)
	reqBody := strings.NewReader(string(bytes))
	req, err := http.NewRequest(method, path, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req = req.WithContext(ctx)
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 获取 OpenTracing 上下文
	span := opentracing.StartSpan(
		"httpDo",
		opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
		ext.SpanKindRPCClient,
	)

	defer span.Finish()

	injectErr := tracer.(opentracing.Tracer).Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	if injectErr != nil {
		span.LogFields(opentracingLog.String("inject-error", err.Error()))
	}

	defer resp.Body.Close()
	buffer := make([]byte, 20480)
	length, _ := resp.Body.Read(buffer)
	return buffer[:length], nil
}
