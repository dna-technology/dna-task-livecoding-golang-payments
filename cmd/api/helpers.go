package main

import (
	"encoding/json"
	"net/http"
)

func responseJson(w http.ResponseWriter, status int, data any) {
	js, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
