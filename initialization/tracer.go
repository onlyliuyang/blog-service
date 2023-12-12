package initialization

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/pkg/tracer"
)

func SetupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service-v2", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	fmt.Println("initialization.SetupTracer Tracer初始化完成...")
	return nil
}
