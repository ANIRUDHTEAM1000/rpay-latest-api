package routes

import (
	"github.com/gin-gonic/gin"
)

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("cash-back/:uid", func(c *gin.Context) {
		// uid := c.Param("uid")
		c.IndentedJSON(200, 87.35)
	})

}
