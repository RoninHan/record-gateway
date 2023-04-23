package router

import (
	"github.com/gin-gonic/gin"
	"sample-go-service/controller"
)

func CategoriesRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("categories")
	{
		UserRouter.POST("create", controller.CreateCategories)
		UserRouter.POST("list", controller.GetCategories)
	}
}
