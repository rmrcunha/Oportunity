package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatepeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
