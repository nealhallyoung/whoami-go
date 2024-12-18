package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	port    string
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.StringVar(&port, "port", getEnv("WHOAMI_PORT_NUMBER", "80"), "give me a port number")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handle(textHandler, verbose))
	mux.Handle("/json", handle(jsonHandler, verbose))
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Server starting on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func handle(next http.HandlerFunc, verbose bool) http.Handler {
	if !verbose {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		log.Printf(`{"remote_addr": "%s", "time": "%s", "method": "%s", "path": "%s", "protocol": "%s"}`,
			r.RemoteAddr,
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
			r.Proto,
		)
	})
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)

	response := map[string]string{
		"ip": ip,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	log.Printf(`{"remote_addr": "%s", "time": "%s", "method": "%s", "path": "%s", "protocol": "%s"}`,
		r.RemoteAddr,
		time.Now().Format(time.RFC3339),
		r.Method,
		r.URL.Path,
		r.Proto,
	)
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	response := ip
	if _, err := w.Write([]byte(response)); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}

	log.Printf(`{"remote_addr": "%s", "time": "%s", "method": "%s", "path": "%s", "protocol": "%s"}`,
		r.RemoteAddr,
		time.Now().Format(time.RFC3339),
		r.Method,
		r.URL.Path,
		r.Proto,
	)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getClientIP(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	// IPv6
	if strings.HasPrefix(remoteAddr, "[") {
		endIdx := strings.LastIndex(remoteAddr, "]")
		if endIdx > 0 {
			return remoteAddr[1:endIdx]
		}
	}

	// IPv4
	parts := strings.Split(remoteAddr, ":")
	return parts[0]
}
