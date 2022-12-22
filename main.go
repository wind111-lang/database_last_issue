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

var ip string

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	//ip = websock.GetIP()
	ctx.SetCookie("user", "", -1, "/", "retranslate-chatroom.com", false, false)

	session.Clear()
	session.Save()
	ctx.Redirect(302, "/redirect")
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("page/*.html")
	router.Static("/source", "page/source/")

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
		log.Println("Password is ", dbPassword)

		formPassword := ctx.PostForm("password")

		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
			log.Println("password is wrong!!")
			ctx.HTML(400, "login.html", gin.H{"err": "ユーザ名かパスワードが間違っています"})
			ctx.Abort()
		} else {
			log.Println("password is correct!!")
			session := sessions.Default(ctx)

			session.Set("username", ctx.PostForm("username"))
			session.Save()

			ctx.SetCookie("user", ctx.PostForm("username"), 3600, "/", "retranslate-chatroom.com", false, false)
			ctx.Redirect(302, "/")
		}
	})

	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	})
	router.POST("/signup", func(ctx *gin.Context) {
		var form structs.User

		if err := ctx.Bind(&form); err != nil {
			ctx.HTML(400, "signup.html", gin.H{"err": err})
			ctx.Abort()
		} else {
			username := ctx.PostForm("username") //usernameを取得
			password := ctx.PostForm("password") //passwordを取得
			birthday := ctx.PostForm("birthday") //birthdayを取得

			if err = db.CreateUser(username, password, birthday); err != nil {
				ctx.HTML(400, "signup.html", gin.H{"err": err})
			} else {
				ctx.SetCookie("user", ctx.PostForm("username"), 3600, "/", "retranslate-chatroom.com", false, false)
				ctx.Redirect(302, "/redirect")
			}
		}
	})

	router.GET("/redirect", func(ctx *gin.Context) {
		ctx.HTML(200, "redirect.html", gin.H{})
	})

	router.GET("/userinfo", func(ctx *gin.Context) {
		ctx.HTML(200, "userinfo.html", gin.H{})

		usr, err := ctx.Cookie("user")
		if err != nil {
			log.Println("cookie is nil!!")
			ctx.Redirect(302, "/login")

			ctx.Abort()
		} else {
			log.Println("cookie is ", usr)
			if db.Getuser(usr); err != nil {
				log.Println("user is nil!!")
				ctx.Redirect(302, "/login")

				ctx.Abort()
			}
		}
	})

	router.GET("/update", func(ctx *gin.Context) {
		ctx.HTML(200, "update.html", gin.H{})

		usr, err := ctx.Cookie("user")
		if err != nil {
			log.Println("cookie is nil!!")
			ctx.Redirect(302, "/login")

			ctx.Abort()
		} else {
			log.Println("cookie is ", usr)
			if db.Getuser(usr); err != nil {
				log.Println("user is nil!!")
				ctx.Redirect(302, "/login")

				ctx.Abort()
			}
		}
	})

	router.POST("/update", func(ctx *gin.Context) {
		usr, err := ctx.Cookie("user")
		if err != nil {
			log.Println("cookie is nil!!")
			ctx.Redirect(302, "/login")

			ctx.Abort()
		} else {
			log.Println("cookie is ", usr)
			if db.Getuser(usr); err != nil {
				log.Println("user is nil!!")
				ctx.Redirect(302, "/login")

				ctx.Abort()
			}
		}

		var form structs.UpdateUser //User struct

		if err := ctx.Bind(&form); err != nil {
			ctx.HTML(400, "update.html", gin.H{"err": err})
			ctx.Abort()
		} else {
			var userinfo structs.User

			username := ctx.PostForm("username")
			password := ctx.PostForm("password")

			userinfo.Username = username
			userinfo.Password = password

			if db.UpdateUser(usr, userinfo); err != nil {
				ctx.HTML(400, "update.html", gin.H{"err": err})
			}

			ctx.SetCookie("user", ctx.PostForm("username"), 3600, "/", "retranslate-chatroom.com", false, false)
			ctx.HTML(302, "userinfo.html", gin.H{"err": "更新しました"})
		}
	})

	router.GET("/delete", func(ctx *gin.Context) {
		ctx.HTML(200, "delete.html", gin.H{})

		usr, err := ctx.Cookie("user")
		if err != nil {
			log.Println("cookie is nil!!")
			ctx.Redirect(302, "/login")

			ctx.Abort()
		} else {
			log.Println("cookie is ", usr)
			if db.Getuser(usr); err != nil {
				log.Println("user is nil!!")
				ctx.Redirect(302, "/login")

				ctx.Abort()
			}
		}
	})

	router.POST("/delete", func(ctx *gin.Context) {
		//fmt.Println("delete")
		usr, err := ctx.Cookie("user")
		if err != nil {
			log.Println("cookie is nil!!")
			ctx.Redirect(302, "/login")
			ctx.Abort()
		} else {
			log.Println("cookie is ", usr)
			if db.Getuser(usr); err != nil {
				log.Println("user is nil!!")
				ctx.Redirect(302, "/login")

				ctx.Abort()
			} else {
				//ctx.SetCookie("user", usr, 3600, "/delete", ip, false, false)
				db.DeleteUser(usr)
			}
		}
		ctx.Redirect(302, "/logout")
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

	ip = websock.GetIP()
	fmt.Println("Server is running on", ":8888")
	router.RunTLS("172.31.12.5:8888","/etc/letsencrypt/live/retranslate-chatroom.com/fullchain.pem","/etc/letsencrypt/live/retranslate-chatroom.com/privkey.pem")
}
