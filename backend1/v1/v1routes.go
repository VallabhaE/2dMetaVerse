package v1

import (
	"database/sql"
	"main/v1/controllers"

	"github.com/gin-gonic/gin"
)

func V1Group(router *gin.Engine,db *sql.DB) {
	v1 := router.Group("/v1")
	{
		
		v1.POST("/signin", controllers.Signin)
		v1.POST("/signup", controllers.SignUp)
	}
}
