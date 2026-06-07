package routes

import (
	"net/http"

	"github.com/nitankagogoi/go-tutorials/webservers/servers/handlers"
)

func RegisterSortingRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/sort/bubble", handlers.BubbleSortHandler)
	mux.HandleFunc("/sort/merge", handlers.MergeSortHandler)
	mux.HandleFunc("/sort/selection", handlers.SelectionSortHandler)
}
