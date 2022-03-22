package routes

import (
	"github.com/gin-gonic/gin"
	"rpay/pkg/transaction/dao"
	tservice "rpay/pkg/transaction/services"
)

func DefineRoutes(transaction *gin.RouterGroup) {
	transaction.POST("/transfer", func(c *gin.Context) {
		a := dao.Transaction_input{}
		c.BindJSON(&a)
		transaction_result := tservice.StartTransaction(a.Sender, a.Receiver, a.Amount)
		c.IndentedJSON(200, transaction_result)
	})
}
