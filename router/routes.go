package router

import (
	docs "github.com/KSKigor/api_rest_go/docs"
	"github.com/KSKigor/api_rest_go/handler"
	"github.com/gin-gonic/gin"
	Swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePAth := "/api/v1"
	docs.SwaggerInfo.BasePath = basePAth
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(Swaggerfiles.Handler))

}
