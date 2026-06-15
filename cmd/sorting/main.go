package main

import (
	"fmt"
	"net/http"

	"github.com/nitankagogoi/go-tutorials/webservers/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterSortingRoutes(mux)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
