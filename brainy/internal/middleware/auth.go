package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	next http.Handler
}

func NewAuthMiddleware(next http.Handler) *AuthMiddleware {
	return &AuthMiddleware{next: next}
}

func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-User")
	if user == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	a.next.ServeHTTP(w, r)
}
