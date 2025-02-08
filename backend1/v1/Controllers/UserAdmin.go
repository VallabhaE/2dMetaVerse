package controllers

import (
	"encoding/json"
	"hash/fnv"
	"main/dbase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32())) // Convert uint32 to string correctly
}

var SESSION_MAP = make(map[string]User)

func Signin(ctx *gin.Context) {
	var req SigninRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `SELECT * FROM users WHERE username = ? AND password = ?`
	row, err := dbase.GLOBAL_DB_CONNECTION.Query(query, req.Username, req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err.Error()})
		return
	}
	defer row.Close()

	user := User{}
	var email string

	if row.Next() {
		err = row.Scan(&user.Id, &user.Username, &email, &user.Password, &user.AvatarId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	byteedUserObj, _ := json.Marshal(&user)
	hash := Hash(string(byteedUserObj))

	SESSION_MAP[hash] = user
	ctx.SetCookie("auth", hash, 60*60*12*1, "/", "", true, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Signed in", "user": req.Username})
}

func SignUp(ctx *gin.Context) {
	var req SignUpRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.ConformPassword != req.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Confirm password does not match password"})
		return
	}

	query := "INSERT INTO users(username, email, password) VALUES (?, ?, ?)"
	_, err := dbase.GLOBAL_DB_CONNECTION.Exec(query, req.Username, req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating User"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account Created Successfully", "user": req.Username})
}

func GetAllAvatars(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("auth")
	if err != nil {
		ctx.SetCookie("auth", "", -1, "/", "", true, true)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Issue with cookie, please login"})
		return
	}

	_, exist := SESSION_MAP[cookie.Value]
	if !exist {
		ctx.SetCookie("auth", "", -1, "/", "", true, true)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session, please login"})
		return
	}

	var res []Avatars
	query := `SELECT * FROM Avatars`

	rows, err := dbase.GLOBAL_DB_CONNECTION.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var avatar Avatars
		if err := rows.Scan(&avatar.Id, &avatar.AvatarName, &avatar.AvatarImg, &avatar.Height, &avatar.Width); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning avatar data", "issue": err.Error()})
			return
		}
		res = append(res, avatar)
	}

	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Row iteration error", "issue": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"avatars": res})
}
