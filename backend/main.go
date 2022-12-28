package main

import (
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/ykinchan/Oshiire/backend/block"
    "github.com/ykinchan/Oshiire/backend/handler"
    "github.com/ykinchan/Oshiire/backend/lib"
    "github.com/joho/godotenv"

    "github.com/gin-contrib/cors"
)

func main() {
	block := block.New()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
		},
		AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
		},
		AllowCredentials: true,
		MaxAge: 24*time.Hour,
	}))

	r.GET("/block", handler.BlocksGet(block))

	r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
}