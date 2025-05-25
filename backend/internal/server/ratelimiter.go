package server

import (
	"encoding/json"
	"net"
	"net/http"
	"socialNetwork/entity"
	"sync"
	"time"
)

type RateLimiter struct {
	Requests   int           // Maximum number of requests allowed
	Interval   time.Duration // Time frame for the rate limit
	sync.Mutex               // Mutex for concurrent access
	clients    map[string]*clientInfo
}

type clientInfo struct {
	requestCount int
	firstRequest time.Time
}

func NewRateLimiter(requests int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		Requests: requests,
		Interval: interval,
		Mutex:    sync.Mutex{},
		clients:  make(map[string]*clientInfo),
	}
}

func (rl *RateLimiter) Limiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get just the host without the ip:
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		clientIP := host
		info, exists := rl.clients[clientIP]

		if !exists {
			info = &clientInfo{
				requestCount: 0,
				firstRequest: time.Now(),
			}
			rl.Lock()
			rl.clients[clientIP] = info
			rl.Unlock()
		}

		// Check if the time interval has passed
		if time.Since(info.firstRequest) > rl.Interval {
			info.requestCount = 0
			info.firstRequest = time.Now()
		}

		// Increment the request count
		info.requestCount++

		// Check if the request count exceeds the limit
		if info.requestCount > rl.Requests {

			w.WriteHeader(http.StatusTooManyRequests)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(entity.ErrorResponse{
				Error: "Rate limit exceeded",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
