package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You are on the homepage!",
	})
}
