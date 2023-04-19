package router

import (
	"sample-go-service/controller"
	"sample-go-service/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoriesRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("categories")
	{
		UserRouter.POST("create", middlewares.JWTAuth(), controller.CreateCategories)
		UserRouter.POST("list", middlewares.JWTAuth(), controller.GetCategories)
	}
}
