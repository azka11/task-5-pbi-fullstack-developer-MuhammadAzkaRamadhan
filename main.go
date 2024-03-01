package main

import (
	// "finaltask-pbi-btpn/controllers"
	initializers "finaltask-pbi-btpn/database"
	// "finaltask-pbi-btpn/middlewares"
	"finaltask-pbi-btpn/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("HelloWorld 2")

	r := gin.Default()

	// user routes
	router.UserRoute(r)

	//photo routes
	router.PhotoRoute(r)

	r.Run()
}
