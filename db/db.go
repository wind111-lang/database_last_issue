package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {

	user := "user"
	password := "hoge"
	hostname := "localhost"
	name := "chatdb"

	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+hostname+")/"+name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")
	return db
} //DBに接続

func GetDatabase(data string) error {
	db := gormConnect()
	defer db.Close()

	return nil
}
