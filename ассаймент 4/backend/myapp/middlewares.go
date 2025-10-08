// В файле middleware.go
package main

import (
	"log"
	"net/http"
)

// Пример middleware для логирования
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Пример CSRF защиты middleware
func CSRFProtectionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CSRF защита
		next.ServeHTTP(w, r)
	})
}

// Пример middleware для заголовков безопасности
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		next.ServeHTTP(w, r)
	})
}
