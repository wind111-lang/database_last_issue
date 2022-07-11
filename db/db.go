package db

import (
	"fmt"

	"chat/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

// type Loc struct {
// 	Location string `json:"location"`
// }

func gormConnect() *gorm.DB {

	user := "user"
	password := "hoge"
	hostname := "localhost:3306"
	name := "chatdb"

	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+hostname+")/"+name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")

	return db
} //DBに接続

func CreateUser(username string, password string, birthday string) error {
	db := gormConnect()
	defer db.Close()

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
	defer db.Close()

	var userinfo structs.User

	db.Table("members").First(&userinfo, "username = ?", username)

	return userinfo
}

func InsertMessage(user string, message string) {
	db := gormConnect()
	defer db.Close()

	var chatlog structs.ChatLog

	chatlog.Username = user
	chatlog.Text = message

	db.Table("chat_log").Create(&chatlog)
}

func UpdateUser(olduser string, newuser structs.User) {
	db := gormConnect()
	defer db.Close()

	var userinfo structs.User

	pass, _ := bcrypt.GenerateFromPassword([]byte(newuser.Password), bcrypt.DefaultCost)

	userinfo.Username = newuser.Username
	userinfo.Password = string(pass)

	db.Table("members").Where("username = ?", olduser).Update(&userinfo)
}

func DeleteUser(user string) {
	db := gormConnect()
	defer db.Close()

	var userinfo structs.User
	fmt.Println(user)

	db.Table("members").Where("username = ?", user).Delete(&userinfo)
	//クエリパラメータで指定したユーザを削除 この場合は指定したuserの情報が全て抹消される
}
