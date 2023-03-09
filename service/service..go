package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cao(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"cao": "haha",
	})
}
