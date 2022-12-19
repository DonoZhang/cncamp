package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nEntering root handler")

	// Get VERSION
	version := os.Getenv("VERSION")
	fmt.Println("Env VERSION: ", version)
	// SET Response Headers
	for key, value := range r.Header {
		w.Header().Set(key, value[0])
	}
	w.Header().Set("VERSION", version)

	// Log Request IP
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	fmt.Println("Request IP: ", ip)
	fmt.Println("Status: ", http.StatusOK)
	w.WriteHeader(http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nEntering health handler")
	fmt.Println("Status: ", http.StatusOK)
	w.WriteHeader(http.StatusOK)
}
