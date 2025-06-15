package main

import (
	"gozam/db"
	"gozam/routers"
	"log"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing...")
	}

	//gin setup
	router := gin.New()

	//redis connection
	redisClient, err := db.NewRedisClient()
	if err != nil {
		log.Fatal("failed to establish redis connection")
	}
	// log.Print(redisClient)

	//postgres connection
	var postgresClient *gorm.DB
	postgresClient, err = db.NewPostgresClient()
	if err != nil {
		log.Printf("failed to establish redis connection: %v", err)
	}
	// log.Print(postgresClient)

	routers.UserRouter(router)
	router.Run(":8000")
}
