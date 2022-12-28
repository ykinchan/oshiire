package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ykinchan/Oshiire/backend/model"
	"github.com/ykinchan/Oshiire/backend/service"
)

// ブロックを追加する
func BlockAdd(c *gin.Context) {
	block := model.Block{}
	err := c.Bind(&block)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	blockService := service.BlockService{}
	err = blockService.SetBlock(&block)

	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// ブロックを取得する
func BlockList(c *gin.Context) {
	blockService := service.BlockService{}
	BlockLists := blockService.GetBlockList()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": BlockLists,
	})
}