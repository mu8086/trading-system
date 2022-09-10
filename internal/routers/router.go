package routers

import (
	"github.com/gin-gonic/gin"
	"trading-system/internal/middleware"
	"trading-system/internal/routers/api/v1"
	_ "trading-system/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	order := v1.NewOrder()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/orders", order.List)
		apiv1.GET("/orders/:id", order.Get)
		apiv1.POST("/orders", order.Create)
		apiv1.PUT("/orders/:id", order.Update)
		apiv1.DELETE("/orders/:id", order.Delete)
	}

	return r
}
