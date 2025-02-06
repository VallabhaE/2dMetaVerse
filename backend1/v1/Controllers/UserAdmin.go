package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"main/dbase"
	
)

func Signin(ctx *gin.Context) {
	var req SigninRequest

	// Unmarshal JSON body into struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	query := `SELECT * FROM users WHERE username = ? AND password = ?`

	
	row,err:= dbase.GLOBAL_DB_CONNECTION.Query(query,req.Username,req.Password)

	if err!=nil{
		ctx.JSON(http.StatusOK, gin.H{"error": "Database error", "issue":err})
	}

	user := User{}
	err = row.Scan(&user.id, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // 404 if no user
			return
	} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err}) // 500 for other errors
			return
	}

	byteedUserObj, err := json.Marshal(&user)

	ctx.SetCookie("auth",string(byteedUserObj),60*60*12*1,"/","",true,true)
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