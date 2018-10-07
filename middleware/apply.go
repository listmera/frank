package middleware

import "net/http"

func ApplyMiddleware(handler http.Handler) http.Handler {
	for _, mid := range stack {
		handler = mid(handler)
	}
	return handler
}
