package routes

import (
	cbServices "rpay/pkg/cash_back/services"

	"github.com/gin-gonic/gin"
)

type CB_redeem_input struct {
	Uid           string  `json:"uid"`
	Redeem_amount float32 `json:"redeem_amount"`
}

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("cash-back/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		res := cbServices.GetCashBack(uid)
		c.IndentedJSON(200, res)
	})

	account.POST("cash-back/redeem", func(c *gin.Context) {
		a := CB_redeem_input{}
		c.BindJSON(&a)
		res := cbServices.RedeemCB(a.Uid, a.Redeem_amount)
		c.IndentedJSON(200, res)
	})

}
