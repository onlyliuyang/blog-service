package routers

import (
	_ "github.com/blog-service/docs"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/controller"
	v1 "github.com/blog-service/internal/controller/v1"
	"github.com/blog-service/internal/middleware"
	"github.com/blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Recovery())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTime * time.Second))
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	upload := v1.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	//r.POST("/auth", v1.GetAuth)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/categories", controller.CategoryController.Create)
		apiv1.DELETE("/categories/:id", controller.CategoryController.Delete)
		apiv1.PUT("/categories/:id", controller.CategoryController.Update)
		apiv1.PATCH("/categories/:id/state", controller.CategoryController.Update)
		apiv1.GET("/categories", controller.CategoryController.List)

		apiv1.POST("/articles", controller.ArticleController.Create)
		apiv1.DELETE("/articles/:id", controller.ArticleController.Delete)
		apiv1.PUT("/articles/:id", controller.ArticleController.Update)
		apiv1.PATCH("/articles/:id/state", controller.ArticleController.Update)
		apiv1.GET("/articles/:id", controller.ArticleController.Get)
		apiv1.GET("/articles", controller.ArticleController.List)

		apiv1.POST("/comments", controller.CommentController.Create)
		apiv1.DELETE("/comments/:id", controller.CommentController.Delete)
	}
	return r
}
