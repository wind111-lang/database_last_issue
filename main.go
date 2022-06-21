package main

import (
	"fmt"

	"chat/structs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	})
	router.POST("/signup", func(ctx *gin.Context) {
		var form structs.User //User struct

		if err := ctx.Bind(&form); err != nil {
			ctx.HTML(400, "signup.html", gin.H{"err": err})
			ctx.Abort()
		} else {
			username := ctx.PostForm("username") //usernameを取得
			password := ctx.PostForm("password") //passwordを取得
			birthday := ctx.PostForm("birthday") //birthdayを取得
			fmt.Println(username, password, birthday)

			// if CreateUser(username, password, birthday); err != nil {
			// 	ctx.HTML(400, "signup.html", gin.H{"err": err})
			// }

			ctx.Redirect(302, "/")
		}
	})

	//db.GetDatabase()

	fmt.Println("Server is running on port 8081")
	router.Run(":8081")

}
