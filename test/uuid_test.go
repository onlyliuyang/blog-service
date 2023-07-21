package test

import (
	"context"
	"fmt"
	"github.com/blog-service/initialization"
	"github.com/blog-service/pkg/util"
	"testing"
)

func TestUUid(t *testing.T)  {
	err := initialization.SetupSetting("/Users/admin/go/src/github.com/go-programing-tour-book/blog-service/configs/app")
	if err != nil {
		t.Fatalf("init.setupSetting err: %v", err)
		return
	}

	err = initialization.SetupRedis()
	if err != nil {
		t.Fatalf("init.setupRedis err: %v", err)
		return
	}

	for i:=0; i<100; i++ {
		fmt.Println(util.Uuid(context.Background()))
	}
}