package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

// rateLimit is a global rate limiter, limits the amount of client requests.
func (app *application) rateLimit(next http.Handler) http.Handler {
	// Allow 1 requests per second with a maximum of 2 requests in a single 'burst'.
	limiter := rate.NewLimiter(1, 2)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Everytime Allow() is called one token will be consumed from the bucket.
		// If no tokens are left in the bucket, allow returns false.
		if !limiter.Allow() {
			http.Error(w, "rate limit exceeded", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
