// api/rate_limiter.go

package api

import (
	"net/http"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

type RateLimiter struct {
	requests     *lru.Cache
	mutex        sync.Mutex
	rateLimit    int
	rateLimitDur time.Duration
}

func NewRateLimiter(size int, rateLimit int, rateLimitDur time.Duration) *RateLimiter {
	cache, _ := lru.New(size)
	return &RateLimiter{
		requests:     cache,
		rateLimit:    rateLimit,
		rateLimitDur: rateLimitDur,
	}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mutex.Lock()
		defer rl.mutex.Unlock()

		ip := r.RemoteAddr
		if val, ok := rl.requests.Get(ip); ok {
			reqInfo := val.(*requestInfo)
			if time.Since(reqInfo.timestamp) < rl.rateLimitDur {
				if reqInfo.count >= rl.rateLimit {
					http.Error(w, "Too many requests", http.StatusTooManyRequests)
					return
				}
				reqInfo.count++
			} else {
				reqInfo.count = 1
				reqInfo.timestamp = time.Now()
			}
		} else {
			rl.requests.Add(ip, &requestInfo{count: 1, timestamp: time.Now()})
		}

		next.ServeHTTP(w, r)
	})
}

type requestInfo struct {
	count     int
	timestamp time.Time
}
