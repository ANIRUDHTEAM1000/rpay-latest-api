package routes

import (
	"github.com/gin-gonic/gin"
	accounts "rpay/pkg/account/routes"
	transaction "rpay/pkg/transaction/routes"
	usercontroller "rpay/pkg/user/routes"
)

func DefineMainRoutes(router *gin.Engine) {

	engine := router.Group("/walletengine")
	{
		usercontroller.DefineRoutes(engine)
		accounts.DefineRoutes(engine)
		transaction.DefineRoutes(engine)
	}
	router.GET("/favicon.ico", func(c *gin.Context) {

	})
}
