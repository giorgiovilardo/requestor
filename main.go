package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func dbStatusCheck() func(w http.ResponseWriter, r *http.Request) {
	var lastTime time.Time
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			response := []byte(fmt.Sprintf("Last seen at %v", lastTime))
			w.Write(response)
		case "POST":
			lastTime = time.Now()
			w.Write([]byte("Seen"))
		default:
			w.Write([]byte("Sorry, only GET and POST methods are supported."))
		}
	}
}

func main() {
	http.HandleFunc("/status", dbStatusCheck())

	fmt.Printf("Starting server!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
