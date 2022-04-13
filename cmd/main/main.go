package main

import (
	"fmt"
	routes "rpay/pkg/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello Start main")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	}))
	routes.DefineMainRoutes(router)
	router.Run()
}
