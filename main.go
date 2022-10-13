package main

import (
	"log"
	"os"

	"github.com/AustinNick/coba-golang/app/config"
	routes "github.com/AustinNick/coba-golang/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectMySQL()

func main() {
	defer config.CloseDB(db)
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	router := NewRouter()
	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func NewRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	if os.Getenv("GO_MODE") == "test" {
		gin.SetMode(gin.TestMode)
	}

	routes.NewUserRoute(db, router)
	router.Use(cors.Default())

	return router
}
