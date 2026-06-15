package main

import (
	"fmt"
	"net/http"

	"github.com/nitankagogoi/go-tutorials/webservers/routes"
)

func main() {
	fmt.Println("Starting the BST server")

	mux := http.NewServeMux()
	routes.RegisterBSTRoutes(mux)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)

}
