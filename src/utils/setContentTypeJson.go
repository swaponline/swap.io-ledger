package utils

import "net/http"

func SetContentTypeJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
