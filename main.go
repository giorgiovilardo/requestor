package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func dbStatusCheck(lastTime *time.Time) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			response := []byte(fmt.Sprintf("<h2>Last seen at %v</h2>", *lastTime))
			w.Write(response)
		case "POST":
			*lastTime = time.Now()
			w.Write([]byte("Seen"))
		default:
			w.Write([]byte("Sorry"))
		}
	}
}

func resetStatusDb(lastTime *time.Time) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			*lastTime = time.Time{}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			response := []byte("<h2>Reset</h2>")
			w.Write(response)
		default:
			w.Write([]byte("Sorry"))
		}
	}
}

func main() {
	var lastTime time.Time
	http.HandleFunc("/status", dbStatusCheck(&lastTime))
	http.HandleFunc("/reset", resetStatusDb(&lastTime))

	fmt.Printf("Starting server!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
