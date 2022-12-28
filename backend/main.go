package	main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ykinchan/Oshiire/backend/controller"
	"github.com/ykinchan/Oshiire/backend/middleware"
)
func main() {
	engine := gin.Default()

	engine.Use(middleware.RecordUaAndTime)

	blockEngine :=engine.Group("/block")
	{
		v1 := blockEngine.Group("/v1")
		{
			v1.POST("/add", controller.BlockAdd)
			v1.GET("/list", controller.BlockList)
		}
	}
	engine.Run(":8000")
}