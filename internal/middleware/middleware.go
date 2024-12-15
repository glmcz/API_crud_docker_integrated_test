package middleware

import "net/http"

func SecureHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent click jacking
		w.Header().Set("X-Frame-Options", "DENY")

		// Prevent MIME type sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// unsafe can be fixed by moving css into file.
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline';")

		// Prevent XSS attacks
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
