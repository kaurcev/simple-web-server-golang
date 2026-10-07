package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		interceptor := &responseWriterInterceptor{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(interceptor, r)
		duration := time.Since(start)
		log.Printf("[%s] %d | %13v | %s", r.Method, interceptor.statusCode, duration, r.URL.Path)
	})
}
