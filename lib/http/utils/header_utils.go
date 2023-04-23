package modularHTTP

import "net/http"

func SetContentJson(w http.ResponseWriter) {
	w.Header().Set("Content-type", "Application/json")

}
