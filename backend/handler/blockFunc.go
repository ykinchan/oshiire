package handler

import (
	"net/http"

	"../block"

	"github.com/gin-gonic/gin"
)

func BlocksGet(blocks *block.Blocks) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := blocks.GetAll()
		c.JSON(http.StatusOK, result)
	}
}