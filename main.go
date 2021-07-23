package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, req *http.Request) {
	delay, err := strconv.Atoi(getenv("DELAY", "0"))
	if err == nil {
		fmt.Println(delay)
	}
	time.Sleep(time.Duration(delay) * 1000 * time.Millisecond)
	w.Header().Add("Server", "nginx/1.21.1") // normal header
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, getenv("CONTENT", "Set CONTENT env var for custom response string"))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/headers", headers)
	port := getenv("PORT", ":5000")
	log.Printf("Listening on port : %s\n", port)
	http.ListenAndServe(port, nil)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
