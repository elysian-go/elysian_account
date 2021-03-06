package main

import (
	"github.com/VictorDebray/elysian_account/account"
	"github.com/elysian-go/redis-sentinel-store/redisstore"
	"github.com/gin-contrib/location"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"os"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host="+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USER")+
			" dbname="+os.Getenv("DB_NAME")+
			" password="+os.Getenv("DB_PWD")+
			" sslmode="+os.Getenv("DB_SSLMODE")+
			" connect_timeout="+os.Getenv("DB_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	//Todo remove this line for production
	db.LogMode(true)
	db.SingularTable(true)
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
		c.Set("user_id", sessionID)
	}
}

func main() {
	db := initDB()
	defer db.Close()

	store := redisstore.InitStore()
	accountAPI := InitAccountAPI(db)
	authAPI := InitAuthAPI(db)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Use(sessions.Sessions("user_session", store))
	router.Use(location.Default())

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		auth.POST("/login", authAPI.Login)
		auth.GET("/logout", authAPI.Logout)

		acc := v1.Group("/account")
		acc.POST("", accountAPI.Create)

		authAcc := v1.Group("/account")
		authAcc.Use(AuthRequired())
		authAcc.GET("", accountAPI.FindAll)
		authAcc.PATCH("", accountAPI.Update)
	}
	err := router.Run(":" + os.Getenv("SVC_PORT"))
	if err != nil {
		panic(err)
	}
}
