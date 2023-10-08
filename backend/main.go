package main

import (
	"log"
	"os"
	"todo/component/appctx"
	"todo/middleware"
	"todo/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=postgres password=0000 dbname=food_delivery port=1444 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize: 10,
		Logger:          logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	secretKey := os.Getenv("APP_SECRET_KEY")

	appCtx := appctx.NewAppContext(db, nil, secretKey)

	r := gin.Default()
	r.Static("/static", "./static")
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	routes.SetupUserRoute(appCtx, v1)

	r.Run()
}
