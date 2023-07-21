package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/initialization"
	"github.com/blog-service/internal/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	_ = setupFlag()

	err := initialization.SetupSetting(config)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
		return
	}

	//如果存在，则覆盖原来配置
	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	err = initialization.SetupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
		return
	}

	err = initialization.SetupMongoDB()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
		return
	}

	err = initialization.SetupRedis()
	if err != nil {
		log.Fatalf("init.setupRedis err: %v", err)
		return
	}

	err = initialization.SetupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
		return
	}

	err = initialization.SetupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
		return
	}
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用配置文件的路径")
	flag.Parse()
	return nil
}


// @title 博客系统
// @version 1.0
// @description Go语言编程之路
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	fmt.Println("Http Listen on :", global.ServerSetting.HttpPort)

	//等待中断信号
	quit := make(chan os.Signal)
	//接收信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	//最大时间控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting...")
}