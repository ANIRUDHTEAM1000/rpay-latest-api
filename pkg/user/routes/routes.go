package routes

import (
	userService "rpay/pkg/user/services"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(user *gin.RouterGroup) {
	user.GET("/user/:userid", func(c *gin.Context) {
		user_id := c.Param("userid")
		user := userService.GetUserByUserId(user_id)
		c.IndentedJSON(200, user)
	})

	user.GET("/user/query/:q", func(c *gin.Context) {
		q := c.Param("q")
		users := userService.GetAllUsersByQuery(q)
		c.IndentedJSON(200, users)
	})

}
