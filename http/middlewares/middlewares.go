package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("R.METHOD =", r.Method)
		log.Println("R.HOST =", r.Host)
		log.Println("R.URL =", r.URL.Path)
		log.Println("R.UserAgent =", r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
