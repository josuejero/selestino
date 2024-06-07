// api/rate_limiter_test.go

package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/josuejero/selestino/api"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rateLimiter := api.NewRateLimiter(100, 2, time.Second)
	limitedHandler := rateLimiter.Limit(handler)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// First request should pass
	limitedHandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Second request should pass
	rr = httptest.NewRecorder()
	limitedHandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Third request should be rate limited
	rr = httptest.NewRecorder()
	limitedHandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)
}
