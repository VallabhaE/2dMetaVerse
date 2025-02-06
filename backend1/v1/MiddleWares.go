package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loginVerifyMiddleWare(ctx *gin.Context) {
	session := sessions.Default(ctx)
	v := session.Get("ssid")

	if v == nil {
		session.Set("ssid", "")
		session.Save()
		ctx.JSON(400, gin.H{"error": "Key Doesnt Exist"})
	} else {
		_, exist := SESSION_DATA[v.(string)]
		if !exist {
			session.Set("ssid", "")
			session.Save()
			ctx.JSON(400, gin.H{"error": "Key Doesnt Exist"})
		}

		ctx.Next()
	}

}
