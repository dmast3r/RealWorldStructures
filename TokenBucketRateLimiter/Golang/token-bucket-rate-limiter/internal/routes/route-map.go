package routes

import (
	"github.com/gin-gonic/gin"
	"token-bucket-rate-limiter/internal/routes/handlers"
)

var RouteMap = map[string]gin.HandlerFunc{
	"/home": handlers.Home,
}
