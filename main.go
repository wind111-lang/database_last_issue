package main

import (
	"fmt"
	"log"

	"chat/db"
	"chat/structs"
	"chat/websock"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	ctx.SetCookie("user", "", -1, "/", "localhost", false, false)

	session.Clear()
	session.Save()
	ctx.Redirect(302, "/redirect")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("page/*.html")

	session := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("chat-session", session))

	router.GET("/logout", func(ctx *gin.Context) {
		Logout(ctx)
	})

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

			//key := ctx.MustGet(gin.AuthUserKey).(string)
			ctx.SetCookie("user", ctx.PostForm("username"), 3600, "/", "localhost", false, false)
			ctx.Next()
			//var usr structs.LoggedInUser
			//usr.Username = ctx.PostForm("username")

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

			log.Println(username, password, birthday)
			//fmt.Println(username, password, birthday)

			if db.CreateUser(username, password, birthday); err != nil {
				ctx.HTML(400, "signup.html", gin.H{"err": err})
			}

			ctx.Redirect(302, "/redirect")
		}
	})

	router.GET("/redirect", func(ctx *gin.Context) {
		ctx.HTML(200, "redirect.html", gin.H{})
	})

	page := router.Group("/")
	page.Use(websock.SessionCheck())
	{
		page.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index.html", gin.H{})
			usr, err := ctx.Cookie("user")
			if err != nil {
				log.Println("cookie is nil!!")
				ctx.Redirect(302, "/login")

				ctx.Abort()
			} else {
				log.Println("cookie is ", usr)
			}
		})
	}

	hub := websock.NewHub()

	router.GET("/ws", func(ctx *gin.Context) {
		websock.ServeWs(hub, ctx.Writer, ctx.Request)
	})

	go hub.Run()

	//db.GetDatabase()

	fmt.Println("Server is running on port 8081")
	router.Run(":8081")

}
