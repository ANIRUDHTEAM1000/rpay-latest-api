package routes

import (
	cardRepo "rpay/pkg/cards/services"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("cards/get/all/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		cards := cardRepo.GetCards(uid)
		// cards := []cardModel.Account{{100, "PostPaid", "12345678", "12/25", "543", "Saily Anderson1", true}, {100, "PayLater", "123456789", "13/25", "543", "Saily Anderson2", true}, {100, "PrePaid", "1234567890", "14/25", "543", "Saily Anderson3", true}, {100, "Freedom", "12345678901", "14/25", "543", "Saily Anderson4", true}}
		c.IndentedJSON(200, cards)
	})

}
