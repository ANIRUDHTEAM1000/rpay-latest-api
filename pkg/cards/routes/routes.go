package routes

import (
	cardModel "rpay/pkg/cards/models"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("cards/get/all/:uid", func(c *gin.Context) {
		// uid := c.Param("uid")
		cards := []cardModel.Card{{"PostPaid", "12345678", "12/25", "543", "Saily Anderson1", true}, {"PayLater", "123456789", "13/25", "543", "Saily Anderson2", true}, {"PrePaid", "1234567890", "14/25", "543", "Saily Anderson3", true}, {"Freedom", "12345678901", "14/25", "543", "Saily Anderson4", true}}
		c.IndentedJSON(200, cards)
	})

}
