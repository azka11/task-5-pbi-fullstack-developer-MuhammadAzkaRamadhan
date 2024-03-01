package router

import (
	"finaltask-pbi-btpn/controllers"
	"finaltask-pbi-btpn/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// Define user routes
	user := router.Group("/users")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
		user.GET("/", middlewares.AuthRequire, controllers.FindAllUser)
		user.GET("/validate", middlewares.AuthRequire, controllers.Validate)
		user.PUT("/update/:id", middlewares.AuthRequire, controllers.UpdateUser)
		user.DELETE("/delete/:id", middlewares.AuthRequire, controllers.DeleteUser)
	}
}
