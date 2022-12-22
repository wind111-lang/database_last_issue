package websock

import (
	"chat/structs"
	//"fmt"

	"github.com/gin-contrib/sessions"

	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"log"
)

var Login structs.Session

//var usr structs.LoggedInUser

func SessionCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		Login.UID = session.Get("username")
		if Login.UID == nil {
			log.Println("user is nil!!")

			ctx.Redirect(302, "/login")
			ctx.Abort()
		} else {
			ctx.Set("username", Login.UID)
			//usr.Username = Login.UID.(string)
			//fmt.Println("user is ", usr.Username)
			ctx.Next()
		}
		log.Println("user is ", Login.UID)
	}
}

// func GetUID(ctx *gin.Context) string {
// 	session := sessions.Default(ctx)
// 	Login.UID = session.Get("username")

// 	return Login.UID.(string)
// }
