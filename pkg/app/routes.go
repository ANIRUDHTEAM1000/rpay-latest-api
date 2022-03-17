package routes

import (
	accounts "rpay/pkg/account/routes"
	usercontroller "rpay/pkg/user/routes"

	"github.com/gin-gonic/gin"
)

func DefineMainRoutes(router *gin.Engine) {

	engine := router.Group("/walletengine")
	{
		usercontroller.DefineRoutes(engine)
		accounts.DefineRoutes(engine)
	}
	router.GET("/favicon.ico", func(c *gin.Context) {

	})
}
