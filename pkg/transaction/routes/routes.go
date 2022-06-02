package routes

import (
	"rpay/pkg/transaction/dao"
	tservice "rpay/pkg/transaction/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(transaction *gin.RouterGroup) {
	transaction.POST("/transfer", func(c *gin.Context) {
		a := dao.Transaction_input{}
		c.BindJSON(&a)
		transaction_result := tservice.StartTransaction(a.Sender, a.Receiver, a.Amount)
		c.IndentedJSON(200, transaction_result)
	})
	transaction.GET("/get/transactions/:userId/:pageNumber", func(c *gin.Context) {
		userId := c.Param("userId")
		pageNumber, _ := strconv.Atoi(c.Param("pageNumber"))
		transactionsList := tservice.GetTransactionList(userId, pageNumber)
		c.IndentedJSON(200, transactionsList)
	})
}
