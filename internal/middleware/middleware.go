package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// ServeHTTP times the duration of the next handler
func ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(time.Since(start))
	})
}
