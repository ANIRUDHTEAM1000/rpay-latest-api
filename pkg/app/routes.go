package routes

import (
	accountsRoutes "rpay/pkg/account/routes"
	cardRoutes "rpay/pkg/cards/routes"
	cashBackRoutes "rpay/pkg/cash_back/routes"
	RakutenPointsRoutes "rpay/pkg/rakuten_points/routes"
	transactionRoutes "rpay/pkg/transaction/routes"
	userRoutes "rpay/pkg/user/routes"

	"github.com/gin-gonic/gin"
)

func DefineMainRoutes(router *gin.Engine) {

	engine := router.Group("/walletengine")
	{
		userRoutes.DefineRoutes(engine)
		accountsRoutes.DefineRoutes(engine)
		transactionRoutes.DefineRoutes(engine)
		cardRoutes.DefineRoutes(engine)
		cashBackRoutes.DefineRoutes(engine)
		RakutenPointsRoutes.DefineRoutes(engine)
	}
	router.GET("/favicon.ico", func(c *gin.Context) {

	})
}
