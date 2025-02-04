package v1

import (
	"main/v1/Controllers"

	"github.com/gin-gonic/gin"
)

func V1Group(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		
		v1.POST("/signin", controllers.Signin)
		v1.POST("/signup", controllers.SignUp)
	}
}
