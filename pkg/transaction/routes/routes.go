package routes

import (
	"github.com/gin-gonic/gin"
	"rpay/pkg/transaction/dao"
	trepo "rpay/pkg/transaction/repository"
	tservice "rpay/pkg/transaction/services"
)

func DefineRoutes(transaction *gin.RouterGroup) {
	transaction.GET("/transfer", func(c *gin.Context) {
		a := dao.Transaction_input{}
		c.BindJSON(&a)
		transaction_result := tservice.StartTransaction(a.Sender, a.Receiver, a.Amount)
		c.IndentedJSON(200, transaction_result)
	})

	transaction.GET("/test/uuid", func(c *gin.Context) {
		c.IndentedJSON(200, trepo.GenUUID())
	})
}
