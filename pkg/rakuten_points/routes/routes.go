package routes

import (
	"github.com/gin-gonic/gin"
)

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("rakuten-points/:uid", func(c *gin.Context) {
		// uid := c.Param("uid")
		c.IndentedJSON(200, 147)
	})

}
