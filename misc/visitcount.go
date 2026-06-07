package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	visitCount int
	mu         sync.Mutex
)

func visitCountHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	visitCount++
	currentCount := visitCount
	fmt.Fprintf(w, "You are visitor number %d\n", currentCount)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", visitCountHandler)
	fmt.Println("Starting the server at 8080")
	http.ListenAndServe(":8080", nil)
}
