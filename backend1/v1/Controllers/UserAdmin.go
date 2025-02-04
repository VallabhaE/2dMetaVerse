package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(ctx *gin.Context) {
	var req SigninRequest

	// Unmarshal JSON body into struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}



	ctx.JSON(http.StatusOK, gin.H{"message": "Signed in", "user": req.Username})

}

func SignUp(ctx *gin.Context) {
	var req SignUpRequest
	// Unmarshal JSON body into struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	

	ctx.JSON(http.StatusOK, gin.H{"message": "Account Created Successfully", "user": req.Username})

}