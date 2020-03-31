package main

import (
	"github.com/VictorDebray/elysian_account/account"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDB() *gorm.DB{
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=docker sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&account.Account{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	accountAPI := InitAccountAPI(db)
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", accountAPI.FindAll)
	router.POST("/somePost", accountAPI.Create)

	err := router.Run()
	if err != nil {
		panic(err)
	}
	// router.Run(":3000") for a hard coded port
}