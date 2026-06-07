package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nitankagogoi/go-tutorials/webservers/internal/sorting"
)

func MergeSortHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var numbers []int

	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()

	err := decoder.Decode(&numbers)

	fmt.Println("Received request to sort numbers ", numbers)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	length := len(numbers)

	fmt.Println("Length of numbers: ", length)

	sorting.MergeSort(numbers, 0, length-1)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(numbers)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
