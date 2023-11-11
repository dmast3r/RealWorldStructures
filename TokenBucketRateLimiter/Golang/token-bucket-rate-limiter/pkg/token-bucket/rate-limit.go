package token_bucket

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rule struct {
	MaxTokens float64
	Rate      float64
}

var rulesMap = map[string]Rule{
	"gen-user": {MaxTokens: 1, Rate: 5},
}

var clientMap = map[string]*TokenBucket{}

/*
At the moment, this function simply fetches the rules from an in-memory hash-map.
However, for a production rate limiter, it could be modified to fetch it from a centralized data-store.
For example, a database.
*/
func getRule(userId string) Rule {
	return rulesMap[userId]
}

func getBucket(userId string) *TokenBucket {
	if _, exists := clientMap[userId]; !exists {
		userRule := getRule(userId)
		clientMap[userId] = newTokeBucket(userRule.Rate, userRule.MaxTokens)
	}

	return clientMap[userId]
}

func RateLimiter(ctx *gin.Context) {
	if getBucket(ctx.ClientIP()).shouldRateLimit(1) {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Hold there, you sparky!",
		})
		return
	}

	ctx.Next()
}
