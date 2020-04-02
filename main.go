package main

import (
	"github.com/VictorDebray/elysian_account/account"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"net/http"
	"os"
)

func initDB() *gorm.DB{
	db, err := gorm.Open("postgres",
		"host="+os.Getenv("DB_HOST")+
		" port="+os.Getenv("DB_PORT")+
		" user="+os.Getenv("DB_USER")+
		" dbname="+os.Getenv("DB_NAME")+
		" password="+os.Getenv("DB_PWD")+
		" sslmode="+os.Getenv("DB_SSLMODE")+
		" connect_timeout=3")
	if err != nil {
		panic(err)
	}

	db.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;")
	db.AutoMigrate(&account.Account{})

	return db
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		if sessionID == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "not authed",
			})
			c.Abort()
		}
	}
}

func main() {
	db := initDB()
	defer db.Close()

	accountAPI := InitAccountAPI(db)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	store, _ := redis.NewStore(10, "tcp",
		os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PWD"), []byte("secret"))

	router.Use(sessions.Sessions("user_session", store))

	acc := router.Group("/account")
	acc.POST("/", accountAPI.Create)

	authAcc := router.Group("/account")
	authAcc.GET("/", accountAPI.FindAll)

	err := router.Run("localhost:"+os.Getenv("SVC_PORT"))
	if err != nil {
		panic(err)
	}
}