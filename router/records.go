package router

import (
	"sample-go-service/controller"
	"sample-go-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RecordRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("record")
	{
		UserRouter.POST("create", middlewares.JWTAuth(), controller.CreateRecord)
		UserRouter.POST("list", middlewares.JWTAuth(), controller.GetRecirds)
	}
}
