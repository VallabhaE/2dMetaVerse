package v1

import (
	"main/v1/controllers"
	"net/http"

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