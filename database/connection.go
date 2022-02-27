package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/richaidos/SupportSection/globals"
)

var db *gorm.DB
var err error

func Init() {
	user := globals.ISettings.Config.Database.User
	password := globals.ISettings.Config.Database.Password
	host := globals.ISettings.Config.Database.Host
	port := globals.ISettings.Config.Database.Port
	database := globals.ISettings.Config.Database.Dbname

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	fmt.Println("Database connected")
	/*
	   if !db.HasTable(&models.Task{}) {
	     err := db.CreateTable(&models.Task{})
	     if err != nil {
	       log.Println("Table already exists")
	     }
	   }

	   db.AutoMigrate(&models.Task{})
	*/
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
