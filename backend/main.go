package	main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"/run/controller"
	"/run/middleware"
)
func main() {
	engine := gin.Default()

	engine.Use(middleware.RecordUaAndTime)

	BlockEngine :=engine.Group("/block"){
		v1 := blockEngine.Group("/v1"){
			v1.POST("/add", controller.BlockAdd)
			v1.GET("/list", controller.BlockList)
		}
	}
	engine.Run(":8000")
}