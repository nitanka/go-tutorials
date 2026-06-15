package routes

import (
	"net/http"

	"github.com/nitankagogoi/go-tutorials/webservers/servers/handlers"
)

func RegisterBSTRoutes(mux *http.ServeMux) {
	bstHandler := handlers.NewBSTHandler()
	mux.HandleFunc("/trees/insert", bstHandler.Insert)
	mux.HandleFunc("/trees/display", bstHandler.Display)
	mux.HandleFunc("/trees/search", bstHandler.Search)
}
