package middleware

import (
	"log"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	next http.Handler
}

func NewLoggingMiddleware(next http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{next: next}
}

func (lm *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lm.next.ServeHTTP(w, r)
	log.Printf("%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
}
