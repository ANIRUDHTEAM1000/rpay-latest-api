package routes

import (
	"github.com/gin-gonic/gin"
	accountServices "rpay/pkg/account/services"
)

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("balance/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		balance := accountServices.GetBalanceByUid(uid)
		c.IndentedJSON(200, balance)
	})

}
