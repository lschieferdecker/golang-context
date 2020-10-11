package main

import (
	"context"
	"net/http"
)

func addContextValueMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "Name", "Lucas")
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
