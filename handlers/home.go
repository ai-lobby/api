package handlers

import "net/http"

func GetHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})
}
