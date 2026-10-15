package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	VERSION = 0.1
)

type PageInformation struct {
	Version float32 `json:"version"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	info := PageInformation {
		Version: VERSION,
	}
	info_json, err := json.Marshal(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Serialization error: %v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(info_json))
	}
}

func main() {		
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	http.HandleFunc("/health", Health)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error server start:", err)
	}
}
