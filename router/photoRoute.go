package router

import (
	"finaltask-pbi-btpn/controllers"
	"finaltask-pbi-btpn/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRoute(router *gin.Engine) {
	// Define user routes
	photo := router.Group("/photo")
	{
		photo.GET("/", middlewares.AuthRequire, controllers.FindAllPhoto)
		photo.POST("/create", middlewares.AuthRequire, controllers.CreatePhoto)
		photo.PUT("/update/:id", middlewares.AuthRequire, controllers.UpdatePhoto)
		photo.DELETE("/delete/:id", middlewares.AuthRequire, controllers.DeletePhoto)
	}
}
