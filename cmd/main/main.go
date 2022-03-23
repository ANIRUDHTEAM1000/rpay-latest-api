package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	routes "rpay/pkg/app"
)

func main() {
	fmt.Println("Hello Strat main")
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
