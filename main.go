package main

import (
	"blog-api/modules/base"
	"blog-api/modules/cronjob"
	"blog-api/modules/post"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appMode := string(os.Getenv("APP_MODE"))

	if len(appMode) > 0 {
		gin.SetMode(appMode)
	}

	session := base.InitDB(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	defer session.Close()

	router := gin.Default()

	modules := []base.ModuleInterface{post.Module{}, cronjob.Module{}}

	for _, module := range modules {
		base.InitModule(router, module)
	}

	router.Run(os.Getenv("APP_PORT"))
}
