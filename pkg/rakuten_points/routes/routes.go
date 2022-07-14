package routes

import (
	rpServices "rpay/pkg/rakuten_points/services"

	"github.com/gin-gonic/gin"
)

type CB_redeem_input struct {
	Uid           string  `json:"uid"`
	Redeem_amount float32 `json:"redeem_amount"`
}

func DefineRoutes(account *gin.RouterGroup) {

	account.GET("rakuten-points/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.IndentedJSON(200, rpServices.GetRakutenPoints(uid))
	})

	account.POST("rakuten-points/redeem", func(c *gin.Context) {
		a := CB_redeem_input{}
		c.BindJSON(&a)
		res := rpServices.RedeemRP(a.Uid, a.Redeem_amount)
		c.IndentedJSON(200, res)
	})

}
