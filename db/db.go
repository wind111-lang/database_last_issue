package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto"
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
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")
	return db
} //DBに接続

func CreateUser(username string, password string, birthday string) {
	db := gormConnect()
	defer db.Close()

	pass, _ := crypto.PasswordEncrypt(password)
}

// func GetDatabase() error {
// 	db := gormConnect()
// 	defer db.Close()

// 	// loc := []Loc{}

// 	// if err := db.Table("azureapi").Find(&loc, "location=?", "japaneast").Error; err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Println(loc)
// 	return nil
// }
