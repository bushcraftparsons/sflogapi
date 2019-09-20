package models

import (
	"fmt"
	"os"

	//See example at https://gorm.io/docs/connecting_to_the_database.html
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database

func init() {

	username := os.Getenv("user")
	password := os.Getenv("password")
	dbName := os.Getenv("dbname")
	dbHost := os.Getenv("host")
	port := os.Getenv("port")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, port, username, dbName, password) //Build connection string
	fmt.Println("Connecting to db")
	//"postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
