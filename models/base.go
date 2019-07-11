package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"os"
)

// Database
var db *gorm.DB

func init() {
	usePostgres := false

	var conn *gorm.DB
	var err error

	if usePostgres {
		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}
	
		username := os.Getenv("db_user")
		password := os.Getenv("db_pass")
		dbName := os.Getenv("db_name")
		dbHost := os.Getenv("db_host")

		dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
		fmt.Println(dbUri)

		conn, err = gorm.Open("postgres", dbUri)
	} else {
		conn, err = gorm.Open("sqlite3", "data.db")
	}
	

	if err != nil {
		fmt.Print("Connect to db failed: " + err.Error())
	}

	db = conn

	// Database migration
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

/*
 GetDB Get a db instance
*/
func GetDB() *gorm.DB {
	return db
}
