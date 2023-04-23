package router

import (
	"github.com/gin-gonic/gin"
	"sample-go-service/controller"
	"sample-go-service/middlewares"
)

func RecordRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("record")
	{
		UserRouter.POST("create", middlewares.JWTAuth(), controller.CreateRecord)
		UserRouter.POST("list", controller.GetRecirds)
	}
}
