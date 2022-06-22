package main

import (
	"fmt"
	"log"

	"chat/db"
	"chat/structs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var Login structs.Session

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
			ctx.Next()
		}
		log.Println("user is ", Login.UID)
	}
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(302, "/login")
}

func main() {
	router := gin.Default()

	session := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("chat-session", session))

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})

	router.POST("/login", func(ctx *gin.Context) {
		dbPassword := db.Getuser(ctx.PostForm("username")).Password
		log.Println("dbPassword is ", dbPassword)

		formPassword := ctx.PostForm("password")

		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
			log.Println("password is wrong!!")
			ctx.HTML(400, "login.html", gin.H{"err": err})
			ctx.Abort()
		} else {
			log.Println("password is correct!!")
			session := sessions.Default(ctx)
			session.Set("username", ctx.PostForm("username"))
			session.Save()

			ctx.Redirect(302, "/")
		}
	})

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

			//fmt.Println(username, password, birthday)

			if db.CreateUser(username, password, birthday); err != nil {
				ctx.HTML(400, "signup.html", gin.H{"err": err})
			}

			ctx.Redirect(302, "/")
		}
	})

	page := router.Group("/")
	page.Use(SessionCheck())
	{
		page.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index.html", gin.H{})
		})
	}

	//db.GetDatabase()

	fmt.Println("Server is running on port 8081")
	router.Run(":8081")

}
