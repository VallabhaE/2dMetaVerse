package v1

import (
	"main/dbase"
	"main/v1/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetUserAvatar(ctx *gin.Context){
	cookie,err:= ctx.Request.Cookie("auth")
	if err!=nil{
		ctx.SetCookie("auth","",-1,"/","",true,true)
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}

	Userinfo,exist := controllers.SESSION_MAP[cookie.Value]
	if !exist{
		ctx.SetCookie("auth","",-1,"/","",true,true)
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success","AvatarId":string(Userinfo.AvatarId)})
	return
}

func SetUserAvatar(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("auth")
	if err != nil {
		ctx.SetCookie("auth", "", -1, "/", "", true, true)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Issue with cookie, please login"})
		return
	}

	Userinfo, exist := controllers.SESSION_MAP[cookie.Value]
	if !exist {
		ctx.SetCookie("auth", "", -1, "/", "", true, true)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Issue with cookie, please login"})
		return
	}

	var req struct {
		AvatarId int `json:"avatar_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.AvatarId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Avatar ID"})
		return
	}

	query := `UPDATE users SET avatar_id = ? WHERE id = ?`
	_, err = dbase.GLOBAL_DB_CONNECTION.Exec(query, req.AvatarId, Userinfo.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "issue": err.Error()})
		return
	}

	Userinfo.AvatarId = strconv.Itoa(req.AvatarId)
	controllers.SESSION_MAP[cookie.Value] = Userinfo

	ctx.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully", "avatar_id": req.AvatarId})
}
