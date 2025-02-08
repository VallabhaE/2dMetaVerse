package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"main/dbase"
	"main/v1/session"
	
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
	var email string

	err = row.Scan(&user.Id, &user.Username,email ,&user.Password,&user.AvatarId)

	if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // 404 if no user
			return
	} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err}) // 500 for other errors
			return
	}

	byteedUserObj, err := json.Marshal(&user)
	hash := hash(string(byteedUserObj))
	session.SESSION_MAP[string(hash)] = user
	ctx.SetCookie("auth",string(hash),60*60*12*1,"/","",true,true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Signed in", "user": req.Username})

}

func SignUp(ctx *gin.Context) {
	var req SignUpRequest
	// Unmarshal JSON body into struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if req.ConformPassword!=req.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Conform password is not equal to password"})
		return
	}



	quarry := "Insert into users(username,email,password) values (?,?,?)"
	


	_,err:= dbase.GLOBAL_DB_CONNECTION.Exec(quarry,req.Username,req.Email,req.Password)

	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating User"})
		return

	}

	

	ctx.JSON(http.StatusOK, gin.H{"message": "Account Created Successfully", "user": req.Username})

}


func getUserAvatar(ctx *gin.Context){
	cookie,err:= ctx.Request.Cookie("auth")
	if err!=nil{
		ctx.SetCookie("auth","",-1,"/","",true,true)
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}

	Userinfo,exist := session.SESSION_MAP[cookie.Value]
	if !exist{
		ctx.SetCookie("auth","",-1,"/","",true,true)
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success","AvatarId":string(Userinfo.AvatarId)})
	return
}


func getAllAvatars(ctx *gin.Context) {
	cookie,err:= ctx.Request.Cookie("auth")
	if err!=nil{
		ctx.SetCookie("auth","",-1,"/","",true,true)
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}

	Userinfo,exist := session.SESSION_MAP[cookie.Value]
	if !exist{
		ctx.SetCookie("auth","",-1,"/","",true,true)
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}


	if err != nil{
		ctx.JSON(http.StatusOK, gin.H{"error": "Cookie not found",})
		return 
	}
	var res []Avatars // Slice to store avatars

	// SQL query to get all avatars
	query := `SELECT * FROM Avatars`

	// Execute the query
	rows, err := dbase.GLOBAL_DB_CONNECTION.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err.Error()})
		return
	}
	defer rows.Close() // Close rows after we're done

	// Iterate over the rows
	for rows.Next() {
		var avatar Avatars
		// Scan row data into avatar struct
		if err := rows.Scan(&avatar.Id, &avatar.AvatarName, &avatar.AvatarImg, &avatar.Height, &avatar.Width); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning avatar data", "issue": err.Error()})
			return
		}
		// Append the avatar to the result slice
		res = append(res, avatar)
	}

	// Check for errors after iterating over rows
	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Row iteration error", "issue": err.Error()})
		return
	}

	// Return the avatars as JSON
	ctx.JSON(http.StatusOK, gin.H{"avatars": res})
}
