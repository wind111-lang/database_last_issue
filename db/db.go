package db

import (
	"fmt"
	"os"

	"chat/structs"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Loc struct {
// 	Location string `json:"location"`
// }

func gormConnect() *gorm.DB {

	user := os.Getenv("dbuser")
	password := os.Getenv("dbpass")
	hostname := os.Getenv("host")
	name := os.Getenv("dbname")
	port := os.Getenv("port")

	dsn := "host=" + hostname + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//fmt.Println("Connected to DB")

	return db
} //DBに接続

func CreateUser(username string, password string, birthday string) error {
	db := gormConnect()
	tempdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer tempdb.Close()

	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//Passwordハッシュ化

	var userinfo structs.User

	userinfo.Username = username
	userinfo.Password = string(pass)
	userinfo.Birthday = birthday

	if err := db.Table("members").Create(&userinfo); err.Error != nil {
		return err.Error
	}

	return nil
}

func Getuser(username string) structs.User {
	db := gormConnect()
	tempdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer tempdb.Close()

	var userinfo structs.User

	db.Table("members").First(&userinfo, "username = ?", username)

	return userinfo
}

func InsertMessage(user string, message string) {
	db := gormConnect()
	tempdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer tempdb.Close()

	var chatlog structs.ChatLog

	chatlog.Username = user
	chatlog.Text = message

	db.Table("chat_log").Create(&chatlog)
}

func UpdateUser(olduser string, newuser structs.User) {
	db := gormConnect()
	tempdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer tempdb.Close()

	var userinfo structs.User

	pass, _ := bcrypt.GenerateFromPassword([]byte(newuser.Password), bcrypt.DefaultCost)

	userinfo.Username = newuser.Username
	userinfo.Password = string(pass)

	db.Table("members").Where("username = ?", olduser).Update(userinfo.Username, userinfo.Password)
}

func DeleteUser(user string) {
	db := gormConnect()
	tempdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer tempdb.Close()

	var userinfo structs.User
	fmt.Println(user)

	db.Table("members").Where("username = ?", user).Delete(&userinfo)
	//クエリパラメータで指定したユーザを削除 この場合は指定したuserの情報が全て抹消される
}
