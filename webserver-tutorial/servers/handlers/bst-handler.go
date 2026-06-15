package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nitankagogoi/go-tutorials/webservers/internal/trees"
)

type BSTHandler struct {
	bst *trees.BST
}

func NewBSTHandler() *BSTHandler {
	return &BSTHandler{bst: &trees.BST{}}
}

func (h *BSTHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var numbers int

	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()

	err := decoder.Decode(&numbers)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	h.bst.AddNode(numbers)
	w.WriteHeader(http.StatusCreated)
}

func (h *BSTHandler) Display(w http.ResponseWriter, r *http.Request) {
	nodes := h.bst.InOrder()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nodes)
}

func (h *BSTHandler) Search(w http.ResponseWriter, r *http.Request) {
	var numbers int

	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Decode(&numbers)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	node := h.bst.Search(numbers)
	w.Header().Set("Content-Type", "application/json")
	if node == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]bool{"found": false})
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"found": true})
}
