package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "4")
	flag.Parse()
}

func main() {
	glog.V(2).Info("Starting server on port 8080")
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		glog.V(2).Info(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("\nEntering root handler")

	// Get VERSION
	version := os.Getenv("VERSION")
	glog.V(4).Info("Env VERSION: ", version)
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
	glog.V(2).Info("Request IP: ", ip)
	glog.V(2).Info("Status: ", http.StatusOK)
	w.WriteHeader(http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("\nEntering health handler")
	glog.V(2).Info("Status: ", http.StatusOK)
	w.WriteHeader(http.StatusOK)
}
