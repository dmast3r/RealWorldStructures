package token_bucket

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	rate                float64
	maxTokens           float64
	currentTokens       float64
	lastRefillTimestamp time.Time
	mutex               sync.Mutex
}

func newTokeBucket(rate, maxTokens float64) *TokenBucket {
	return &TokenBucket{rate: rate, maxTokens: maxTokens, currentTokens: maxTokens, lastRefillTimestamp: time.Now()}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTimestamp).Seconds()
	tokensToAdd := elapsed * tb.rate

	tb.currentTokens = math.Min(tb.currentTokens+tokensToAdd, tb.maxTokens)
	tb.lastRefillTimestamp = now
}

func (tb *TokenBucket) shouldRateLimit(requiredTokens float64) bool {
	tb.refill()

	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	if tb.currentTokens < requiredTokens {
		return true
	}

	tb.currentTokens -= requiredTokens
	return false
}
