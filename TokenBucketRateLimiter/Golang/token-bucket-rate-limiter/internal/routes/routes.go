package routes

import (
	"github.com/gin-gonic/gin"
	tokenbucket "token-bucket-rate-limiter/pkg/token-bucket"
)

func initialiseRouter(router *gin.Engine) {
	router.GET("/home", tokenbucket.RateLimiter, RouteMap["/home"])
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	initialiseRouter(router)
	return router
}
